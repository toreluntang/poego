package poego

import (
	"net/url"
)

//GetLadder returns a ladder for the supplied id.
//Optional value "limit" which is an integer between 20 and 200 specifies how many entries to pull.
//Optional value "offset" specifies how far from 0 to begin pulling entries.
//The max amount of entries in a ladder is 15000 which means you need to make 75 requests to pull an entire set.
func (p *Poego) GetLadder(id string, v url.Values) (ladder []Ladder, err error) {

	r := p.buildRequest("GET", "/ladders/"+id, v)
	err = p.makeRequest(r, &ladder)

	if err != nil {
		return ladder, err
	}

	return ladder, err
}