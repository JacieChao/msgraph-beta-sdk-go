package models
type AuthenticationContextDetail int

const (
    REQUIRED_AUTHENTICATIONCONTEXTDETAIL AuthenticationContextDetail = iota
    PREVIOUSLYSATISFIED_AUTHENTICATIONCONTEXTDETAIL
    NOTAPPLICABLE_AUTHENTICATIONCONTEXTDETAIL
    UNKNOWNFUTUREVALUE_AUTHENTICATIONCONTEXTDETAIL
)

func (i AuthenticationContextDetail) String() string {
    return []string{"required", "previouslySatisfied", "notApplicable", "unknownFutureValue"}[i]
}
func ParseAuthenticationContextDetail(v string) (any, error) {
    result := REQUIRED_AUTHENTICATIONCONTEXTDETAIL
    switch v {
        case "required":
            result = REQUIRED_AUTHENTICATIONCONTEXTDETAIL
        case "previouslySatisfied":
            result = PREVIOUSLYSATISFIED_AUTHENTICATIONCONTEXTDETAIL
        case "notApplicable":
            result = NOTAPPLICABLE_AUTHENTICATIONCONTEXTDETAIL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONCONTEXTDETAIL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthenticationContextDetail(values []AuthenticationContextDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationContextDetail) isMultiValue() bool {
    return false
}
