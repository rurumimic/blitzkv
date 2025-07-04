package config

import "fmt"

func Display(conf *Config) {
	fmt.Println("⚡ Starting BlitzKV")
	fmt.Println("------------------------------------------")
	fmt.Printf("App        : %s\n", conf.App.Name)
	fmt.Printf("Version    : %s\n", conf.App.Version)
	fmt.Printf("Environment: %s\n", conf.App.Env)
	fmt.Printf("Debug Mode : %t\n", conf.App.Debug)
	fmt.Println()
	fmt.Printf("Port       : %d\n", conf.Server.Port)
	fmt.Println()
	fmt.Printf("Max Keys   : %d\n", conf.Store.MaxKeys)
	fmt.Printf("Max Value  : %d bytes\n", conf.Store.MaxValueSize)
	fmt.Println()
	fmt.Printf("Log Level  : %s\n", conf.Log.Level)
	fmt.Println("------------------------------------------")
	fmt.Println("✅ Configuration loaded successfully.")
}
