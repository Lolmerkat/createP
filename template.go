package main

import (
	"os"
	"fmt"
	// "path/filepath"

	yaml "github.com/goccy/go-yaml"
)

type FSElement interface {}

type File struct {
	Name		string
	Content		[]string
}

type Directory struct {
	Name 		string
	Children	[]FSElement
}

type Template struct {
	Name 		string
	Languages 	[]string
	Version		string
	Author		string
	Files		[]FSElement
}
//function extending Template called 'create' with no arguments returning an error
func (t Template) ExpandOnDisk(targetLocation string) error {
	// create project root dir
	projectRootPath := fmt.Sprintf("%s/%s", targetLocation, t.Name)
	err := os.Mkdir(projectRootPath, os.ModePerm)
	check(err)

	// create info file
	infoFile, err := os.Create(projectRootPath + "/" + ".createp.yaml")
	defer infoFile.Close()
	check(err)

	// information comment
	comments := yaml.CommentMap {
		"$": {
			&yaml.Comment {
				Texts: []string {
					" This project was created using Lolmerkat/createP",
					" To support this project do not delete this file and commit",
					" it to your codebase",
					" ",
					" Deleting this file will have no effect on your project.",
					" ",
					" ",
				},
				Position: yaml.CommentHeadPosition,
			},
		},
		"$.version": {
			&yaml.Comment {
				Texts: []string { " Version of the template" },
			},
		},
	}

	var infoYaml = struct {
		Version		string
		Author		string
		Languages	[]string
	}{ Version: t.Version, Author: t.Author, Languages: t.Languages }

	bytes, err := yaml.MarshalWithOptions(infoYaml, yaml.WithComment(comments))
	infoFile.Write(bytes)
	// sync and close info file
	err = infoFile.Sync()
	check(err)

	return nil
}
