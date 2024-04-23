package edgetts

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
)

const (
	trustedClientToken = "6A5AA1D4EAFF4E9FB37E23D68491D6F4"
	getSSSAPI          = "wss://speech.platform.bing.com/consumer/speech/synthesize/readaloud/edge/v1?TrustedClientToken=" + trustedClientToken
	getVoiceAPI        = "https://speech.platform.bing.com/consumer/speech/synthesize/readaloud/voices/list?trustedclienttoken=" + trustedClientToken
)

type Client struct {
	cfg *Config
}

func New(cfg *Config) *Client {
	return &Client{cfg: cfg}
}

func (sev *Client) GetVoice() ([]Voice, error) {
	client := sev.getHTTPCli()
	resp, err := client.R().
		EnableTrace().
		SetHeaders(map[string]string{
			"Authority":        "speech.platform.bing.com",
			"Sec-CH-UA":        `" Not;A Brand";v="99", "Microsoft Edge";v="91", "Chromium";v="91"`,
			"Sec-CH-UA-Mobile": "?0",
			"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36 Edg/91.0.864.41",
			"Accept":           "*/*",
			"Sec-Fetch-Site":   "none",
			"Sec-Fetch-Mode":   "cors",
			"Sec-Fetch-Dest":   "empty",
			"Accept-Encoding":  "gzip, deflate, br",
			"Accept-Language":  "en-US,en;q=0.9",
		}).
		Get(getVoiceAPI)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to list voices, http status code: %s", resp.Status())
	}

	var list []Voice
	if err := json.Unmarshal(resp.Body(), &list); err != nil {
		return nil, err
	}

	return list, nil
}

func (sev *Client) TTS(text, voice string, options ...SpeechOption) ([]byte, error) {
	if voice == "" {
		return nil, ErrTTSVoiceEmpty
	}
	speech := &Speech{
		text:           text,
		voice:          voice,
		rate:           "+0%",
		volume:         "+0%",
		pitch:          "+0Hz",
		receiveTimeout: 10,
	}
	for _, option := range options {
		option(speech)
	}

	return sev.tts(speech)
}

func (sev *Client) getHTTPCli() *resty.Client {
	client := resty.New()
	if proxyURL := sev.getProxy(); proxyURL != "" {
		client.SetProxy(proxyURL)
	}
	return client
}

func (sev *Client) tts(req *Speech) ([]byte, error) {
	conn, err := sev.newConn(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()

	done := make(chan struct{})
	failed := make(chan error)
	audioData := make([]byte, 0)

	go func() {
		defer func() {
			close(done)
			close(failed)
		}()

		for {
			msgType, data, err := conn.ReadMessage()
			if msgType == -1 && data == nil && err != nil { // 已经断开链接
				failed <- err
				return
			}

			switch msgType {
			case websocket.TextMessage:
				textHeader, err := getHeadersAndData(data)
				if err != nil {
					failed <- err
					return
				}
				if string(textHeader["Path"]) == "turn.end" {
					return
				}
			case websocket.BinaryMessage:
				if len(data) < 2 {
					failed <- errors.New("we received a binary message, but it is missing the header length")
					return
				}

				length := binary.BigEndian.Uint16(data[:2])
				if len(data) < int(length+2) {
					failed <- errors.New("we received a binary message, but it is missing the audio data")
					return
				}
				audioData = append(audioData, data[2+length:]...)
			default:
				log.Println("recv:", data)
			}
		}
	}()

	for _, msg := range req.ToMessage() {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			return nil, err
		}
	}

	select {
	case err := <-failed:
		return nil, err
	case <-done:
		return audioData, err
	}
}

func (sev *Client) newConn(req *Speech) (*websocket.Conn, error) {
	dialer := &websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: true,
	}

	if proxy := sev.getProxy(); proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}
		dialer.Proxy = http.ProxyURL(proxyURL)
	}

	header := http.Header{
		"Pragma":          []string{"no-cache"},
		"Cache-Control":   []string{"no-cache"},
		"Origin":          []string{"chrome-extension://jdiccldimpdaibmpdkjnbmckianbfold"},
		"Accept-Encoding": []string{"gzip, deflate, br"},
		"Accept-Language": []string{"en-US,en;q=0.9"},
		"User-Agent":      []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36 Edg/91.0.864.41"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(req.receiveTimeout)*time.Second)
	defer func() {
		cancel()
	}()
	reqURL := getSSSAPI + "&ConnectionId=" + getRequestID()
	conn, _, err := dialer.DialContext(ctx, reqURL, header) //nolint
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (sev *Client) getProxy() string {
	proxy := ""
	if sev.cfg != nil {
		proxy = sev.cfg.Proxy
	}
	return proxy
}
