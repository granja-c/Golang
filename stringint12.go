package main
import "fmt"

func main() {
  var og, w, res string
	fmt.Print("String: ")
  fmt.Scanln(&og)
  og += " "
  w = ""
  res = ""
  
  for i := 0; i < len(og); i++{
    if i != 0 && og[i] != ' ' {
      w += string(og[i])
      fmt.print(og )
      // w += og[i-1:i]
      // w += string([]rune(og)[i])
      // fmt.Print(string([]rune(og)[i]))
    } else {
      res = w + " " + res
      w = ""
    }
  }
  fmt.Printf(res)
}
