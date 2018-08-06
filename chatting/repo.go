package chatting

type Repo interface {
	// Set adds or overwrites the value for the given key
	Set(key string, val interface{}) error

	// Has tests if the given key is in the repo
	Has(key string) bool

	// Get retrieves the generic value from the db
	Get(key string) interface{}
}
