package models
// Android Device Owner Play Store mode type.
type AndroidDeviceOwnerPlayStoreMode int

const (
    // Not Configured
    NOTCONFIGURED_ANDROIDDEVICEOWNERPLAYSTOREMODE AndroidDeviceOwnerPlayStoreMode = iota
    // Only apps that are in the policy are available and any app not in the policy will be automatically uninstalled from the device.
    ALLOWLIST_ANDROIDDEVICEOWNERPLAYSTOREMODE
    // All apps are available and any app that should not be on the device should be explicitly marked as 'BLOCKED' in the applications policy.
    BLOCKLIST_ANDROIDDEVICEOWNERPLAYSTOREMODE
)

func (i AndroidDeviceOwnerPlayStoreMode) String() string {
    return []string{"notConfigured", "allowList", "blockList"}[i]
}
func ParseAndroidDeviceOwnerPlayStoreMode(v string) (any, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERPLAYSTOREMODE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERPLAYSTOREMODE
        case "allowList":
            result = ALLOWLIST_ANDROIDDEVICEOWNERPLAYSTOREMODE
        case "blockList":
            result = BLOCKLIST_ANDROIDDEVICEOWNERPLAYSTOREMODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerPlayStoreMode(values []AndroidDeviceOwnerPlayStoreMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerPlayStoreMode) isMultiValue() bool {
    return false
}
