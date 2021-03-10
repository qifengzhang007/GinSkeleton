package casbin_v2

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"strings"
	"time"
)

//创建 casbin Enforcer(执行器)
func InitCasbinEnforcer() (*casbin.SyncedEnforcer, error) {
	var tmpDbConn *gorm.DB
	var Enforcer *casbin.SyncedEnforcer
	switch strings.ToLower(variable.ConfigGormv2Yml.GetString("Gormv2.UseDbType")) {
	case "mysql":
		if variable.GormDbMysql == nil {
			return nil, errors.New(my_errors.ErrorCasbinCanNotUseDbPtr)
		}
		tmpDbConn = variable.GormDbMysql
	case "sqlserver", "mssql":
		if variable.GormDbSqlserver == nil {
			return nil, errors.New(my_errors.ErrorCasbinCanNotUseDbPtr)
		}
		tmpDbConn = variable.GormDbSqlserver
	case "postgre", "postgresql", "postgres":
		if variable.GormDbPostgreSql == nil {
			return nil, errors.New(my_errors.ErrorCasbinCanNotUseDbPtr)
		}
		tmpDbConn = variable.GormDbPostgreSql
	default:
	}

	prefix := variable.ConfigYml.GetString("Casbin.TablePrefix")
	tbName := variable.ConfigYml.GetString("Casbin.TableName")

	a, err := gormadapter.NewAdapterByDBUseTableName(tmpDbConn, prefix, tbName)
	if err != nil {
		return nil, errors.New(my_errors.ErrorCasbinCreateAdaptFail)
	}
	modelConfig := variable.ConfigYml.GetString("Casbin.ModelConfig")

	if m, err := model.NewModelFromString(modelConfig); err != nil {
		return nil, errors.New(my_errors.ErrorCasbinNewModelFromStringFail + err.Error())
	} else {
		if Enforcer, err = casbin.NewSyncedEnforcer(m, a); err != nil {
			return nil, errors.New(my_errors.ErrorCasbinCreateEnforcerFail)
		}
		_ = Enforcer.LoadPolicy()
		AutoLoadSeconds := variable.ConfigYml.GetDuration("Casbin.AutoLoadPolicySeconds")
		Enforcer.StartAutoLoadPolicy(time.Second * AutoLoadSeconds)
		return Enforcer, nil
	}
}
