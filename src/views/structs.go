package views

type Response struct {
  Code int
  Body interface{}
}

type PostRequest struct {
  Name string
  Todo string
}
