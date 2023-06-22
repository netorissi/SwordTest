package shared

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func init() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	Logger = log.Sugar()
}
