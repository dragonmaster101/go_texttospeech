package watson_tts

import (
	"bytes"
	"os"
	"github.com/IBM/go-sdk-core/v5/core"
	api "github.com/watson-developer-cloud/go-sdk/v2/texttospeechv1"
)

func CreateService(API_KEY , API_URL string) *api.TextToSpeechV1{
	authenticator := &core.IamAuthenticator{
		ApiKey:    API_KEY,
	}

	options := &api.TextToSpeechV1Options{
		Authenticator: authenticator,
	}

	textToSpeech, textToSpeechErr := api.NewTextToSpeechV1(options)

	if textToSpeechErr != nil {
		panic(textToSpeechErr)
	}

	textToSpeech.SetServiceURL(API_URL)

	return textToSpeech;
}

func Synthesize(service *api.TextToSpeechV1 ,text string , fileName string){
	result, _, responseErr := service.Synthesize(
		&api.SynthesizeOptions{
		Text:   core.StringPtr(text),
		Accept: core.StringPtr("audio/mp3"),
		Voice:  core.StringPtr("en-US_AllisonV3Voice"),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	if result != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(result)
		file, _ := os.Create(fileName)
		file.Write(buff.Bytes())
		file.Close()
	}
}