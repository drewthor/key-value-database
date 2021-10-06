package service

import "github.com/drewthor/key-value-database/repository"

type DatastoreService struct {
	DatastoreDAO *repository.DatastoreDAO
}

/*
Return:
string - stored value in db, otherwise "" the zero value
bool - true if key exists in db, else false
 */
func (ds *DatastoreService) Get(key string) (string, bool) {
	return ds.DatastoreDAO.Get(key)
}

func (ds *DatastoreService) Set(key, value string) {
	ds.DatastoreDAO.Set(key, value)
}

/*
Return:
bool - true if key value was deleted from db, else false
 */
func (ds *DatastoreService) Delete(key string) bool {
	_, exists := ds.DatastoreDAO.Get(key)
	ds.DatastoreDAO.Delete(key)

	return exists
}