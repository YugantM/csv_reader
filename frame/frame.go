package frame

import("fmt"
       "io/ioutil"
       "strings"
      // "./frame"
       )

type Frame struct{
  Head []string
  Data []string
}

func (f *Frame) Init(file_name string){

  data, err := ioutil.ReadFile(file_name)
      if err != nil {
          fmt.Println("File reading error", err)
          return
      }

    rows:= strings.Split(string(data),"\n")
    f.Head = strings.Split(rows[0],",")
    f.Data = rows[1:]
}
