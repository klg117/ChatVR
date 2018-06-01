package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)

	printError(err, "Failed to create client: %v")

	filename := "5b0fad594ddd94.30760614.mp3"

	data, err := ioutil.ReadFile(filename)

	printError(err, "Failed to read file: %v")

	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	printError(err, "Failed to recognize: %v")

	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
}

func printError(err error, errMsg string) {
	if err != nil {
		log.Fatalf(errMsg, err)
	}
}
