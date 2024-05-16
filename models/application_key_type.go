package models
type ApplicationKeyType int

const (
    CLIENTSECRET_APPLICATIONKEYTYPE ApplicationKeyType = iota
    CERTIFICATE_APPLICATIONKEYTYPE
    UNKNOWNFUTUREVALUE_APPLICATIONKEYTYPE
)

func (i ApplicationKeyType) String() string {
    return []string{"clientSecret", "certificate", "unknownFutureValue"}[i]
}
func ParseApplicationKeyType(v string) (any, error) {
    result := CLIENTSECRET_APPLICATIONKEYTYPE
    switch v {
        case "clientSecret":
            result = CLIENTSECRET_APPLICATIONKEYTYPE
        case "certificate":
            result = CERTIFICATE_APPLICATIONKEYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPLICATIONKEYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeApplicationKeyType(values []ApplicationKeyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApplicationKeyType) isMultiValue() bool {
    return false
}
