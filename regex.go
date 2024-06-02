package main

import (
     "fmt"
     "regexp"
)

func main() {
	
   var myString = `<img src='img1single.JPG'><img src="img2double.jpg">`
   
   
   
   var myRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)\.(jpg|JPG)["']`)
   var imgTags = myRegex.FindAllStringSubmatch(myString, -1)
   out := make([]string, len(imgTags))
  for i := range out {
    fmt.Println(imgTags[i][1] +"." imgTags[i][2])
   }
   
   
   
   
   
 }
 