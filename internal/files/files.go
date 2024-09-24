package files

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"encoding/json"

	"github.com/AsierAlaminos/NoteShell/internal/model"
)

func CreateConfDir() {

	confDirPath := fmt.Sprintf("%s/.noteshell", CheckUser())

	if _, err := os.Stat(confDirPath); os.IsNotExist(err) {
		os.Mkdir(confDirPath, 0755)
	}

	ideasPath := fmt.Sprintf("%s/ideas", confDirPath)

	if _, err := os.Stat(ideasPath); os.IsNotExist(err) {
		os.Mkdir(ideasPath, 0755)
	}
}

func CreateIdeaFiles(idea string, categories []string) {
	homeUser := CheckUser()

	ideaDirPath := fmt.Sprintf("%s/.noteshell/ideas/%s", homeUser, strings.ToLower(idea))

	if _, err := os.Stat(ideaDirPath); os.IsNotExist(err) {
		os.Mkdir(ideaDirPath, 0755)
		ideaPath := fmt.Sprintf("%s/.noteshell/ideas/%s/%s.json", homeUser, strings.ToLower(idea), strings.ToLower(idea))
		mdPath := fmt.Sprintf("%s/.noteshell/ideas/%s/%s.md", homeUser, strings.ToLower(idea), strings.ToLower(idea))
		createFile(ideaPath)
		createFile(mdPath)
		writeIdeaJson(ideaPath, idea, categories)
	}
}

func createFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, errFile := os.Create(filePath)
		if errFile == nil {
			return
		}
		fmt.Printf("[!] Error creating %s\n", filePath)
	}
}

func writeIdeaJson(ideaPath string, idea string, categories []string) {
	ideaModel := model.Idea{
		Name: idea,
		DescFile: idea + ".md",
		Categories: categories,
	}
	squeleton := fmt.Sprintf("{\"name\": \"%s\", \"descfile\": \"%s.md\", \"categories\": [%s]}", ideaModel.Name, ideaModel.Name, ideaModel.ParseCategoriesJson())
	err := os.WriteFile(ideaPath, []byte(squeleton), 0777)
	if err != nil {
		fmt.Println("[!] Error writing the file")
	} else {
		fmt.Println("[*] File written succesfully")
	}
}

func ReadDirs(path string) []string {
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}
	var dirs []string
	for _, e := range entries {
		dirs = append(dirs, e.Name())
	}

	return dirs
}

func ReadJsonIdea(path string) model.Idea {
	byteValue, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("[!] Exiting... (invalid json file)")
		os.Exit(1)
	}
	var idea model.Idea

	if err := json.Unmarshal(byteValue, &idea); err != nil {
		fmt.Println("[!] error json file")
	}

	return idea
}

func Banner() string {
	homeDir := CheckUser()
	byteValue, err := os.ReadFile(fmt.Sprintf("%s/.noteshell/banner.txt", homeDir))
	if err != nil {
		fmt.Println("[!] Exiting... (invalid file)")
		os.Exit(1)
	}
	return string(byteValue)
}

func CheckUser() string {
	currentUser, err := user.Current()

	if err != nil {
		log.Fatal("[!] User %s doesn't exist\n", currentUser.Username)
	}

	return currentUser.HomeDir
}
