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
