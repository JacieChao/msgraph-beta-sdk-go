package windowsupdates
type CveSeverityLevel int

const (
    CRITICAL_CVESEVERITYLEVEL CveSeverityLevel = iota
    IMPORTANT_CVESEVERITYLEVEL
    MODERATE_CVESEVERITYLEVEL
    UNKNOWNFUTUREVALUE_CVESEVERITYLEVEL
)

func (i CveSeverityLevel) String() string {
    return []string{"critical", "important", "moderate", "unknownFutureValue"}[i]
}
func ParseCveSeverityLevel(v string) (any, error) {
    result := CRITICAL_CVESEVERITYLEVEL
    switch v {
        case "critical":
            result = CRITICAL_CVESEVERITYLEVEL
        case "important":
            result = IMPORTANT_CVESEVERITYLEVEL
        case "moderate":
            result = MODERATE_CVESEVERITYLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CVESEVERITYLEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCveSeverityLevel(values []CveSeverityLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CveSeverityLevel) isMultiValue() bool {
    return false
}
