package main

import (
	"final-project/database"
	"final-project/router"
)

func init() {

}

func main() {
    database.InitializeDb()
    router.InitializeRouter()
}
