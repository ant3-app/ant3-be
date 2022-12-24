package saveimage

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func SaveImage(imgByte []byte) (image.Image, error) ***REMOVED***

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil ***REMOVED***
			log.Fatalln(err)
			return nil, err
	***REMOVED***
	

	out, _ := os.Create("./img.jpeg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil ***REMOVED***
			log.Println(err)
			return nil, err
	***REMOVED***
	return img, nil
***REMOVED***


func imgByteToIoReader(imgByte []byte) (*bytes.Reader) ***REMOVED***
	return bytes.NewReader(imgByte)
***REMOVED***