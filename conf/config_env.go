package conf

type AppConfig struct {
	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"5432"`
	DBUser   string `env:"DB_USER" envDefault:"postgres"`
	DBPass   string `env:"DB_PASS" envDefault:"postgres"`
	DBName   string `env:"DB_NAME" envDefault:"postgres"`
	DBSchema string `env:"DB_SCHEMA" envDefault:"public"`

	// LOGIN SOCIAL CONFIG
	FBClientID       string `env:"FB_CLIENT_ID" envDefault:"3223945524335918"`
	FbClientSecret   string `env:"FB_CLIENT_SECRET" envDefault:"66188f8766be2d60ade18b4246068c31"`
	GGClientID       string `env:"GG_CLIENT_ID" envDefault:"198497907727-8ku6dfuq0f9unt2p4i4saed5n0biqtqq.apps.googleusercontent.com"`
	GGClientSecret   string `env:"GG_CLIENT_SECRET" envDefault:"vTX2b6WrAG1xeFIi6KSkC7QN"`
	AppleTeamID      string `env:"APPLE_TEAM_ID" envDefault:"9N54GUYHL3"`
	AppleClientID    string `env:"APPLE_CLIENT_ID" envDefault:"app.bipbipcentral.com"`
	AppleCallbackURL string `env:"APPLE_CALLBACK_URL" envDefault:"https://api-stg-cen.jx.com.vn/api/login/applecallback"`
	AppleKeyID       string `env:"APPLE_KEY_ID" envDefault:"87WSSPNVDZ"`
	BeBaseUrl        string `env:"BE_BASE_URL" envDefault:"https://localhost:8080/"`
	FeBaseUrl        string `env:"FE_BASE_URL" envDefault:""`
	URLMappingCart   string `env:"URL_MAPPING_CART" envDefault:"http://ms-cart-management/api/mapping-after-login"`
}
