package models
// Authoring source of a policy
type SecurityBaselinePolicySourceType int

const (
    DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE SecurityBaselinePolicySourceType = iota
    DEVICEINTENT_SECURITYBASELINEPOLICYSOURCETYPE
)

func (i SecurityBaselinePolicySourceType) String() string {
    return []string{"deviceConfiguration", "deviceIntent"}[i]
}
func ParseSecurityBaselinePolicySourceType(v string) (any, error) {
    result := DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE
    switch v {
        case "deviceConfiguration":
            result = DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE
        case "deviceIntent":
            result = DEVICEINTENT_SECURITYBASELINEPOLICYSOURCETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityBaselinePolicySourceType(values []SecurityBaselinePolicySourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityBaselinePolicySourceType) isMultiValue() bool {
    return false
}
