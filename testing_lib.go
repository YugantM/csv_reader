package main

import ("fmt"
        "./frame")

func main(){
      //fmt.Println("fmt used!")
      df := new(frame.Frame)
      df.Init("T.csv")

      fmt.Println(df.Rows(0,6))
}
