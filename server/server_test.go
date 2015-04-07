package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/alexlyulkov/conf/conf"
)

func AssertEqual(t *testing.T, value1, value2 interface{}) {
	if !reflect.DeepEqual(value1, value2) {
		panic("AsserEqual failed: " + fmt.Sprint(value1) +
			" != " + fmt.Sprint(value2))
	}
}

func PostForm(method func(w http.ResponseWriter, r *http.Request),
	params url.Values) (code int, body string, decodedJSONValues interface{}) {

	req, err := http.NewRequest("POST", "", strings.NewReader(params.Encode()))
	if err != nil {
		panic(err.Error())
	}

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded; param=value",
	)

	w := httptest.NewRecorder()

	(http.HandlerFunc)(method).ServeHTTP(w, req)

	code = w.Code
	body = w.Body.String()

	if (code == 200) && (len(body) != 0) {
		err := json.Unmarshal(([]byte)(body), &decodedJSONValues)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}

func TestAll(t *testing.T) {
	conf.InitRootDirectory("/var/tmp/alex_config_tmp")
	conf.DeleteNode("")
	defer conf.DeleteNode("")

	code, _, _ := PostForm(Insert,
		url.Values{"name": {"n1"}, "value": {`"v1"`}})
	AssertEqual(t, code, 200)

}

func TestBadNames(t *testing.T) {
	conf.InitRootDirectory("/var/tmp/alex_config_tmp")
	conf.DeleteNode("")
	defer conf.DeleteNode("")

	code, body, _ := PostForm(Insert,
		url.Values{"name": {"!&*rkt"}, "value": {`"dknfgergt"`}})
	AssertEqual(t, code, 400)
	AssertEqual(t, ([]byte)(body), ([]byte)("Name should consist only of English letters and numbers separated by dots.\x0A"))

	code, body, _ = PostForm(Update,
		url.Values{"name": {"!&*erte"}, "value": {`"dknfgergt"`}})
	AssertEqual(t, code, 400)
	AssertEqual(t, ([]byte)(body), ([]byte)("Name should consist only of English letters and numbers separated by dots.\x0A"))

	code, body, _ = PostForm(Delete,
		url.Values{"name": {"wqe qwwqeq"}})
	AssertEqual(t, code, 400)
	AssertEqual(t, ([]byte)(body), ([]byte)("Name should consist only of English letters and numbers separated by dots.\x0A"))

	code, body, _ = PostForm(Read,
		url.Values{"name": {"!&*erte"}})
	AssertEqual(t, code, 400)
	AssertEqual(t, ([]byte)(body), ([]byte)("Name should consist only of English letters and numbers separated by dots.\x0A"))
}

func TestBadValues(t *testing.T) {
	conf.InitRootDirectory("/var/tmp/alex_config_tmp")
	conf.DeleteNode("")
	defer conf.DeleteNode("")

	code, body, _ := PostForm(Insert, url.Values{"name": {"erte"}})
	AssertEqual(t, code, 400)
	AssertEqual(t, ([]byte)(body), ([]byte)("Node value is not specified\x0A"))

	code, body, _ = PostForm(Update, url.Values{"name": {"erte"}})
	AssertEqual(t, code, 400)
	AssertEqual(t, ([]byte)(body), ([]byte)("Node value is not specified\x0A"))

	code, body, _ = PostForm(Insert,
		url.Values{"name": {"erte"}, "value": {"dknfgergt"}})
	AssertEqual(t, code, 400)
	AssertEqual(t, strings.HasPrefix(body, "Node value should be proper json. Can't parse node value:"), true)

	code, body, _ = PostForm(Update,
		url.Values{"name": {"erte"}, "value": {"dknfgergt"}})
	AssertEqual(t, code, 400)
	AssertEqual(t, strings.HasPrefix(body, "Node value should be proper json. Can't parse node value:"), true)

}

func TestValueIsNotMapsAndStrings(t *testing.T) {
	conf.InitRootDirectory("/var/tmp/alex_config_tmp")
	conf.DeleteNode("")
	defer conf.DeleteNode("")

	code, _, _ := PostForm(Insert,
		url.Values{"name": {"erte"}, "value": {`{"x":5}`}})
	AssertEqual(t, code, 400)

	code, _, _ = PostForm(Update,
		url.Values{"name": {"erte"}, "value": {`{"x":5, "y":"esserf"}`}})
	AssertEqual(t, code, 400)

}

func TestInsertUpdateReadDelete(t *testing.T) {
	conf.InitRootDirectory("/var/tmp/alex_config_tmp")
	conf.DeleteNode("")
	defer conf.DeleteNode("")

	tree := make(map[string]interface{})
	tree["subtree01"] = make(map[string]interface{})
	tree["subtree01"].(map[string]interface{})["i01"] = "v1"
	tree["subtree01"].(map[string]interface{})["i02"] = "v2"
	tree["i03"] = "v3"

	code, _, _ := PostForm(Insert,
		url.Values{"name": {"tree01"}, "value": {`{"subtree01":{"i01":"v1", "i02":"v2"}, "i03":"v3"}`}})
	AssertEqual(t, code, 200)

	code, _, loadedTree := PostForm(Read,
		url.Values{"name": {"tree01"}, "depth": {"2"}})
	AssertEqual(t, code, 200)
	AssertEqual(t, tree, loadedTree)

	code, _, _ = PostForm(Update,
		url.Values{"name": {"tree01"}, "value": {`{"subtree01":{"i01":"v1", "i02":"v2"}, "i03":"v3_2"}`}})
	AssertEqual(t, code, 200)

	tree["i03"] = "v3_2"

	code, _, loadedTree = PostForm(Read,
		url.Values{"name": {"tree01"}, "depth": {"-1"}})
	AssertEqual(t, code, 200)
	AssertEqual(t, tree, loadedTree)

	code, _, _ = PostForm(Delete,
		url.Values{"name": {"tree01"}})
	AssertEqual(t, code, 200)

	code, body, _ := PostForm(Read,
		url.Values{"name": {"tree01"}, "depth": {"-1"}})
	AssertEqual(t, code, 400)
	AssertEqual(t, body, "Node with such name does not exist.\x0A")
}
