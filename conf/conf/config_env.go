package conf

type AppConfig struct {
	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"1995"`
	DBUser   string `env:"DB_USER" envDefault:"test"`
	DBPass   string `env:"DB_PASS" envDefault:"Abc@12345"`
	DBName   string `env:"DB_NAME" envDefault:"test"`
	DBSchema string `env:"DB_SCHEMA" envDefault:"public"`

	NumberOfCharacter int    `env:"NUMBER_OF_CHARACTER" envDefault:"4"`
	TimeLife          int64  `env:"TIME_LIFE" envDefault:"120"`
	FailedTimes       int    `env:"FAILED_TIMES" envDefault:"3"`
	LockTime          int64  `env:"LOCK_TIME" envDefault:"3"`
	URLGetOTP         string `env:"URL_GET_OTP" envDefault:"https://api.twilio.com/2010-04-01/Accounts/"`

	// LOGIN SOCIAL CONFIG
	FBClientID       string `env:"FB_CLIENT_ID" envDefault:"3223945524335918"`
	FbClientSecret   string `env:"FB_CLIENT_SECRET" envDefault:"66188f8766be2d60ade18b4246068c31"`
	GGClientID       string `env:"GG_CLIENT_ID" envDefault:"255826360579-qg6j8tfrpddsa2u9bol443uf0duv0aji.apps.googleusercontent.com"`
	GGClientSecret   string `env:"GG_CLIENT_SECRET" envDefault:"OmhwhkbOKSmTQOtWUpMu5PoQ"`
	AppleTeamID      string `env:"APPLE_TEAM_ID" envDefault:"9N54GUYHL3"`
	AppleClientID    string `env:"APPLE_CLIENT_ID" envDefault:"app.bipbipcentral.com"`
	AppleCallbackURL string `env:"APPLE_CALLBACK_URL" envDefault:"https://api-stg-cen.jx.com.vn/api/login/applecallback"`
	AppleKeyID       string `env:"APPLE_KEY_ID" envDefault:"87WSSPNVDZ"`
	BeBaseUrl        string `env:"BE_BASE_URL" envDefault:"https://localhost:8080/"`
	FeBaseUrl        string `env:"FE_BASE_URL" envDefault:""`
	URLMappingCart   string `env:"URL_MAPPING_CART" envDefault:"http://ms-cart-management/api/mapping-after-login"`
}
