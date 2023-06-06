package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"belajar-protobuf/pb"

	"github.com/golang/protobuf/proto"
)

func main() {
	datas := &pb.CovidDataList{
		Data: []*pb.CovidData{
			{
				Year:  2020,
				Cases: 1000,
			},
			{
				Year:  2021,
				Cases: 1350,
			},
			{
				Year:  2020,
				Cases: 3000,
			},
		},
	}

	dataBytes, err := proto.Marshal(datas)
	if err != nil {
		panic(err)
	}

	dataJson, err := json.Marshal(datas)
	if err != nil {
		panic(err)
	}
	fmt.Println(dataBytes)
	fmt.Println(string(dataJson))

	err = ioutil.WriteFile("covid_data.bin", dataBytes, 0o644)
	if err != nil {
		log.Fatal("error saving covid data to file: ", err)
	}

	err = ioutil.WriteFile("covid_data.json", dataJson, 0o644)
	if err != nil {
		log.Fatal("error saving covid data to file: ", err)
	}

	dataBytes, err = ioutil.ReadFile("covid_data.bin")
	if err != nil {
		log.Fatal("error reading covid data from file: ", err)
	}

	dataJson, err = ioutil.ReadFile("covid_data.json")
	if err != nil {
		log.Fatal("error reading covid data from file: ", err)
	}

	fmt.Println(hex.EncodeToString(dataJson))

	decodedCovidDataList := &pb.CovidDataList{}
	err = proto.Unmarshal(dataBytes, decodedCovidDataList)
	if err != nil {
		log.Fatal("error decoding covid data: ", err)
	}

	for _, data := range decodedCovidDataList.GetData() {
		fmt.Printf("Year: %d, Cases: %d\n", data.GetYear(), data.GetCases())
	}
}
