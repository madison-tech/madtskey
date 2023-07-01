package madtskey

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"time"
)

type TSResponse struct {
	Id           string       `json:"id"`
	Key          string       `json:"key"`
	Created      time.Time    `json:"created"`
	Expires      time.Time    `json:"expires"`
	Revoked      time.Time    `json:"revoked"`
	Capabilities Capabilities `json:"capabilities"`
	Description  string       `json:"description"`
}

type Capabilities struct {
	Devices Devices `json:"devices"`
}

type Devices struct {
	Create Create `json:"create"`
}

type Create struct {
	Reusable      bool     `json:"reusable"`
	Ephemeral     bool     `json:"ephemeral"`
	Preauthorized bool     `json:"preauthorized"`
	Tags          []string `json:"tags"`
}

type Req struct {
	Capabilities  Capabilities `json:"capabilities"`
	ExpirySeconds int          `json:"expirySeconds"`
	Description   string       `json:"description"`
}

func (r *Req) AsReader() io.Reader {
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(data)
}
