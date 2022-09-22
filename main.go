package main

import (
	"crawler/xsky"
	"log"
)

func main()  {
	token, err := xsky.GetToken()
	if err != nil {
		log.Printf("GetToken err: %v", err)
		return
	}

	jobList, err := xsky.GetJobList(token)
	if err != nil {
		log.Printf("GetJobList err: %v", err)
		return
	}

	err = xsky.SaveJson(jobList)
	if err != nil {
		log.Printf("SaveJson err: %v", err)
		return
	}
}