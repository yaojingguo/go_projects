package main

import (
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"log"
	"yao/gin-sw/v2/kv"

	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
	zapplugin "github.com/SkyAPM/go2sky-plugins/zap"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var zapLog *zap.SugaredLogger

var tracer *go2sky.Tracer
func main() {
	// Set up SkyWalking
	re, err := reporter.NewGRPCReporter("localhost:11800")
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer re.Close()

	tracer, err = go2sky.NewTracer("gin-server", go2sky.WithReporter(re))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}

	r := gin.Default()
	r.Use(v3.Middleware(r, tracer))

	r.GET("/ping", func(c *gin.Context) {
		ctx := c.Request.Context()
		one(ctx)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/etcd", func(c *gin.Context) {
		ctx := c.Request.Context()
		cli := kv.NewClient()
		defer cli.Close()
		val, err := kv.Get(ctx, cli, "TestPut")
		if err != nil {
			c.JSON(500, gin.H{
				"message": "interval error",
			})
		}
		fmt.Println(val)
		zapLog.Infof("val: ", val)
		c.JSON(200, gin.H{
			"no": val,
		})
	})

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8888")
}

func one(ctx context.Context) {
	subSpan, newCtx, err := tracer.CreateLocalSpan(ctx, go2sky.WithOperationName("one-span"))
	if err != nil {
		panic(err)
	}
	defer subSpan.End()

	logger := zap.NewExample()
	logger.With(zapplugin.TraceContext(ctx)...).Info("test")

	fmt.Println("one func")
	two(newCtx)
}

func two(ctx context.Context) {
	subSpan, newCtx, err := tracer.CreateLocalSpan(ctx, go2sky.WithOperationName("two-span"))
	if err != nil {
		panic(err)
	}
	defer subSpan.End()
	fmt.Println("two func")
	three(newCtx)
}

func three(ctx context.Context) {
	fmt.Println("three func")
}

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	zapLog = logger.Sugar()
}