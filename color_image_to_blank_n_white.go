package main
 
import (
    "image"
    "image/color"
  _ "image/jpeg"
    "image/png"
    "math"
    "log"
    "os"
	"time"
)
 
func main() {
//input file which you want to convert into blank_and_white or gray scale
    filename := "input1.png"  //you can use png
    //filename := "input2.jpg" //you can use jpg
    infile, err := os.Open(filename)
     
    if err != nil { 
        log.Printf("failed opening %s: %s", filename, err)
        panic(err.Error())
    }
    defer infile.Close()
 
    imgSrc, _, err := image.Decode(infile)
    if err != nil {
        panic(err.Error())
    }
 
    // Create a new grayscale image
    bounds := imgSrc.Bounds()
    w, h := bounds.Max.X, bounds.Max.Y
    grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
    for x := 0; x < w; x++ {
        for y := 0; y < h; y++ {
            imageColor := imgSrc.At(x, y)              
            rr, gg, bb, _ := imageColor.RGBA()
            r := math.Pow(float64(rr), 2.2)
            g := math.Pow(float64(gg), 2.2)
            b := math.Pow(float64(bb), 2.2)
            m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
            Y := uint16(m + 0.5)
            grayColor := color.Gray{uint8(Y >> 8)}
            grayScale.Set(x, y, grayColor)
        }
    }
 
    // convert black and white or grayscale image to the new file and append current time and sec for uniqueness
    newFileName := "blank_and_white_"+time.Now().Format("20060102150405")+".png"
    newfile, err := os.Create(newFileName)
    if err != nil {
        log.Printf("failed creating %s: %s", newfile, err)
        panic(err.Error())
    }
    defer newfile.Close()
    png.Encode(newfile,grayScale)
}