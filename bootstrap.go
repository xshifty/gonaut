package gonaut

import (
	"fmt"
	"net/http"
	"reflect"
)

// Bootstrap struct to holds and initialize application
type Bootstrap struct {
	controllersPath string
}

// NewBootstrap create a Bootstrap instance
func NewBootstrap(initializer func(b *Bootstrap)) *Bootstrap {
	b := Bootstrap{controllersPath: "./controllers"}
	initializer(&b)
	return &b
}

// Run will initialize the application
func (b *Bootstrap) Run(port int) {
	parsingMap := buildParsingMapFromControllersDir(b.controllersPath)
	routingPreload := buildRoutingPreload(parsingMap)

	endpoints := map[string]map[string]func(http.ResponseWriter, *http.Request){}
	for _, r := range routingPreload {
		uri := r.a.uri
		method := r.a.method
		if _, ok := endpoints[uri]; !ok {
			endpoints[uri] = map[string]func(http.ResponseWriter, *http.Request){}
		}
		if _, ok := endpoints[uri][method]; !ok {
			endpoints[uri][method] = Get(fmt.Sprintf("controllers.%s:%s", r.c, r.f)).(func(http.ResponseWriter, *http.Request))
		}
	}
	for ep, methods := range endpoints {
		http.HandleFunc(ep, func(w http.ResponseWriter, r *http.Request) {
			if method, ok := methods[r.Method]; ok {
				method(w, r)
			}
		})
	}

	LogDebug("Starting server on port %d...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		LogFatal("Server error: %s", err)
	}
}

func (b *Bootstrap) RegisterController(ctrl interface{}) {
	if ctrl == nil {
		LogFatal("Controller can't be nil")
	}

	cto := reflect.TypeOf(ctrl)
	if cto.Field(0).Name != "BaseController" {
		LogFatal("Struct %s must extends BaseController", cto.Name())
	}
	controllerDep := fmt.Sprintf("controllers.%s", cto.Name())
	Set(controllerDep, func(c Container) interface{} {
		return ctrl
	})

	for i := 0; i < cto.NumMethod(); i++ {
		method := cto.Method(i)
		depName := fmt.Sprintf("controllers.%s:%s", cto.Name(), method.Name)
		Set(depName, func(c Container) interface{} {
			return func(w http.ResponseWriter, r *http.Request) {
				LogInfo("%s %s", r.Method, r.RequestURI)
				method.Func.Call([]reflect.Value{reflect.ValueOf(Get(controllerDep)), reflect.ValueOf(w), reflect.ValueOf(r)})
			}
		})
		LogDebug("Registering controller action: %s", depName)
	}
}

// WithControllers set default controllers path for application bootstrap
func (b *Bootstrap) WithControllers(p string) {
	b.controllersPath = p
}
