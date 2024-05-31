package postgres

import (
	"my_module/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) Create(user model.User) {
	u.DB.AutoMigrate(model.User{})

	u.DB.Create(&user)
}

func (u *UserRepo) GetAllUsers() ([]model.User, error) {
	user := []model.User{}
	res := u.DB.Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (u *UserRepo) Update(user model.User) error{
	res := u.DB.Model(&user).Where("id = ?", user.ID).Updates(
		model.User{
			FirstName:  user.FirstName,
			Lastname:   user.Lastname,
			Email:      user.Email,
			Password:   user.Password,
			Age:        user.Age,
			Field:      user.Field,
			Gender:     user.Gender,
			IsEmployee: user.IsEmployee,
		},
	)
	if res.Error != nil{
		return res.Error
	}
	return nil
}

func (u *UserRepo) Delete(id uint) error{
	res := u.DB.Where("ID = ?",id).Delete(&model.User{})
	if res.Error != nil{
		return res.Error
	}
	return nil
}
