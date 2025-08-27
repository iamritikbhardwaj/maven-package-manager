package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Project struct {
	XMLName      xml.Name      `xml:"project"`
	Dependencies *Dependencies `xml:"dependencies"`
}

type Dependencies struct {
	DependencyList []Dependency `xml:"dependency"`
}

type Dependency struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"artifactId"`
	Version    string `xml:"version"`
}

const pomFilePath = "pom.xml"

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <groupId> <artifactId> <version>")
		return
	}

	groupId := os.Args[1]
	artifactId := os.Args[2]
	version := os.Args[3]

	file, err := os.Open("pom.xml")
	if err != nil {
		fmt.Println("Error opening pom.xml:", err)
		return
	}
	defer file.Close()

	data, err := os.ReadFile(pomFilePath)
	if err != nil {
		fmt.Println("Error reading pom.xml:", err)
		return
	}

	var project Project
	if err := xml.Unmarshal(data, &project); err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	if project.Dependencies == nil {
		project.Dependencies = &Dependencies{}
	}

	found := false
	for i, dep := range project.Dependencies.DependencyList {
		if dep.GroupID == groupId && dep.ArtifactID == artifactId {
			project.Dependencies.DependencyList[i].Version = version
			found = true
			fmt.Println("Dependency exists. Updating version.")
			break
		}
	}

	if !found {
		newDep := Dependency{
			GroupID:    groupId,
			ArtifactID: artifactId,
			Version:    version,
		}
		project.Dependencies.DependencyList = append(project.Dependencies.DependencyList, newDep)
		fmt.Println("Dependency added.")
	}

	output, err := xml.MarshalIndent(project, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}

	// Add XML header
	finalOutput := []byte(xml.Header + string(output))

	if err := os.WriteFile(pomFilePath, finalOutput, 0644); err != nil {
		fmt.Println("Error writing pom.xml:", err)
		return
	}

	fmt.Println("pom.xml updated.")
}
