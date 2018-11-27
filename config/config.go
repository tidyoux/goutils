package config

type Config struct {
	*Pair
}

func New() *Config {
	return &Config{
		Pair: NewPair("root"),
	}
}

func (c *Config) Parse(raw string) error {
	data := []rune("(" + raw + ")")
	_, _, err := parsePair(data, c.Pair)
	if err != nil {
		return err
	}
	return nil
}
