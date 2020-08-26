package user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistVerifyCodeByEmail(vc model.VerifyCode) (bool, error) {
	if err := global.DB.Where("email = ?", vc.Email).First(&model.VerifyCode{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ExistVerifyCodeByID(vc model.VerifyCode) (bool, error) {
	if err := global.DB.Where("id = ?", vc.ID).First(&vc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddVerifyCode(vc model.VerifyCode) error {
	if err := global.DB.Create(&vc).Error; err != nil {
		return err
	}

	return nil
}

func DeleteVerifyCode(vc model.VerifyCode) error {
	if err := global.DB.Delete(model.VerifyCode{}, "id = ?", vc.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteVerifyCode() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.VerifyCode{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateVerifyCode(vc model.VerifyCode) (*model.VerifyCode, error) {
	if err := global.DB.Model(model.VerifyCode{}).Where("id = ?", vc.ID).Updates(vc).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", vc.ID).First(&vc).Error; err != nil {
		return nil, err
	}

	return &vc, nil
}

func GetVerifyCode(vc model.VerifyCode) (*model.VerifyCode, error) {
	if err := global.DB.Where("id = ?", vc.ID).First(&vc).Error; err != nil {
		return nil, err
	}

	return &vc, nil
}

func GetVerifyCodes(vc model.VerifyCode, pageNum, pageSize int) ([]model.VerifyCode, error) {
	var vcs []model.VerifyCode

	err := global.DB.Where(vc).Offset(pageNum).Limit(pageSize).Find(&vcs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return vcs, nil
}

func GetVerifyCodeTotal(vc model.VerifyCode) (int, error) {
	var count int

	if err := global.DB.Model(&model.VerifyCode{}).Where(vc).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func CheckVerifyCode(vc model.VerifyCode) (bool, error) {
	if err := global.DB.Where("code = ? AND email = ?", vc.Code, vc.Email).First(&model.VerifyCode{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
