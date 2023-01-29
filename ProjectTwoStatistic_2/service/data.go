package service

import (
	"net/http"
	"projects/ProjectTwoStatistic_2/repository"
	"time"
)

func AddData(w http.ResponseWriter, r *http.Request) {
	var err error
	for true {
		err = repository.SaveNewData()

		time.Sleep(time.Hour * 3)
	}
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
	}
}
