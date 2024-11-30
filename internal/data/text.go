package data

import (
	"log/slog"
	"math/rand/v2"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"zway/internal/data/models"
)

func setUniqKey() string {
	file, _ := os.Open("logs/log.txt")
	loger := slog.New(slog.NewJSONHandler(file, nil))

	var result string
	simbols := "WERTYUIOPLKLJHGFDSAZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890!£$%^&*(#"
	for i := 0; i < 50; i++ {
		result = result + string(simbols[rand.IntN(len(simbols)-1)])
	}

	loger.Info("token generated %s :3", result, "")

	return result
}

func AddText(text models.Text) error {
	file, _ := os.Open("logs/log.txt")
	loger := slog.New(slog.NewJSONHandler(file, nil))

	db, err := gorm.Open(sqlite.Open("storage/main.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	id := setUniqKey()

	text.ID = id

	if err := db.Create(&text).Error; err != nil {
		return err
	}

	loger.Info("text added: %s :3", text.Title,"")

	return nil
}

func GetText(uniq string) (models.Text, error) {
	file,_ := os.Open("logs/log.txt")
	loger := slog.New(slog.NewJSONHandler(file, nil))

	db, err := gorm.Open(sqlite.Open("storage/main.db"), &gorm.Config{})
	if err != nil {
		return models.Text{}, err
	}

	var text models.Text

	if err := db.Where(models.Text{ID: uniq}).Find(&text).Error; err != nil {
		return models.Text{}, err
	}

	loger.Info("text sended: Title : %s \n ID : %s :3", text.Title, text.ID)

	return text, nil
}
