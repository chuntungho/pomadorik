package main

const APP_NAME = "Pomadorik"

const APP_WIDTH = 250
const APP_HEIGHT = 250
const SOUND_FILE = "timer.mp3"

// pause name: seconds
var DEFAULT_TIMERS = map[string]int{
	"TOMATO": 1500, // 1500 sec = 25 min
	"SHORT":  300,
	"LONG":   600,
}
