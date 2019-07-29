package frame

import("fmt"
       "io/ioutil"
       "strings"
      // "./frame"
       )

type Frame struct{
  Store []string
}

func (f *Frame) Head() string{

  for index := 1;  index <= 5; index++ {
      fmt.Println("[",index-1,"]",f.Store[index])
  }

return f.Store[0]
}

func (f *Frame) Height() int{

return len(f.Store)
}

func (f *Frame) Data() []string{

  for index := 1;  index < len(f.Store)-1; index++ {
      fmt.Println("[",index-1,"]",f.Store[index])
  }

return f.Store[1:]
}


func (f *Frame) Init(file_name string){

  data, err := ioutil.ReadFile(file_name)
      if err != nil {
          fmt.Println("File reading error", err)
          return
      }

      f.Store = strings.Split(string(data),"\n")

}
