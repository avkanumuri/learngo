package urlshortner

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type UrlPaths struct {
	Path string `yaml:"path,omitempty"`
	Url  string `yaml:"url,omitempty"`
}

func MapsHandler(UrlMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := UrlMap[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YamlHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var up []UrlPaths
	err := yaml.Unmarshal(data, &up)
	if err != nil {
		fmt.Println(err)
	}
	mup := make(map[string]string)
	for _, item := range up {
		mup[item.Path] = item.Url
	}
	return MapsHandler(mup, fallback), nil
}
