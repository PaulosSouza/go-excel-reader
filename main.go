package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/PaulosSouza/go-excel-reader/repository"
	storage "github.com/PaulosSouza/go-excel-reader/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonDTO struct {
	Age       int
	Gender    string
	Born      string
	CreatedAt time.Time
}

func main() {

	filePath := filepath.Join("assets", "adult10m")
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mongoConnection := storage.MongoConnect()
	personRepository := repository.NewPersonRepository(
		mongoConnection.GetCollection(storage.PersonCollection),
	)

	csvReader := csv.NewReader(file)

	start := time.Now()

	var lineCounter int
	personArr := make([]interface{}, 0)
	person := make(primitive.D, 1)

	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		lineCounter++

		if err != nil {
			log.Fatal(err)
		}

		age, err := strconv.Atoi(line[0])

		if err != nil {
			panic(err)
		}

		person = bson.D{
			primitive.E{Key: "age", Value: age},
			primitive.E{Key: "gender", Value: line[9]},
			primitive.E{Key: "born", Value: line[13]},
			primitive.E{Key: "createdAt", Value: time.Now()},
		}

		personArr = append(personArr, person)

		if lineCounter == 1000 {
			channel := make(chan bool, 4)

			quarter, half, third, last :=
				personArr[0:250], personArr[250:500], personArr[500:750], personArr[750:1000]

			go personRepository.InsertMany(quarter, channel)
			go personRepository.InsertMany(half, channel)
			go personRepository.InsertMany(third, channel)
			go personRepository.InsertMany(last, channel)

			<-channel

			lineCounter = 0
			personArr = nil
			person = nil
		}
	}

	elpased := time.Since(start)

	fmt.Printf("The duration is: %s", elpased)
}
