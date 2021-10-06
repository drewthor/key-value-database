package repository

type DatastoreDAO struct {
	DB map[string]string
}

func (d *DatastoreDAO) Get(key string) (string, bool) {
	value, ok := d.DB[key]
	return value, ok
}

func (d *DatastoreDAO) Set(key, value string) {
	d.DB[key] = value
}

// deletes key in DB, otherwise no-op
func (d *DatastoreDAO) Delete(key string) {
	delete(d.DB, key)
}