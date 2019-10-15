package main

import (
	"github.com/MohitArora1/student/controller"
	"github.com/MohitArora1/student/utils"
)

func main() {
	utils.InitConfig()
	controller.RunController(":8080")
}
