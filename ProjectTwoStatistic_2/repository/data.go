package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"projects/ProjectTwoStatistic_2/config"
	"projects/ProjectTwoStatistic_2/model"
	"time"
)

func SaveNewData() error {
	var regionCodes model.RegionCodeList

	body, err := os.ReadFile("./file/regionCode.json")
	if err != nil {
		fmt.Println("****")
		return err
	}

	err = json.Unmarshal(body, &regionCodes)
	if err != nil {
		fmt.Println("********")
		return err
	}

	var allVideos model.Videos
	var ids model.Ids
	for _, code := range regionCodes.List {
		videos, er := GetVideo("20", code.Code)
		if er != nil {
			fmt.Println("************")
			return fmt.Errorf("error occure when videos pulling from %v", code.Country)
		}
		for _, video := range videos.Items {
			ids.Ids = append(ids.Ids, model.Id{Id: video.Snippet.ChannelId})
			allVideos.Items = append(allVideos.Items, video)
		}
	}

	channels, err := GetChannel(ids)
	if err != nil {
		fmt.Println("****************")
		return err
	}

	var database, videoCollection, channelCollection string
	database, err = config.Reader("DATABASE")
	if err != nil {
		fmt.Println("********************")
		return err
	}
	channelCollection, err = config.Reader("CHANNELCOLLECTION")
	if err != nil {
		fmt.Println("************************")
		return err
	}
	videoCollection, err = config.Reader("VIDEOCOLLECTION")
	if err != nil {
		fmt.Println("****************************")
		return err
	}
	dbVideoClient, err := config.ConnectionDB(database, videoCollection)
	dbChannelVideo, err := config.ConnectionDB(database, channelCollection)

	videoRepo := Client{Collection: dbVideoClient}
	channelRepo := Client{Collection: dbChannelVideo}

	err = videoRepo.AddVideo(allVideos)
	if err != nil {
		return err
	}

	err = channelRepo.AddChannel(channels)
	if err != nil {
		return err
	}

	/*
		errVideo := make(chan error)
			errChannel := make(chan error)

		go func() {
				errVideo <- videoRepo.AddVideo(allVideos)
			}()
			if cErr := <-errVideo; cErr != nil {
				return cErr
			}

			go func() {
				errChannel = channelRepo.AddChannel(channels)
			}()
			if cErr := <-errChannel; cErr != nil {
				return cErr
			}
	*/
	return nil
}

func GetVideo(maxResult string, regionCode string) (model.Videos, error) {
	var videos model.Videos

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/videos", nil)
	if err != nil {
		return videos, err
	}

	q := req.URL.Query()

	key, err := config.Reader("YOUTUBEKEY")
	if err != nil {
		return videos, err
	}

	q.Add("key", key)
	q.Add("chart", "mostPopular")
	q.Add("part", "snippet,statistics,contentDetails,status")
	if regionCode != "" {
		q.Add("regionCode", regionCode)
	}
	q.Add("maxResult", maxResult)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return videos, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return videos, err
	}

	err = json.Unmarshal(body, &videos)
	if err != nil {
		return model.Videos{}, err
	}

	return videos, nil
}

func GetChannel(ids model.Ids) (model.Channels, error) {
	var channels model.Channels
	var allChannels model.Channels

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println("--")
		return channels, err
	}

	key, err := config.Reader("YOUTUBEKEY")
	if err != nil {
		fmt.Println("----")
		return channels, err
	}

	var channelIds string
	var n int
	for i, id := range ids.Ids {
		if channelIds == "" {
			channelIds = id.Id
		} else {
			channelIds = fmt.Sprintf("%v,%v", channelIds, id.Id)
		}
		n += 1
		if i == len(ids.Ids) || n == 10 {
			q := req.URL.Query()
			q.Add("key", key)
			q.Add("part", "snippet,statistics,brandingSettings,status")
			q.Add("id", channelIds)

			req.URL.RawQuery = q.Encode()

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				fmt.Println("------")
				return channels, err
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println("--------")
				return channels, err
			}

			err = json.Unmarshal(body, &channels)
			for _, c := range channels.Items {
				allChannels.Items = append(allChannels.Items, c)
			}
			n = 0
			channelIds = ""
		}
	}

	return allChannels, nil

}

func (dbC *Client) AddVideo(videos model.Videos) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	var ids model.Ids

	body, err := os.ReadFile("./file/videosId.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &ids)
	if err != nil {
		return err
	}

	for _, c := range videos.Items {
		var isIn bool = true
		for _, id := range ids.Ids {
			if c.Id == id.Id {
				isIn = false
			}
		}
		if isIn {
			_, err = dbC.Collection.InsertOne(ctx, c)
			if err != nil {
				return fmt.Errorf("video has %v id can not ınsert: %v", c.Id, err.Error())
			}
			ids.Ids = append(ids.Ids, model.Id{
				Id: c.Id,
			})
		}
	}

	body, err = json.Marshal(ids)
	if err != nil {
		return err
	}
	err = os.WriteFile("./file/videosId.json", body, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (dbC *Client) AddChannel(channels model.Channels) error {
	errChan := make(chan error)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	var ids model.Ids

	body, err := os.ReadFile("./file/channelsId.json")
	if err != nil {
		errChan <- err
		return <-errChan
	}

	err = json.Unmarshal(body, &ids)

	for _, c := range channels.Items {
		var isIn bool = true
		for _, id := range ids.Ids {
			if c.Id == id.Id {
				isIn = false
			}
		}
		if isIn {
			_, err = dbC.Collection.InsertOne(ctx, c)
			if err != nil {
				errChan <- fmt.Errorf("channel has %v id can not ınsert: %v", c.Id, err.Error())
				return <-errChan
			}
			ids.Ids = append(ids.Ids, model.Id{
				Id: c.Id,
			})
		}
	}
	body, err = json.Marshal(ids)
	if err != nil {
		return err
	}
	err = os.WriteFile("./file/channelsId.json", body, 0644)
	if err != nil {
		return err
	}

	return nil
}
