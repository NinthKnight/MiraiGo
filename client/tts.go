package client

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/binary"
	"github.com/Mrs4s/MiraiGo/client/pb/richmedia"
	"github.com/Mrs4s/MiraiGo/utils"
	"github.com/golang/protobuf/proto"
)

func (c *QQClient) GetTts(text string) ([]byte, error) {
	url := "https://textts.qq.com/cgi-bin/tts"
	data := "{\"appid\": \"201908021016\",\"text\": \"" + text + "\"}"
	rsp, err := utils.HttpPostBytesWithCookie(url, []byte(data), c.getCookies())
	if err != nil {
		return nil, err
	}
	ttsReader := binary.NewReader(rsp)
	ttsWriter := binary.NewWriter()
	for {
		// 数据格式 69e(字符串)  十六进制   数据长度  0 为结尾
		// 0D 0A (分隔符) payload  0D 0A
		var dataLen []byte
		for b := ttsReader.ReadByte(); b != byte(0x0d); b = ttsReader.ReadByte() {
			dataLen = append(dataLen, b)
		}
		ttsReader.ReadByte()
		var length int
		_, _ = fmt.Sscan("0x"+string(dataLen), &length)
		if length == 0 {
			break
		}
		ttsRsp := &richmedia.TtsRspBody{}
		err := proto.Unmarshal(ttsReader.ReadBytes(length), ttsRsp)
		if err != nil {
			return nil, err
		}
		for _, voiceItem := range ttsRsp.VoiceData {
			ttsWriter.Write(voiceItem.Voice)
		}
		ttsReader.ReadBytes(2)
	}
	return ttsWriter.Bytes(), nil
}
