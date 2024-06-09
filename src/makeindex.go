package main

import (
	"fmt"
	"log"
	"os"
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

	s := `<html>
<head>
<title>Welcome to Our Website</title>
</head>
<body>
<h2>Hello ` + v + ` the Tomcat server is running!</h2>
<h1>The current server time is ${java.time.LocalDateTime.now()}</h1>
</body>
</html>`

	fmt.Fprintln(file, s)
}
