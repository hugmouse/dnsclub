package dnsclub

import (
    "github.com/hugmouse/doRequest"
    "time"
)

type URL int

const (
	Latest URL = iota
    Digest
    Review
    Blog
    Discussions
)

var ClubURL = map[URL]string{
	Latest: "https://club.dns-shop.ru/club-exp-api/v1/on-air/get-latest/",
    Digest: "https://club.dns-shop.ru/club-exp-api/v1/blog/getList/digest/",
    Review: "https://club.dns-shop.ru/club-exp-api/v1/blog/getList/review/",
    Blog: "https://club.dns-shop.ru/club-exp-api/v1/blog/getList/blog/",
    Discussions: "https://club.dns-shop.ru/club-exp-api/v1/communicator/getList/discussions/",
}

type LatestInfo struct {
	Data []struct {
		GUID      string `json:"guid"`
		TypeID    int    `json:"typeId"`
		Title     string `json:"title"`
		PageID    int    `json:"pageId"`
		Content   string `json:"content"`
		ContentID int    `json:"contentId"`
		User      struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			AvatarURL string `json:"avatarUrl"`
		} `json:"user"`
		Date time.Time `json:"date"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
}

type CheckAuth struct {
	Data struct {
		IsRedactor bool   `json:"isRedactor"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		ID         int    `json:"id"`
		Name       string `json:"name"`
		AvatarURL  string `json:"avatarUrl"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
}

type BlogPosts struct {
	Data    DataOfPosts `json:"data"`
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
}

type AuthorOfPost struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

type ImageInPost struct {
	DesktopURL string `json:"desktopUrl"`
	MobileURL  string `json:"mobileUrl"`
}

type ItemsOfPosts struct {
	ID             int           `json:"id"`
	Title          string        `json:"title"`
	RootRubricType string        `json:"rootRubricType"`
	Author         AuthorOfPost  `json:"author"`
	CreateDate     time.Time     `json:"createDate"`
	Rating         int           `json:"rating"`
	ViewCount      int           `json:"viewCount"`
	CommentCount   int           `json:"commentCount"`
	Teaser         string        `json:"teaser"`
	Image          ImageInPost   `json:"image"`
	IsIndexable    bool          `json:"isIndexable"`
}

type DataOfPosts struct {
	Items      []ItemsOfPosts `json:"items"`
	TotalCount int            `json:"totalCount"`
}

type CommunicatorPosts struct {
	Data    CommunicatorData `json:"data"`
	Message interface{}      `json:"message"`
	Code    int              `json:"code"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

type CommunicatorProducts struct {
	ID            string      `json:"id"`
	Code          string      `json:"code"`
	Title         string      `json:"title"`
	ImageURL      string      `json:"imageUrl"`
	ShopURL       string      `json:"shopUrl"`
	Rating        interface{} `json:"rating"`
	OpinionsCount interface{} `json:"opinionsCount"`
}

type CommunicatorItems struct {
	ID            int                    `json:"id"`
	User          User                   `json:"user"`
	ImageURL      interface{}            `json:"imageUrl"`
	ImageID       interface{}            `json:"imageId"`
	PublishStamp  string                 `json:"publishStamp"`
	IsPublished   interface{}            `json:"isPublished"`
	Title         string                 `json:"title"`
	Teaser        string                 `json:"teaser"`
	Description   string                 `json:"description"`
	Views         int                    `json:"views"`
	AnswersCount  int                    `json:"answersCount"`
	CommentsCount int                    `json:"commentsCount"`
	Rating        int                    `json:"rating"`
	Status        int                    `json:"status"`
	Products      []CommunicatorProducts `json:"products"`
	Templates     []interface{}          `json:"templates"`
	Brands        []interface{}          `json:"brands"`
	Comments      []interface{}          `json:"comments"`
	Images        interface{}            `json:"images"`
	ImagesThumbs  interface{}            `json:"imagesThumbs"`
}

type CommunicatorData struct {
	Items      []CommunicatorItems `json:"items"`
	TotalCount int                 `json:"totalCount"`
}

type AwardedUsers struct {
	Data    []AwardsData `json:"data"`
	Message interface{}  `json:"message"`
	Code    int          `json:"code"`
}

type Achievement struct {
	ID       string `json:"id"`
	ImageURL string `json:"imageUrl"`
	Title    string `json:"title"`
}

type AwardsData struct {
	User        User        `json:"user"`
	LevelImgURL string      `json:"levelImgUrl"`
	ProfileURL  string      `json:"profileUrl"`
	Achievement Achievement `json:"achievement"`
	Stamp       time.Time   `json:"stamp"`
}

func GetLatestInfo() (info *LatestInfo, err error) {
    err = doRequest.RequestDecode("GET", ClubURL[Latest], "", nil, &info)
    if err != nil {
        return info, err
    }
    return info, nil
}

func GetDigestPosts() (info *BlogPosts, err error) {
    err = doRequest.RequestDecode("GET", ClubURL[Digest], "", nil, &info)
    if err != nil {
        return info, err
    }
    return info, nil
}

func GetBlogPosts() (info *BlogPosts, err error) {
    err = doRequest.RequestDecode("GET", ClubURL[Blog], "", nil, &info)
    if err != nil {
        return info, err
    }
    return info, nil
}

func GetReviewPosts() (info *BlogPosts, err error) {
    err = doRequest.RequestDecode("GET", ClubURL[Review], "", nil, &info)
    if err != nil {
        return info, err
    }
    return info, nil
}