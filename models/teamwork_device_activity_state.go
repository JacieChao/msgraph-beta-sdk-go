package models
type TeamworkDeviceActivityState int

const (
    UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE TeamworkDeviceActivityState = iota
    BUSY_TEAMWORKDEVICEACTIVITYSTATE
    IDLE_TEAMWORKDEVICEACTIVITYSTATE
    UNAVAILABLE_TEAMWORKDEVICEACTIVITYSTATE
    UNKNOWNFUTUREVALUE_TEAMWORKDEVICEACTIVITYSTATE
)

func (i TeamworkDeviceActivityState) String() string {
    return []string{"unknown", "busy", "idle", "unavailable", "unknownFutureValue"}[i]
}
func ParseTeamworkDeviceActivityState(v string) (any, error) {
    result := UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE
        case "busy":
            result = BUSY_TEAMWORKDEVICEACTIVITYSTATE
        case "idle":
            result = IDLE_TEAMWORKDEVICEACTIVITYSTATE
        case "unavailable":
            result = UNAVAILABLE_TEAMWORKDEVICEACTIVITYSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICEACTIVITYSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTeamworkDeviceActivityState(values []TeamworkDeviceActivityState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamworkDeviceActivityState) isMultiValue() bool {
    return false
}
