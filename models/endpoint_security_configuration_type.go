package models
// The endpoint security policy type.
type EndpointSecurityConfigurationType int

const (
    // Unknown.
    UNKNOWN_ENDPOINTSECURITYCONFIGURATIONTYPE EndpointSecurityConfigurationType = iota
    // Antivirus.
    ANTIVIRUS_ENDPOINTSECURITYCONFIGURATIONTYPE
    // Disk encryption.
    DISKENCRYPTION_ENDPOINTSECURITYCONFIGURATIONTYPE
    // Firewall.
    FIREWALL_ENDPOINTSECURITYCONFIGURATIONTYPE
    // Endpoint detection and response.
    ENDPOINTDETECTIONANDRESPONSE_ENDPOINTSECURITYCONFIGURATIONTYPE
    // Attack surface reduction.
    ATTACKSURFACEREDUCTION_ENDPOINTSECURITYCONFIGURATIONTYPE
    // Account protection.
    ACCOUNTPROTECTION_ENDPOINTSECURITYCONFIGURATIONTYPE
)

func (i EndpointSecurityConfigurationType) String() string {
    return []string{"unknown", "antivirus", "diskEncryption", "firewall", "endpointDetectionAndResponse", "attackSurfaceReduction", "accountProtection"}[i]
}
func ParseEndpointSecurityConfigurationType(v string) (any, error) {
    result := UNKNOWN_ENDPOINTSECURITYCONFIGURATIONTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_ENDPOINTSECURITYCONFIGURATIONTYPE
        case "antivirus":
            result = ANTIVIRUS_ENDPOINTSECURITYCONFIGURATIONTYPE
        case "diskEncryption":
            result = DISKENCRYPTION_ENDPOINTSECURITYCONFIGURATIONTYPE
        case "firewall":
            result = FIREWALL_ENDPOINTSECURITYCONFIGURATIONTYPE
        case "endpointDetectionAndResponse":
            result = ENDPOINTDETECTIONANDRESPONSE_ENDPOINTSECURITYCONFIGURATIONTYPE
        case "attackSurfaceReduction":
            result = ATTACKSURFACEREDUCTION_ENDPOINTSECURITYCONFIGURATIONTYPE
        case "accountProtection":
            result = ACCOUNTPROTECTION_ENDPOINTSECURITYCONFIGURATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEndpointSecurityConfigurationType(values []EndpointSecurityConfigurationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EndpointSecurityConfigurationType) isMultiValue() bool {
    return false
}
