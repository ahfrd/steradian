package helpers

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func LinkToBase64(link string) (string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding, nil
}

func Base64ToImage(b64 string, id int) string {
	stringId := strconv.Itoa(id)

	datax := strings.Split(b64, ",")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(datax[1]))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)
	pathFile := "app/upload/"
	//Encode from image format to writer
	pngFilename := pathFile + stringId + "." + formatString
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	switch formatString {
	case "png":
		err = png.Encode(f, m)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Println("Png file", pngFilename, "created")
	case "jpeg":
		err = jpeg.Encode(f, m, nil)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Println("Jpeg file", pngFilename, "created")
	case "gif":
		err = gif.Encode(f, m, nil)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Println("Jpeg file", pngFilename, "created")
	}
	return pngFilename
}

func FormatDateString(date string) string {

	// Parse the string date into a time.Time object
	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	formattedDate := parsedTime.Format("2006-01-02")
	fmt.Println("Formatted date:", formattedDate)
	return formattedDate
}
