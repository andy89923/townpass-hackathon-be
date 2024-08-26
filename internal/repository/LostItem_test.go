package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"go.uber.org/zap"
	"go-cleanarch/pkg/domain"

)

type  LostItemRepoSuite struct {
	suite.Suite
	db *gorm.DB
	repo domain.LostItemRepository
}



func (suite *LostItemRepoSuite) SetupSuite() {
	logger, _ := zap.NewDevelopment()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	} 
		
	suite.T().Log("Database connection established")

	db.AutoMigrate(&LostItem{})
	suite.db = db
	suite.repo = NewPostgresLostItemRepository(db, logger) //let the interface hold the repository object
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

func (suite *LostItemRepoSuite) TestGetAll() {
	lostItems := []domain.LostItem{
		{
			LostTime: "2021-01-01",
			Kind: "Phone",
			PropertyName: "Samsung Galaxy S21",
			Location: "Dining Room",
			PhoneNumber: "081234567890",
		},
		{
			LostTime: "2021-01-02",
			Kind: "Wallet",
			PropertyName: "Leather Wallet",
			Location: "Living Room",
			PhoneNumber: "081234567891",
		},
	}

	for _, lostItem := range lostItems {
		_, err := suite.repo.Create(&lostItem)
		assert.Nil(suite.T(), err)
	}

	allLostItems, err := suite.repo.GetAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), allLostItems)
	assert.Equal(suite.T(), len(lostItems), len(allLostItems))
	
	for i, lostItem := range lostItems {
		assert.Equal(suite.T(), lostItem.LostTime, allLostItems[i].LostTime)
		assert.Equal(suite.T(), lostItem.Kind, allLostItems[i].Kind)
		assert.Equal(suite.T(), lostItem.PropertyName, allLostItems[i].PropertyName)
		assert.Equal(suite.T(), lostItem.Location, allLostItems[i].Location)
		assert.Equal(suite.T(), lostItem.PhoneNumber, allLostItems[i].PhoneNumber)
	}
}

func (suite *LostItemRepoSuite) TestGetByID() {
	lostItem := domain.LostItem{
		LostTime: "2021-01-01",
		Kind: "Phone",
		PropertyName: "Samsung Galaxy S21",
		Location: "Dining Room",
		PhoneNumber: "081234567890",
	}

	createdLostItem, err := suite.repo.Create(&lostItem)
	assert.Nil(suite.T(), err)

	Id := createdLostItem.Id
	suite.T().Log("ID: ", Id)
	foundLostItem, err := suite.repo.GetByID(Id)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), foundLostItem)
	assert.Equal(suite.T(), lostItem.LostTime, foundLostItem.LostTime)
	assert.Equal(suite.T(), lostItem.Kind, foundLostItem.Kind)
	assert.Equal(suite.T(), lostItem.PropertyName, foundLostItem.PropertyName)
	assert.Equal(suite.T(), lostItem.Location, foundLostItem.Location)
	assert.Equal(suite.T(), lostItem.PhoneNumber, foundLostItem.PhoneNumber)
}

func (suite *LostItemRepoSuite) TestUpdate() {
	lostItem := domain.LostItem{
		LostTime: "2021-01-01",
		Kind: "Phone",
		PropertyName: "Samsung Galaxy S21",
		Location: "Dining Room",
		PhoneNumber: "081234567890",
	}

	createdLostItem, err := suite.repo.Create(&lostItem)
	assert.Nil(suite.T(), err)

	Id := createdLostItem.Id
	suite.T().Log("ID: ", Id)

	updatedLostItem := domain.LostItem{
		Id: Id,
		LostTime: "2021-01-02",
		Kind: "Wallet",
		PropertyName: "Leather Wallet",
		Location: "Living Room",
		PhoneNumber: "081234567891",
	}
	updatedLostItemCopy := updatedLostItem

	updateErr := suite.repo.Update(&updatedLostItem)
	assert.Nil(suite.T(), updateErr)
	assert.Nil(suite.T(), updateErr)
	assert.Equal(suite.T(), updatedLostItemCopy.LostTime, updatedLostItem.LostTime)
	assert.Equal(suite.T(), updatedLostItemCopy.Kind, updatedLostItem.Kind)
	assert.Equal(suite.T(), updatedLostItemCopy.PropertyName, updatedLostItem.PropertyName)
	assert.Equal(suite.T(), updatedLostItemCopy.Location, updatedLostItem.Location)	
	assert.Equal(suite.T(), updatedLostItemCopy.PhoneNumber, updatedLostItem.PhoneNumber)
}

func (suite *LostItemRepoSuite) TestDelete() {
	lostItem := domain.LostItem{
		LostTime: "2021-01-01",
		Kind: "Phone",
		PropertyName: "Samsung Galaxy S21",
		Location: "Dining Room",
		PhoneNumber: "081234567890",
	}

	createdLostItem, err := suite.repo.Create(&lostItem)
	assert.Nil(suite.T(), err)

	Id := createdLostItem.Id
	suite.T().Log("ID: ", Id)

	deleteErr := suite.repo.Delete(Id)
	assert.Nil(suite.T(), deleteErr)

	_, err = suite.repo.GetByID(Id)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrNotFound, err)
}