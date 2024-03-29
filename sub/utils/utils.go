package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	genome "msk-sub/genome"
	msk_protobuf "msk-sub/protobuf"
	"os"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Get argument flags
func GetFlags() (string, string, string, string) {
	masterJSONFilePtr := flag.String("masterJSONFile", "sub.json", "MasterJSON filename")
	channelPtr := flag.String("channel", "channels.*", "Channel to subscribe to")
	allChannelsPtr := flag.String("allChannels", "channels.*", "All channels")
	natsServerPtr := flag.String("natsServer", "nats://localhost:4222", "NATS server connection URL")

	flag.Parse()

	masterJSONFile := *masterJSONFilePtr
	channel := *channelPtr
	allChannels := *allChannelsPtr
	natsServer := *natsServerPtr

	return masterJSONFile, channel, allChannels, natsServer
}

// Check that the master JSON file name exists
// Deletes contents if it exits and creates new file if not
func CheckFile(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		err := os.Truncate(filename, 0)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// Deserialize the data passed from publisher and print it to terminal
func DeserializeAndPrintData(message *nats.Msg) *msk_protobuf.Result {
	newStruct := &msk_protobuf.Result{}
	err := proto.Unmarshal(message.Data, newStruct)
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	marshalOpts := protojson.MarshalOptions{
		Indent:          "    ",
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
	data, err = marshalOpts.Marshal(newStruct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	return newStruct
}

// Read the master JSON file into struct
func ReadMasterJSON(masterJSONFile string) []msk_protobuf.Result {
	// Read master JSON
	file, err := ioutil.ReadFile(masterJSONFile)
	if err != nil {
		log.Fatal(err)
	}
	data := []msk_protobuf.Result{}
	json.Unmarshal(file, &data)

	return data
}

// Add sample if it is new, else update the existing one
func AddOrUpdateSample(idMap map[string]int, dmpID string, JSONData []msk_protobuf.Result, newStruct *msk_protobuf.Result) ([]byte, map[string]int) {
	_, sampleExistsInFile := idMap[dmpID]
	
	// Check whether another version of the sample is already in master JSON file
	if sampleExistsInFile {
		fmt.Println("Another version of this sample was already in JSON file and was rewritten")
		indexOfExistingSample := idMap[dmpID]
		JSONData = append(JSONData[:indexOfExistingSample], JSONData[indexOfExistingSample+1:]...)
	} else {
		fmt.Println("Sample added to JSON file")
	}
	JSONData = append(JSONData, *newStruct)
	idMap[dmpID] = len(JSONData) - 1

	marshalOpts := protojson.MarshalOptions{
		Indent:          "    ",
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
	data, err := marshalOpts.Marshal(newStruct)
	if err != nil {
		log.Fatal(err)
	}

	return data, idMap
}

// Write the new sample data into the master JSON file
func WriteToMasterJSON(data []byte, masterJSONFile string) {
	err := ioutil.WriteFile(masterJSONFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToGenomeJSON(data []genome.Genome, genomeFile string) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(genomeFile, dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
