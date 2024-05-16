package models
type WindowsDeviceUsageType int

const (
    // Single User Device Type
    SINGLEUSER_WINDOWSDEVICEUSAGETYPE WindowsDeviceUsageType = iota
    // Shared Device Type
    SHARED_WINDOWSDEVICEUSAGETYPE
)

func (i WindowsDeviceUsageType) String() string {
    return []string{"singleUser", "shared"}[i]
}
func ParseWindowsDeviceUsageType(v string) (any, error) {
    result := SINGLEUSER_WINDOWSDEVICEUSAGETYPE
    switch v {
        case "singleUser":
            result = SINGLEUSER_WINDOWSDEVICEUSAGETYPE
        case "shared":
            result = SHARED_WINDOWSDEVICEUSAGETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsDeviceUsageType(values []WindowsDeviceUsageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsDeviceUsageType) isMultiValue() bool {
    return false
}
