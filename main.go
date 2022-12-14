package main

import (
	"final-project/database"
	"final-project/router"

	"github.com/asaskevich/govalidator"
	"github.com/k0kubun/pp"
)

func init() {
	pp.ColoringEnabled = false

	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	db := database.InitializeDb()
	r := router.InitializeRouter(db)
	r.Run(":8080")
}
