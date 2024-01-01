package configs

type Application struct {
	Name        string `mapstructure:"SERVICE_NAME"`
	Description string `mapstructure:"DESCRIPTION"`
	Endpoint    string `mapstructure:"ENDPOINT"`
	Version     string `mapstructure:"VERSION"`
	Port        int    `mapstructure:"PORT"`
}

type Services []struct {
	Name     string `mapstructure:"NAME"`
	Endpoint string `mapstructure:"ENDPOINT"`
}

type DynamoDB struct {
	Region    string `mapstructure:"REGION"`
	TableName string `mapstructure:"TABLE_NAME"`
}

type AWS struct {
	AccessKeyID     string `mapstructure:"ACCESS_KEY_ID"`
	SecretAccessKey string `mapstructure:"SECRET_ACCESS_KEY"`
	DefaultRegion   string `mapstructure:"DEFAULT_REGION"`
}
