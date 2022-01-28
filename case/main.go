package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/kylma/kad"
)

func main() {
	log.Println("Reading layout")
	l, err := os.ReadFile("../lasercut.json")
	if err != nil {

		log.Fatalf("Failed to load file\nError: %s", err.Error())
		return

	}

	// you can define settings and the layout in JSON
	c_pre := []byte(`{
		"layout":`)
	c_post := []byte(`,
		"case": {
			"case-type":"sandwich",
			"mount-holes-num":12,
			"mount-holes-size":2.1,
			"mount-holes-edge":5
		},
		"top-padding":5,
		"left-padding":5,
		"right-padding":5,
		"bottom-padding":5,
	}`)
	c := [][]byte{c_pre, l, c_post}
	config := bytes.Join(c, []byte(" "))
	//log.Printf("Config: \n %s", config)

	// create a new KAD instance
	cad := kad.New()

	// populate the 'cad' instance with the JSON contents
	err = json.Unmarshal(config, cad)
	if err != nil {
		log.Fatalf("Failed to parse json data into the KAD file\nError: %s", err.Error())
	}

	// and you can define settings via the KAD instance
	cad.Hash = "Lasercut"           // the name of the design
	cad.FileStore = kad.STORE_LOCAL // store the files locally
	cad.FileDirectory = "./"        // the path location where the files will be saved
	cad.FileServePath = "/"         // the url path for the 'results' (don't worry about this)

	// here are some more settings defined for this case
	cad.Case.UsbWidth = 12 // all dimension are in 'mm'
	cad.Fillet = 3         // 3mm radius on the rounded corners of the case
	cad.SwitchType = 1
	cad.StabType = 1
	cad.Result.Formats = []string{"dxf"}

	// lets draw the SVG files now
	err = cad.Draw()
	if err != nil {
		log.Fatal("Failed to Draw the KAD file\nError: ", err.Error())
	}
}
