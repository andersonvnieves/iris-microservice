package handlers

// import (
// 	"encoding/json"
// 	"log"
// 	"main/models"
// 	"main/repositories"
// 	"net/http"
// )

// type AddIrisHandler *repositories.IrisRepository

// func (addIris AddIrisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	var iris models.Iris
// 	dec := json.NewDecoder(r.Body)
// 	err := dec.Decode(&iris)
// 	if err != nil {
// 		log.Fatal(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	w.Write(nil)
// }
