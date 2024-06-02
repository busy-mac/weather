package main
import (
   
    
    "fmt"
    "time"
    "strconv"
    "math"
   
    )
func main() {

 t1 := time.Now().UTC()
 t2 := time.Now()
fmt.Println(t1.Year())
//date = strconv.Itoa(t.Year() ) + strconv.Itoa(t.Month() ) + strconv.Itoa(t.Day())
fmt.Println(t1.String())
fmt.Println(t1.Format("2006-01-02 15:04:05"))
fmt.Println(t2.String())
fmt.Println(t2.Format("2006-01-02 15:04:05"))
 


var mth int = int(t1.Month() )

 m := fmt.Sprintf("%02d", mth)
d :=fmt.Sprintf("%02d",t1.Day())

folder_date := strconv.Itoa(t1.Year()) + m + d 
fmt.Println(folder_date)

folder_time :=fmt.Sprintf("%02d",t1.Hour()) + fmt.Sprintf("%02d",t1.Minute())
fmt.Println(folder_time)


name :="WV201807101440.JPG"
year := name[2:6]
month :=name[6:8]
day :=name[8:10]
hour :=name[10:12]
minute :=name[12:14]
val:=year+" "+ month +" "+ day + " "+ hour +" " + minute 
fmt.Println(year,month,day,hour,minute)
t, err := time.Parse("2006 01 02 15 04",val)
	fmt.Println(t, err)
	
	loc, err := time.LoadLocation("Asia/Kathmandu")
	fmt.Println(loc, err)
	
	t = t.In(loc)
	fmt.Println(t.Format("2006 01 02 15 04"))


  hs := t1.Sub(t).Hours()

 
 
    hs, mf := math.Modf(hs)
    ms := mf * 60

    ms, sf := math.Modf(ms)
    ss := sf * 60

    fmt.Println(hs, "hours", ms, "minutes", ss, "seconds")

        








}