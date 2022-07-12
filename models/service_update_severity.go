package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ServiceUpdateSeverity int

const (
    NORMAL_SERVICEUPDATESEVERITY ServiceUpdateSeverity = iota
    HIGH_SERVICEUPDATESEVERITY
    CRITICAL_SERVICEUPDATESEVERITY
    UNKNOWNFUTUREVALUE_SERVICEUPDATESEVERITY
)

func (i ServiceUpdateSeverity) String() string {
    return []string{"normal", "high", "critical", "unknownFutureValue"}[i]
}
func ParseServiceUpdateSeverity(v string) (interface{}, error) {
    result := NORMAL_SERVICEUPDATESEVERITY
    switch v {
        case "normal":
            result = NORMAL_SERVICEUPDATESEVERITY
        case "high":
            result = HIGH_SERVICEUPDATESEVERITY
        case "critical":
            result = CRITICAL_SERVICEUPDATESEVERITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SERVICEUPDATESEVERITY
        default:
            return 0, errors.New("Unknown ServiceUpdateSeverity value: " + v)
    }
    return &result, nil
}
func SerializeServiceUpdateSeverity(values []ServiceUpdateSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
