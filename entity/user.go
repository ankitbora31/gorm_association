package entity

type User struct {
	ID      uint
	Name    string
	Gender  string
	Age     int
	Address Address `gorm:"foreignKey:AddressID"`
}
