package data

type MyApplicationRegistryResponse struct {
	Status bool        `json:"status"`
	Result interface{} `json:"result"`
}

type Test2Request struct {
	InputText string `json:"InputText"`
}

type Test2Response struct {
	Response []Pair `json:"response"`
}

type Pair struct {
	Word  string `json: "Word"`
	Count int    `json: "Count"`
}

const MangtasTest1ServiceUrl = "http://localhost:9090"
const MangtasTest1WordCountEndpoint = "/mangtas/test1"
