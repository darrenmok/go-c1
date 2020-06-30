package visitor_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/darrenmok/go-c1/visitor"
)

func TestGet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	visitor.Get(w, r)

	bytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	if !strings.Contains(string(bytes), "visitor") {
		t.Errorf("got %s", string(bytes))
	}

	count, err := strconv.ParseInt(strings.Split(string(bytes), " ")[3], 10, 64)
	if err != nil {
		t.Errorf("error parsing int: %v", err)
	}

	w = httptest.NewRecorder()
	visitor.Get(w, r)

	bytes, err = ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	nextCount, err := strconv.ParseInt(strings.Split(string(bytes), " ")[3], 10, 64)
	if err != nil {
		t.Errorf("error parsing int: %v", err)
	}

	if expCount := count + 1; expCount != nextCount {
		t.Errorf("want %d got %d", expCount, nextCount)
	}
}
