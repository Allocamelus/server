package server

type Config struct {
	// Address to listen on
	// Defaults to :8080
	Address string `json:"port"`
}

func (c *Config) Default() {
	if c.Address == "" {
		c.Address = ":8080"
	}
}
