package models
type ApplicationMode int

const (
    MANUAL_APPLICATIONMODE ApplicationMode = iota
    AUTOMATIC_APPLICATIONMODE
    RECOMMENDED_APPLICATIONMODE
)

func (i ApplicationMode) String() string {
    return []string{"manual", "automatic", "recommended"}[i]
}
func ParseApplicationMode(v string) (any, error) {
    result := MANUAL_APPLICATIONMODE
    switch v {
        case "manual":
            result = MANUAL_APPLICATIONMODE
        case "automatic":
            result = AUTOMATIC_APPLICATIONMODE
        case "recommended":
            result = RECOMMENDED_APPLICATIONMODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeApplicationMode(values []ApplicationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApplicationMode) isMultiValue() bool {
    return false
}
