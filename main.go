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
	UserId int    `json:"user_id"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getPosts(c *gin.Context) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	defer func() { _ = reader.Close() }()

	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func createPost(c *gin.Context) {
	var form PostForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"userId": form.UserId,
		"title":  form.Title,
		"body":   form.Body,
	})

	responseBody := bytes.NewBuffer(postBody)
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func() { _ = response.Body.Close() }()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	view := make(map[string]interface{})

	err = json.Unmarshal(body, &view)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, view)
}

func getOnePost(c *gin.Context) {
	id := c.Param("id")
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	defer func() { _ = reader.Close() }()

	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func getComments(c *gin.Context) {
	postId := c.Query("postId")
	response, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=" + postId)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	defer func() { _ = reader.Close() }()

	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func updatePost(c *gin.Context) {
	var form PostForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	putBody, _ := json.Marshal(map[string]interface{}{
		"userId": form.UserId,
		"id": form.Id,
		"title":  form.Title,
		"body":   form.Body,
	})

	responseBody := bytes.NewBuffer(putBody)
	id := c.Param("id")
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/posts/" + id, responseBody)
	if err != nil {
		log.Println(err)
		return
	}

	// set the request header Content-Type for json
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	reader := resp.Body
	defer func() { _ = reader.Close() }()

	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

//func updateFieldPost(c *gin.Context) {
//	var form PostForm
//	if err := c.ShouldBindJSON(&form); err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	fmt.Println(form)
//
//	var update map[string]interface{}
//
//	if err := mapstructure.Decode(form, &update); err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	patchBody, _ := json.Marshal(update)
//	fmt.Println(update)
//	fmt.Println(patchBody)
//
//	responseBody := bytes.NewBuffer(patchBody)
//	id := c.Param("id")
//	req, err := http.NewRequest(http.MethodPatch, "https://jsonplaceholder.typicode.com/posts/" + id, responseBody)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	client := &http.Client{}
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	reader := resp.Body
//	defer func() { _ = reader.Close() }()
//
//	contentLength := resp.ContentLength
//	contentType := resp.Header.Get("Content-Type")
//	extraHeaders := map[string]string{}
//	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
//}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/" + id, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post was deleted successfully",
	})
}


func main() {
	r := gin.Default()
	r.GET("/post/list", getPosts)
	r.POST("/post/create", createPost)
	r.GET("/post/:id", getOnePost)
	r.GET("/comments", getComments)
	r.PUT("/post/:id", updatePost)
	//r.PATCH("/post/:id", updateFieldPost)
	r.DELETE("/post/:id", deletePost)

	err := r.Run(":80")
	if err != nil {
		return
	}
}
