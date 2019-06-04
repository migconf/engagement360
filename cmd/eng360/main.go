package main

import (
	"enagement360/internal/server"
	"fmt"
)


func main(){
	fmt.Println("Starting EngagementHelper server...")

	// setup router
	cntlr := server.PssRouter{}
	cntlr.Init()

}