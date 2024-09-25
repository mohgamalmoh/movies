package movies_sync_status

type MoviesSyncStatus struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}

func (MoviesSyncStatus) TableName() string {
	return "movies_sync_status"
}
