package main

import ( 
	"go.uber.org/zap"
	"time"
)

func nosugar(url string) {
	//logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
)

}

func main() {
//	logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL", "attempt", 3)
	sugar.Infof("Failed to fetch URL: %s", "www.innodep.com")

	nosugar("www.google.com")
}


