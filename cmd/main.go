package main

import (
	"go.uber.org/zap"
	"go.x2ox.com/THz"
	"go.x2ox.com/sorbifolia/rogu"
	"os"
	"os/signal"
	"syscall"
	"video-server/api"
)

func init() {
	rogu.MustReplaceGlobals(rogu.DefaultZapConfig(rogu.DefaultZapEncoderConfig(),
		[]string{"stdout"},
		[]string{"stderr"}))
}

func main() {
	r := api.Router()
	r.SetLog(zap.L().Named("THZ"))
	handleSignal(r)

	if err := r.ListenAndServe(":8080"); err != nil {
		zap.L().Fatal("server exit", zap.Error(err))
	}
}

func handleSignal(server *THz.THz) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		sig := <-c
		zap.L().Info("server exit signal", zap.Any("signal notify", sig))

		_ = server.Stop()

		os.Exit(0)
	}()
}
