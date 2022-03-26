package json

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}
type Bike struct {
	Id       string   `json:"id"`
	Location Location `json:"location"`
	Time     string   `json:"time"`
}

// The same json tags will be used to encode data into JSON
type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func marshal() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))
}

func marshal2() {
	location := &Location{
		Type:        "Point",
		Coordinates: []float32{2.2861460, 48.826802},
	}
	pigeon := &Bike{
		Id:       "bb2cdb1l52n4oiuufrig",
		Time:     "2018-04-04T14:40:05+02:00",
		Location: *location,
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))
}

func unmarshal() {
	birdJson := `{
					"birds":{
						"pigeon":"likes to perch on rocks",
						"eagle":"bird of prey",
						"eagle2":"2.3"
						
					},
					"animals":"none"
				}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{})

	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, ":", value.(string))
	}
	//pigeon likes to perch on rocks
	//eagle bird of prey
}

func unmarshal2(_json string) {

	var bikeJson string
	if len(_json) == 0 {
		bikeJson = _json
	} else {
		bikeJson = `{
			"id":"bb2cdb1l52n4oiuufrig",
			"location":
				{"type": "Point","coordinates": [2.2861460, 48.8268020]
				},
			"time": "2018-04-04T14:40:05+02:00"}`
	}
	//bikeJson := `{"id":"bb2cdb1l52n4oiuufrig","location":{"type":"Point","coordinates":[2.286146,48.8268]},"time":"2018-04-04T14:40:05+02:00"}`

	var result map[string]interface{}
	json.Unmarshal([]byte(bikeJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	/*
		birds := result["location"].(map[string]interface{})

		for key, value := range birds {
			// Each value is an interface{} type, that is type asserted as a string
			fmt.Println(key, ":", value.(string))
		}
	*/
	fmt.Println("unmarshal 2")
	//fmt.Println(result["location"].coordinates) //map[string]) //["coordinates"][1])
	fmt.Println(result)
	//pigeon likes to perch on rocks
	//eagle bird of prey
}

func unmarshalBike() {
	bikeJson := `{"id":"bb2cdb1l52n4oiuufrig","location": {"type": "Point","coordinates": [2.2861460, 48.8268020]},"time": "2018-04-04T14:40:05+02:00",}`
	var bike Bike

	//json.Unmarshal([]byte(bikeJson), &bike )struct{ *bike }{&bike})
	json.Unmarshal([]byte(bikeJson), struct{ *Bike }{&bike})

	//fmt.Printf("%+v\n%+v\n", post, location)
	fmt.Printf("%+v\n", bike)
	fmt.Println(bike.Location.Coordinates)
	bike.Location.Coordinates = []float32{1.2, 3.2}
	fmt.Println(bike.Location, bikeJson)

}

func launch() {

	marshal()
	unmarshal()
	unmarshal2(`{
		"id":"bb2cdb1l52n4oiuufrig",
		"location":
			{"type": "Point","coordinates": [2.2861460, 48.8268020]
			},
		"time": "2018-04-04T14:40:05+02:00"}`)

	marshal2()
}
