package models
type AuthenticationMethodPlatform int

const (
    UNKNOWN_AUTHENTICATIONMETHODPLATFORM AuthenticationMethodPlatform = iota
    WINDOWS_AUTHENTICATIONMETHODPLATFORM
    MACOS_AUTHENTICATIONMETHODPLATFORM
    IOS_AUTHENTICATIONMETHODPLATFORM
    ANDROID_AUTHENTICATIONMETHODPLATFORM
    LINUX_AUTHENTICATIONMETHODPLATFORM
    UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODPLATFORM
)

func (i AuthenticationMethodPlatform) String() string {
    return []string{"unknown", "windows", "macOS", "iOS", "android", "linux", "unknownFutureValue"}[i]
}
func ParseAuthenticationMethodPlatform(v string) (any, error) {
    result := UNKNOWN_AUTHENTICATIONMETHODPLATFORM
    switch v {
        case "unknown":
            result = UNKNOWN_AUTHENTICATIONMETHODPLATFORM
        case "windows":
            result = WINDOWS_AUTHENTICATIONMETHODPLATFORM
        case "macOS":
            result = MACOS_AUTHENTICATIONMETHODPLATFORM
        case "iOS":
            result = IOS_AUTHENTICATIONMETHODPLATFORM
        case "android":
            result = ANDROID_AUTHENTICATIONMETHODPLATFORM
        case "linux":
            result = LINUX_AUTHENTICATIONMETHODPLATFORM
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODPLATFORM
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthenticationMethodPlatform(values []AuthenticationMethodPlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationMethodPlatform) isMultiValue() bool {
    return false
}
