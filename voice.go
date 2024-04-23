package edgetts

// Voice 声音
type Voice struct {
	Name           string   `json:"Name"`           // 名称
	ShortName      string   `json:"ShortName"`      // 简称
	Gender         string   `json:"Gender"`         // 性别
	Locale         string   `json:"Locale"`         // 语言代码
	SuggestedCodec string   `json:"SuggestedCodec"` // 编码
	FriendlyName   string   `json:"FriendlyName"`   // 描述
	Status         string   `json:"Status"`         // 状态
	VoiceTag       VoiceTag `json:"VoiceTag"`       // 标签
}

type VoiceTag struct {
	ContentCategories  []string `json:"ContentCategories"`  // 分类
	VoicePersonalities []string `json:"VoicePersonalities"` // 特性
}

func MatchVoice(voice Voice, local, gender string) bool {
	if local != "" && voice.Locale != local {
		return false
	}
	if gender != "" && voice.Gender != gender {
		return false
	}
	return true
}
