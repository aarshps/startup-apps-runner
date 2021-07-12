package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Starting Apps...")

	startApp("Teams", "C:\\Users\\aps\\AppData\\Local\\Microsoft\\Teams\\Update.exe", ([]string{"--processStart", "Teams.exe"})...)
	startApp("Outlook", "C:\\Program Files (x86)\\Microsoft Office\\root\\Office16\\OUTLOOK.EXE")
	startApp("Docker Desktop", "C:\\Program Files\\Docker\\Docker\\Docker Desktop.exe")
	startApp("ConEmu", "C:\\Program Files\\ConEmu\\ConEmu64.exe")

	fmt.Println("Apps Started...")
}

func startApp(appName string, exePath string, args ...string) {
	fmt.Println("Starting" + appName + "...")

	command := exec.Command(exePath, args...)
	command.Start()

	time.Sleep(5)
}
