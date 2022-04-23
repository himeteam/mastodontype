package mastodontype

const (
	// AttachmentTypeImage Static image
	AttachmentTypeImage = "image"
	// AttachmentTypeVideo   Video clip
	AttachmentTypeVideo = "video"
	// AttachmentTypeAudio Audio track
	AttachmentTypeAudio = "audio"
	// AttachmentTypeGifv  Looping, soundless animation
	AttachmentTypeGifv = "gifv"
	// AttachmentTypeUnknown  unsupported or unrecognized file type
	AttachmentTypeUnknown = "unknown"
)

// Attachment Represents a file or media attachment that can be added to a status.
// @doc https://docs.joinmastodon.org/entities/attachment/
type Attachment struct {
	// The ID of the attachment in the database.
	ID string `json:"id"`
	//  The type of the attachment.
	Type string `json:"type"`
	// The location of the original full-size attachment.
	URL string `json:"url"`
	// The location of a scaled-down preview of the attachment.
	PreviewURL string `json:"preview_url"`
	// The location of the full-size original attachment on the remote website.
	RemoteURL *string `json:"remote_url"`
	// Metadata returned by Paperclip.
	Meta any `json:"meta"`
	// Alternate text that describes what is in the media attachment, to be used for the visually impaired or when media attachments do not load.
	Description *string `json:"description"`
	// A hash computed by the BlurHash algorithm, for generating colorful preview thumbnails when media has not been downloaded yet.
	BlurHash *string `json:"blurhash"`
}

// ImageAttachmentMeta  Static image attachment meta.
type ImageAttachmentMeta struct {
	Original struct {
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Size   string  `json:"size"`
		Aspect float64 `json:"aspect"`
	} `json:"original"`

	Small struct {
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Size   string  `json:"size"`
		Aspect float64 `json:"aspect"`
	} `json:"small"`

	Focus struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"focus"`
}

// VideoAttachmentMeta Video clip attachment meta.
type VideoAttachmentMeta struct {
	Length        string  `json:"length"`
	Duration      float64 `json:"duration"`
	FPS           int     `json:"fps"`
	Size          string  `json:"size"`
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	Aspect        float64 `json:"aspect"`
	AudioEncode   string  `json:"audio_encode"`
	AudioBitrate  string  `json:"audio_bitrate"`
	AudioChannels string  `json:"audio_channels"`

	Original struct {
		Width     int     `json:"width"`
		Height    int     `json:"height"`
		FrameRate string  `json:"frame_rate"`
		Duration  float64 `json:"duration"`
		Bitrate   int     `json:"bitrate"`
	} `json:"original"`

	Small struct {
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Size   string  `json:"size"`
		Aspect float64 `json:"aspect"`
	} `json:"small"`
}

// GifvAttachmentMeta Looping, soundless animation attachment metadata.
type GifvAttachmentMeta struct {
	Length   string  `json:"length"`
	Duration float64 `json:"duration"`
	FPS      int     `json:"fps"`
	Size     string  `json:"size"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	Aspect   float64 `json:"aspect"`

	Original struct {
		Width     int     `json:"width"`
		Height    int     `json:"height"`
		FrameRate string  `json:"frame_rate"`
		Duration  float64 `json:"duration"`
		Bitrate   int     `json:"bitrate"`
	} `json:"original"`

	Small struct {
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Size   string  `json:"size"`
		Aspect float64 `json:"aspect"`
	} `json:"small"`
}

// AudioAttachmentMeta Audio track attachment metadata.
type AudioAttachmentMeta struct {
	Length        string  `json:"length"`
	Duration      float64 `json:"duration"`
	AudioEncode   string  `json:"audio_encode"`
	AudioBitrate  string  `json:"audio_bitrate"`
	AudioChannels string  `json:"audio_channels"`

	Original struct {
		Duration float64 `json:"duration"`
		Bitrate  int     `json:"bitrate"`
	} `json:"original"`
}
