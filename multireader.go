package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	head := `<html>
 <head>
  <title>My Awesome Web Page</title>
 </head>
 `

	body, _ := os.Open("body.html")

	foot := `</html>`

	// START OMIT
	io.Copy(os.Stdout,
		io.MultiReader(
			strings.NewReader(head),  // string data
			body,                     // *os.File returned by opening body.html
			strings.NewReader(foot))) // string stat
	// END OMIT
}
