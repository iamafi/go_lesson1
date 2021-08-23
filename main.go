package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type PostForm struct {
	userId   int
	id 		 int
	title    string
	body 	 string
}

func getPosts(c *gin.Context) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	defer reader.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func createPost(c *gin.Context) {
	var form PostForm
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	postBody, _ := json.Marshal(map[string]string{
		"userId": string(form.userId),
		"id":     string(form.id),
		"title":  form.title,
		"body":  form.body,
	})

	responseBody := bytes.NewBuffer(postBody)
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func main() {
	r := gin.Default()
	r.GET("/post/list", getPosts)
	r.POST("post/create", createPost)

	r.Run(":80")
}
