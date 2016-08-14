package golinks

import (
	//"html/template"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"
	//"appengine/log"
	//"appengine/user"
)

type Alias struct {
	Name 	string
	Url 	string
	Date    time.Time
}

func init() {
	http.HandleFunc("/", root)
}

func userName(r *http.Request) string {
	return ".beta"
}

func root(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doGet(w, r)
	case http.MethodPost:
		doPost(w, r)
	case http.MethodPut:
		doPut(w, r)
	default:
		// fallback to GET
		doGet(w, r)
	}
}

func doGet(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	path := strings.TrimPrefix(r.URL.Path, "/")
	k := datastore.NewKey(c, "Alias", path, 0, nil)

	a := new(Alias)
	if err := datastore.Get(c, k, a); err != nil {
		c.Infof("failed to Get %v", k)
		http.Error(w, "Alias doesn't exist", http.StatusNotFound)
		return
	}
	c.Debugf("Loaded Alias %v with key %v", a, k)

	http.Redirect(w, r, a.Url, http.StatusFound)
}

func doPost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	alias, err := parseBody(r)
	if err != nil {
		c.Errorf("failed to parse request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := datastore.NewKey(c, "Alias", ".beta/" + alias.Name, 0, nil)

	c.Debugf("Storing Alias %v with key %v", alias, key)
	if _, err := datastore.Put(c, key, &alias); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/.show/.beta/" + alias.Name, http.StatusFound)
}

func doPut(w http.ResponseWriter, r *http.Request) {
}

func createAlias(name string, url string) Alias {
	return Alias{ name, url, time.Now() }
}

func parseBody(r *http.Request) (Alias, error) {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		return parseForm(r)
	}
	if contentType == "application/json" {
		return parseJson(r)
	}

	return Alias{}, fmt.Errorf("Unknown Content-Type: %s", contentType)
}

func parseForm(r *http.Request) (Alias, error) {
	alias := Alias{ r.FormValue("name"), r.FormValue("url"), time.Now() }
	return alias, nil
}

func parseJson(r *http.Request) (Alias, error) {
	decoder := json.NewDecoder(r.Body)
	var alias Alias
	err := decoder.Decode(&alias)
	if err != nil {
		return alias, err
	}
	return alias, nil
}
