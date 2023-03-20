package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/application-research/delta-db/db_models"
	"github.com/application-research/delta-db/messaging"
	db2 "github.com/application-research/delta-events-consumer/db"
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Instantiate a new consumer that subscribes to the specified topic
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	dbDsn := viper.Get("DBDSN").(string)
	db, err := db2.OpenDatabase(dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	// Instantiate a new consumer that subscribes to the specified topic
	consumer := messaging.NewDeltaMetricsMessageConsumer()

	// add a handler to consume messages
	consumer.Consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var baseMessage messaging.DeltaMetricsBaseMessage
		json.Unmarshal(message.Body, &baseMessage)

		// parse the message body and check the db_models type using reflection
		fmt.Println(baseMessage.ObjectType)
		fmt.Println(baseMessage.Object)

		switch baseMessage.ObjectType {
		case "DeltaStartupLogs":
			baseString := baseMessage.Object.(map[string]interface{})
			var deltaStartupLogs db_models.DeltaStartupLogs
			transcode(baseString, &deltaStartupLogs)
			db.Create(&deltaStartupLogs)
		case "ContentLog":
			baseContentLog := baseMessage.Object.(map[string]interface{})
			var contentLog db_models.ContentLog
			transcode(baseContentLog, &contentLog)
			db.Create(&contentLog)
		case "ContentDealLog":
			baseContentDealLog := baseMessage.Object.(map[string]interface{})
			var contentDealLog db_models.ContentDealLog
			transcode(baseContentDealLog, &contentDealLog)
			db.Create(&contentDealLog)
		case "ContentDealProposalLog":
			baseContentDealLog := baseMessage.Object.(map[string]interface{})
			var contentDealProposalLog db_models.ContentDealProposalLog
			transcode(baseContentDealLog, &contentDealProposalLog)
			db.Create(&contentDealProposalLog)
		case "ContentDealProposalParametersLog":
			baseContentDealLog := baseMessage.Object.(map[string]interface{})
			var contentDealProposalParametersLog db_models.ContentDealProposalParametersLog
			transcode(baseContentDealLog, &contentDealProposalParametersLog)
			db.Create(&contentDealProposalParametersLog)
		case "ContentMinerLog":
			baseContentDealLog := baseMessage.Object.(map[string]interface{})
			var contentMinerLog db_models.ContentMinerLog
			transcode(baseContentDealLog, &contentMinerLog)
			db.Create(&contentMinerLog)
		case "ContentWalletLog":
			baseContentDealLog := baseMessage.Object.(map[string]interface{})
			var contentWalletLog db_models.ContentWalletLog
			transcode(baseContentDealLog, &contentWalletLog)
			db.Create(&contentWalletLog)
		case "PieceCommitmentLog":
			baseContentDealLog := baseMessage.Object.(map[string]interface{})
			var pieceCommitmentLog db_models.PieceCommitmentLog
			transcode(baseContentDealLog, &pieceCommitmentLog)
			db.Create(&pieceCommitmentLog)
		case "WalletLog":
			baseContentDealLog := baseMessage.Object.(map[string]interface{})
			var walletLog db_models.WalletLog
			transcode(baseContentDealLog, &walletLog)
			db.Create(&walletLog)
		case "LogEvent":
			baseLogEvent := baseMessage.Object.(map[string]interface{})
			var logEvent messaging.LogEvent
			transcode(baseLogEvent, &logEvent)
			db.Create(&logEvent)
		case "InstanceMetaLog":
			baseInstanceMetaLog := baseMessage.Object.(map[string]interface{})
			var instanceMetaLog db_models.InstanceMetaLog
			transcode(baseInstanceMetaLog, &instanceMetaLog)
			db.Create(&instanceMetaLog)
		default:
			log.Printf("Unknown message type: %v", baseMessage.ObjectType)
		}

		return nil
	}))

	err = consumer.Consumer.ConnectToNSQD(messaging.MetricsTopicUrl)
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	// disconnect
	consumer.Consumer.Stop()

}

func transcode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
