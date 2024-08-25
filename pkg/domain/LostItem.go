package domain


type LostItem struct {
    LostTime string `json:"lost_time" form:"lost_time"`
    Kind      string `json:"kind" form:"kind"`
    PropertyName string `json:"property_name" form:"property_name"`
    Location   string `json:"location" form:"location"`
    PhoneNumber string `json:"phone_number" form:"phone_number"`
}

type LostItemRepository interface {
    Create(LostItem *LostItem) (*LostItem, error)
    // GetAll() ([]*LostItem, error)
    // GetByID(id uint) (*LostItem, error)
    // Update(LostItem *LostItem) (*LostItem, error)
    // Delete(id uint) error
}

type LostItemService interface {
    AddNewLostItem(lostItem *LostItem) (*LostItem, error)
    // GetAllLostItems() ([]*LostItem, error)
    // GetLostItemById(id uint) (*LostItem, error)
    // UpdateLostItem(LostItem *LostItem) (*LostItem, error)
    // DeleteLostItem(id uint) error
}

