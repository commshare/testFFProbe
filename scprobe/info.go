package scprobe

import ("time"

	"encoding/json"
)

func (format *Format) StartTime()(duration time.Duration){
	return time.Duration(format.StartTimeSeconds * float64(time.Second))
}

func (format *Format) Duration()(duration time.Duration) {
	return time.Duration(format.DurationSecons * float64(time.Second))
}
/*
Returns a []Stream slice, even if it's empty
*/
func (p *FFProbeInfo) GetStreams(streamType StreamType)(streams []Streams){
	if  p.Streams != nil {
		for _,s := range p.Streams {
			if s == nil {
				continue
			}
			switch streamType {
			case STREAM_ANY :
				streams = append(streams,*s)
			default:
				if string(streamType) == s.CodecType {
					streams = append(streams,*s)
				}
			}
		}
	}
	return
}

func (p *FFProbeInfo) GetFirstVideoStream() (vStream *Streams){
	if p.Streams == nil {
		//logger.Error("ffprobe streams is nil")
		return nil
	}
	for _, s := range p.Streams {
		if s== nil {
			continue
		}
		if s.CodecType == STREAM_VIDEO {
			vStream = s
			return
		}
	}
	return
}

func (p *FFProbeInfo)GetFirstAudioStream()(aStream *Streams) {
	if p.Streams != nil {
		for _,s :=range p.Streams {
			if s == nil {
				continue
			}
			if s.CodecType == STREAM_AUDIO {
				aStream = s
				return
			}
		}
	}
	return
}

func (p * FFProbeInfo)GetDimension() *Dimension {
	if p.Streams != nil {
		for _,s := range p.Streams {
			if s == nil {
				continue
			}
			if s.CodecType == STREAM_VIDEO {
				return &Dimension{
					Width:s.Width,
					Height:s.Height,
					CodedWidth: s.CodedWidth,
					CodedHeight:s.CodedHeight,
				}
			}
		}
	}
	return nil
}

func (p *FFProbeInfo)GetDuration(){
	p.Format.Duration()
}

func (p *FFProbeInfo)GetStartTime(){
	p.Format.StartTime()
}

func NewInfo(jsondata string )(*FFProbeInfo, error){
	info := new(FFProbeInfo)
	err := json.Unmarshal([]byte(jsondata),info)
	if err != nil {
		return nil ,err
	}
	return info,nil
}