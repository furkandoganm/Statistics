package model

type Videos struct {
	Kind  string      `json:"kind"`
	Etag  string      `json:"etag"`
	Items []VideoItem `json:"items"`
}

type VideoItem struct {
	Version        int             `json:"version"`
	Kind           string          `json:"kind"`
	Etag           string          `json:"etag"`
	Id             string          `json:"id"`
	Snippet        VSnippet        `json:"snippet"`
	Statistics     Statistics      `json:"statistics"`
	ContentDetails VContentDetails `json:"contentDetails"`
	Status         Status          `json:"status"`
}

type VSnippet struct {
	PublishedAt          string    `json:"publishedAt"`
	ChannelId            string    `json:"channelId"`
	Title                string    `json:"title"`
	Thumbnails           Thumbnail `json:"thumbnails"`
	ChannelTitle         string    `json:"channelTitle"`
	Tags                 []string  `json:"tags"`
	CategoryId           string    `json:"categoryId"`
	LiveBroadcastContent string    `json:"liveBroadcastContent"`
	DefaultLanguage      string    `json:"defaultLanguage"`
	Localized            Localized `json:"localized"`
	DefaultAudioLanguage string    `json:"defaultAudioLanguage"`
}

type Thumbnail struct {
	Keys []Key `json:"keys" json:"default"`
}

type Key struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Statistics struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

type VContentDetails struct {
	Duration           string            `json:"duration"`
	Dimension          string            `json:"dimension"`
	Definition         string            `json:"definition"`
	Caption            string            `json:"caption"`
	LicensedContent    bool              `json:"licensedContent"`
	RegionRestriction  RegionRestriction `json:"regionRestriction"`
	Projection         string            `json:"projection"`
	HasCustomThumbnail bool              `json:"hasCustomThumbnail"`
}

type RegionRestriction struct {
	Allowed []string `json:"allowed"`
	Blocked []string `json:"blocked"`
}

type Status struct {
	MadeForKids             bool `json:"madeForKids"`
	SelfDeclaredMadeForKids bool `json:"selfDeclaredMadeForKids"`
}

//Channel

type Channels struct {
	Kind  string        `json:"kind"`
	Etag  string        `json:"etag"`
	Items []ChannelItem `json:"items"`
}

type ChannelItem struct {
	Version          int               `json:"version"`
	Kind             string            `json:"kind"`
	Etag             string            `json:"etag"`
	Id               string            `json:"id"`
	Snippet          CSnippet          `json:"snippet"`
	Statistics       Statistics        `json:"statistics"`
	BrandingSettings CBrandingSettings `json:"brandingSettings"`
	Status           Status            `json:"status"`
}

type CSnippet struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	CustomUrl       string    `json:"customUrl"`
	PublishedAt     string    `json:"publishedAt"`
	DefaultLanguage string    `json:"defaultLanguage"`
	Localized       Localized `json:"localized"`
	Country         string    `json:"country"`
}

type CBrandingSettings struct {
	Channel Channel `json:"channel"`
	Watch   Watch   `json:"watch"`
}

type Channel struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Keywords        string `json:"keywords"`
	DefaultLanguage string `json:"defaultLanguage"`
	Country         string `json:"country"`
}

type Watch struct {
	TextColor          string `json:"textColor"`
	BackGroundColor    string `json:"backGroundColor"`
	FeaturedPlaylistId string `json:"featuredPlaylistId"`
}
