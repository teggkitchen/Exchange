package main

import (
	_ "exchange/database"
	orm "exchange/database"
	"exchange/router"
)

func main() {
	defer orm.GormOpen.Close()
	router := router.InitRouter()
	router.Run(":8000")
}
