package internal

func TestDB() {

	DB.AutoMigrate(&User{}, &Message{})
	testMsg := Message{
		Content: "test message",
	}

	DB.Create(&testMsg) // pass pointer of data to Create
}

func Migrate() {
	DB.AutoMigrate(&User{}, &Message{})
}
