package repository

import (
	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/domain"
	interfaces "github.com/abhinandkakkadi/moviesgo-admin-service/pkg/repository/interface"
	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/utils/models"
	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDatabase{
		DB: DB,
	}
}

func (a *AdminDatabase) CreateAdmin(admin models.AdminSignUp) (int, error) {

	var id int
	if err := a.DB.Raw("insert into admins (name,email,password) values (?, ?, ?) RETURNING id", admin.Name, admin.Email, admin.Password).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil

}

func (a *AdminDatabase) CheckAdminAvailability(admin models.AdminSignUp) bool {

	var count int
	if err := a.DB.Raw("select count(*) from admins where email = ?", admin.Email).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (a *AdminDatabase) LoginHandler(adminDetails models.AdminLogin) (domain.Admin,error) {
	
	var adminCompareDetails domain.Admin
	if err := a.DB.Raw("select * from admins where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil

}