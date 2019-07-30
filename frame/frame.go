package frame

import("fmt"
       "io/ioutil"
       "strings"
      // "./frame"
       )

// structure of the object Frame
type Frame struct{
  Store []string
}

// Head function returns first five rows of the frame
// Takes Nil, Returns string
func (f *Frame) Head() string{

  for index := 1;  index <= 5; index++ {
      fmt.Println("[",index-1,"]",f.Store[index])
  }

return f.Store[0]
}

// Row returns the row for the given index
// Takes integer, returns string
func (f *Frame) Row(number int) string{

return f.Store[1+number]
}

// Rows returns rows from given starting index to ending index
// Takes integers, returns string
func (f *Frame) Rows(start int,end int) string{
  temp := ""

  for index:=start; index< end; index++{
    temp = temp + f.Row(index) + "\n"
  }
return temp
}

//Height returns number of rows
//Takes nil, returns integer
func (f *Frame) Height() int{

return len(f.Store)
}

//It returns array containing data
func (f *Frame) Data() []string{

  for index := 1;  index < len(f.Store)-1; index++ {
      fmt.Println("[",index-1,"]",f.Store[index])
  }

return f.Store[1:]
}


//Init method takes file_name as argument, reads the file and loads the data in the object
//Returns Nil
func (f *Frame) Init(file_name string){

  data, err := ioutil.ReadFile(file_name)
      if err != nil {
          fmt.Println("File reading error", err)
          return
      }

      f.Store = strings.Split(string(data),"\n")
}
