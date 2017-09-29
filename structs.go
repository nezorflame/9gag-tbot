package main

// LoginResponse describes login API response
type LoginResponse struct {
	Meta MetaData  `json:"meta"`
	Data LoginData `json:"data"`
}

// PostListResponse describes post list API response
type PostListResponse struct {
	Meta MetaData     `json:"meta"`
	Data PostListData `json:"data"`
}

// MetaData describes meta data from every API response
type MetaData struct {
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
	SID       string `json:"sid"`
}

// LoginData describes login data from LoginResponse
type LoginData struct {
	DummyField        string   `json:"dummyField"` // always equals to "dummyValue"
	MinVersion        string   `json:"minVersion"`
	UserToken         string   `json:"userToken"`
	TokenExpiry       int64    `json:"tokenExpiry"`
	SecondsTillExpiry int64    `json:"secondsTillExpiry"`
	AlgoliaToken      string   `json:"algoliaToken"`
	User              UserData `json:"user"`
	// CommentAuth struct
	// Noti struct
}

// UserData describes user from LoginData
type UserData struct {
	UserID            string `json:"userID"`
	AccountID         string `json:"accountId"`
	ProfileURL        string `json:"profileUrl"`
	Email             string `json:"email"`
	HasPassword       byte   `json:"hasPassword"`
	LoginName         string `json:"loginName"`
	FullName          string `json:"fullName"`
	Gender            string `json:"gender"`
	Birthday          string `json:"birthday"`
	About             string `json:"about"`
	Website           string `json:"website"`
	Lang              string `json:"lang"`
	Location          string `json:"location"`
	ProfileColor      string `json:"profileColor"`
	IsFollowing       byte   `json:"isFollowing"`
	SafeMode          byte   `json:"safeMode"`
	HideUpvote        string `json:"hideUpvote"`
	TimezoneGMTOffset int    `json:"timezoneGmtOffset"`
	AvatarURLLarge    string `json:"avatarUrlLarge"`
	AvatarURLMedium   string `json:"avatarUrlMedium"`
	AvatarURLSmall    string `json:"avatarUrlSmall"`
	AvatarURLTiny     string `json:"avatarUrlTiny"`
	CanPostToFB       byte   `json:"canPostToFB"`
	FBDisplayName     string `json:"fbDisplayName"`
	FBUserID          string `json:"fbUserId"`
	FBTimeline        byte   `json:"fbTimeline"`
	FBPublish         byte   `json:"fbPublish"`
	FBLikeAction      byte   `json:"fbLikeAction"`
	FBCreateAction    byte   `json:"fbCreateAction"`
	FBCommentAction   byte   `json:"fbCommentAction"`
	GPlusAccountName  string `json:"gplusAccountName"`
	GPlusUserID       string `json:"gplusUserId"`
	// Permissions struct
}

// PostListData describes post list data from PostListResponse
type PostListData struct {
	DummyField   string     `json:"dummyField"` // always equals to "dummyValue"
	DidEndOfList byte       `json:"didEndOfList"`
	Posts        []PostData `json:"posts"`
}

// PostData describes post from PostListData
type PostData struct {
	ID             string `json:"id"`
	URL            string `json:"url"`
	Status         string `json:"status"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	Version        byte   `json:"version"`
	NSFW           byte   `json:"nsfw"`
	UpVoteCount    int    `json:"upVoteCount"`
	DownVoteCount  int    `json:"downVoteCount"`
	TotalVoteCount int    `json:"totalVoteCount"`
	ViewsCount     int    `json:"viewsCount"`
	Score          int    `json:"score"`
	ReportedStatus byte   `json:"reportedStatus"`
	CreationTs     int64  `json:"creationTs"`
	AlbumWebURL    string `json:"albumWebUrl"`
	HasImageTile   byte   `json:"hasImageTile"`
	// PostTile struct
	SortTs           int  `json:"sortTs"`
	OrderID          int  `json:"orderId"`
	HasLongPostCover byte `json:"hasLongPostCover"`
	Images           struct {
		Image700         Image `json:"image700"`
		Image460         Image `json:"image460"`
		Image220x145     Image `json:"image220x145"`
		ImageFbThumbnail Image `json:"imageFbThumbnail"`
	} `json:"images"`
	SourceDomain       string `json:"sourceDomain"`
	SourceURL          string `json:"sourceUrl"`
	ExternalURL        string `json:"externalUrl"`
	Channel            string `json:"channel"`
	IsVoted            string `json:"isVoted"`
	UserScore          int    `json:"userScore"`
	CommentOpClientID  string `json:"commentOpClientId"`
	CommentOpSignature string `json:"commentOpSignature"`
	// Creator struct
	CommentsCount int    `json:"commentsCount"`
	FbShares      int    `json:"fbShares"`
	TweetCount    int    `json:"tweetCount"`
	Created       string `json:"created"`
	CommentSystem string `json:"commentSystem"`
	// TopComments struct
	// TargetedAdTags struct
	Sections []string `json:"sections"`
	Tags     []Tag    `json:"tags"`
}

// Image describes image data from PostData
type Image struct {
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	URL         string `json:"url"`
	Mask        string `json:"mask"`
	Placeholder string `json:"placeholder"`
}

// Tag describes tag data from PostData
type Tag struct {
	Key string `json:"key"`
	URL string `json:"url"`
}
