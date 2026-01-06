package model
 import ("gorm.io/gorm")

 type Booking struct {
	gorm.Model
	CustomerName string `valid:"required~Customer name is required"`
	RoomNumber   int	 `valid:"range(101|999),required~Invalid room number"`
	GuestCount   int	`valid:"range(1|4),required~Guest count must be between 1 and 4"`
	PhoneNumber  string `valid:"stringlength(10|10),required~Phone number is required"`
 }