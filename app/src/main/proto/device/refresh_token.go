package device

import (
	"github.com/hashicorp/go-version"
	"gitlab.sweet.tv/proto/application"
)

func IsRefreshTokenSupported(dev *DeviceInfo) bool {
	var supported = false

	v, err := version.NewVersion(dev.GetFirmware().GetVersionString())
	if err != nil {
		return false
	}

	switch dev.GetType() {
	case DeviceInfo_DT_AndroidTV, DeviceInfo_DT_iNext:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.2.1"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_GRIZLI_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.0.6"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_TRINITY_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("3.0.8"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_Bestlink, application.ApplicationInfo_AT_Streamnetwork,
			application.ApplicationInfo_AT_Homenet, application.ApplicationInfo_AT_IMPERIAL_TV:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("3.0.8"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_DiaNet:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("3.0.9"))) {
				supported = true
			}
		}
	case DeviceInfo_DT_SmartTV:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player, application.ApplicationInfo_AT_TRINITY_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.2.0"))) {
				supported = true
			}
		}
	case DeviceInfo_DT_Android_Player:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("2.1.0"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_TRINITY_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("3.1.8"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_GRIZLI_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.0.4"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_Apelsin:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.0.2"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_DiaNet, application.ApplicationInfo_AT_IMPERIAL_TV:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.0.6"))) {
				supported = true
			}
		}
	case DeviceInfo_DT_IOS_Player:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.7.1"))) {
				supported = true
			}
		}
	case DeviceInfo_DT_AppleTV:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.7.6"))) {
				supported = true
			}
		}
	case DeviceInfo_DT_Web_Browser:
		switch dev.GetApplication().GetType() {
		case application.ApplicationInfo_AT_SWEET_TV_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.6.0"))) {
				supported = true
			}
		case application.ApplicationInfo_AT_TRINITY_Player:
			if v.GreaterThanOrEqual(version.Must(version.NewVersion("1.0.3"))) {
				supported = true
			}
		}
	case DeviceInfo_DT_Xbox:
		supported = true
	}

	return supported
}
