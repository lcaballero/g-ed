package params

type ValContext interface {
	String(string) string
	Int(string) int
	Float64(string) float64
	IsSet(string) bool
	Bool(string) bool
}

type ServeParams struct {
	Root string
}

// Load uses ValContext to create ServeParams from which the server
// is configured.
func Load(val ValContext) ServeParams {
	return ServeParams{
		Root: val.String("root"),
	}
}
