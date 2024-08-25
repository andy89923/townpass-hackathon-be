package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"go-cleanarch/pkg/domain"
)

type  LostItemRepoSuite struct {
	suite.Suite
	db *gorm.DB
	repo domain.LostItemRepository
}

func (suite *LostItemRepoSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	} 
		
	suite.T().Log("Database connection established")

	db.AutoMigrate(&LostItem{})
	suite.db = db
	suite.repo = NewPostgresLostItemRepository(db) //let the interface hold the repository object
}

func  TestLostItemRepoSuite(t *testing.T) {
	suite.Run(t, new(LostItemRepoSuite))
}

// -------------- Test Function --------------
func (suite *LostItemRepoSuite) TestCreate() {
	lostItem := domain.LostItem{
		LostTime: "2021-01-01",
		Kind: "Phone",
		PropertyName: "Samsung Galaxy S21",
		Location: "Dining Room",
		PhoneNumber: "081234567890",
	}

	createdLostItem, err := suite.repo.Create(&lostItem) //get the repository object from the underlying interface
	
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), createdLostItem)
	assert.Equal(suite.T(), lostItem.LostTime, createdLostItem.LostTime)
	assert.Equal(suite.T(), lostItem.Kind, createdLostItem.Kind)
	assert.Equal(suite.T(), lostItem.PropertyName, createdLostItem.PropertyName)
	assert.Equal(suite.T(), lostItem.Location, createdLostItem.Location)
	assert.Equal(suite.T(), lostItem.PhoneNumber, createdLostItem.PhoneNumber)
}