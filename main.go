package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	qr, err := qrcode.New(fmt.Sprintf("https://lkygcy.xyz/qrcode?t=%d", time.Now().Unix()))
	if err != nil {
		log.Fatalf("failed to generate Qrcode :%v", err)
	}
	opt1 := standard.WithCircleShape()
	opt2 := standard.WithFgColor(color.RGBA{250, 128, 114, 1.000})
	
	w, err := standard.New("qrcode.jpeg", opt1, opt2)
	if err != nil {
		log.Fatalf("standard.New fialed :%v", err)
	}

	if err = qr.Save(w); err != nil {
		log.Fatalf("failed to save image:%v", err)
	}
}
