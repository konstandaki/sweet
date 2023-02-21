package device

import (
	"github.com/hashicorp/go-version"
	"gitlab.sweet.tv/proto/application"
)

func IsTimeshiftSupported(dev *DeviceInfo) bool {
	var supported = false

	v, err := version.NewVersion(dev.GetFirmware().GetVersionString())
	if err != nil {
		return false
	}

	switch dev.GetType() {
	case DeviceInfo_DT_AndroidTV, DeviceInfo_DT_iNext:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.6.1"))) {
				supported = true
			}
			break
		case application.ApplicationInfo_AT_TRINITY_Player:
			if v.GreaterThan(version.Must(version.NewVersion("3.2.1"))) {
				supported = true
			}
			break
		case application.ApplicationInfo_AT_GRIZLI_Player:
			if v.GreaterThan(version.Must(version.NewVersion("2.1.9"))) {
				supported = true
			}
			break
		}
		break

	case DeviceInfo_DT_SmartTV:
		supported = true
		if v.Equal(version.Must(version.NewVersion("2.2.0"))) {
			supported = false
		}
		break
	case DeviceInfo_DT_Android_Player:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.3.8"))) {
				supported = true
			}
			break
		case application.ApplicationInfo_AT_TRINITY_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("3.3.1"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_IOS_Player:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.8.5"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_AppleTV:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.8.6"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_Web_Browser:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.7.40"))) {
				supported = true
			}
			break
		case application.ApplicationInfo_AT_TRINITY_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.0.0"))) {
				supported = true
			}
			break
		}
		break
	case DeviceInfo_DT_Xbox:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.1.27"))) {
				supported = true
			}
			break
		}
		break
	}

	return supported
}
