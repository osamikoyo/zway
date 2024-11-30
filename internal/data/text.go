package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand/v2"
	"zway/internal/data/models"
)

func setUniqKey() string {
	var result string
	simbols := "WERTYUIOPLKLJHGFDSAZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890!£$%^&*(#"
	for i := 0; i < 50; i++ {
		result = result + string(simbols[rand.IntN(len(simbols)-1)])
	}

	return result
}

func AddText(text models.Text) error {
	db, err := gorm.Open(sqlite.Open("storage/main.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	id := setUniqKey()

	text.ID = id

	if err := db.Create(&text).Error; err != nil {
		return err
	}

	return nil
}

func GetText(uniq string) (models.Text, error) {
	db, err := gorm.Open(sqlite.Open("storage/main.db"), &gorm.Config{})
	if err != nil {
		return models.Text{}, err
	}

	var text models.Text

	if err := db.Where(models.Text{ID: uniq}).Find(&text).Error; err != nil {
		return models.Text{}, err
	}

	return text, nil
}
