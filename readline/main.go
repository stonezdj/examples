package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	DeploymentName = "deployment123"
)

func main() {

	file, err := os.Open("/Users/daojunz/Downloads/k8s_dep/private-reg-pod.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		indent := GetIndent(line)

		if strings.HasPrefix(strings.TrimSpace(line), "image:") {
			line := strings.TrimSpace(line)
			image := strings.TrimPrefix(line, "image:")
			image = strings.TrimSpace(image)
			newName := GetNewName(image)
			AddReplicationJob(image, newName)
			fmt.Printf("%simage: %s\n", indent, newName)
		} else {
			fmt.Println(line)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetIndent(line string) string {
	i := 0
	var chars []byte
	for _, c := range line {
		if c == ' ' {
			i++
			chars = append(chars, ' ')
		} else {
			break
		}
	}
	return string(chars)
}

func GetNewName(imageName string) string {
	registryHostName := "hub.docker.com"
	projectName := "library"
	var imageShortName string
	results := strings.Split(imageName, "/")
	if len(results) == 1 {
		imageShortName = results[0]
	} else if (len(results)) == 2 {
		projectName = results[0]
		imageShortName = results[1]
	} else if len(results) == 3 {
		registryHostName = results[0]
		projectName = results[1]
		imageShortName = results[2]
	}
	log.Printf("Image short name:%v", imageShortName)
	return fmt.Sprintf("%s/%s_%s_%s", DeploymentName, registryHostName, projectName, imageShortName)
}

func AddReplicationJob(source, target string) error {
	log.Printf("Replication image from %s to %s \n", source, target)
	return nil
}
