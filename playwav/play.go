package playwav

import (
	"errors"
	"fmt"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

/* func Play() {

	// 选择要播放的音频buffer索引，例如播放索引为0的音频
	playIndex := 3

	// 检查索引是否有效
	if playIndex < 0 || playIndex >= len(buffers) {
		log.Fatalf("playIndex out of range: %d", playIndex)
	}

	// 从buffers获取streamer
	shot := buffers[playIndex].Streamer(0, buffers[playIndex].Len())

	// 播放音频并等待播放完毕
	playAndWait(shot)
} */

func PlayWithMandarinSingle(number int) error {

	fileNames := SingleWithMandarin(number)
	var streamers []beep.Streamer
	fmt.Println("mandarin buffers", mandarinBuffers)
	fmt.Println("names", fileNames)
	for _, fileName := range fileNames {
		fmt.Println("file name", fileName)
		v, exist := mandarinBuffers[fileName]
		if !exist {
			return errors.New("wrong file name")
		}

		streamer := v.Streamer(0, v.Len())
		streamers = append(streamers, streamer)
	}

	playAndWait(streamers)
	return nil
}

// playAndWait 播放streamer并等待其播放完毕
func playAndWait(streamers []beep.Streamer) {

	done := make(chan bool)
	sequence := beep.Seq(streamers...) // 展开 streamers 切片
	sequence = beep.Seq(sequence, beep.Callback(func() {

		go func() {
			t := time.NewTimer(1 * time.Second)
			<-t.C
			done <- true
		}()
	}))

	// 播放音频并在播放完毕后通过回调发送信号

	speaker.Play(sequence)
	// 等待播放完毕
	<-done
	// 等待 speaker 处理完所有的音频数据

}
