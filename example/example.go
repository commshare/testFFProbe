package main

import (
	"encoding/json"
	"log"
	"time"
	"../ffprobe"
)

func main() {
	path := "http://live10034.bs2dl.yy.com/crs_612dc769e21e411383ed30136aa52363.ts?s=178976&e=275795&r=800x600&tc=0&ts=1503298493"
	log.Printf("path %s \n",path)
	data, err := ffprobe.GetProbeData(path,2*time.Second ) /*要500ms么500*time.Millisecond*/
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}

	buf, err := json.MarshalIndent(data, "", "  ")
	log.Println("--------probe data---------\n")
	log.Print(string(buf))

	buf, err = json.MarshalIndent(data.GetFirstVideoStream(), "", "  ")
	log.Println("--------------first video stream-----------")
	log.Print(string(buf))
	log.Println("----------------data format info--------------------")
	log.Printf("\nDuration: %v\n", data.Format.Duration())
	log.Printf("\nStartTime: %v\n", data.Format.StartTime())
	log.Printf("---video w %v and h %v\n",data.GetFirstVideoStream().Width,data.GetFirstVideoStream().Height)

	//start := time.Now()
	//for i := 0; i < 100; i++ {
	//	_, err = ffprobe.GetProbeData(path)
	//	if err != nil {
	//		log.Panicf("Error getting data: %v", err)
	//	}
	//}
	//log.Printf("100 times get time: %v", time.Now().Sub(start))
}
