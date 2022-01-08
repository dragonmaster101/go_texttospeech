package main

import (
	watson_tts "github.com/dragonmaster101/go_texttospeech/watson_tts"
)

func main(){
	service := watson_tts.CreateService("WATSON-TEXT-TO-SPEECH-KEY" , "URL");
	watson_tts.Synthesize(service , 
		"Hi I am Watson brought to you by the go text to speech library." , "watson.mp3");
	
}