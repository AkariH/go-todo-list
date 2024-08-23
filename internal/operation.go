package internal

func TestDB() {

	DB.AutoMigrate(&User{})
	chino := User{Name: "chino", Age: 14}

	DB.Create(&chino) // pass pointer of data to Create
}
