package main

import (
	"fmt"
	"os/exec"
)

func listarprocesos() string {

	// ps -aux | awk '{ print $1, $2, $3, $4, $9, $11 }'

	listado_procesos, error := exec.Command("ps", "-aux").Output()
	// listado_procesos := string(comando)
	if error != nil {
		fmt.Printf("%s", error)
	}
	return string(listado_procesos)
}
func main() {
	fmt.Println("Listado de procesos")
	// Listamos todos los procesos que coincidan con los introducidos por parámetro
	fmt.Println(listarprocesos())
}
