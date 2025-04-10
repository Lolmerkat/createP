package main

import (
	"fmt"
	"os"
	"log"
	"github.com/goccy/go-yaml"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	template := Template {
		Name: "HTML",
		Languages: []string{ "html", "javascript", "css" },
		Version: "0.1-alpha",
		Author: "Lolmerkat",
		Files: []FSElement {
			File {
				Name: "index.html",
				Content: []string{
					"<!DOCTYPE html>",
					"<html>",
					"<head>",
					"\t<title>Document</title>",
					"</head>",
					"<body>",
					"\t",
					"</body>",
					"</html>",
				},
			},
			Directory {
				Name: "styles",
				Children: []FSElement {
					Directory {
						Name: "css",
						Children: []FSElement {
							File {
								Name: "colors.css",
								Content: []string {},
							},
							File {
								Name: "main.css",
								Content: []string {
									"body {",
									"\tmargin: none;",
									"}",
								},
							},
						},
					},
				},
			},
			Directory {
				Name: "scripts",
				Children: []FSElement {
					File {
						Name: "main.js",
						Content: []string {},
					},
				},
			},
		},
	}

	var fileName string = template.Name + ".yml"
	bytes, err := yaml.Marshal(template)
	file, err := os.Create(fileName)
	check(err)
	defer file.Close()
	n, err := file.Write(bytes)
	check(err)
	err = file.Sync()
	check(err)

	fmt.Printf("Wrote %d bytes to %s", n, fileName)

	template.ExpandOnDisk("./")
	}
