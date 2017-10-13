package configuration

type Config struct {
	Count   uint
	Length  uint
	Charset string
}

var charsets = map[string]string{
	"numbers":      "0123456789",
	"alphabetic":   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"alphanumeric": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
}

type option func(*Config)

// getCharset, get character set (numbers, alphabetic or alphanumeric) //
// input string
// output string
func setCharset(name string) string {
	return charsets[name]
}

// Config, sets the option specified //
// input option
func (o *Config) Config(opts ...option) {
	for _, opt := range opts {
		opt(o)
	}
}

// charset, set charset value //
// input string
// output option
func charset(charset string) option {
	return func(o *Config) {
		o.Charset = charset
	}
}

// count, set count value //
// input uint
// output option
func count(count uint) option {
	return func(o *Config) {
		o.Count = count
	}
}

// length, set length value //
// input uint
// output option
func length(length uint) option {
	return func(o *Config) {
		o.Length = length
	}
}

// New, create new configuration //
// input option
// output *Config
// output error
func New(options ...option) (*Config, error) {
	c := Config{
		Charset: setCharset("alphanumeric"),
		Length:  8,
		Count:   1,
	}

	for _, opt := range options {
		opt(&c)
	}

	return &c, nil
}
