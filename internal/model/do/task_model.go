package do

type Task struct {
	ID      int    `gorm:"column:id; primary_key; not null" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Status  int    `gorm:"column:status" json:"status"`
	OwnerID int    `gorm:"column:owner_id" json:"owner_id"`
	BaseModel
}
