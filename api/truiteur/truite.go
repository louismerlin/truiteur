package truiteur

type Truite struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

func addTruite(data url.Values) {
	fmt.Printf("Post!")
}

func getTruite() Truite {
	return Truite{Id: 1, Body: "Hello, world!"}
}
