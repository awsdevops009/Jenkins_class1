package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func generateHTML() {
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	v := "docker"
	s := `<html>
<script>
function startTime() {
    var today = new Date();
    var h = today.getHours();
    var m = today.getMinutes();
    var s = today.getSeconds();
    m = checkTime(m);
    s = checkTime(s);
    document.getElementById('clock').innerHTML =
    h + ":" + m + ":" + s;
    var t = setTimeout(startTime, 500);
}
function checkTime(i) {
    if (i < 10) {i = "0" + i};  
    return i;
}
</script>
<body onload="startTime()">
    <h2>Hello from ` + v + ` !</h2>
    <h1>The current server time is</h1>
    <div id="clock"></div>
</body>
</html>`

	fmt.Fprintf(file, s)
}

func main() {
	generateHTML()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
