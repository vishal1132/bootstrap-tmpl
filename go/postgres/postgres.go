{% if postgres_enabled %}
package postgres

import (
	"context"

	"{{ module_name }}/config"
	"{{ module_name }}/errors"
	"{{ module_name }}/observability/log"
	"{{ module_name }}/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(ctx context.Context, databaseConfig *config.DatabaseConfig) *gorm.DB {
	log.Info(ctx, "Connecting to postgres")
	db := utils.Must(gorm.Open(postgres.Open(databaseConfig.ConnectionString), &gorm.Config{}))
	log.Info(ctx, "Connected to postgres")
	return db
}

func isRecordNotFoundError(resp *gorm.DB) bool {
	return resp.Error != nil && resp.Error.Error() == "record not found"
}

func HandleError(resp *gorm.DB) error {

	if isRecordNotFoundError(resp) && resp.RowsAffected == 0 {
		return errors.NewRecordNotFoundError(resp.Error)
	} else if resp.Error == nil {
		return nil
	} else {
		return errors.NewDatabaseError(resp.Error)
	}
}
{% endif %}