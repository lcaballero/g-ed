package hub

var id = 0

// New creates a new unique id.
func NewId() int {
	id++
	return id
}
