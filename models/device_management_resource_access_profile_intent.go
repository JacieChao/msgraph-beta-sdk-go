package models
// The administrator intent for the assignment of the profile.
type DeviceManagementResourceAccessProfileIntent int

const (
    // Apply the profile.
    APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT DeviceManagementResourceAccessProfileIntent = iota
    // Remove the profile from devices that have installed the profile.
    REMOVE_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
)

func (i DeviceManagementResourceAccessProfileIntent) String() string {
    return []string{"apply", "remove"}[i]
}
func ParseDeviceManagementResourceAccessProfileIntent(v string) (any, error) {
    result := APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
    switch v {
        case "apply":
            result = APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
        case "remove":
            result = REMOVE_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceManagementResourceAccessProfileIntent(values []DeviceManagementResourceAccessProfileIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceManagementResourceAccessProfileIntent) isMultiValue() bool {
    return false
}
