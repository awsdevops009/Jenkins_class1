package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Create("index.jsp")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	v := os.Getenv("name")
	if v == "" {
		v = "docker"
	}

	// Get the current date
	currentDate := time.Now().Local()

	// Get the day of the week (e.g., Monday, Tuesday, etc.)
	dayOfWeek := currentDate.Weekday().String()

	s := fmt.Sprintf(`<html>
<head>
<title>Welcome to Our Website</title>
</head>
<body>
<h2>Hello %s the Tomcat server is running!</h2>
<h1>The current server time is %s</h1>
<h1>Today's day is %s</h1>
</body>
</html>`, v, currentDate.Format("2006-01-02 15:04:05"), dayOfWeek)

	fmt.Fprintln(file, s)
}
