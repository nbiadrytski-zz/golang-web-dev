package main

import "fmt"

func main() {
	name := "Mikalai Biadrytski"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)

	// go run main.go > index.html will dump the output to .html file
}
