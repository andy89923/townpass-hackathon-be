package repository

import (
	"errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)


type LocationTable struct {
	gorm.Model

	MM domain.MajorMinor `gorm:"column:item_id"`
	location int `gorm:"column:loc_id"`
	subLocation int `gorm:"column:sub_loc_id"`
}

func (l *LocationTable) TableName() string {
	return "m_m_list"
}

//--------------------------------------

type postgresLocationRepository struct {
	db *gorm.DB
	logger *zap.Logger
}

func NewPostgresLocationRepository(db *gorm.DB, logger *zap.Logger) domain.LocationRepository {
	return &postgresLocationRepository{
		db: db,
		logger: logger,
	}
}

func (r *postgresLocationRepository) GetLocationByMM(mm domain.MajorMinor) (locationId int, subLocationId int, err error) {
	var location LocationTable
	result := r.db.Where(&LocationTable{MM: mm}).Find(&location)
	
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, 0, domain.ErrNotFound	
	}
	
	return location.location, location.subLocation, nil
}

func (r *postgresLocationRepository) Create(location *domain.Location, locationId int, subLocationId int) error {
	locationModel := LocationTable{
		MM: location.MajorMinor,
		location: locationId,
		subLocation: subLocationId,
	}


	result := r.db.Create(&locationModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}