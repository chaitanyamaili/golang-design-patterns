package main

import "fmt"

func main() {
	// Create a NotificationBuilder and use it to set properties
	bldr := newNotificationBuilder()

	// Use the builder to set some properties
	bldr.SetTitle("This is notification title")
	bldr.SetSubTitle("This is subtitle")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image,jpg")
	bldr.SetMessage("This is the notification message")
	bldr.SetPriority(5)
	bldr.SetType("alert")

	// Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error building notification object", err)
	} else {
		fmt.Println("Finished Notification: %+v", notif)
	}
}
