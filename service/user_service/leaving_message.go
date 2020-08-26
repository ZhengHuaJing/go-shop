package user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistLeavingMessageByID(lm model.LeavingMessage) (bool, error) {
	if err := global.DB.Where("id = ?", lm.ID).First(&lm).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddLeavingMessage(lm model.LeavingMessage) (*model.LeavingMessage, error) {
	if err := global.DB.Create(&lm).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("user_id = ? AND tile = ? AND content = ?", lm.UserID, lm.Title, lm.Content).
		First(&lm).Error; err != nil {
		return nil, err
	}

	return &lm, nil
}

func DeleteLeavingMessage(lm model.LeavingMessage) error {
	if err := global.DB.Delete(model.LeavingMessage{}, "id = ?", lm.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteLeavingMessage() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.LeavingMessage{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateLeavingMessage(lm model.LeavingMessage) (*model.LeavingMessage, error) {
	if err := global.DB.Model(model.LeavingMessage{}).Where("id = ?", lm.ID).Updates(lm).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", lm.ID).First(&lm).Error; err != nil {
		return nil, err
	}

	return &lm, nil
}

func GetLeavingMessage(lm model.LeavingMessage) (*model.LeavingMessage, error) {
	if err := global.DB.Where("id = ?", lm.ID).First(&lm).Error; err != nil {
		return nil, err
	}

	return &lm, nil
}

func GetLeavingMessages(lm model.LeavingMessage, pageNum, pageSize int) ([]model.LeavingMessage, error) {
	var lms []model.LeavingMessage

	err := global.DB.Where(lm).Offset(pageNum).Limit(pageSize).Find(&lms).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return lms, nil
}

func GetLeavingMessageTotal(lm model.LeavingMessage) (int, error) {
	var count int

	if err := global.DB.Model(&model.LeavingMessage{}).Where(lm).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
