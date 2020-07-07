package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang"})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another fake message."})
	return s.Get(url)
}

func main() {

	var r Retriever
	r = &mock.Retriever{Contents: "This is a fake message."}
	fmt.Println(download(r))
	fmt.Printf("%T %v\n", r, r)
	fmt.Println()

	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute}
	//fmt.Println(download(r))
	fmt.Printf("%T %v\n", r, r)

	inspect(r)

	// type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println("TimeOut:", realRetriever.TimeOut)

	fmt.Println("Try session")
	s := &mock.Retriever{Contents: "This is a fake message."}
	fmt.Println(session(s))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(">%T, %v\n", r, r)
	fmt.Print("> Type Switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("pointer Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
