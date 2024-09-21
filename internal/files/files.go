package files

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/AsierAlaminos/NoteShell/internal/utils"
)

func CreateConfDir() {
	currentUser, err := user.Current()

	if err != nil {
		fmt.Printf("[!] User %s doesn't exist\n", currentUser.Username)
	}

	confDirPath := fmt.Sprintf("%s/.noteshell", currentUser.HomeDir)

	if _, err := os.Stat(confDirPath); os.IsNotExist(err) {
		os.Mkdir(confDirPath, 0755)
		fmt.Printf("[*] %s created\n", confDirPath)
	}

	ideasPath := fmt.Sprintf("%s/ideas", confDirPath)

	if _, err := os.Stat(ideasPath); os.IsNotExist(err) {
		os.Mkdir(ideasPath, 0755)
		fmt.Printf("[*] %s created\n", ideasPath)
	}
}

func CreateIdeaFiles(idea string) {
	currentUser, err := user.Current()

	if err != nil {
		fmt.Printf("[!] User %s doesn't exist", currentUser.Username)
	}

	ideaDirPath := fmt.Sprintf("%s/.noteshell/ideas/%s", currentUser.HomeDir, strings.ToLower(idea))

	if _, err := os.Stat(ideaDirPath); os.IsNotExist(err) {
		os.Mkdir(ideaDirPath, 0755)
		fmt.Printf("[*] %s created\n", ideaDirPath)
		ideaPath := fmt.Sprintf("%s/.noteshell/ideas/%s/%s.json", currentUser.HomeDir, strings.ToLower(idea), strings.ToLower(idea))
		mdPath := fmt.Sprintf("%s/.noteshell/ideas/%s/%s.md", currentUser.HomeDir, strings.ToLower(idea), strings.ToLower(idea))
		createFile(ideaPath)
		createFile(mdPath)
		writeIdeaJson(ideaPath, idea, []string{"uno", "dos", "tres"})
	}
}

func createFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, errFile := os.Create(filePath)
		if errFile == nil {
			fmt.Printf("[*] %s created\n", filePath)
			return
		}
		fmt.Printf("[!] Error creating %s\n", filePath)
	}
}

func writeIdeaJson(ideaPath string, idea string, categories []string) {
	squeleton := fmt.Sprintf("{\"name\": %s, \"descfile\": %s.md, \"categories\": [%s]}", idea, idea, main.ParseCategories(categories))
	err := os.WriteFile(ideaPath, []byte(squeleton), 0777)
	if err != nil {
		fmt.Println("[!] Error writing the file")
	} else {
		fmt.Println("[*] File written succesfully")
	}
}

