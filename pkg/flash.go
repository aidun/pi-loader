package pkg

func Flash() error {
	InitializeRepo()

	return downloadImage()
}
