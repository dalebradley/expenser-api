package conf

import (
	"log"

	"github.com/ian-kent/gofigure"
)

type Config struct {
	ExpenserPort string `env:"EXPENSER_PORT"                    flag:"expenser-port"                    flagDesc:"Primary Expnser Server Port"`
	Collection   string `env:"MONGODB_COLLECTION"           flag:"mongodb-collection"           flagDesc:"MongoDB collection for data"`
	Database     string `env:"MONGODB_DATABASE"             flag:"mongodb-database"             flagDesc:"MongoDB database for data"`
	MongoDBURL   string `env:"MONGODB_URL"                  flag:"mongodb-url"                  flagDesc:"MongoDB server URL"`
}

// DefaultConfig returns a pointer to a Config instance that has been populated
// with default values.
func DefaultConfig() Config {
	return Config{
		ExpenserPort: ":8888",
		Database:     "transactions",
		Collection:   "registered_office_address",
		MongoDBURL:   "mongodb://chs-mongo:27017",
	}
}

// Get returns a pointer to a config instance that has been populated with
// values provided by the environment or command-line flags, or with default
// values if none are provided.
func Get() *Config {
	var cfg Config
	cfg = DefaultConfig()
	err := gofigure.Gofigure(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
