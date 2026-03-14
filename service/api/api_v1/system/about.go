package system

import (
	"vox-homepage/api/api_v1/common/apiReturn"
	"vox-homepage/lib/cmn"

	"github.com/gin-gonic/gin"
)

type About struct {
}

func (a *About) Get(c *gin.Context) {
	version := cmn.GetSysVersionInfo()
	apiReturn.SuccessData(c, gin.H{
		"versionName": version.Version,
		"versionCode": version.Version_code,
	})
}
