package files

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"encoding/json"

	"github.com/AsierAlaminos/NoteShell/internal/model"
	"github.com/charmbracelet/bubbles/list"
)

func CreateConfDir() {

	confDirPath := fmt.Sprintf("%s/.noteshell", CheckUser())

	if _, err := os.Stat(confDirPath); os.IsNotExist(err) {
		os.Mkdir(confDirPath, 0755)
	}

	docsPath := fmt.Sprintf("%s/docs", confDirPath)

	if _, err := os.Stat(docsPath); os.IsNotExist(err) {
		os.Mkdir(docsPath, 0755)
	}
	createFile(fmt.Sprintf("%s/ideas.json", confDirPath), true, "[]")
}

func CreateIdea(idea string, categories []string) {
	homeUser := CheckUser()
	ideasPath := fmt.Sprintf("%s/.noteshell/ideas.json", homeUser)
	fmt.Printf("ideasPath: %s\n", ideasPath)

	if exist, id := checkIdea(idea, ideasPath); !exist {
		mdPath := fmt.Sprintf("%s/.noteshell/docs/%s.md", homeUser, strings.ToLower(idea))
		addIdea(id, idea, categories, mdPath, ideasPath)
		createFile(mdPath, true, fmt.Sprintf("# %s", idea))
	}
}

func checkIdea(name string, path string) (bool, int) {
	byteValue, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("[!] Exiting... (invalid json file)")
		os.Exit(1)
	}
	var ideas []model.Idea

	if err := json.Unmarshal(byteValue, &ideas); err != nil {
		fmt.Println("[!] error json file")
	}
	for _, idea := range ideas {
		if idea.Name == name {
			return true, -1
		}
	}
	return false, len(ideas)
}

func addIdea(id int, name string, categories []string, descFilePath string , path string) {
	newIdea := model.Idea {
		Id: id,
		Name: name,
		DescFile: descFilePath,
		Categories: categories,
	}

	byteValue, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("[!] Exiting... (invalid json file)")
		os.Exit(1)
	}
	var ideas []model.Idea

	if err := json.Unmarshal(byteValue, &ideas); err != nil {
		fmt.Println("[!] error json file")
	}

	ideas = append(ideas, newIdea)
	jsonIdeas, err := json.Marshal(ideas)
	if err != nil {
		fmt.Println("[!] Error parsing ideas to json")
	}
	if err := os.WriteFile(path, jsonIdeas, 0755); err != nil {
		fmt.Println("[!] Error adding idea")
	}
}

func createFile(filePath string, write bool, msg string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, errFile := os.Create(filePath)
		if errFile != nil {
			fmt.Printf("[!] Error creating %s\n", filePath)
		}
		if write {
			file, err := os.OpenFile(filePath, os.O_WRONLY, 0777)
			if err != nil {
				return
			}
			file.Write([]byte(msg))
		}
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

func DeleteIdea(id int) []list.Item {
	filePath := fmt.Sprintf("%s/.noteshell/ideas.json", CheckUser())
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("[!] Exiting... (invalid json file)")
		os.Exit(1)
	}
	var ideas []model.Idea
	var newIdeas []model.Idea

	if err := json.Unmarshal(byteValue, &ideas); err != nil {
		fmt.Println("[!] error json file")
	}

	for i, idea := range ideas {
		if i != id {
			newIdeas = append(newIdeas, idea)
		}
	}

	var items []list.Item
	for _, i := range newIdeas {
		items = append(items, i)
	}
	
	jsonIdeas, err := json.Marshal(newIdeas)
	if err != nil {
		fmt.Println("[!] Error parsing ideas to json")
	}
	if err := os.WriteFile(filePath, jsonIdeas, 0775); err != nil {
		fmt.Println("[!] Error removing idea")
	}

	return items
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

func WriteDescription(idea, text string) {
	homeDir := CheckUser()
	filepath := fmt.Sprintf("%s/.noteshell/docs/%s.md", homeDir, idea)
	if err := os.WriteFile(filepath, []byte(text), 0755); err != nil {
		fmt.Println("[!] Error writing description")
	}
}

func ReadDescription(idea string) string {
	homedir := CheckUser()
	filepath := fmt.Sprintf("%s/.noteshell/docs/%s.md", homedir, idea)

	byteValue, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("[!] Exiting... (error reading description)\n%s", err)
	}
	return string(byteValue)
}
