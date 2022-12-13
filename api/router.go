package api

import (
	"go.x2ox.com/THz"
	"go.x2ox.com/THz/middleware/cors"
	"go.x2ox.com/THz/middleware/recovery"
)

func Router() *THz.THz {
	t := THz.New()
	if err := t.SetTrustedProxies(
		"10.0.0.0/8", "127.0.0.1/8", "172.16.0.0/12", "192.168.0.0/16", "fc00::/7",
	); err != nil {
		panic(err)
	}
	t.SetTrustedHeaders("Cf-Connecting-Ip", "X-Forwarded-For")

	t.AddIntercept(recovery.New().Middleware())
	t.AddIntercept(cors.New().Middleware())

	t.GET("/test", func(c *THz.Context) { c.JSON("hello world") })

	{
		upload := t.Group("/api/v1/upload")
		upload.POST("/media", UploadMediaAction)
		upload.POST("/chunk", UploadChunkAction)
		upload.POST("/merge", MergeChunkAction)
		upload.POST("/check", CheckChunkAction)
	}

	return t
}
