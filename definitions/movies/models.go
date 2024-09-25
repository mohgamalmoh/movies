package movies

type Movie struct {
	Id       string `gorm:"column:id;autoIncrement;primary_key"`
	Name     string `gorm:"type:varchar(200)"`
	Genre    string `gorm:"type:varchar(255)"`
	Year     string `gorm:"type:varchar(255)"`
	Overview string `gorm:"type:text"`
}
