package models
// ManagedInstallerStatus
type ManagedInstallerStatus int

const (
    // Managed Installer is Disabled
    DISABLED_MANAGEDINSTALLERSTATUS ManagedInstallerStatus = iota
    // Managed Installer is Enabled
    ENABLED_MANAGEDINSTALLERSTATUS
)

func (i ManagedInstallerStatus) String() string {
    return []string{"disabled", "enabled"}[i]
}
func ParseManagedInstallerStatus(v string) (any, error) {
    result := DISABLED_MANAGEDINSTALLERSTATUS
    switch v {
        case "disabled":
            result = DISABLED_MANAGEDINSTALLERSTATUS
        case "enabled":
            result = ENABLED_MANAGEDINSTALLERSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeManagedInstallerStatus(values []ManagedInstallerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManagedInstallerStatus) isMultiValue() bool {
    return false
}
