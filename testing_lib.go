package main

import ("fmt"
        "./frame")

func main(){
      fmt.Println("fmt used!")
      df := new(frame.Frame)
      df.Init("T.csv")

      for index := 0;  index < len(df.Data)-1; index++ {
          fmt.Println("Head:",df.Data[index])
      }


}
