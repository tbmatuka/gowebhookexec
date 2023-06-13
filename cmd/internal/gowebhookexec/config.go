package gowebhookexec

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	Host    string `mapstructure:"host"`
	Port    string `mapstructure:"port"`
	SslKey  string `mapstructure:"ssl_key"`
	SslCert string `mapstructure:"ssl_cert"`
	Handler map[string]*ViperHandlerConfig
}

type ViperHandlerConfig struct {
	Key     string `mapstructure:"key"`
	CmdName string `mapstructure:"cmd"`
}

func LoadConfig() ViperConfig {
	pflag.String("host", "", "IP address or interface to listen on. Listens everywhere if empty.")
	pflag.String("port", "1234", "TCP port to listen on.")
	pflag.String("sslkey", "", "Path to the SSL key file.")
	pflag.String("sslcert", "", "Path to the SSL certificate file.")
	pflag.String("default-key", "", "Secret key for the default config.")
	pflag.String("default-cmd", "date", "Command for the default config.")
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Panicf("Error binding command line options: %v\n", err)
	}

	var config ViperConfig

	config.Host = viper.GetString("host")
	config.Port = viper.GetString("port")
	config.SslKey = viper.GetString("sslkey")
	config.SslCert = viper.GetString("sslcert")

	config.Handler = make(map[string]*ViperHandlerConfig)

	if viper.GetString("default-key") != "" {
		config.Handler["default"] = new(ViperHandlerConfig)
		config.Handler["default"].Key = viper.GetString("default-key")
		config.Handler["default"].CmdName = viper.GetString("default-cmd")
	}

	viper.SetConfigName("webhook-exec")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/.config/webhook-exec/")

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: %v\n", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config
}
