package gorm_v2

import (
	"goskeleton/app/utils/yml_config"
	"goskeleton/app/utils/yml_config/interf"
)

var gormv2Conf interf.YmlConfigInterf

func init() {
	gormv2Conf = yml_config.CreateYamlFactory("gorm_v2")
}
