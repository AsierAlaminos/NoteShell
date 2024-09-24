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
	currentUser, err := user.Current()

	if err != nil {
		fmt.Printf("[!] User %s doesn't exist\n", currentUser.Username)
	}

	confDirPath := fmt.Sprintf("%s/.noteshell", currentUser.HomeDir)

	if _, err := os.Stat(confDirPath); os.IsNotExist(err) {
		os.Mkdir(confDirPath, 0755)
	}

	ideasPath := fmt.Sprintf("%s/ideas", confDirPath)

	if _, err := os.Stat(ideasPath); os.IsNotExist(err) {
		os.Mkdir(ideasPath, 0755)
	}
}

func CreateIdeaFiles(idea string, categories []string) {
	currentUser, err := user.Current()

	if err != nil {
		fmt.Printf("[!] User %s doesn't exist", currentUser.Username)
	}

	ideaDirPath := fmt.Sprintf("%s/.noteshell/ideas/%s", currentUser.HomeDir, strings.ToLower(idea))

	if _, err := os.Stat(ideaDirPath); os.IsNotExist(err) {
		os.Mkdir(ideaDirPath, 0755)
		ideaPath := fmt.Sprintf("%s/.noteshell/ideas/%s/%s.json", currentUser.HomeDir, strings.ToLower(idea), strings.ToLower(idea))
		mdPath := fmt.Sprintf("%s/.noteshell/ideas/%s/%s.md", currentUser.HomeDir, strings.ToLower(idea), strings.ToLower(idea))
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
		/*jsonPath := fmt.Sprintf("%s/%s/%s.json", path, e.Name(), e.Name())
		idea := readJsonIdea(jsonPath)
		fmt.Println("\n" + jsonPath)
		fmt.Printf("name: %s\ndescfile: %s\ncategories: %s\n", idea.Name, idea.DescFile, utils.ParseCategories(idea.Categories))*/
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
	byteValue, err := os.ReadFile("/home/asmus/.noteshell/banner.txt")
	if err != nil {
		fmt.Println("[!] Exiting... (invalid file)")
		os.Exit(1)
	}
	return string(byteValue)
}
