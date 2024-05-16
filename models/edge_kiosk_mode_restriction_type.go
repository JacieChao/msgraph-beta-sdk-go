package models
// Specify how the Microsoft Edge settings are restricted based on kiosk mode.
type EdgeKioskModeRestrictionType int

const (
    // Not configured (unrestricted).
    NOTCONFIGURED_EDGEKIOSKMODERESTRICTIONTYPE EdgeKioskModeRestrictionType = iota
    // Interactive/Digital signage in single-app mode.
    DIGITALSIGNAGE_EDGEKIOSKMODERESTRICTIONTYPE
    // Normal mode (full version of Microsoft Edge).
    NORMALMODE_EDGEKIOSKMODERESTRICTIONTYPE
    // Public browsing in single-app mode.
    PUBLICBROWSINGSINGLEAPP_EDGEKIOSKMODERESTRICTIONTYPE
    // Public browsing (inPrivate) in multi-app mode.
    PUBLICBROWSINGMULTIAPP_EDGEKIOSKMODERESTRICTIONTYPE
)

func (i EdgeKioskModeRestrictionType) String() string {
    return []string{"notConfigured", "digitalSignage", "normalMode", "publicBrowsingSingleApp", "publicBrowsingMultiApp"}[i]
}
func ParseEdgeKioskModeRestrictionType(v string) (any, error) {
    result := NOTCONFIGURED_EDGEKIOSKMODERESTRICTIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_EDGEKIOSKMODERESTRICTIONTYPE
        case "digitalSignage":
            result = DIGITALSIGNAGE_EDGEKIOSKMODERESTRICTIONTYPE
        case "normalMode":
            result = NORMALMODE_EDGEKIOSKMODERESTRICTIONTYPE
        case "publicBrowsingSingleApp":
            result = PUBLICBROWSINGSINGLEAPP_EDGEKIOSKMODERESTRICTIONTYPE
        case "publicBrowsingMultiApp":
            result = PUBLICBROWSINGMULTIAPP_EDGEKIOSKMODERESTRICTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEdgeKioskModeRestrictionType(values []EdgeKioskModeRestrictionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EdgeKioskModeRestrictionType) isMultiValue() bool {
    return false
}
