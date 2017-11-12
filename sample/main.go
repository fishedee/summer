package main

import (
	"encoding/json"
	_ "github.com/fishedee/summer/sample/service"
	"github.com/fishedee/summer/sample/util"
	"net/http"
)

func logMiddle(handlers []interface{}) interface{} {
	preHandler := handlers[len(handlers)-1].(func(http.ResponseWriter, *http.Request))

	return func(writer http.ResponseWriter, request *http.Request) {
		util.MyLog.Debug("Request In %v", request.URL.Path)
		preHandler(writer, request)
		util.MyLog.Debug("Request Out %v", request.URL.Path)
	}
}

func jsonMiddle(handlers []interface{}) interface{} {
	preHandler := handlers[len(handlers)-1].(func() interface{})
	return func(writer http.ResponseWriter, request *http.Request) {
		result := preHandler()
		jsonMap := map[string]interface{}{
			"code": 0,
			"msg":  "",
			"data": result,
		}
		resultByte, err := json.Marshal(jsonMap)
		if err != nil {
			panic(err)
		}
		writer.Write(resultByte)
	}
}

func handleA() interface{} {
	return "Hello Fish!AA"
}

func handleB() interface{} {
	return "Hello Fish!BB"
}

func main() {
	util.MyLog.Debug("Server is running...")
	server := util.NewServer()
	server.Use(logMiddle)
	server.Use(jsonMiddle)
	server.Route("/a", handleA)
	server.Route("/b", handleB)
	server.Run(":8073")
}
