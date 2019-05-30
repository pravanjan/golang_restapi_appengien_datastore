package controllers

import (
	"encoding/json"
	"fmt"
	"golang_restapi_appengien_datastore/data_access/model"
	"golang_restapi_appengien_datastore/data_access/modelmanager"

	"net/http"
	"time"
)

var CreateNewEntry = func(writer http.ResponseWriter, request *http.Request) {
	layout := "2006-01-02T15:04:05.000Z"

	decoder := json.NewDecoder(request.Body)
	//creating a map data to store the json temp json data
	entryMap := make(map[string]string)

	err := decoder.Decode(&entryMap)
	if err != nil {
		panic(err)
	}
	//converting string date to data time

	t, err := time.Parse(layout, string(entryMap["dateAdded"]))
	datastoreEntry := &model.Entry{
		Name:      entryMap["name"],
		Role:      entryMap["role"],
		DateAdded: t,
	}
	result := modelmanager.SaveEntry(*datastoreEntry, request)
	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	out, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(writer, string(out))

}

var DefaultView = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("default view ")
	fmt.Fprint(w, "Hello there!! ")

	//var e2 model.Entry
	// if err = datastore.Get(ctx, key, &e2); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

}
