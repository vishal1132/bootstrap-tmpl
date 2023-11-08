{% if postgres_enabled %}
package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"{{ module_name }}/utils"
)

type DatabaseConfig struct {
	DatabaseUser       string
	DatabasePassword   string
	DatabaseName       string
	DatabaseHost       string
	DatabasePort       string
	MaxIdleConnections int
	MaxOpenConnections int
	LogMode            bool
	ConnectionString   string
}

const (
	ConfigDatabaseName        = "database.name"
	ConfigDatabasePassword    = "database.password"
	ConfigDatabaseHost        = "database.host"
	ConfigDatabasePort        = "database.port"
	ConfigDatabaseUsername    = "database.username"
	ConfigDatabaseMaxIdleConn = "database.maxIdleConn"
	ConfigDatabaseMaxOpenConn = "database.maxOpenConn"
	ConfigDatabaseLogMode     = "database.debugEnabled"
)

func loadDBConfig() *DatabaseConfig {
	user := viper.GetString(ConfigDatabaseUsername)
	password := viper.GetString(ConfigDatabasePassword)
	databaseName := viper.GetString(ConfigDatabaseName)
	host := viper.GetString(ConfigDatabaseHost)
	port := viper.GetString(ConfigDatabasePort)
	maxIdleConn := viper.GetInt(ConfigDatabaseMaxIdleConn)
	maxOpenConn := viper.GetInt(ConfigDatabaseMaxOpenConn)
	logMode := viper.GetBool(ConfigDatabaseLogMode)
	dbConfig := &DatabaseConfig{
		DatabaseUser:       user,
		DatabasePassword:   password,
		DatabaseName:       databaseName,
		DatabaseHost:       host,
		DatabasePort:       port,
		MaxIdleConnections: maxIdleConn,
		MaxOpenConnections: maxOpenConn,
		LogMode:            logMode,
	}
	dbConfig.ConnectionString = getDatabaseConnectionString(dbConfig)
	return dbConfig
}

func getDatabaseConnectionString(databaseConfig *DatabaseConfig) string {
	user := databaseConfig.DatabaseUser
	password := databaseConfig.DatabasePassword
	database := databaseConfig.DatabaseName
	host := databaseConfig.DatabaseHost
	port := databaseConfig.DatabasePort

	psqlConf := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s %s",
		host,
		port,
		user,
		database,
		getSSlConnString(),
	)
	if password != "" {
		psqlConf = fmt.Sprintf("%s password=%s", psqlConf, password)
	}
	return psqlConf
}

func getSSlConnString() string {
	sslMode := viper.GetString("database.sslMode")
	switch sslMode {
	case "disable":
		return "sslmode=disable"
	case "verify-ca":
		path := "/dev/shm"
		writeToFile(path+"/server-ca.pem", viper.GetString("database.sslRootCert"))
		writeToFile(path+"/client-cert.pem", viper.GetString("database.sslCert"))
		writeToFile(path+"/client-key.pem", viper.GetString("database.sslKey"))
		sslStringFmter := "sslmode=verify-ca sslrootcert=%s/server-ca.pem sslcert=%s/client-cert.pem sslkey=%s/client-key.pem"
		var args = []interface{}{path, path, path}
		return fmt.Sprintf(sslStringFmter, args...)
	default:
		return ""
	}
}

func writeToFile(filename string, content string) {
	f := utils.Must(os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755))
	_, err := f.WriteString(content)
	f.Close()
	utils.Must(0, err)
}
{% endif %}