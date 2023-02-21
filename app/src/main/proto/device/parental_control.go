package device

import (
	"github.com/hashicorp/go-version"
	"gitlab.sweet.tv/proto/application"
)

func IsParentalControlSupported(dev *DeviceInfo) bool {
	var supported = false

	v, err := version.NewVersion(dev.GetFirmware().GetVersionString())
	if err != nil {
		return false
	}

	switch dev.GetType() {
	case DeviceInfo_DT_AndroidTV, DeviceInfo_DT_iNext:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.6.6"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_Android_Player:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.4.1"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_IOS_Player:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.8.9"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_SmartTV:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			switch dev.GetSubType() {
			case DeviceInfo_DST_SAMSUNG:
				if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.4.3"))) {
					supported = true
				}
				break
			case DeviceInfo_DST_LG:
				if v.Equal(version.Must(version.NewVersion("2.4.31"))) {
					supported = true
				}
				break
			case DeviceInfo_DST_SONY:
				if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.4.3"))) {
					supported = true
				}
				break
			case DeviceInfo_DST_HISENSE:
				if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.4.3"))) {
					supported = true
				}
				break
			}
			break
		}
		break
	case DeviceInfo_DT_AppleTV:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.9.0"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_Web_Browser:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.8.79"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_Xbox:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.2.17"))) {
				supported = true
			}
			break
		}
		break
	}

	return supported
}
