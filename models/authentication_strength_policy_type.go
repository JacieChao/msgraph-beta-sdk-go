package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type AuthenticationStrengthPolicyType int

const (
    BUILTIN_AUTHENTICATIONSTRENGTHPOLICYTYPE AuthenticationStrengthPolicyType = iota
    CUSTOM_AUTHENTICATIONSTRENGTHPOLICYTYPE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHPOLICYTYPE
)

func (i AuthenticationStrengthPolicyType) String() string {
    return []string{"builtIn", "custom", "unknownFutureValue"}[i]
}
func ParseAuthenticationStrengthPolicyType(v string) (interface{}, error) {
    result := BUILTIN_AUTHENTICATIONSTRENGTHPOLICYTYPE
    switch v {
        case "builtIn":
            result = BUILTIN_AUTHENTICATIONSTRENGTHPOLICYTYPE
        case "custom":
            result = CUSTOM_AUTHENTICATIONSTRENGTHPOLICYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHPOLICYTYPE
        default:
            return 0, errors.New("Unknown AuthenticationStrengthPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationStrengthPolicyType(values []AuthenticationStrengthPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}