package ginHttp

import (
	"time"

	cmodel "github.com/Cepave/open-falcon-backend/common/model"

	"github.com/gin-gonic/gin"
	"github.com/masato25/open_lambda_query/g"
	"github.com/masato25/open_lambda_query/gin_http/computeFunc"
	grahttp "github.com/masato25/open_lambda_query/gin_http/grafana"
	"github.com/masato25/open_lambda_query/gin_http/openFalcon"
)

type QueryInput struct {
	StartTs       time.Time
	EndTs         time.Time
	ComputeMethod string
	Endpoint      string
	Counter       string
}

//this function will generate query string obj for QueryRRDtool
func getq(q QueryInput) cmodel.GraphQueryParam {
	request := cmodel.GraphQueryParam{
		Start:     q.StartTs.Unix(),
		End:       q.EndTs.Unix(),
		ConsolFun: q.ComputeMethod,
		Endpoint:  q.Endpoint,
		Counter:   q.Counter,
	}
	return request
}

//accept cross domain request
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func StartWeb() {
	handler := gin.Default()
	handler.Use(CORSMiddleware())
	conf := g.Config()

	compute := handler.Group("/func")
	compute.GET("/compute", computeFunc.Compute)
	compute.GET("/funcations", computeFunc.GetAvaibleFun)
	compute.GET("/smapledata", computeFunc.GetTestData)

	openfalcon := handler.Group("/owl")
	openfalcon.GET("/endpoints", openFalcon.GetEndpoints)
	openfalcon.GET("/queryrrd", openFalcon.QueryData)

	grafana := handler.Group("/api/grafana")
	grafana.GET("/", grahttp.GrafanaMain)
	grafana.GET("/metrics/find", grahttp.GrafanaMain)
	grafana.POST("/render", grahttp.GetQueryTargets)
	handler.Run(conf.GinHttp.Listen)
}
