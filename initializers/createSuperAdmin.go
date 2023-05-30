package initializers

import (
	"os"

	"github.com/Cedar-81/swype/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateSuperAdmin()  {
	var user models.User
	var err error

	err = DB.Where("email = ?", os.Getenv("SUPER_EMAIL")).
		First(&user).Error
	if err == nil {
		return
	}

	//hash password
	pass, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SUPER_PASS")), 10)
	if err != nil {
		panic("Unable to hash password")
	}

	// create user
	user = models.User{
		Username: "superadmins",
		Email:    os.Getenv("SUPER_EMAIL"),
		Password: string(pass),
		Role:     "SUPER",
	}

	result := DB.Create(&user)

	if result.Error != nil {
		panic(result.Error)
	}

}