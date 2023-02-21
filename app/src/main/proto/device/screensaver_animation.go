package device

import (
	"github.com/hashicorp/go-version"
	"gitlab.sweet.tv/proto/application"
)

func IsScreensaverAnimationSupported(dev *DeviceInfo) bool {
	var supported = false

	v, err := version.NewVersion(dev.GetFirmware().GetVersionString())
	if err != nil {
		return false
	}
	switch dev.GetType() {
	case DeviceInfo_DT_AndroidTV, DeviceInfo_DT_iNext:
		if v.GreaterThan(version.Must(version.NewVersion("2.7.4"))) {
			supported = true
		}
		break
	case DeviceInfo_DT_Android_Player:
		supported = true
		break
	case DeviceInfo_DT_IOS_Player:
		supported = true
		break
	case DeviceInfo_DT_SmartTV:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			switch dev.GetSubType() {
			case DeviceInfo_DST_SAMSUNG:
				if dev.GetFirmware().GetVersionCode() == 25108 {
					supported = true
				}
				break
			case DeviceInfo_DST_LG:
				if dev.GetFirmware().GetVersionCode() == 24201 || dev.GetFirmware().GetVersionCode() == 24206 {
					supported = true
				}
				break
			case DeviceInfo_DST_HISENSE:
				if dev.GetFirmware().GetVersionCode() == 243101 || dev.GetFirmware().GetVersionCode() == 24306 {
					supported = true
				}
				break
			}
			break
		}
		break
	case DeviceInfo_DT_AppleTV:
		supported = true
		break
	case DeviceInfo_DT_Web_Browser:
		break
	case DeviceInfo_DT_Xbox:
		supported = true
		break
	}

	return supported
}
