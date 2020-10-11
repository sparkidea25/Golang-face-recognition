package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Kagami/go-face"
)

const dataDir = ""

func main() {
	fmt.Println("Welcome to face recognition")

	rec, err := face.NewRecognizer(dataDir)

	if err != nil {
		fmt.Println("cannot initialize face recognition")
	}

	defer rec.Close()
	fmt.Println("Recognizer initialized")

	avengersImage := filepath.Join(dataDir, "tony-stark.jpg")
	faces, err := rec.RecognizeFile(avengersImage)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	fmt.Println("Number of Faces in Images: ", len(faces))
}
