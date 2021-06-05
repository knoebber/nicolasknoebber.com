package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"microblog"
	"mime"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const defaultAddress = ":4000"

func main() {
	addr := flag.String("addr", defaultAddress, "HTTP address to listen on")

	r := mux.NewRouter()
	s := &http.Server{
		Addr:         *addr,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	r.HandleFunc("/posts", listPosts).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	r.HandleFunc("/posts/{id}/image", attachImage).Methods("POST")
	r.HandleFunc("/posts/{id}/image/{filename}", updateImage).Methods("PUT")
	r.HandleFunc("/posts/{id}/image/{filename}", deleteImage).Methods("DELETE")
	s.Handler = handlers.LoggingHandler(os.Stdout, r)

	log.Printf("devserver listening at %s", *addr)
	log.Panicf("microblog devserver: %v", s.ListenAndServe())
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	svc, err := session.NewSession()
	if err != nil {
		svcError(w, err)
		return
	}

	setJSON(w, microblog.NewPostResponse(microblog.ListPosts(svc)))
}

func getPost(w http.ResponseWriter, r *http.Request) {
	svc, err := session.NewSession()
	if err != nil {
		svcError(w, err)
		return
	}

	postID := mux.Vars(r)["id"]

	setJSON(w, microblog.NewPostResponse(microblog.GetPost(svc, postID)))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var p *microblog.Post

	svc, err := session.NewSession()
	if err != nil {
		svcError(w, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		badRequest(w, fmt.Errorf("decoding post: %w", err))
		return
	}

	setJSON(w, microblog.NewPostResponse(microblog.CreatePost(svc, p.Text)))
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	var p *microblog.Post

	svc, err := session.NewSession()
	if err != nil {
		svcError(w, err)
		return
	}

	postID := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		badRequest(w, fmt.Errorf("decoding post: %w", err))
		return
	}

	setJSON(w, microblog.NewPostResponse(microblog.UpdatePost(svc, postID, p.Text)))
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	svc, err := session.NewSession()
	if err != nil {
		svcError(w, err)
		return
	}

	postID := mux.Vars(r)["id"]

	setJSON(w, microblog.NewPostResponse(nil, microblog.DeletePost(svc, postID)))
}

func attachImage(w http.ResponseWriter, r *http.Request) {
	svc, err := session.NewSession()
	if err != nil {
		svcError(w, err)
		return
	}

	_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		badRequest(w, err)
	}
	boundry := params["boundry"]

	postID := mux.Vars(r)["id"]

	image, err := microblog.UnmarshalImage(r.Body, boundry)
	if err != nil {
		badRequest(w, err)
		return
	}

	setJSON(w, microblog.NewPostResponse(microblog.AttachImage(svc, postID, image)))
}

func updateImage(w http.ResponseWriter, r *http.Request) {
	var i *microblog.Image

	svc, err := session.NewSession()
	if err != nil {
		svcError(w, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(i); err != nil {
		badRequest(w, fmt.Errorf("decoding image: %w", err))
		return
	}

	postID := mux.Vars(r)["id"]
	filename := mux.Vars(r)["filename"]

	setJSON(w, microblog.NewPostResponse(microblog.UpdateImage(svc, postID, filename, i.Caption, i.Alt)))
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
	svc, err := session.NewSession()
	if err != nil {
		svcError(w, err)
		return
	}

	postID := mux.Vars(r)["id"]
	filename := mux.Vars(r)["filename"]

	setJSON(w, microblog.NewPostResponse(microblog.DeleteImage(svc, postID, filename)))
}

func badRequest(w http.ResponseWriter, err error) {
	log.Printf("bad request: %s", err)
	http.Error(w, "bad request", http.StatusBadRequest)
}

func svcError(w http.ResponseWriter, err error) {
	log.Printf("failed to get aws session: %s", err)
	http.Error(w, "bad request", http.StatusInternalServerError)
}

func setJSON(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("encoding json: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
