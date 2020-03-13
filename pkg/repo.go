package pkg

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const loaderDir = ".pi-loader"

const imageUrl = "https://downloads.raspberrypi.org/raspbian_lite/images/raspbian_lite-2020-02-14/2020-02-13-raspbian-buster-lite.zip"
const imageName = "raspbian_lite-latest"

func InitializeRepo() {

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

func downloadImage() error {

	zipFilename := fmt.Sprintf("%s/%s.zip",
		getRepoDir(),
		imageName,
	)

	imageFilename := fmt.Sprintf("%s/%s.img",
		getRepoDir(),
		imageName,
	)

	// Download
	if _, err := os.Stat(zipFilename); os.IsNotExist(err) {

		resp, err := http.Get(imageUrl)
		if err != nil {
			return fmt.Errorf("Error downloading the image: %s", err)
		}

		out, err := os.Create(zipFilename)
		if err != nil {
			return fmt.Errorf("Error downloading the image: %s", err)
		}

		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return fmt.Errorf("Error downloading the image: %s", err)
		}
	}

	// Extract
	if _, err := os.Stat(imageFilename); os.IsNotExist(err) {
		zipReader, err := zip.OpenReader(zipFilename)
		if err != nil {
			return fmt.Errorf("Error extracting the image: %s", err)
		}

		defer zipReader.Close()

		out, err := os.Create(imageFilename)
		if err != nil {
			return fmt.Errorf("Error extracting the image: %s", err)
		}

		// There is only one image
		imageFileInZip := zipReader.File[0]

		imageFileReader, err := imageFileInZip.Open()
		if err != nil {
			return fmt.Errorf("Error extracting the image: %s", err)
		}

		_, err = io.Copy(out, imageFileReader)
	}
	return nil
}
