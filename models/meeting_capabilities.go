package models
type MeetingCapabilities int

const (
    QUESTIONANDANSWER_MEETINGCAPABILITIES MeetingCapabilities = iota
    UNKNOWNFUTUREVALUE_MEETINGCAPABILITIES
)

func (i MeetingCapabilities) String() string {
    return []string{"questionAndAnswer", "unknownFutureValue"}[i]
}
func ParseMeetingCapabilities(v string) (any, error) {
    result := QUESTIONANDANSWER_MEETINGCAPABILITIES
    switch v {
        case "questionAndAnswer":
            result = QUESTIONANDANSWER_MEETINGCAPABILITIES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MEETINGCAPABILITIES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMeetingCapabilities(values []MeetingCapabilities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MeetingCapabilities) isMultiValue() bool {
    return false
}
