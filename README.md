# edgetts

基于微软的 Azure Cognitive Services来实现文本到语音转换（TTS）。该库提供了一个简单的API，可以将文本转换为语音，并且支持多种语言和声音。

## 使用

### 命令行

```bash
go install github.com/wychl/edgetts/cmd@latest


#支持的语音模型
edgetts voice
edgetts voice --local zh-CN
edgetts voice --gender Male

#文本生成语音
edgetts speech--text "这是个测试" --voice "zh-CN-YunxiaNeural"
```

### 支持的语音模型

```go
package main

import(
	"fmt"
	"log"

	"github.com/wychl/edgetts"
)

funcmain(){
	cli:=edgetts.New(nil)
	data,err:=cli.GetVoice()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v",data)
}

```

### 文本生成语音

```go
package main

import(
	"log"
	"os"

	"github.com/wychl/edgetts"
)

func main(){
	cli:=edgetts.New(nil)
	data,err:=cli.TTS("这是个测试","zh-CN-YunxiaNeural")
	if err!=nil{
		log.Fatal(err)
	}
	WriteToFile("test.mp3",data)
}

func WriteToFile(filestring,data[]byte)error{
	f,err:=os.OpenFile(file,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	if err!=nil{
		return err
	}
	_,err=f.Write(data)
	return err
}
```

### 音速/音调/音量/

- `rate`音速

## 开发环境

- `golangci-lint`lint工具
- `gofumpt`格式化工具

```bash
#安装lint工具
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
#安装格式化工具
go install mvdan.cc/gofumpt@v0.6.0
#配置hook
git config core.hookspath.githooks
```

## EdgeTTS语音模型

|名称|语言|地区|性别|
|---|---|---|---|
|af-ZA-AdriNeural|阿非利卡语|南非|女性|
|af-ZA-WillemNeural|阿非利卡语|南非|男性|
|am-ET-AmehaNeural|阿姆哈拉语|埃塞俄比亚|男性|
|am-ET-MekdesNeural|阿姆哈拉语|埃塞俄比亚|女性|
|ar-AE-FatimaNeural|阿拉伯语|阿联酋|女性|
|ar-AE-HamdanNeural|阿拉伯语|阿联酋|男性|
|ar-BH-AliNeural|阿拉伯语|巴林|男性|
|ar-BH-LailaNeural|阿拉伯语|巴林|女性|
|ar-DZ-AminaNeural|阿拉伯语|阿尔及利亚|女性|
|ar-DZ-IsmaelNeural|阿拉伯语|阿尔及利亚|男性|
|ar-EG-SalmaNeural|阿拉伯语|埃及|女性|
|ar-EG-ShakirNeural|阿拉伯语|埃及|男性|
|ar-IQ-BasselNeural|阿拉伯语|伊拉克|男性|
|ar-IQ-RanaNeural|阿拉伯语|伊拉克|女性|
|ar-JO-SanaNeural|阿拉伯语|约旦|女性|
|ar-JO-TaimNeural|阿拉伯语|约旦|男性|
|ar-KW-FahedNeural|阿拉伯语|科威特|男性|
|ar-KW-NouraNeural|阿拉伯语|科威特|女性|
|ar-LB-LaylaNeural|阿拉伯语|黎巴嫩|女性|
|ar-LB-RamiNeural|阿拉伯语|黎巴嫩|男性|
|ar-LY-ImanNeural|阿拉伯语|利比亚|女性|
|ar-LY-OmarNeural|阿拉伯语|利比亚|男性|
|ar-MA-JamalNeural|阿拉伯语|摩洛哥|男性|
|ar-MA-MounaNeural|阿拉伯语|摩洛哥|女性|
|ar-OM-AbdullahNeural|阿拉伯语|阿曼|男性|
|ar-OM-AyshaNeural|阿拉伯语|阿曼|女性|
|ar-QA-AmalNeural|阿拉伯语|卡塔尔|女性|
|ar-QA-MoazNeural|阿拉伯语|卡塔尔|男性|
|ar-SA-HamedNeural|阿拉伯语|沙特阿拉伯|男性|
|ar-SA-ZariyahNeural|阿拉伯语|沙特阿拉伯|女性|
|ar-SY-AmanyNeural|阿拉伯语|叙利亚|女性|
|ar-SY-LaithNeural|阿拉伯语|叙利亚|男性|
|ar-TN-HediNeural|阿拉伯语|突尼斯|男性|
|ar-TN-ReemNeural|阿拉伯语|突尼斯|女性|
|ar-YE-MaryamNeural|阿拉伯语|也门|女性|
|ar-YE-SalehNeural|阿拉伯语|也门|男性|
|az-AZ-BabekNeural|阿塞拜疆语|阿塞拜疆|男性|
|az-AZ-BanuNeural|阿塞拜疆语|阿塞拜疆|女性|
|bg-BG-BorislavNeural|保加利亚语|保加利亚|男性|
|bg-BG-KalinaNeural|保加利亚语|保加利亚|女性|
|bn-BD-NabanitaNeural|孟加拉语|孟加拉国|女性|
|bn-BD-PradeepNeural|孟加拉语|孟加拉国|男性|
|bn-IN-BashkarNeural|孟加拉语|印度|男性|
|bn-IN-TanishaaNeural|孟加拉语|印度|女性|
|bs-BA-GoranNeural|波斯尼亚语|波斯尼亚和黑塞哥维那|男性|
|bs-BA-VesnaNeural|波斯尼亚语|波斯尼亚和黑塞哥维那|女性|
|ca-ES-EnricNeural|加泰罗尼亚语|西班牙|男性|
|ca-ES-JoanaNeural|加泰罗尼亚语|西班牙|女性|
|cs-CZ-AntoninNeural|捷克语|捷克共和国|男性|
|cs-CZ-VlastaNeural|捷克语|捷克共和国|女性|
|cy-GB-AledNeural|威尔士语|英国|男性|
|cy-GB-NiaNeural|威尔士语|英国|女性|
|da-DK-ChristelNeural|丹麦语|丹麦|女性|
|da-DK-JeppeNeural|丹麦语|丹麦|男性|
|de-AT-IngridNeural|德语|奥地利|女性|
|de-AT-JonasNeural|德语|奥地利|男性|
|de-CH-JanNeural|德语|瑞士|男性|
|de-CH-LeniNeural|德语|瑞士|女性|
|de-DE-AmalaNeural|德语|德国|女性|
|de-DE-ConradNeural|德语|德国|男性|
|de-DE-FlorianMultilingualNeural|德语|德国|男性|
|de-DE-KatjaNeural|德语|德国|女性|
|de-DE-KillianNeural|德语|德国|男性|
|de-DE-SeraphinaMultilingualNeural|德语|德国|女性|
|el-GR-AthinaNeural|希腊语|希腊|女性|
|el-GR-NestorasNeural|希腊语|希腊|男性|
|en-AU-NatashaNeural|英语|澳大利亚|女性|
|en-AU-WilliamNeural|英语|澳大利亚|男性|
|en-CA-ClaraNeural|英语|加拿大|女性|
|en-CA-LiamNeural|英语|加拿大|男性|
|en-GB-LibbyNeural|英语|英式英语英国|女性|
|en-GB-MaisieNeural|英语|英式英语英国|女性|
|en-GB-RyanNeural|英语|英式英语英国|男性|
|en-GB-SoniaNeural|英语|英式英语英国|女性|
|en-GB-ThomasNeural|英语|英式英语英国|男性|
|en-HK-SamNeural|英语|港式英语香港|男性|
|en-HK-YanNeural|英语|港式英语香港|女性|
|en-IE-ConnorNeural|英语|爱尔兰|男性|
|en-IE-EmilyNeural|英语|爱尔兰|女性|
|en-IN-NeerjaExpressiveNeural|英语|印度|女性|
|en-IN-NeerjaNeural|英语|印度|女性|
|en-IN-PrabhatNeural|英语|印度|男性|
|en-KE-AsiliaNeural|英语|肯尼亚|女性|
|en-KE-ChilembaNeural|英语|肯尼亚|男性|
|en-NG-AbeoNeural|英语|尼日利亚|男性|
|en-NG-EzinneNeural|英语|尼日利亚|女性|
|en-NZ-MitchellNeural|英语|新西兰|男性|
|en-NZ-MollyNeural|英语|新西兰|女性|
|en-PH-JamesNeural|英语|菲律宾|男性|
|en-PH-RosaNeural|英语|菲律宾|女性|
|en-SG-LunaNeural|英语|新加坡|女性|
|en-SG-WayneNeural|英语|新加坡|男性|
|en-TZ-ElimuNeural|英语|坦桑尼亚|男性|
|en-TZ-ImaniNeural|英语|坦桑尼亚|女性|
|en-US-AnaNeural|英语|美式英语美国|女性|
|en-US-AndrewNeural|英语|美式英语美国|男性|
|en-US-AriaNeural|英语|美式英语美国|女性|
|en-US-AvaNeural|英语|美式英语美国|女性|
|en-US-BrianNeural|英语|美式英语美国|男性|
|en-US-ChristopherNeural|英语|美式英语美国|男性|
|en-US-EmmaNeural|英语|美式英语美国|女性|
|en-US-EricNeural|英语|美式英语美国|男性|
|en-US-GuyNeural|英语|美式英语美国|男性|
|en-US-JennyNeural|英语|美式英语美国|女性|
|en-US-MichelleNeural|英语|美式英语美国|女性|
|en-US-RogerNeural|英语|美式英语美国|男性|
|en-US-SteffanNeural|英语|美式英语美国|男性|
|en-ZA-LeahNeural|英语|南非|女性|
|en-ZA-LukeNeural|英语|南非|男性|
|es-AR-ElenaNeural|西班牙语|阿根廷|女性|
|es-AR-TomasNeural|西班牙语|阿根廷|男性|
|es-BO-MarceloNeural|西班牙语|玻利维亚|男性|
|es-BO-SofiaNeural|西班牙语|玻利维亚|女性|
|es-CL-CatalinaNeural|西班牙语|智利|女性|
|es-CL-LorenzoNeural|西班牙语|智利|男性|
|es-CO-GonzaloNeural|西班牙语|哥伦比亚|男性|
|es-CO-SalomeNeural|西班牙语|哥伦比亚|女性|
|es-CR-JuanNeural|西班牙语|哥斯达黎加|男性|
|es-CR-MariaNeural|西班牙语|哥斯达黎加|女性|
|es-CU-BelkysNeural|西班牙语|古巴|女性|
|es-CU-ManuelNeural|西班牙语|古巴|男性|
|es-DO-EmilioNeural|西班牙语|多米尼加共和国|男性|
|es-DO-RamonaNeural|西班牙语|多米尼加共和国|女性|
|es-EC-AndreaNeural|西班牙语|厄瓜多尔|女性|
|es-EC-LuisNeural|西班牙语|厄瓜多尔|男性|
|es-ES-AlvaroNeural|西班牙语|西班牙|男性|
|es-ES-ElviraNeural|西班牙语|西班牙|女性|
|es-ES-XimenaNeural|西班牙语|西班牙|女性|
|es-GQ-JavierNeural|西班牙语|赤道几内亚|男性|
|es-GQ-TeresaNeural|西班牙语|赤道几内亚|女性|
|es-GT-AndresNeural|西班牙语|危地马拉|男性|
|es-GT-MartaNeural|西班牙语|危地马拉|女性|
|es-HN-CarlosNeural|西班牙语|洪都拉斯|男性|
|es-HN-KarlaNeural|西班牙语|洪都拉斯|女性|
|es-MX-DaliaNeural|西班牙语|墨西哥|女性|
|es-MX-JorgeNeural|西班牙语|墨西哥|男性|
|es-NI-FedericoNeural|西班牙语|尼加拉瓜|男性|
|es-NI-YolandaNeural|西班牙语|尼加拉瓜|女性|
|es-PA-MargaritaNeural|西班牙语|巴拿马|女性|
|es-PA-RobertoNeural|西班牙语|巴拿马|男性|
|es-PE-AlexNeural|西班牙语|秘鲁|男性|
|es-PE-CamilaNeural|西班牙语|秘鲁|女性|
|es-PR-KarinaNeural|西班牙语|波多黎各|女性|
|es-PR-VictorNeural|西班牙语|波多黎各|男性|
|es-PY-MarioNeural|西班牙语|巴拉圭|男性|
|es-PY-TaniaNeural|西班牙语|巴拉圭|女性|
|es-SV-LorenaNeural|西班牙语|萨尔瓦多|女性|
|es-SV-RodrigoNeural|西班牙语|萨尔瓦多|男性|
|es-US-AlonsoNeural|西班牙语|美国|男性|
|es-US-PalomaNeural|西班牙语|美国|女性|
|es-UY-MateoNeural|西班牙语|乌拉圭|男性|
|es-UY-ValentinaNeural|西班牙语|乌拉圭|女性|
|es-VE-PaolaNeural|西班牙语|委内瑞拉|女性|
|es-VE-SebastianNeural|西班牙语|委内瑞拉|男性|
|et-EE-AnuNeural|爱沙尼亚语|爱沙尼亚|女性|
|et-EE-KertNeural|爱沙尼亚语|爱沙尼亚|男性|
|fa-IR-DilaraNeural|波斯语|伊朗|女性|
|fa-IR-FaridNeural|波斯语|伊朗|男性|
|fi-FI-HarriNeural|芬兰语|芬兰|男性|
|fi-FI-NooraNeural|芬兰语|芬兰|女性|
|fil-PH-AngeloNeural|菲律宾语|菲律宾|男性|
|fil-PH-BlessicaNeural|菲律宾语|菲律宾|女性|
|fr-BE-CharlineNeural|法语|比利时（瓦隆语区）|女性|
|fr-BE-GerardNeural|法语|比利时（瓦隆语区）|男性|
|fr-CA-AntoineNeural|法语|加拿大（魁北克）|男性|
|fr-CA-JeanNeural|法语|加拿大（魁北克）|男性|
|fr-CA-SylvieNeural|法语|加拿大（魁北克）|女性|
|fr-CA-ThierryNeural|法语|加拿大（魁北克）|男性|
|fr-CH-ArianeNeural|法语|瑞士|女性|
|fr-CH-FabriceNeural|法语|瑞士|男性|
|fr-FR-DeniseNeural|法语|法国|女性|
|fr-FR-EloiseNeural|法语|法国|女性|
|fr-FR-HenriNeural|法语|法国|男性|
|fr-FR-RemyMultilingualNeural|法语|法国|男性|
|fr-FR-VivienneMultilingualNeural|法语|法国|女性|
|ga-IE-ColmNeural|爱尔兰语|爱尔兰|男性|
|ga-IE-OrlaNeural|爱尔兰语|爱尔兰|女性|
|gl-ES-RoiNeural|加利西亚语|加利西亚|男性|
|gl-ES-SabelaNeural|加利西亚语|加利西亚|女性|
|gu-IN-DhwaniNeural|古吉拉特语|印度|女性|
|gu-IN-NiranjanNeural|古吉拉特语|印度|男性|
|he-IL-AvriNeural|希伯来语|以色列|男性|
|he-IL-HilaNeural|希伯来语|以色列|女性|
|hi-IN-MadhurNeural|印地语|印度|男性|
|hi-IN-SwaraNeural|印地语|印度|女性|
|hr-HR-GabrijelaNeural|克罗地亚语|克罗地亚|女性|
|hr-HR-SreckoNeural|克罗地亚语|克罗地亚|男性|
|hu-HU-NoemiNeural|匈牙利语|匈牙利|女性|
|hu-HU-TamasNeural|匈牙利语|匈牙利|男性|
|id-ID-ArdiNeural|印尼语|印度尼西亚|男性|
|id-ID-GadisNeural|印尼语|印度尼西亚|女性|
|is-IS-GudrunNeural|冰岛语|冰岛|女性|
|is-IS-GunnarNeural|冰岛语|冰岛|男性|
|it-IT-DiegoNeural|意大利语|意大利|男性|
|it-IT-ElsaNeural|意大利语|意大利|女性|
|it-IT-GiuseppeNeural|意大利语|意大利|男性|
|it-IT-IsabellaNeural|意大利语|意大利|女性|
|ja-JP-KeitaNeural|日语|日语日本|男性|
|ja-JP-NanamiNeural|日语|日语日本|女性|
|jv-ID-DimasNeural|爪哇语|印度尼西亚|男性|
|jv-ID-SitiNeural|爪哇语|印度尼西亚|女性|
|ka-GE-EkaNeural|格鲁吉亚语|格鲁吉亚|女性|
|ka-GE-GiorgiNeural|格鲁吉亚语|格鲁吉亚|男性|
|kk-KZ-AigulNeural|哈萨克语|哈萨克斯坦|女性|
|kk-KZ-DauletNeural|哈萨克语|哈萨克斯坦|男性|
|km-KH-PisethNeural|柬埔寨语|柬埔寨|男性|
|km-KH-SreymomNeural|柬埔寨语|柬埔寨|女性|
|kn-IN-GaganNeural|卡纳达语|印度|男性|
|kn-IN-SapnaNeural|卡纳达语|印度|女性|
|ko-KR-HyunsuNeural|韩语|韩语韩国|男性|
|ko-KR-InJoonNeural|韩语|韩语韩国|男性|
|ko-KR-SunHiNeural|韩语|韩语韩国|女性|
|lo-LA-ChanthavongNeural|老挝语|老挝|男性|
|lo-LA-KeomanyNeural|老挝语|老挝|女性|
|lt-LT-LeonasNeural|立陶宛语|立陶宛|男性|
|lt-LT-OnaNeural|立陶宛语|立陶宛|女性|
|lv-LV-EveritaNeural|拉脱维亚语|拉脱维亚|女性|
|lv-LV-NilsNeural|拉脱维亚语|拉脱维亚|男性|
|mk-MK-AleksandarNeural|马其顿语|北马其顿|男性|
|mk-MK-MarijaNeural|马其顿语|北马其顿|女性|
|ml-IN-MidhunNeural|马拉雅拉姆语|印度|男性|
|ml-IN-SobhanaNeural|马拉雅拉姆语|印度|女性|
|mn-MN-BataaNeural|蒙古语|蒙古|男性|
|mn-MN-YesuiNeural|蒙古语|蒙古|女性|
|mr-IN-AarohiNeural|马拉地语|印度|女性|
|mr-IN-ManoharNeural|马拉地语|印度|男性|
|ms-MY-OsmanNeural|马来语|马来西亚|男性|
|ms-MY-YasminNeural|马来语|马来西亚|女性|
|mt-MT-GraceNeural|马耳他语|马耳他|女性|
|mt-MT-JosephNeural|马耳他语|马耳他|男性|
|my-MM-NilarNeural|缅甸语|缅甸|女性|
|my-MM-ThihaNeural|缅甸语|缅甸|男性|
|nb-NO-FinnNeural|挪威语（书面）|挪威|男性|
|nb-NO-PernilleNeural|挪威语（书面）|挪威|女性|
|ne-NP-HemkalaNeural|尼泊尔语|尼泊尔|女性|
|ne-NP-SagarNeural|尼泊尔语|尼泊尔|男性|
|nl-BE-ArnaudNeural|荷兰语|比利时（弗拉芒语区）|男性|
|nl-BE-DenaNeural|荷兰语|比利时（弗拉芒语区）|女性|
|nl-NL-ColetteNeural|荷兰语|荷兰|女性|
|nl-NL-FennaNeural|荷兰语|荷兰|女性|
|nl-NL-MaartenNeural|荷兰语|荷兰|男性|
|pl-PL-MarekNeural|波兰语|波兰|男性|
|pl-PL-ZofiaNeural|波兰语|波兰|女性|
|ps-AF-GulNawazNeural|普什图语|阿富汗|男性|
|ps-AF-LatifaNeural|普什图语|阿富汗|女性|
|pt-BR-AntonioNeural|葡萄牙语|巴西|男性|
|pt-BR-FranciscaNeural|葡萄牙语|巴西|女性|
|pt-BR-ThalitaNeural|葡萄牙语|巴西|女性|
|pt-PT-DuarteNeural|葡萄牙语|葡萄牙|男性|
|pt-PT-RaquelNeural|葡萄牙语|葡萄牙|女性|
|ro-RO-AlinaNeural|罗马尼亚语|罗马尼亚|女性|
|ro-RO-EmilNeural|罗马尼亚语|罗马尼亚|男性|
|ru-RU-DmitryNeural|俄语|俄罗斯|男性|
|ru-RU-SvetlanaNeural|俄语|俄罗斯|女性|
|si-LK-SameeraNeural|辛哈拉语|斯里兰卡|男性|
|si-LK-ThiliniNeural|辛哈拉语|斯里兰卡|女性|
|sk-SK-LukasNeural|斯洛伐克语|斯洛伐克|男性|
|sk-SK-ViktoriaNeural|斯洛伐克语|斯洛伐克|女性|
|sl-SI-PetraNeural|斯洛文尼亚语|斯洛文尼亚|女性|
|sl-SI-RokNeural|斯洛文尼亚语|斯洛文尼亚|男性|
|so-SO-MuuseNeural|索马里语|索马里|男性|
|so-SO-UbaxNeural|索马里语|索马里|女性|
|sq-AL-AnilaNeural|阿尔巴尼亚语|阿尔巴尼亚|女性|
|sq-AL-IlirNeural|阿尔巴尼亚语|阿尔巴尼亚|男性|
|sr-RS-NicholasNeural|塞尔维亚语|塞尔维亚|男性|
|sr-RS-SophieNeural|塞尔维亚语|塞尔维亚|女性|
|su-ID-JajangNeural|巽他语|印度尼西亚|男性|
|su-ID-TutiNeural|巽他语|印度尼西亚|女性|
|sv-SE-MattiasNeural|瑞典语|瑞典|男性|
|sv-SE-SofieNeural|瑞典语|瑞典|女性|
|sw-KE-RafikiNeural|斯瓦希里语|肯尼亚|男性|
|sw-KE-ZuriNeural|斯瓦希里语|肯尼亚|女性|
|sw-TZ-DaudiNeural|斯瓦希里语|坦桑尼亚|男性|
|sw-TZ-RehemaNeural|斯瓦希里语|坦桑尼亚|女性|
|ta-IN-PallaviNeural|泰米尔语|印度|女性|
|ta-IN-ValluvarNeural|泰米尔语|印度|男性|
|ta-LK-KumarNeural|泰米尔语|斯里兰卡|男性|
|ta-LK-SaranyaNeural|泰米尔语|斯里兰卡|女性|
|ta-MY-KaniNeural|泰米尔语|马来西亚|女性|
|ta-MY-SuryaNeural|泰米尔语|马来西亚|男性|
|ta-SG-AnbuNeural|泰米尔语|新加坡|男性|
|ta-SG-VenbaNeural|泰米尔语|新加坡|女性|
|te-IN-MohanNeural|泰卢固语|印度|男性|
|te-IN-ShrutiNeural|泰卢固语|印度|女性|
|th-TH-NiwatNeural|泰语|泰国|男性|
|th-TH-PremwadeeNeural|泰语|泰国|女性|
|tr-TR-AhmetNeural|土耳其语|土耳其|男性|
|tr-TR-EmelNeural|土耳其语|土耳其|女性|
|uk-UA-OstapNeural|乌克兰语|乌克兰|男性|
|uk-UA-PolinaNeural|乌克兰语|乌克兰|女性|
|ur-IN-GulNeural|乌尔都语|印度|女性|
|ur-IN-SalmanNeural|乌尔都语|印度|男性|
|ur-PK-AsadNeural|乌尔都语|巴基斯坦|男性|
|ur-PK-UzmaNeural|乌尔都语|巴基斯坦|女性|
|uz-UZ-MadinaNeural|乌兹别克语|乌兹别克斯坦|女性|
|uz-UZ-SardorNeural|乌兹别克语|乌兹别克斯坦|男性|
|vi-VN-HoaiMyNeural|越南语|越南|女性|
|vi-VN-NamMinhNeural|越南语|越南|男性|
|zh-CN-XiaoxiaoNeural|汉语（简体中文）|普通话中国|女性|
|zh-CN-XiaoyiNeural|汉语（简体中文）|普通话中国|女性|
|zh-CN-YunjianNeural|汉语（简体中文）|普通话中国|男性|
|zh-CN-YunxiNeural|汉语（简体中文）|普通话中国|男性|
|zh-CN-YunxiaNeural|汉语（简体中文）|普通话中国|男性|
|zh-CN-YunyangNeural|汉语（简体中文）|普通话中国|男性|
|zh-CN-liaoning|汉语（简体中文）|辽宁方言中国|女性|
|zh-CN-shaanxi|汉语（简体中文）|陕西方言中国|女性|
|zh-HK-HiuGaaiNeural|汉语（繁体中文）|粤语香港|女性|
|zh-HK-HiuMaanNeural|汉语（繁体中文）|粤语香港|女性|
|zh-HK-WanLungNeural|汉语（繁体中文）|粤语香港|男性|
|zh-TW-HsiaoChenNeural|汉语（繁体中文）|台湾|女性|
|zh-TW-HsiaoYuNeural|汉语（繁体中文）|台湾|女性|
|zh-TW-YunJheNeural|汉语（繁体中文）|台湾|男性|
|zu-ZA-ThandoNeural|祖鲁语|南非|女性|
|zu-ZA-ThembaNeural|祖鲁语|南非|男性|

## 参考资料

- [SSML](https://cloud.google.com/text-to-speech/docs/ssml?hl=zh-cn)
- [prosody](https://www.w3.org/TR/speech-synthesis11/#S3.2.4)
