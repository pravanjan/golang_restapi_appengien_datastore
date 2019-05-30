package modelmanager

import (
	"fmt"
	"golang_restapi_appengien_datastore/data_access/model"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func SaveEntry(data model.Entry, request *http.Request) model.Entry {
	ctx := appengine.NewContext(request)
	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "entry", nil), &data)

	var entry model.Entry
	if err = datastore.Get(ctx, key, &entry); err != nil {
		fmt.Println("Error while processing ")
	}
	return entry

}
