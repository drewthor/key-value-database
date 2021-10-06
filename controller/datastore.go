package controller

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/drewthor/key-value-database/api"
	"github.com/drewthor/key-value-database/service"
	"github.com/drewthor/key-value-database/util"
	"github.com/go-chi/chi/v5"
)

type DatastoreController struct {
	DatastoreService *service.DatastoreService
}

func (dc DatastoreController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/get", dc.Get)
	r.Post("/set", dc.Set)
	r.Post("/delete", dc.Delete)

	return r
}

func (dc DatastoreController) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		util.WriteJSON(http.StatusBadRequest, api.Error{Error: "invalid request: missing required query parameter: key"}, w)
		return
	}

	value, exists := dc.DatastoreService.Get(key)
	if !exists {
		util.WriteJSON(http.StatusNotFound, api.Error{Error: fmt.Sprintf("could not find value for key: %s", key)}, w)
		return
	}

	response := api.KeyValue{Key: key, Value: value}
	util.WriteJSON(http.StatusOK, response, w)
}

func (dc DatastoreController) Set(w http.ResponseWriter, r *http.Request) {
	keyValue := api.KeyValue{}

	err := json.NewDecoder(r.Body).Decode(&keyValue)
	if err != nil {
		util.WriteJSON(http.StatusBadRequest, api.Error{Error: "invalid request body"}, w)
		return
	}

	dc.DatastoreService.Set(keyValue.Key, keyValue.Value)

	util.WriteJSON(http.StatusOK, keyValue, w)
}

func (dc DatastoreController) Delete(w http.ResponseWriter, r *http.Request) {
	request := api.DeleteKeyRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		util.WriteJSON(http.StatusBadRequest, api.Error{Error: "invalid request body"}, w)
		return
	}

	deleted := dc.DatastoreService.Delete(request.Key)
	if deleted {
		util.WriteJSON(http.StatusOK, request, w)
	} else {
		util.WriteJSON(http.StatusOK, map[string]string{"message": "key not found"}, w)
	}
}
