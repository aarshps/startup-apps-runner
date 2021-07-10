package main

import "fmt"
import "os/exec"

func main() {
	fmt.Println("Started...")

	// goPath, err := exec.LookPath("docker-desktop")

	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// } else {
	// 	fmt.Println("Path : ", goPath)
	// }

	args := []string { "--processStart", "Teams.exe" }
	command := exec.Command("C:\\Users\\aps\\AppData\\Local\\Microsoft\\Teams\\Update.exe", args...)
	command.Start();

	fmt.Println("Teams Started...")
}