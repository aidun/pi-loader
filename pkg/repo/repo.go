package repo

import (
	"fmt"
	"log"
	"os"
)

const loaderDir = ".pi-loader"

func Initialize() {
	if _, err := os.Stat(getRepoDir()); os.IsNotExist(err) {
		err := os.Mkdir(getRepoDir(), 0755)
		if err != nil {
			log.Fatalf("could not create repository %s: %v", getRepoDir(), err)
		}
	}
}

func getRepoDir() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/%s", homeDir, loaderDir)
}
