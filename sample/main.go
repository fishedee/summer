package main

import (
	"encoding/json"
	"github.com/fishedee/summer/ioc"
	"github.com/fishedee/summer/sample/api"
	_ "github.com/fishedee/summer/sample/service"
	"github.com/fishedee/summer/sample/util"
	"net/http"
	"strconv"
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
	preHandler := handlers[len(handlers)-1].(func(map[string][]string) interface{})
	return func(writer http.ResponseWriter, request *http.Request) {
		query := map[string][]string(request.URL.Query())
		var jsonMap interface{}
		func() {
			defer func() {
				err := recover()
				if err != nil {
					jsonMap = map[string]interface{}{
						"code": 1,
						"msg":  err,
						"data": nil,
					}
				}
			}()
			result := preHandler(query)
			jsonMap = map[string]interface{}{
				"code": 0,
				"msg":  "",
				"data": result,
			}
		}()
		resultByte, err := json.Marshal(jsonMap)
		if err != nil {
			panic(err)
		}
		writer.Write(resultByte)
	}
}

type Controller struct {
	userAo api.UserAo
}

func (this *Controller) Get(query map[string][]string) interface{} {
	id := query["id"]
	if id == nil {
		panic("Unknown Id")
	}
	idInt, err := strconv.Atoi(id[0])
	if err != nil {
		panic(err)
	}
	return this.userAo.Get(idInt)
}

func (this *Controller) Add(query map[string][]string) interface{} {
	name := query["name"]
	if name == nil {
		panic("Unknown name")
	}
	return this.userAo.Add(api.User{
		Name: name[0],
	})
}

func NewController(userAo api.UserAo) *Controller {
	result := &Controller{}
	result.userAo = userAo
	return result
}

func main() {
	controller := ioc.New((*Controller)(nil)).(*Controller)

	util.MyLog.Debug("Server is running...")
	server := util.NewServer()
	server.Use(logMiddle)
	server.Use(jsonMiddle)
	server.Route("/get", controller.Get)
	server.Route("/add", controller.Add)
	server.Run(":8073")
}

func init() {
	ioc.Register(NewController)
}
