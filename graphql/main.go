package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sandbox-graphql/modules/music"
	"time"
)

func main() {
	cm := chainMiddleware(withLogging, withTracing)
	http.HandleFunc("/graphql", cm(func(w http.ResponseWriter, r *http.Request) {
		music.GetMusic(w, r)
	}))

	fmt.Println("listening on port 12345...")
	//tét
	t := reflect.TypeOf(music.Album{})
	f, _ := t.FieldByName("Artist")
	fmt.Println(f.Tag.Lookup("json"))

	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal(err)
	}
}

type middleware func(next http.HandlerFunc) http.HandlerFunc

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		defer log.Println("logged end")
		time.Sleep(1 * time.Second)
		next.ServeHTTP(w, r)
	}
}

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
		defer log.Println("trace end")
		time.Sleep(1 * time.Second)
		next.ServeHTTP(w, r)
	}
}

func chainMiddleware(mw ...middleware) middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}
