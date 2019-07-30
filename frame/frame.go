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

func (f *Frame) Row(number int) string{

return f.Store[1+number]
}

func (f *Frame) Rows(start int,end int) string{
  temp := ""

  for index:=start; index< end; index++{
    temp = temp + f.Row(index) + "\n"
  }
return temp
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
