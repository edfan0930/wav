package playwav

import (
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

// 預嵌入的音頻文件
//
//go:embed sounds/english/*.wav
var englishSounds embed.FS

//go:embed sounds/mandarin/*.wav
var mandarinSounds embed.FS

//go:embed sounds/hokkien/*.wav
var hokkienSounds embed.FS

//go:embed sounds/hakka/*.wav
var hakkaSounds embed.FS

//go:embed sounds/other/*.wav
var otherSounds embed.FS

// 全局变量，存储所有的 beep.Buffer
var englishBuffers, mandarinBuffers, hokkienBuffers, hakkaBuffers, otherBuffers map[string]*beep.Buffer

func InitSpeaker() {
	/* 	LoadBuffersFromFS(&englishSounds, &englishBuffers)
	   	LoadBuffersFromFS(&mandarinSounds, &mandarinBuffers)
	   	LoadBuffersFromFS(&hokkienSounds, &hokkienBuffers)
	   	LoadBuffersFromFS(&hakkaSounds, &hakkaBuffers)
	   	LoadBuffersFromFS(&otherSounds, &otherBuffers) */
	fileName := "1.wav" // 示例文件名

	// 初始化 speaker，仅在开始时调用一次
	var format beep.Format
	path := fmt.Sprintf("%s/%s/%s", "sounds", "english", fileName)
	f, err := englishSounds.Open(path)
	if err != nil {
		log.Fatalf("unable to open file %s: %v", path, err)
	}
	_, format, err = wav.Decode(f)
	if err != nil {
		log.Fatalf("unable to decode file %s: %v", path, err)
	}
	f.Close()

	//initial speaker
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// 加载所有音频文件到对应的 buffers
	//LoadBuffersFromFS("english", &englishSounds, englishBuffers, englishFileNames)
}

func InitMandarinBuffers(fileNames []string) error {
	mandarinBuffers = make(map[string]*beep.Buffer)

	err := LoadBuffersFromFS("mandarin", &mandarinSounds, mandarinBuffers, fileNames)

	return err
}

func InitEnglishBuffers(fileNames []string) error {
	englishBuffers = make(map[string]*beep.Buffer)

	err := LoadBuffersFromFS("english", &englishSounds, englishBuffers, fileNames)

	return err
}
func InitHokkienBuffers(fileNames []string) error {

	hokkienBuffers = make(map[string]*beep.Buffer)

	err := LoadBuffersFromFS("hokkien", &hokkienSounds, hokkienBuffers, fileNames)

	return err
}

func InitHakkaBuffers(fileNames []string) error {

	hakkaBuffers = make(map[string]*beep.Buffer)
	err := LoadBuffersFromFS("hakka", &hakkaSounds, hakkaBuffers, fileNames)

	return err
}

func LoadBuffersFromFS(language string, fs *embed.FS, buffers map[string]*beep.Buffer, fileNames []string) error {

	for _, fileName := range fileNames {

		path := fmt.Sprintf("%s/%s/%s", "sounds", language, fileName)

		f, err := fs.Open(path)
		if err != nil {
			return err
		}

		streamer, format, err := wav.Decode(f)
		if err != nil {
			return err
		}

		buffer := beep.NewBuffer(format)
		buffer.Append(streamer)
		streamer.Close()

		buffers[fileName] = buffer
		f.Close()
	}

	return nil
}

/* func LoadSoundFiles() {

	for i := 0; i <= 9; i++ {
		fileName := fmt.Sprintf("sounds/mandarin/%d.wav", i)

		// 打开音频文件
		f, err := mandarinSounds.Open(fileName)
		if err != nil {
			log.Fatalf("unable to open file %s: %v", fileName, err)
		}

		// 解码音频文件
		streamer, format, err := wav.Decode(f)
		if err != nil {
			log.Fatalf("unable to decode file %s: %v", fileName, err)
		}

		// 在第一个文件时初始化 speaker
		if i == 0 {
			initSpeaker(format)
		}

		// 创建 Buffer 并加载音频数据
		buffer := beep.NewBuffer(format)
		buffer.Append(streamer)
		streamer.Close()

		// 将 Buffer 添加到全局变量 buffers 中
		buffers = append(buffers, buffer)

		// 关闭文件
		f.Close()
	}
} */
