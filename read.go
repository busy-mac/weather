package main

import (
    "fmt"
    "regexp"
     "io/ioutil"
    "net/http"
    "os"
    "bufio"
)


func main() {
   response, err := http.Get("http://www.sattmet.tmd.go.th/satellite_Data_Image/asia/20180710/")
    
   if err != nil {
        fmt.Println(err)
    }
    
    defer response.Body.Close()
   var bodystring string ; 
    if response.StatusCode == http.StatusOK {
    bodyBytes, err2 := ioutil.ReadAll(response.Body)
     if err2 != nil {
        fmt.Println(err2)
    }
    bodystring = string(bodyBytes)
}

//println(bodystring);
    
  
   
   
   
   var myRegex = regexp.MustCompile(`<a[^>]+\bhref=["']([^"']+)\.(jpg|JPG)["']`)
   var imgTags = myRegex.FindAllStringSubmatch(bodystring, -1)
   out := make([]string, len(imgTags))
  contents := make([]string, len(imgTags))
   for  i := range out {
    //fmt.Println(imgTags[i][1] +".jpg")
    contents[i] = fmt.Sprintf(imgTags[i][1] +".JPG" )
   }
   
   
   
   var filename_new = "images_new.txt"
   //var filename_old = "images_old.txt"
   if _, err := os.Stat(filename_new); err != nil {
  os.Remove(filename_new)
       }
   
     f, err:= os.Create(filename_new)
if err != nil {
        fmt.Println(err)
         }
    
  defer f.Close()

  

   w := bufio.NewWriter(f)
  for i := range out {
    fmt.Fprintln(w, contents[i])
  }
   
  w.Flush() 
    
   
  

	
	
	
	
    }//main