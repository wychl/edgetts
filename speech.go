package edgetts

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

// https://cloud.google.com/text-to-speech/docs/ssml?hl=zh-cn#prosody
// https://www.w3.org/TR/speech-synthesis11/#S3.2.4
const ssmlTpl = `
<speak xmlns="http://www.w3.org/2001/10/synthesis" version="1.0" xml:lang="en-US">
   <voice name="%s">
      <prosody pitch="%s" rate="%s" volume="%s">%s</prosody>
   </voice>
</speak>
`

// Speech is a struct representing speech with the service.
type Speech struct {
	text  string // 文本
	voice string // 语音模型

	rate   string // 速率
	volume string // 音量
	pitch  string // 音调

	receiveTimeout int // 生成语音的超时时间，单位秒
}

func (req *Speech) ToMessage() []string {
	return []string{
		req.getCommand(),
		req.getSSML(),
	}
}

func (req *Speech) getCommand() string {
	var builder strings.Builder

	// 拼接X-Timestamp部分
	builder.WriteString(fmt.Sprintf("X-Timestamp:%s\r\n", getDate()))

	// 拼接Content-Type部分
	builder.WriteString("Content-Type:application/json; charset=utf-8\r\n")

	// 拼接Path部分
	builder.WriteString("Path:speech.config\r\n\r\n")

	// 拼接JSON部分
	builder.WriteString(`{"context":{"synthesis":{"audio":{"metadataoptions":{`)
	builder.WriteString(`"sentenceBoundaryEnabled":false,"wordBoundaryEnabled":true},`)
	builder.WriteString(`"outputFormat":"audio-24khz-48kbitrate-mono-mp3"`)
	builder.WriteString("}}}}\r\n")

	return builder.String()
}

func (req *Speech) getSSML() string {
	log.Printf("speech: %s %s %s %s %s", req.text, req.voice, req.rate, req.volume, req.pitch)

	requestID := getRequestID()
	timestamp := getDate()
	ssmlValue := getSSML(req.text, req.voice, req.rate, req.volume, req.pitch)

	ssml := fmt.Sprintf(
		"X-RequestId:%s\r\n"+
			"Content-Type:application/ssml+xml\r\n"+
			"X-Timestamp:%sZ\r\n"+
			"Path:ssml\r\n\r\n"+
			"%s",
		requestID, timestamp, ssmlValue)
	return ssml
}

type SpeechOption func(*Speech)

// WithRate sets the rate for speech.
func WithRate(rate string) SpeechOption {
	return func(c *Speech) {
		c.rate = rate
	}
}

// WithVolume sets the volume for speech.
func WithVolume(volume string) SpeechOption {
	return func(c *Speech) {
		c.volume = volume
	}
}

// WithPitch sets the pitch for speech.
func WithPitch(pitch string) SpeechOption {
	return func(c *Speech) {
		c.pitch = pitch
	}
}

// WithReceiveTimeout sets the receive timeout for speech.
func WithReceiveTimeout(receiveTimeout int) SpeechOption {
	return func(c *Speech) {
		c.receiveTimeout = receiveTimeout
	}
}

// getHeadersAndData returns the headers and data from the given data.
func getHeadersAndData(data interface{}) (map[string][]byte, error) {
	var dataBytes []byte
	switch v := data.(type) {
	case string:
		dataBytes = []byte(v)
	case []byte:
		dataBytes = v
	default:
		return nil, errors.New("data must be string or []byte")
	}

	headers := make(map[string][]byte)
	headerEnd := bytes.Index(dataBytes, []byte("\r\n\r\n"))
	if headerEnd == -1 {
		return nil, errors.New("invalid data format: no header end")
	}

	headerLines := bytes.Split(dataBytes[:headerEnd], []byte("\r\n"))
	for _, line := range headerLines {
		parts := bytes.SplitN(line, []byte(":"), 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid header format")
		}
		key := string(bytes.TrimSpace(parts[0]))
		value := bytes.TrimSpace(parts[1])
		headers[key] = value
	}

	return headers, nil
}

// getRequestID generates a UUID without dashes.
func getRequestID() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(), "-", "")
}

// getSSML creates an SSML string from the given parameters.
func getSSML(text string, voice string, rate string, volume string, pitch string) string {
	ssml := fmt.Sprintf(ssmlTpl, voice, pitch, rate, volume, text)
	return ssml
}

// getDate returns a JavaScript-style date string.
func getDate() string {
	date := time.Now().UTC()
	value := date.Format("Mon Jan 02 2006 15:04:05 GMT+0000 (Coordinated Universal Time)")
	return value
}
