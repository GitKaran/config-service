package database

import (
	"github.com/hellofreshdevtests/GitKaran-devops-test/models"
	"github.com/jinzhu/gorm"
	"strings"
)

func GetAllConfigs(db *gorm.DB) ([]models.Config, error) {
	configs := []models.Config{}
	query := db.Select("configs.*").
		Group("configs.id")
	if err := query.Find(&configs).Error; err != nil {
		return configs, err
	}

	return configs, nil
}

func GetConfigByName(name string, db *gorm.DB) (models.Config, bool, error) {
	config := models.Config{}

	query := db.Select("configs.*")
	query = query.Group("configs.id")
	err := query.Where("configs.name = ?", name).Find(&config).Error
	if err != nil || gorm.IsRecordNotFoundError(err) {
		return config, false, err
	}
	return config, true, nil
}

func SearchConfigByKey(key string, value string, db *gorm.DB) ([]models.Config, bool, error) {
	configs := []models.Config{}
	pathkey := "$" + strings.TrimPrefix(key, "metadata")

	err := db.Raw("select * from configs where json_extract(metadata, $1) = $2", pathkey, value).Scan(&configs).Error
	if err != nil || gorm.IsRecordNotFoundError(err) {
		return configs, false, err
	}

	return configs, true, nil
}

func DeleteConfig(name string, db *gorm.DB) error {
	var config models.Config
	if err := db.Where("name = ? ", name).Delete(&config).Error; err != nil {
		return err
	}
	return nil
}

func UpdateConfig(existingConfig *models.Config, updatedConfig *models.Config, db *gorm.DB) error {
	if err := db.First(&existingConfig).Update(&updatedConfig).Error; err != nil {
		return err
	}
	return nil
}
