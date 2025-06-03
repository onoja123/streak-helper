package main

import (
	"fmt"
	"time"
)

type StreakData struct {
	Friends []string
	Logs    map[string]string
}

var data = StreakData{
	Friends: []string{},
	Logs:    map[string]string{},
}

func main() {
	for {
		fmt.Println("\n📸 STREAK HELPER")
		fmt.Println("1. View streak friends")
		fmt.Println("2. Add a friend")
		fmt.Println("3. Log a snap sent")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option (1-4): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			viewFriends()
		case 2:
			addFriend()
		case 3:
			logSnap()
		case 4:
			fmt.Println("Goodbye 👋")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func viewFriends() {
	if len(data.Friends) == 0 {
		fmt.Println("No friends added yet.")
		return
	}
	for _, name := range data.Friends {
		lastSent := data.Logs[name]
		if lastSent == "" {
			lastSent = "Never"
		}
		fmt.Printf("- %s (last sent: %s)\n", name, lastSent)
	}
}

func addFriend() {
	fmt.Print("Enter friend's name: ")
	var name string
	fmt.Scanln(&name)
	data.Friends = append(data.Friends, name)
	fmt.Println("Added:", name)
}

func logSnap() {
	if len(data.Friends) == 0 {
		fmt.Println("No friends to log.")
		return
	}
	fmt.Println("Select a friend to log snap:")
	for i, name := range data.Friends {
		fmt.Printf("%d. %s\n", i+1, name)
	}
	var idx int
	fmt.Print("Enter number: ")
	fmt.Scanln(&idx)

	if idx < 1 || idx > len(data.Friends) {
		fmt.Println("Invalid selection.")
		return
	}

	friend := data.Friends[idx-1]
	data.Logs[friend] = time.Now().Format("Jan 2 15:04")
	fmt.Printf("✅ Logged snap sent to %s at %s\n", friend, data.Logs[friend])
}
