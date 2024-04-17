package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"main/models"
	"net/http"
	"os"
	"strconv"
	"time"
)

func readIsis(filePath string) ([]models.Iris, error) {

	var irisCollection []models.Iris

	irisFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer irisFile.Close()

	csvReader := csv.NewReader(irisFile)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for i, line := range data {
		if i > 0 {
			var iris models.Iris
			sepalLength, err := strconv.ParseFloat(line[0], 32)
			if err == nil {
				iris.SepalLength = float32(sepalLength)
			}

			eepalWidth, err := strconv.ParseFloat(line[1], 32)
			if err == nil {
				iris.SepalWidth = float32(eepalWidth)
			}

			petalLength, err := strconv.ParseFloat(line[2], 32)
			if err == nil {
				iris.PetalLength = float32(petalLength)
			}

			petalWidth, err := strconv.ParseFloat(line[3], 32)
			if err == nil {
				iris.PetalWidth = float32(petalWidth)
			}

			iris.Variety = line[4]
			irisCollection = append(irisCollection, iris)
		}
	}
	return irisCollection, nil
}

func main() {

	//http.Handle("/iris/add", handlers.AddIrisHandler(""))

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		irisCollection, err := readIsis("./data/iris.csv")
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(irisCollection)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		w.Write(data)
	})

	http.Handle("/", myHandler("Contoso"))

	s := http.Server{
		Addr: ":3000",
	}

	go func() {
		log.Fatal(s.ListenAndServeTLS("./cert.pem", "./key.pem"))
	}()

	fmt.Printf("Server started, press <Enter> to shutdown.")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

type myHandler string

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Tina", "Lotus")
	http.SetCookie(w, &http.Cookie{
		Name:    "biscoiteira",
		Value:   "0923984",
		Expires: time.Now().Add(24 * time.Hour),
	})
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(mh))
}
