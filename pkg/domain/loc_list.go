package domain


type LocList struct {
	templeId 		int 	`gorm:"column:temple_id"`
	templeName 		string  `gorm:"column:temple_name"`
	loc				string	`gorm:"column:loc"`
	mainDeity 		int		`gorm:"column:main_deity"`
	history 		string	`gorm:"column:history"`
	worshipOrder 	string	`gorm:"column:worship_order"`
	inCharge 		string	`gorm:"column:in_charge"`
	linkRef 		string	`gorm:"column:link_ref"`
	numsOfSubId 	int		`gorm:"column:nums_of_sub_id"`
}
type LocListRepository interface {
	// TODO
}
// type LocListRepository interface {
//     Create(LostItem *LostItem) (*LostItem, error)
//     GetAll() ([]*LostItem, error)
//     GetByID(id uint) (*LostItem, error)
//     Update(LostItem *LostItem) error
//     Delete(id uint) error
// }

// type LostItemService interface {
//     AddNewLostItem(lostItem *LostItem) (*LostItem, error)
//     GetAllLostItems() ([]*LostItem, error)
//     GetLostItemById(id uint) (*LostItem, error)
//     UpdateLostItem(LostItem *LostItem) (*LostItem, error)
//     DeleteLostItem(id uint) error
// }

