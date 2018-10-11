package main

import (
	_ "exchange/database"
	orm "exchange/database"
	"exchange/router"
)

func main() {
	// cmd := exec.Command("./script.sh")
	// cmd.Stdout = os.Stdout
	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)

	defer orm.GormOpen.Close()
	router := router.InitRouter()
	router.Run(":8000")
}
