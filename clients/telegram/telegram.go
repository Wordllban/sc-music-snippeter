package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"sc-music-snippeter/lib/errutil"
	"strconv"
)

const (
	getUpdatesMethod = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host: host,
		basePath: basePath(token),
		client: http.Client{},
	}
}


func (c *Client) SendMessage(chatID int, text string) error {
	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("text", text)

	_,err := c.doRequest(sendMessageMethod, query)
	if err != nil {
		return errutil.Wrap("can't send message", err)
	}

	return nil
}

func (c *Client) Updates(offset int, limit int)([]Update, error) {
	query := url.Values{}
	query.Add("offset", strconv.Itoa(offset))
	query.Add("limit", strconv.Itoa(limit))

	// do request
	data, err := c.doRequest(getUpdatesMethod, query)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); 
	err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) doRequest(method string, query url.Values) (data[]byte, err error) {
	defer func() { err = errutil.Wrap("can't do request", err) }()

	url := url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method),
	}

	req,err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp,err:= c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {_=resp.Body.Close()}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return body, nil
}

func basePath(token string) string {
	return "bot" + token
}