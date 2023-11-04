package connect

import (
	"modules/hellswitch"

	"github.com/herb-go/util"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var Switch *hellswitch.Hellswitch

func initSwitch() {
	Switch = hellswitch.New()
	Switch.OnSwitchStatusChange = OnSwitchStatusChange
	Switch.OnGlobalMessage = OnGlobalMessage
	Switch.Start()
}
func OnSwitchStatusChange(status int) {
	switch status {
	case hellswitch.StatusConnected:
		println("Connected to hellswitch network.")
	case hellswitch.StatusDisconnected:
		println("Disconnected from hellswitch network.")
	}
}
func Broadcast(msg []byte) {
	defer util.Recover()
	d := DecodeGBK(msg)
	if d != nil {
		Switch.Broadcast(d)
	}
}
func OnGlobalMessage(msg []byte) {
	defer util.Recover()
	data := EncodeGBK(msg)
	if data != nil {
		DefaultManager.Broadcast(data)
	}
}

func EncodeGBK(data []byte) []byte {
	encoder := simplifiedchinese.GBK.NewEncoder()
	d, err := encoder.Bytes(data)
	if err != nil {
		util.LogError(err)
		return nil
	}
	return d
}

func DecodeGBK(data []byte) []byte {
	decoder := simplifiedchinese.GBK.NewDecoder()
	d, err := decoder.Bytes(data)
	if err != nil {
		util.LogError(err)
		return nil
	}
	return d
}
