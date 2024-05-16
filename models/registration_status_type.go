package models
type RegistrationStatusType int

const (
    REGISTERED_REGISTRATIONSTATUSTYPE RegistrationStatusType = iota
    ENABLED_REGISTRATIONSTATUSTYPE
    CAPABLE_REGISTRATIONSTATUSTYPE
    MFAREGISTERED_REGISTRATIONSTATUSTYPE
    UNKNOWNFUTUREVALUE_REGISTRATIONSTATUSTYPE
)

func (i RegistrationStatusType) String() string {
    return []string{"registered", "enabled", "capable", "mfaRegistered", "unknownFutureValue"}[i]
}
func ParseRegistrationStatusType(v string) (any, error) {
    result := REGISTERED_REGISTRATIONSTATUSTYPE
    switch v {
        case "registered":
            result = REGISTERED_REGISTRATIONSTATUSTYPE
        case "enabled":
            result = ENABLED_REGISTRATIONSTATUSTYPE
        case "capable":
            result = CAPABLE_REGISTRATIONSTATUSTYPE
        case "mfaRegistered":
            result = MFAREGISTERED_REGISTRATIONSTATUSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REGISTRATIONSTATUSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRegistrationStatusType(values []RegistrationStatusType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RegistrationStatusType) isMultiValue() bool {
    return false
}
