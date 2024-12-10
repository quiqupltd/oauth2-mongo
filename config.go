package mongo

// Config mongodb configuration parameters
type Config struct {
	URL          string
	DB           string
	Username     string
	Password     string
	Service      string
	IsReplicaSet bool
}

func NewConfig(url, db, username, password, service string, isReplicaSet bool) *Config {
	config := &Config{
		URL:          url,
		DB:           db,
		Username:     username,
		Password:     password,
		Service:      service,
		IsReplicaSet: isReplicaSet,
	}
	return config
}
