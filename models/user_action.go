package models
type UserAction int

const (
    REGISTERSECURITYINFORMATION_USERACTION UserAction = iota
    REGISTERORJOINDEVICES_USERACTION
    UNKNOWNFUTUREVALUE_USERACTION
)

func (i UserAction) String() string {
    return []string{"registerSecurityInformation", "registerOrJoinDevices", "unknownFutureValue"}[i]
}
func ParseUserAction(v string) (any, error) {
    result := REGISTERSECURITYINFORMATION_USERACTION
    switch v {
        case "registerSecurityInformation":
            result = REGISTERSECURITYINFORMATION_USERACTION
        case "registerOrJoinDevices":
            result = REGISTERORJOINDEVICES_USERACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USERACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUserAction(values []UserAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserAction) isMultiValue() bool {
    return false
}
