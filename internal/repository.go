package internal

import (
	"fmt"
	"gorm.io/gorm/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type User struct {
	Name      string `json:"name"`
	Age       uint
	Birthday  time.Time `gorm:"default:2000-07-12"`
	Money     float32
	ID        int       `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}

type Message struct {
	Content   string    `json:"content"`
	ID        int       `json:"id";gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}

type RequestBody struct {
	User User
}

func InitDB() {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	Migrate()
}
