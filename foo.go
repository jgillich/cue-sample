package main

import (
	"log"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func main() {

	overlay := map[string]load.Source{}

	config := &load.Config{
		Overlay: overlay,
	}

	config.Tags = []string{"foo={\"foo\":1}"}
	instances := load.Instances([]string{"./bar", "./foo"}, config)
	for _, b := range instances {
		if b.Err != nil {
			log.Fatal(b.Err)
		}
	}

	cueContext := cuecontext.New()

	defs, err := cueContext.BuildInstances(instances)
	if err != nil {
		log.Fatal(err)
	}

	json, err := defs[0].MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("json: %v", string(json))

}
