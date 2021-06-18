package domains

import "encoding/xml"

type ResponseHangup struct {
	XMLName xml.Name `xml:"Response"`
	Pause   Pause    `xml:"Pause,omitempty"`
	Hangup  Hangup   ``
}

type Hangup struct {
}
