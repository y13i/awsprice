// this file implements API query via HTTPS and cache handling

package awsprice

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

type backendClient struct {
	CacheTTL string
}

func (b backendClient) getData(p APIPath, v interface{}) error {
	bytes, e := getCache(p, b.CacheTTL)

	if e != nil {
		if bytes, e = getAPI(p); e != nil {
			return e
		}

		if e = putCache(p, bytes); e != nil {
			return e
		}
	}

	if e = json.Unmarshal(bytes, &v); e != nil {
		return e
	}

	return nil
}

func getAPI(p APIPath) ([]byte, error) {
	// print("calling api")

	var b []byte

	r, e := http.Get(APIBaseURL + string(p))

	if e != nil {
		return b, e
	}

	defer r.Body.Close()

	if b, e = ioutil.ReadAll(r.Body); e != nil {
		return b, e
	}

	return b, nil
}

func getCache(p APIPath, ttl string) ([]byte, error) {
	if cacheIsValid(p, ttl) {
		return ioutil.ReadFile(cachePath(p))
	}

	return []byte{}, errors.New("cache is not valid")
}

func putCache(p APIPath, b []byte) error {
	cp := cachePath(p)

	if e := os.MkdirAll(path.Dir(cp), 0755); e != nil {
		return e
	}

	return ioutil.WriteFile(cp, b, 0644)
}

func cachePath(p APIPath) string {
	return path.Join(os.TempDir(), "awsprice", string(p))
}

func cacheIsValid(p APIPath, durationString string) bool {
	s, e := os.Stat(cachePath(p))

	if e != nil {
		return false
	}

	d, e := time.ParseDuration(durationString)

	if e != nil {
		return false
	}

	return s.ModTime().Add(d).After(time.Now())
}

func clearCache() error {
	return os.RemoveAll(cachePath(""))
}
