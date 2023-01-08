package user

import "gorm.io/gorm"

// interface atau kontrak R besar artinya publik
type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// membuat instance
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// membuat function Save untuk repository
func (r *repository) Save(user User) (User, error) {
	// &user->pointer dari fungsi Save user
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
