package device

import (
	"gitlab.sweet.tv/proto/application"
	"google.golang.org/protobuf/proto"
	"testing"
)

func Test_IsRefreshTokenSupported(t *testing.T) {
	type args struct {
		dev *DeviceInfo
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"DT_AndroidTV 2.5.6 AT_SWEET_TV_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AndroidTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("2.5.6"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_SWEET_TV_Player,
					},
				},
			},
			true,
		},
		{
			"DT_iNext 2.5.6 AT_SWEET_TV_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_iNext.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("2.5.6"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_SWEET_TV_Player,
					},
				},
			},
			true,
		},
		{
			"DT_AndroidTV 3.2.1 AT_TRINITY_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AndroidTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("3.2.1"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_TRINITY_Player,
					},
				},
			},
			true,
		},
		{
			"DT_iNext 3.2.1 AT_TRINITY_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_iNext.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("3.2.1"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_TRINITY_Player,
					},
				},
			},
			true,
		},
		{
			"DT_AndroidTV 2.1.9 AT_GRIZLI_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AndroidTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("2.1.9"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_GRIZLI_Player,
					},
				},
			},
			true,
		},
		{
			"DT_iNext 2.1.9 AT_GRIZLI_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_iNext.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("2.1.9"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_GRIZLI_Player,
					},
				},
			},
			true,
		},
		{
			"DT_AndroidTV 3.0.8 AT_Bestlink",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AndroidTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("3.0.8"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_Bestlink,
					},
				},
			},
			true,
		},
		{
			"DT_AndroidTV 3.0.8 Streamnetwork",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AndroidTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("3.0.8"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_Streamnetwork,
					},
				},
			},
			true,
		},
		{
			"DT_AndroidTV 3.0.8 AT_Homenet",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AndroidTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("3.0.8"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_Homenet,
					},
				},
			},
			true,
		},
		{
			"DT_AndroidTV 3.0.8 AT_IMPERIAL_TV",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AndroidTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("3.0.8"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_IMPERIAL_TV,
					},
				},
			},
			true,
		},
		{
			"DT_AndroidTV 3.0.9 AT_DiaNet",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AndroidTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("3.0.9"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_DiaNet,
					},
				},
			},
			true,
		},
		{
			"DT_Android_Player 1.0.0",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Android_Player.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.0.0"),
					},
				},
			},
			false,
		},
		{
			"DT_Android_Player 1.0.0 AT_Apelsin",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Android_Player.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.0.0"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_Apelsin,
					},
				},
			},
			false,
		},
		{
			"DT_Android_Player 2.1.0 AT_SWEET_TV_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Android_Player.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("2.1.0"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_SWEET_TV_Player,
					},
				},
			},
			true,
		},
		{
			"DT_Android_Player 2.0.0 AT_SWEET_TV_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Android_Player.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("2.0.0"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_SWEET_TV_Player,
					},
				},
			},
			false,
		},
		{
			"DT_Web_Browser 1.6.0 AT_SWEET_TV_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Web_Browser.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.6.0"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_SWEET_TV_Player,
					},
				},
			},
			true,
		},
		{
			"DT_Web_Browser 1.6.0 AT_Kahovka",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Web_Browser.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.6.0"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_Kahovka,
					},
				},
			},
			false,
		},
		{
			"DT_Web_Browser 1.6.0",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Web_Browser.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.6.0"),
					},
				},
			},
			false,
		},
		{
			"DT_Web_Browser 1.5.9",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Web_Browser.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.5.9"),
					},
				},
			},
			false,
		},
		{
			"DT_AppleTV 1.6.0 AT_SWEET_TV_Player",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AppleTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.6.0"),
					},
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_SWEET_TV_Player,
					},
				},
			},
			false,
		},
		{
			"DT_AppleTV 1.6.0",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_AppleTV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.6.0"),
					},
				},
			},
			false,
		},
		{
			"DT_Kivi_TV 1.6.0",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Kivi_TV.Enum(),
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.6.0"),
					},
				},
			},
			false,
		},
		{
			"DT_SmartTV",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_SmartTV.Enum(),
				},
			},
			false,
		},
		{
			"DT_Xbox",
			args{
				dev: &DeviceInfo{
					Type: DeviceInfo_DT_Xbox.Enum(),
				},
			},
			false,
		},
		{
			name: "DT_Android_Player 1.0.6 AT_DiaNet",
			args: args{
				dev: &DeviceInfo{
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_DiaNet,
					},
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.0.6"),
					},
					Type: DeviceInfo_DT_Android_Player.Enum(),
				},
			},
			want: true,
		},
		{
			name: "DT_Android_Player 1.0.5 AT_DiaNet",
			args: args{
				dev: &DeviceInfo{
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_DiaNet,
					},
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.0.5"),
					},
					Type: DeviceInfo_DT_Android_Player.Enum(),
				},
			},
			want: false,
		},
		{
			name: "DT_Android_Player 1.0.6 AT_IMPERIAL_TV",
			args: args{
				dev: &DeviceInfo{
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_IMPERIAL_TV,
					},
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.0.6"),
					},
					Type: DeviceInfo_DT_Android_Player.Enum(),
				},
			},
			want: true,
		},
		{
			name: "DT_Android_Player 1.0.5 AT_IMPERIAL_TV",
			args: args{
				dev: &DeviceInfo{
					Application: &application.ApplicationInfo{
						Type: application.ApplicationInfo_AT_IMPERIAL_TV,
					},
					Firmware: &DeviceInfo_FirmwareInfo{
						VersionString: proto.String("1.0.5"),
					},
					Type: DeviceInfo_DT_Android_Player.Enum(),
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRefreshTokenSupported(tt.args.dev); got != tt.want {
				t.Errorf("IsRefreshTokenSupported() = %v, want %v", got, tt.want)
			}
		})
	}
}
