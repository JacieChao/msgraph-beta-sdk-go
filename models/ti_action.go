package models
type TiAction int

const (
    UNKNOWN_TIACTION TiAction = iota
    ALLOW_TIACTION
    BLOCK_TIACTION
    ALERT_TIACTION
    UNKNOWNFUTUREVALUE_TIACTION
)

func (i TiAction) String() string {
    return []string{"unknown", "allow", "block", "alert", "unknownFutureValue"}[i]
}
func ParseTiAction(v string) (any, error) {
    result := UNKNOWN_TIACTION
    switch v {
        case "unknown":
            result = UNKNOWN_TIACTION
        case "allow":
            result = ALLOW_TIACTION
        case "block":
            result = BLOCK_TIACTION
        case "alert":
            result = ALERT_TIACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TIACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTiAction(values []TiAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TiAction) isMultiValue() bool {
    return false
}
