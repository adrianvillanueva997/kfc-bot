package commands

import (
	randomgeneration "adrianvillanueva997/kfcbot/src/utilities/randomGeneration"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/diamondburned/arikawa/v2/gateway"
)

type ChanThreads struct {
	Page    int `json:"page"`
	Threads []struct {
		No           int `json:"no"`
		LastModified int `json:"last_modified"`
		Replies      int `json:"replies"`
	} `json:"threads"`
}

type ChanThread struct {
	Posts []struct {
		No          int    `json:"no"`
		Now         string `json:"now"`
		Name        string `json:"name"`
		Sub         string `json:"sub,omitempty"`
		Com         string `json:"com,omitempty"`
		Filename    string `json:"filename,omitempty"`
		Ext         string `json:"ext,omitempty"`
		W           int    `json:"w,omitempty"`
		H           int    `json:"h,omitempty"`
		TnW         int    `json:"tn_w,omitempty"`
		TnH         int    `json:"tn_h,omitempty"`
		Tim         int64  `json:"tim,omitempty"`
		Time        int    `json:"time"`
		Md5         string `json:"md5,omitempty"`
		Fsize       int    `json:"fsize,omitempty"`
		Resto       int    `json:"resto"`
		Bumplimit   int    `json:"bumplimit,omitempty"`
		Imagelimit  int    `json:"imagelimit,omitempty"`
		SemanticURL string `json:"semantic_url,omitempty"`
		Replies     int    `json:"replies,omitempty"`
		Images      int    `json:"images,omitempty"`
		UniqueIps   int    `json:"unique_ips,omitempty"`
	} `json:"posts"`
}

type replyMessage struct {
	filename string
	webmURL  string
}

func get4ChanThreads() ([]ChanThreads, error) {
	resp, err := http.Get("https://a.4cdn.org/wsg/threads.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("bad request, try again later")
	}
	decoder := json.NewDecoder(resp.Body)
	var data []ChanThreads
	err = decoder.Decode(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data, nil
}

func get4ChanThreadData(threads ChanThreads) (ChanThread, error) {
	randomThread := randomgeneration.RandomInteger(len(threads.Threads)-1, 0)
	var data ChanThread

	resp, err := http.Get(fmt.Sprintf("https://a.4cdn.org/wsg/thread/%s.json", strconv.Itoa(threads.Threads[randomThread].No)))
	if err != nil {
		return data, err
	}
	if resp.StatusCode != 200 {
		return data, errors.New("bad request, try again later")
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

func getRandomWebm(thread ChanThread) replyMessage {
	randomReply := randomgeneration.RandomInteger(len(thread.Posts)-1, 0)
	if !checkWebmStatus(thread, randomReply) {
		loop := true
		for loop {
			randomReply = randomgeneration.RandomInteger(len(thread.Posts)-1, 0)
			if checkWebmStatus(thread, randomReply) {
				loop = false
			}
		}
	}
	var data replyMessage
	data.filename = thread.Posts[randomReply].Filename
	data.webmURL = fmt.Sprintf("https://i.4cdn.org/wsg/%s.webm", strconv.Itoa(int(thread.Posts[randomReply].Tim)))
	return data

}

func checkWebmStatus(thread ChanThread, numberOfReply int) bool {
	return thread.Posts[numberOfReply].Ext == ".webm"
}

func (b *Bot) Meme(*gateway.MessageCreateEvent) (string, error) {
	chanData, err := get4ChanThreads()
	if err != nil {
		return "", err
	}
	randomPage := randomgeneration.RandomInteger(len(chanData)-1, 0)
	threadData, err := get4ChanThreadData(chanData[randomPage])
	if err != nil {
		return "", err
	}
	url := getRandomWebm(threadData)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s \n %s", url.filename, url.webmURL), nil
}
