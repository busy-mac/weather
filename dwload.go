package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "os"
    "regexp"
     "io/ioutil"
     "time"
     "math"
     "strconv"
    "net/http"
)




func readFileWithReadLine(fn string) ( error , []string) {
    fmt.Println("readFileWithReadLine")

    file, err := os.Open(fn)
    defer file.Close()

    

    // Start reading from the file with a reader.
    reader := bufio.NewReader(file)
 var lines []string;
    for {
    	
    	
        var buffer bytes.Buffer

        var l []byte
        var isPrefix bool
        for {
            l, isPrefix, err = reader.ReadLine()
            buffer.Write(l)

            // If we've reached the end of the line, stop reading.
            if !isPrefix {
                break
            }

            // If we're just at the EOF, break
            if err != nil {
                break
            }
        }

        if err == io.EOF {
            break
        }

        line := buffer.String()
   lines = append(lines, line)
       //fmt.Printf(" > Read %d characters\n", len(line))
//println(line)
        
        
        // Process the line here.
        //fmt.Println(" > > " + limitLength(line, 50))
    }

    if err != io.EOF {
        fmt.Printf(" > Failed!: %v\n", err)
    }
  /*  
    
 for i := range lines{
 	 println(lines[i])
 	 
 }
    
 */   
    
    return err,lines
}


func read_server(url string) ( []string) {
	
	
 client := &http.Client{}

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                fmt.Println(err)
        }

        req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:60.0) Gecko/20100101 Firefox/60.0")

        resp, err := client.Do(req)
        if err != nil {
                fmt.Println(err)
        }

        defer resp.Body.Close()
	
	
	
	
	
	
	/*
	response, err := http.Get(url)
    
   if err != nil {
        fmt.Println(err)
    }
    
    defer response.Body.Close()
   */ 
    
    
   var bodystring string ; 
    if resp.StatusCode == http.StatusOK {
    bodyBytes, err2 := ioutil.ReadAll(resp.Body)
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

return contents

}


 func download_image(url string ,download_path string,image_name string)  {


 	 client := &http.Client{}

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                fmt.Println(err)
        }

        req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:60.0) Gecko/20100101 Firefox/60.0")

        resp, err := client.Do(req)
        if err != nil {
                fmt.Println(err)
        }

        defer resp.Body.Close()
	
 	 
 	 
 	 
 	 
 	 
 	 
 	 /*
 	 
 	  
    // don't worry about errors
    response, e := http.Get(url)
    if e != nil {
       fmt.Println(e)
    }

    defer response.Body.Close()
    
  */  
    
    
  save_path:= download_path+image_name ;
    //open a file for writing
    file, err := os.Create(save_path )
    if err != nil {
       fmt.Println(err)
    }
    // Use io.Copy to just dump the response body to the file. This supports huge files
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    file.Close()
    fmt.Println(" Download Success!")
 	 
 	 
 	 
 	 

}


func time_diff(previous_image string ,new_image string)(int){
	
	
name :=previous_image
year := name[2:6]
month :=name[6:8]
day :=name[8:10]
hour :=name[10:12]
minute :=name[12:14]


  val:=year+" "+ month +" "+ day + " "+ hour +" " + minute 
   //fmt.Println(year,month,day,hour,minute)

t1, _ := time.Parse("2006 01 02 15 04",val)
	//fmt.Println(t, err)
	
name2 :=new_image
year2 := name2[2:6]
month2 :=name2[6:8]
day2 :=name2[8:10]
hour2 :=name2[10:12]
minute2 :=name2[12:14]


  val2:=year2+" "+ month2 +" "+ day2 + " "+ hour2 +" " + minute2 
   //fmt.Println(year2,month2,day2,hour2,minute2)

t2, _ := time.Parse("2006 01 02 15 04",val2)


  mns := t2.Sub(t1).Minutes()

 absmns:= math.Abs(mns)
  
 return  int(absmns)
   

	
	
}






func main() {
	
	
	_, err := http.Get("https://www.google.com")
    
   if err != nil {
        fmt.Println(err)
        fmt.Println("No Inernet connection ! Try Again!")
        os.Exit(1);
        
    }
	filename_old := "images_old.txt"
	_,contents_old := readFileWithReadLine(filename_old)
	
	
	
	oldFile_last :=contents_old[len(contents_old)-2]
	
	
	//read new images list from server
	
	t1 := time.Now().UTC()
	var mth int = int(t1.Month() )

 m := fmt.Sprintf("%02d", mth)
d :=fmt.Sprintf("%02d",t1.Day())

folder_date := strconv.Itoa(t1.Year()) + m + d 
		
	
	url:="http://www.sattmet.tmd.go.th/satellite_Data_Image/asia/IR/" +folder_date + "/"
	contents_new := read_server(url)
	
	newFile_last :=contents_new[len(contents_new)-1]
	
	var download_ready_contents []string;
	var previous_image string;
	
	var diff int = 0;
	var adiff int =0;
	dnload_root_folder :="images/"
	number_of_downloads:=20
	
	dnload_folder_path:= dnload_root_folder +folder_date
	
	//download folder
	if _, err := os.Stat(dnload_folder_path); os.IsNotExist(err) {
   er1:= os.Mkdir(dnload_folder_path, 0700)
    
    
    if er1 != nil {
        fmt.Println(er1)
        fmt.Println("Cannot create Directory under images!")
        os.Exit(1);
     }
}
	
dnload_folder :=dnload_folder_path + "/"



	if newFile_last != oldFile_last{
		
		
		
		
		
		
		
		if _, err := os.Stat(filename_old); err != nil {
  os.Remove(filename_old)
  		}
      f, err:= os.Create(filename_old)
if err != nil {
        fmt.Println(err)
    }
    
  defer f.Close()

  

   w := bufio.NewWriter(f)
  for i := range contents_new {
    fmt.Fprintln(w, contents_new[i])
           }
   
  w.Flush() 
  
  
  
	
	download_ready_contents =	contents_new[len(contents_new)- number_of_downloads:]
	
	
	
	
	
	
	for i := range download_ready_contents{
		
		
	if 	previous_image != ""{
		
	diff = time_diff(previous_image ,download_ready_contents[i])
	//fmt.Println("diff:",diff)	
	adiff = diff + adiff;
	fmt.Println("added diff:",adiff)
	
		
	}
	if 	previous_image == ""{
		fmt.Println("old file last :",oldFile_last)
		fmt.Println("ew file first :",download_ready_contents[i])
	diff = time_diff(oldFile_last ,download_ready_contents[i])
	//fmt.Println("diff:",diff)	
	adiff = diff + adiff;
	
	fmt.Println("added diff from last image:",adiff)
	
		
	}
	
	
	
	
	
	
	
	
	
	
	filename := dnload_folder + download_ready_contents[i]
	
	
	if (adiff > 25) || (adiff == 0 ){
	
	         if _, err := os.Stat(filename);  os.IsNotExist(err) {
			//fmt.Println(err)
			    
			
	
			
	fmt.Println(download_ready_contents[i])
	adiff=0	
	  image_url:= url + download_ready_contents[i]
		download_image(image_url,dnload_folder,download_ready_contents[i])
			
			
			
		
	      }//if not exists
		}//if diff
		
	
	previous_image =download_ready_contents[i]	
		
	
	
	
	
		
	         }//for
	
	
//always download last image
fn:=dnload_folder+previous_image
img_url:=url+previous_image
	  if _, err := os.Stat(fn);  os.IsNotExist(err) {
	  	  fmt.Println("last image",previous_image)
	download_image(img_url,dnload_folder,previous_image)
	  }         
	         
	
	
}//if old file != new file
	
	
	
	
   
}//main






