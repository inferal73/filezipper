package filezipper

var Version string

// Config ...
type Config struct {
	Entry string
	Out   string
}

// NewConfig ...
func NewConfig(entry, out string) *Config {
	return &Config{
		Entry: entry,
		Out:   out,}
}
