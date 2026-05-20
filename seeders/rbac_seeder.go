package seeders

import (
	"log"

	"github.com/nazala-mafa/go-crud/middlewares"
	"gorm.io/gorm"
)

func RbacSeeder(db *gorm.DB) {
	enforcer := middlewares.NewRBAC(db).GetEnforcer()

	_, err := enforcer.AddPermissionsForUser(
		"admin",
		[]string{"products", "create"},
		[]string{"products", "update"},
		[]string{"products", "delete"},
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = enforcer.AddRoleForUser(
		"1",
		"admin",
	)

	if err != nil {
		log.Fatal(err)
	}
}
