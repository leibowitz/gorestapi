package main

import (
  "encoding/json"

  "github.com/go-martini/martini"
  "github.com/leibowitz/halgo"
)

type MyItem struct {
  halgo.Links
  Reference string `json:"reference"`
}

type MyResource struct {
  halgo.Links
  halgo.Embedded
  Name string `json:"name"`
}

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    res := MyResource{
      Links: halgo.Links{}.
        Self("/orders").
        Next("/orders?page=2").
        Link("ea:find", "/orders{?id}").
        Add("ea:admin", halgo.Link{Href: "/admins/2", Title: "Fred"}, halgo.Link{Href: "/admins/5", Title: "Kate"}),
      Embedded: halgo.Embedded{}.Add("ea:order", MyItem{Reference: "ABC"}),
      Name: "James",
    }
    bytes, err := json.Marshal(res)
    if err != nil {
      return err.Error()
    }
    return string(bytes)
  })
  m.RunOnAddr(":8080")
}
