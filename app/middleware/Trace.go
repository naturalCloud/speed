package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/syyongx/php2go"
	"speed/app/lib"
	"speed/app/lib/log"
	app "speed/bootstrap"
)

func Trace(ctx *gin.Context) {
	ctx.Set("traceId", uuid.New().String())
	uri := ctx.Request.URL.String()
	if uri != "/" {
		uri = php2go.Ltrim(php2go.StrReplace("/", "_", uri, -2), "_")
	} else {
		uri = "root"
	}
	ctx.Set("uri", uri)
	ctx.Set("startTime", lib.MsTime())
	log.WithCtx(ctx).Info("service start")

	//todo 未经测试
	if app.Config.GetString("appEnv") == "prod" {
		sqlLog := new(log.SqlLog)
		sqlLog.SetCtx(ctx)
		app.Db.SetLogger(sqlLog)
	}

	ctx.Next()
	log.WithCtx(ctx).Infof("service end --- 耗时 %v ms", lib.MsTime()-ctx.Value("startTime").(int64))

}
