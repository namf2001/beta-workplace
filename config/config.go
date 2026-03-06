package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// EnvConfig defines the expected environment variables for the application.
type EnvConfig struct {
	AppPort     string `envconfig:"APP_PORT" default:"8080"`
	AppEnv      string `envconfig:"APP_ENV" default:"dev"`
	AppDebug    bool   `envconfig:"APP_DEBUG" default:"true"`
	AppTimezone string `envconfig:"APP_TIMEZONE" default:"Asia/Ho_Chi_Minh"`

	DBHost         string `envconfig:"DB_HOST" required:"true"`
	DBPort         string `envconfig:"DB_PORT" required:"true"`
	DBUser         string `envconfig:"DB_USER" required:"true"`
	DBPassword     string `envconfig:"DB_PASSWORD" required:"true"`
	DBName         string `envconfig:"DB_NAME" required:"true"`
	DBSSLMode      string `envconfig:"DB_SSL_MODE" required:"true"`
	DBMaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS"`

	JWTSecret         string        `envconfig:"JWT_SECRET" required:"true"`
	JWTAccessDuration time.Duration `envconfig:"JWT_ACCESS_DURATION"`

	GoogleRedirectURL  string `envconfig:"GOOGLE_REDIRECT_URL"`
	GoogleClientID     string `envconfig:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `envconfig:"GOOGLE_CLIENT_SECRET"`

	MailSMTPHost     string `envconfig:"MAIL_SMTP_HOST"`
	MailSMTPPort     int    `envconfig:"MAIL_SMTP_PORT"`
	MailSMTPUser     string `envconfig:"MAIL_SMTP_USER" required:"true"`
	MailSMTPPassword string `envconfig:"MAIL_SMTP_PASSWORD"`
	MailEmailFrom    string `envconfig:"MAIL_EMAIL_FROM"`
}

var c *EnvConfig

// findProjectRoot walks up from CWD to find the directory containing go.mod
func findProjectRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	cur := wd
	for {
		if _, err := os.Stat(filepath.Join(cur, "go.mod")); err == nil {
			return cur
		}
		parent := filepath.Dir(cur)
		if parent == cur {
			// Reached filesystem root
			return wd
		}
		cur = parent
	}
}

// Init initializes config
func Init(env string) {

	root := findProjectRoot()

	// Load base .env (optional) from the project root
	_ = godotenv.Load(filepath.Join(root, ".env"))

	// Load env-specific file from the project root (optional)
	envFile := filepath.Join(root, fmt.Sprintf(".env.%s", env))
	if _, err := os.Stat(envFile); err == nil {
		if err := godotenv.Overload(envFile); err != nil {
			log.Printf("warning: could not load env file %s: %v", envFile, err)
		}
	} else if os.IsNotExist(err) {
		log.Printf("warning: env file %s not found; relying on environment variables", envFile)
	} else {
		log.Printf("warning: cannot stat env file %s: %v", envFile, err)
	}

	// Validate environment variables using envconfig
	var envCfg EnvConfig
	if err := envconfig.Process("", &envCfg); err != nil {
		log.Fatalf("[env] validation failed: %v", err)
	}

	c = &envCfg
}

// GetConfig returns config
func GetConfig() *EnvConfig {
	return c
}
