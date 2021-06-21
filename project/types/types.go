package types

type T struct {
	Back      string `yaml:"back"`
	Port      int    `yaml:"port"`
	CacheSize int    `yaml:"cache"`
}
type Page struct {
	Page string
	Err  string
}
