package scprobe

type StreamType string

const (
	STREAM_ANY   StreamType = ""
	STREAM_VIDEO            = "video"
	STREAM_AUDIO            = "audio"
)
/*learn from https://github.com/xdave/ff/blob/master/info.go*/
type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails"`
}
// Stream side data. More fields may be added later as needed.
type SideData struct {
	Displaymatrix string `json:"displaymatrix"`
	Rotation      int    `json:"rotation"`
	Size          int    `json:"side_data_size"`
	Type          string `json:"side_data_type"`
}
type Streams struct {
	Index          int    `json:"index"`
	CodecName      string `json:"codec_name"`
	CodecLongName  string `json:"codec_long_name"`
	Profile        string `json:"profile"`
	CodecType      string `json:"codec_type"`
	CodecTimeBase  string `json:"codec_time_base"`
	CodecTagString string `json:"codec_tag_string"`
	CodecTag       string `json:"codec_tag"`
	SampleFmt      string `json:"sample_fmt,omitempty"`
	SampleRate     string `json:"sample_rate,omitempty"`
	Channels       int    `json:"channels,omitempty"`
	ChannelLayout  string `json:"channel_layout,omitempty"`
	BitsPerSample  int    `json:"bits_per_sample,omitempty"`
	ID             string `json:"id"`
	RFrameRate     string `json:"r_frame_rate"`
	AvgFrameRate   string `json:"avg_frame_rate"`
	TimeBase       string `json:"time_base"`
	StartPts       int    `json:"start_pts"`
	StartTime      string `json:"start_time"`
	DurationTs     int    `json:"duration_ts"`
	Duration       string `json:"duration"`
	BitRate        string `json:"bit_rate,omitempty"`
	Disposition    Disposition `json:"disposition"`
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	CodedWidth         int    `json:"coded_width,omitempty"`
	CodedHeight        int    `json:"coded_height,omitempty"`
	HasBFrames         int    `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
	PixFmt             string `json:"pix_fmt,omitempty"`
	Level              int    `json:"level,omitempty"`
	ColorRange         string `json:"color_range,omitempty"`
	ColorSpace         string `json:"color_space,omitempty"`
	ColorTransfer      string `json:"color_transfer,omitempty"`
	ColorPrimaries     string `json:"color_primaries,omitempty"`
	ChromaLocation     string `json:"chroma_location,omitempty"`
	FieldOrder         string `json:"field_order,omitempty"`
	Refs               int    `json:"refs,omitempty"`
	IsAvc              string `json:"is_avc,omitempty"`
	NalLengthSize      string `json:"nal_length_size,omitempty"`
	BitsPerRawSample   string `json:"bits_per_raw_sample,omitempty"`
	SideDataList       []SideData  `json:"side_data_list"`
}


type Format struct {
	Filename       string `json:"filename"`
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTimeSeconds float64 `json:"start_time,string"`
	DurationSecons   float64 `json:"duration,string"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
}

type FFProbeInfo struct {
	Streams []*Streams`json:"streams"`
	Format Format `json:"format"`
}

type Dimension struct {
	Width  int
	Height int
	CodedWidth int
	CodedHeight int
}

/*
##out###{
    "streams": [
        {
            "index": 0,
            "codec_name": "aac",
            "codec_long_name": "AAC (Advanced Audio Coding)",
            "profile": "HE-AACv2",
            "codec_type": "audio",
            "codec_time_base": "1/44100",
            "codec_tag_string": "[15][0][0][0]",
            "codec_tag": "0x000f",
            "sample_fmt": "fltp",
            "sample_rate": "44100",
            "channels": 2,
            "channel_layout": "stereo",
            "bits_per_sample": 0,
            "id": "0x100",
            "r_frame_rate": "0/0",
            "avg_frame_rate": "0/0",
            "time_base": "1/90000",
            "start_pts": 3811548678,
            "start_time": "42350.540867",
            "duration_ts": 16763259,
            "duration": "186.258433",
            "bit_rate": "23944",
            "disposition": {
                "default": 0,
                "dub": 0,
                "original": 0,
                "comment": 0,
                "lyrics": 0,
                "karaoke": 0,
                "forced": 0,
                "hearing_impaired": 0,
                "visual_impaired": 0,
                "clean_effects": 0,
                "attached_pic": 0
            }
        },
        {
            "index": 1,
            "codec_name": "h264",
            "codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
            "codec_type": "video",
            "codec_time_base": "1/180000",
            "codec_tag_string": "[27][0][0][0]",
            "codec_tag": "0x001b",
            "width": 0,
            "height": 0,
            "coded_width": 0,
            "coded_height": 0,
            "has_b_frames": 0,
            "sample_aspect_ratio": "0:1",
            "display_aspect_ratio": "0:0",
            "level": -99,
            "chroma_location": "left",
            "refs": 1,
            "is_avc": "0",
            "nal_length_size": "0",
            "id": "0x101",
            "r_frame_rate": "180000/2",
            "avg_frame_rate": "0/0",
            "time_base": "1/90000",
            "start_pts": 3811548678,
            "start_time": "42350.540867",
            "duration_ts": 16763259,
            "duration": "186.258433",
            "disposition": {
                "default": 0,
                "dub": 0,
                "original": 0,
                "comment": 0,
                "lyrics": 0,
                "karaoke": 0,
                "forced": 0,
                "hearing_impaired": 0,
                "visual_impaired": 0,
                "clean_effects": 0,
                "attached_pic": 0
            }
        }
    ],
    "format": {
        "filename": "http://live10034.bs2dl.yy.com/crs_612dc769e21e411383ed30136aa52363.ts",
        "nb_streams": 2,
        "nb_programs": 1,
        "format_name": "mpegts",
        "format_long_name": "MPEG-TS (MPEG-2 Transport Stream)",
        "start_time": "42350.540867",
        "duration": "186.258433",
        "size": "17866580",
        "bit_rate": "767388",
        "probe_score": 100
    }
}


*/