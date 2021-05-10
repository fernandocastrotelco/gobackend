package config

import (
	"fmt"
	"os"
)

var (
	// HOST de la base de datos
	HOST = getEnv("HOST", "localhost")
	// PORT returns the server listening port
	PORT = getEnv("PORT", "33060")
	// USER return name of username database
	USER = getEnv("USER", "root")
	// PASSW return password of database
	PASSW = getEnv("PASSW", "telcodb")
	// DB returns the name of the database
	DB = getEnv("DB", "datosmaestros")
	// RENAPER endpoint HOST
	RENAPER = getEnv("RENAPER", "http://45.63.111.223:6105")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
