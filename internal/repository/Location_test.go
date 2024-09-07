package repository

import (
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"go.uber.org/zap"
	"go-cleanarch/pkg/domain"

)

type LocationRepoSuite struct {
	suite.Suite
	db *gorm.DB
	repo domain.LocationRepository
}

func (suite *LocationRepoSuite) SetupSuite() {
	logger, _ := zap.NewDevelopment()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	} 
		
	suite.T().Log("Database connection established")

	db.AutoMigrate(&LocationTable{})
	suite.db = db
	suite.repo = NewPostgresLocationRepository(db, logger) //let the interface hold the repository object
}

func  TestLocationRepoSuite(t *testing.T) {
	suite.Run(t, new(LocationRepoSuite))
}

func (suite *LocationRepoSuite) TestCreate(t *testing.T) {
	location := domain.Location{
		MajorMinor: domain.MajorMinor(1),
	}

	_ = suite.repo.Create(&location, 1, 2) //get the repository object from the underlying interface
	
	// assert.Nil(t, err)
}