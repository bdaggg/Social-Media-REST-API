package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
		os.Exit(1)
	}
	return os.Getenv(key)
}

/*
	func EnvCloudName() string {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_CLOUD_NAME")
	}

	func EnvCloudAPIKey() string {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_API_KEY")
	}

	func EnvCloudAPISecret() string {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_API_SECRET")
	}

func EnvCloudUploadFolder() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
*/
