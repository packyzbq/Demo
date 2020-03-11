package frame

import (
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, _ := r.getRoute("GET", "/hello/b/c")

	if n == nil {
		t.Fail()
	}

	if n.pattern != "/hello/b/c" {
		t.Fail()
	}

	//if ps["name"] != "name" {
	//	t.Fail()
	//}
}

func TestOutputRoute(t *testing.T) {
	r := newTestRouter()
	n := r.roots["GET"]
	Output(n, 0)
}
