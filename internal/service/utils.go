package service

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {

	logger, err := zap.NewProduction()
	defer logger.Sync()

	/*{"level":"error",
	"ts":1684425318.217807,
	"caller":"service/clienthandlefuncs.go:85",
	"msg":"Failed to decode request: %!w(*json.SyntaxError=&{invalid character 'e' looking for beginning of value 12})",
	"stacktrace":"github.com/NotYourAverageFuckingMisery/animello/internal/service.(*Service).EditEmail.func1\n\t/Users/diflapuna/Documents/My programms/Golang/eblan/internal/service/clienthandlefuncs.go:85\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2109\ngithub.com/gorilla/mux.(*Router).ServeHTTP\n\t/Users/diflapuna/go/pkg/mod/github.com/gorilla/mux@v1.8.0/mux.go:210\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2947\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1991"}
	*/

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	sugar := logger.Sugar()
	sugar.Infow(
		"Logger is started and ready to log.",
		"url", "http://localhost:6969",
	)
	return sugar
}
