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
	s := `<%@ page import="java.io.*,java.util.*,javax.servlet.*" %>
<html>
<body>
<h2>Hello ` + v + ` the Tomcat server is running!</h2>
<h1>The current server time is</h1>
<%
    Date date = new Date();
    out.print("<h4 align=\"center\">" + date.toString() + "</h4>");
%>
</body>
</html>`
	fmt.Fprintln(file, s)
}
