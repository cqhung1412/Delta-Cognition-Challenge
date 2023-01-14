package aws_util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Label struct {
	Name       string   `json:"name"`
	Confidence float64  `json:"confidence"`
	Parents    []string `json:"parents"`
}

type DogJson struct {
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
}

func (maker *AwsMaker) ReadJsonFile(bucket string, filename string) (DogJson, error) {
	client := s3.NewFromConfig(*maker.Config)

	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	}

	result, err := client.GetObject(context.Background(), requestInput)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()
	body1, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString1 := fmt.Sprintf("%s", body1)

	var s3data DogJson
	decoder := json.NewDecoder(strings.NewReader(bodyString1))
	err = decoder.Decode(&s3data)
	if err != nil {
		fmt.Println("twas an error")
	}

	return s3data, err
}
