package db

import (
	"github.com/application-research/delta-db/db_models"
	_ "github.com/application-research/delta-db/db_models"
	"github.com/application-research/delta-db/messaging"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase(dbDsn string) (*gorm.DB, error) {
	// use postgres
	var DB *gorm.DB
	var err error

	if dbDsn[:8] == "postgres" {
		DB, err = gorm.Open(postgres.Open(dbDsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(sqlite.Open(dbDsn), &gorm.Config{})
	}

	// generate new models.
	ConfigureModels(DB) // create models.

	if err != nil {
		return nil, err
	}
	return DB, nil
}

func ConfigureModels(db *gorm.DB) {
	db.AutoMigrate(&db_models.DeltaStartupLogs{}, &messaging.LogEvent{}, &db_models.ContentLog{}, &db_models.PieceCommitmentLog{}, &db_models.ContentDealLog{}, &db_models.ContentDealProposalLog{}, &db_models.ContentDealProposalParametersLog{}, &db_models.ContentWalletLog{}, &db_models.ContentMinerLog{}, &db_models.PieceCommitmentLog{}, &db_models.ContentWalletLog{}, &db_models.WalletLog{}, &db_models.InstanceMetaLog{})
}
