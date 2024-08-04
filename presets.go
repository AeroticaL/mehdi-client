package azuretls

import (
	"math"

	"github.com/Noooste/fhttp/http2"
)

const (
	Chrome  = "chrome"
	Firefox = "firefox"
	Opera   = "opera"
	Safari  = "safari"
	Edge    = "edge"
	Ios     = "ios"
	Android = "android" //deprecated
)

// defaultHeaderSettings returns HTTP/2 settings for a given navigator with customizable values.
func defaultHeaderSettings(navigator string, customSettings map[http2.SettingID]uint32) (map[http2.SettingID]uint32, []http2.SettingID) {
	defaultSettings := map[http2.SettingID]uint32{
		http2.SettingHeaderTableSize:   65536,
		http2.SettingEnablePush:        0,
		http2.SettingInitialWindowSize: 6291456,
		http2.SettingMaxHeaderListSize: 262144,
	}

	switch navigator {
	case Firefox:
		defaultSettings = map[http2.SettingID]uint32{
			http2.SettingMaxFrameSize:      16384,
			http2.SettingInitialWindowSize: 131072,
			http2.SettingHeaderTableSize:   65536,
		}

	case Ios:
		defaultSettings = map[http2.SettingID]uint32{
			http2.SettingHeaderTableSize:      4096,
			http2.SettingMaxConcurrentStreams: 100,
			http2.SettingInitialWindowSize:    2097152,
			http2.SettingMaxFrameSize:         16384,
			http2.SettingMaxHeaderListSize:    math.MaxUint32,
		}
	}

	// Apply custom settings if provided
	for k, v := range customSettings {
		defaultSettings[k] = v
	}

	var ids []http2.SettingID
	for id := range defaultSettings {
		ids = append(ids, id)
	}

	return defaultSettings, ids
}

// defaultWindowsUpdate returns the default window update value for a given navigator.
func defaultWindowsUpdate(navigator string, customValue uint32) uint32 {
	switch navigator {
	case Firefox:
		return customValue
	case Ios:
		return customValue
	default:
		return customValue
	}
}

// defaultStreamPriorities returns the default stream priorities for a given navigator with customizable values.
func defaultStreamPriorities(navigator string, customPriorities []http2.Priority) []http2.Priority {
	defaultPriorities := []http2.Priority{}

	switch navigator {
	case Firefox:
		defaultPriorities = []http2.Priority{
			{
				StreamID: 3,
				PriorityParam: http2.PriorityParam{
					Weight: 200,
				},
			},
			{
				StreamID: 5,
				PriorityParam: http2.PriorityParam{
					Weight: 100,
				},
			},
			{
				StreamID: 7,
				PriorityParam: http2.PriorityParam{
					Weight: 0,
				},
			},
			{
				StreamID: 9,
				PriorityParam: http2.PriorityParam{
					Weight:    0,
					StreamDep: 7,
				},
			},
			{
				StreamID: 11,
				PriorityParam: http2.PriorityParam{
					Weight:    0,
					StreamDep: 3,
				},
			},
			{
				StreamID: 13,
				PriorityParam: http2.PriorityParam{
					Weight: 240,
				},
			},
		}
	}

	if len(customPriorities) > 0 {
		return customPriorities
	}

	return defaultPriorities
}

// defaultHeaderPriorities returns the default header priority parameters for a given navigator with customizable values.
func defaultHeaderPriorities(navigator string, customPriority *http2.PriorityParam) *http2.PriorityParam {
	defaultPriority := &http2.PriorityParam{
		Weight:    255,
		StreamDep: 0,
		Exclusive: true,
	}

	switch navigator {
	case Firefox:
		defaultPriority = &http2.PriorityParam{
			Weight:    41,
			StreamDep: 13,
			Exclusive: false,
		}
	}

	if customPriority != nil {
		return customPriority
	}

	return defaultPriority
}
