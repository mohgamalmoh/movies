package users

type User struct {
	Id    int    `gorm:"type:int;primary_key"`
	Name  string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255)"`
}

type UsersMovies struct {
	Id    int `gorm:"type:int;primary_key"`
	User  int `gorm:"type:int"`
	Movie int `gorm:"type:int"`
}
