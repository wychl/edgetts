package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wychl/edgetts"
)

var client *edgetts.Client

// global flag.
var proxy string // 代理

// voice cmd glag.
var (
	local  string // 语言代码
	gender string // 性别
)

// audio cmd flag.
var (
	text   string // 文本
	output string // 输出文件
	voice  string // 语音模型
	rate   string // 语音速率
	volume string // 语音音量
	pitch  string // 语音音调
)

func init() {
	// global flag
	rootCmd.PersistentFlags().StringVar(&proxy, "proxy", "", "proxy address")

	// voice cmd flag
	voiceCmd.Flags().StringVar(&local, "local", "", "language code")
	voiceCmd.Flags().StringVar(&gender, "gender", "", "voice gender Male/Female")

	// audioCmd flag
	speechCmd.Flags().StringVar(&text, "text", "", "the content of the text")
	speechCmd.Flags().StringVar(&output, "output", "output.mp3", "output file(mp3)")
	speechCmd.Flags().StringVar(&voice, "voice", "", "speech model")
	speechCmd.Flags().StringVar(&rate, "rate", "", "speech rate")
	speechCmd.Flags().StringVar(&volume, "volume", "", "speech volume")
	speechCmd.Flags().StringVar(&pitch, "pitch", "", "speech pitch")
	// _ = speechCmd.MarkFlagRequired("text")
	_ = speechCmd.MarkFlagRequired("voice")

	rootCmd.AddCommand(voiceCmd, speechCmd)
}

func main() {
	client = edgetts.New(&edgetts.Config{Proxy: proxy})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootExp = `
# 文本到语音
edgetts voice

# 文字转语音
edgetts speech --voice zh-CN-XiaoxiaoNeural --text "hello world"
`

var rootCmd = &cobra.Command{
	Use:     "edgetts",
	Short:   "edgetts 文本生成语音",
	Long:    `edgetts 调用Edge TTS服务,将文本文本生成语音`,
	Example: rootExp,
	Args:    cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, _ []string) {
	},
}

var voiceExp = `
# 列出所有语音的模型
edgetts voice

# 列出指定语言的语音的模型
edgetts voice --local "zh-CN"

# 列出指定性别语音的模型
edgetts voice --gender "Male"
edgetts voice --gender "Female"
`

var voiceCmd = &cobra.Command{
	Use:     "voice",
	Short:   "列出Edge TTS服务支持的语音模型",
	Long:    `列出Edge TTS服务支持的语音模型`,
	Example: voiceExp,
	Run: func(cmd *cobra.Command, args []string) {
		voices, err := client.GetVoice()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		list := make([]string, 0, len(voices))
		for _, voice := range voices {
			if edgetts.MatchVoice(voice, local, gender) {
				list = append(list, fmt.Sprintf("%s %s %s", voice.ShortName, voice.Locale, voice.Gender))
			}
		}
		value := strings.Join(list, "\n")
		_, _ = cmd.OutOrStdout().Write([]byte(value))
	},
}

var speechExp = `
# hello world
edgetts speech --text "hello world" --voice "zh-CN-YunxiNeural"

# 指定输出文件
edgetts speech --text "hello world" --voice "zh-CN-YunxiNeural" --output hello.mp3

# 指定语音速率
edgetts speech --text "hello world" --voice "zh-CN-YunxiNeural" --rate "-50%"

# 指定语音音量
edgetts speech --text "hello world" --voice "zh-CN-YunxiNeural" --volume "-50%"

# 指定语音音调
edgetts speech --text "hello world" --voice "zh-CN-YunxiNeural" --pitch "-50Hz"

# 使用os.Stdin 来从管道（pipeline）中读取输入文字
echo "hello world" | edgetts speech --voice "zh-CN-YunxiNeural"
`

var speechCmd = &cobra.Command{
	Use:       "speech",
	Short:     "speech 调用Edge TTS服务生成语音",
	Long:      `调用Edge TTS服务生成语音"`,
	Example:   speechExp,
	ValidArgs: []string{"text", "voice"},
	Run: func(cmd *cobra.Command, args []string) {
		voiceText := text
		if voiceText == "" {
			// 默认行为
			reader := bufio.NewReader(os.Stdin)
			for {
				input, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				voiceText = input
			}
		}
		if voiceText == "" {
			fmt.Println("请输入文字")
			os.Exit(1)
		}

		opts := make([]edgetts.SpeechOption, 0)
		if rate != "" {
			opts = append(opts, edgetts.WithRate(rate))
		}
		if volume != "" {
			opts = append(opts, edgetts.WithVolume(volume))
		}
		if pitch != "" {
			opts = append(opts, edgetts.WithPitch(pitch))
		}
		if pitch != "" {
			opts = append(opts, edgetts.WithPitch(pitch))
		}
		data, err := client.TTS(voiceText, voice, opts...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		f, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = f.Write(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
