package main

import("fmt"
       "io/ioutil"
       "strings"
      // "./frame"
       )

type frame struct{
  head []string
  data []string
}

func (f *frame) Init(file_name string){

  data, err := ioutil.ReadFile(file_name)
      if err != nil {
          fmt.Println("File reading error", err)
          return
      }

    rows:= strings.Split(string(data),"\n")
    f.head = strings.Split(rows[0],",")
    f.data = rows[1:]
}

func main(){
fmt.Println("fmt used!")
df := new(frame)
df.Init("text.txt")

fmt.Println("Head:",df.head[0],"\nData:",df.data[1])

}
