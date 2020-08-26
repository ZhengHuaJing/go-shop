package user_service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistUserByUserName(u model.User) (bool, error) {
	if err := global.DB.Where("user_name = ?", u.UserName).First(&model.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ExistUserByID(u model.User) (bool, error) {
	if err := global.DB.Where("id = ?", u.ID).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddUser(u model.User) (*model.User, error) {
	if err := global.DB.Create(&u).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("user_name = ?", u.UserName).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func DeleteUser(u model.User) error {
	if err := global.DB.Delete(model.User{}, "id = ?", u.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteUser() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.User{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(u model.User) (*model.User, error) {
	if err := global.DB.Model(model.User{}).Where("id = ?", u.ID).Updates(u).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", u.ID).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUser(u model.User) (*model.User, error) {
	if err := global.DB.Where("id = ?", u.ID).First(&u).Error; err != nil {
		return nil, err
	}

	names, err := global.Enforcer.GetRolesForUser(u.UserName)
	if err != nil {
		return nil, err
	}

	u.RoleName = names[0]

	return &u, nil
}

func GetUsers(u model.User, pageNum, pageSize int) ([]model.User, error) {
	var users []model.User

	err := global.DB.Where(u).Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, u := range users {
		names, err := global.Enforcer.GetRolesForUser(u.UserName)
		if err != nil {
			return nil, err
		}

		u.RoleName = names[0]
	}

	return users, nil
}

func GetUserTotal(u model.User) (int, error) {
	var count int

	if err := global.DB.Model(&model.User{}).Where(u).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func UserLogin(u model.User) (*model.User, error) {
	var newUser model.User

	if err := global.DB.Where("user_name = ?", u.UserName).First(&newUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	if u.Password != u.Password {
		return nil, nil
	}

	return &newUser, nil
}
