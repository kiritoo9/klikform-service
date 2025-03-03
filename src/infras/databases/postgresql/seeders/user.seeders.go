package seeders

import (
	"klikform/src/applications/models"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) {
	// check data already seeded
	var count int64
	db.Model(&models.Users{}).Count(&count)
	if count > 0 {
		return // stop process so seeder won't run
	}

	// get data role where name is `admin`
	var role models.Roles
	resultRole := db.Model(&models.Roles{}).Where(&models.Roles{Deleted: false, Name: "admin"}).First(&role)
	if resultRole.Error != nil {
		log.Println("Role admin is not found, user not created")
		return
	}

	// perform to insert data `user`
	adminID := uuid.New()
	password := "admin123" // default value for admin
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while encrypting password for admin user")
		return
	}

	user := models.Users{
		ID:       adminID,
		Email:    "admin@admin.com",
		Password: string(hashedBytes),
		Fullname: "Administrator",
		Phone:    "08721312731",
		Address:  "The Avenue Street, No. 312, Side river, Jakarta",
		Status:   "S2", // set as active
	}
	db.Create(&user)

	// perform to insert data `user_role`
	userRole := models.UserRoles{
		ID:     uuid.New(),
		UserID: adminID,
		RoleID: role.ID,
	}
	db.Create(&userRole)
}
