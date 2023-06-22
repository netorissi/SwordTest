package config

type ConfigGlobal struct {
	Server        Server   `mapstructure:"server"`
	Swagger       Swagger  `mapstructure:"swagger"`
	MySQL         Database `mapstructure:"mysql"`
	MessageBroker Broker   `mapstructure:"broker"`
}

type Server struct {
	BasePath string `mapstructure:"base_path"`
	Port     string `mapstructure:"port"`
}

type Swagger struct {
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	Version     string `mapstructure:"version"`
	Host        string `mapstructure:"host"`
}

type Database struct {
	WriterURL string `mapstructure:"writer_url"`
	ReaderURL string `mapstructure:"reader_url"`
}

type Broker struct {
	DefaultURL string `mapstructure:"url"`
}
