package graph
import (
    "strings"
    "errors"
)
type WorkforceIntegrationSupportedEntities int

const (
    NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES WorkforceIntegrationSupportedEntities = iota
    SHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    SWAPREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    USERSHIFTPREFERENCES_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OPENSHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OPENSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OFFERSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    TIMECARD_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    TIMEOFFREASON_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    TIMEOFF_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    TIMEOFFREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
)

func (i WorkforceIntegrationSupportedEntities) String() string {
    return []string{"NONE", "SHIFT", "SWAPREQUEST", "USERSHIFTPREFERENCES", "OPENSHIFT", "OPENSHIFTREQUEST", "OFFERSHIFTREQUEST", "UNKNOWNFUTUREVALUE", "TIMECARD", "TIMEOFFREASON", "TIMEOFF", "TIMEOFFREQUEST"}[i]
}
func ParseWorkforceIntegrationSupportedEntities(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "SHIFT":
            return SHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "SWAPREQUEST":
            return SWAPREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "USERSHIFTPREFERENCES":
            return USERSHIFTPREFERENCES_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "OPENSHIFT":
            return OPENSHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "OPENSHIFTREQUEST":
            return OPENSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "OFFERSHIFTREQUEST":
            return OFFERSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "TIMECARD":
            return TIMECARD_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "TIMEOFFREASON":
            return TIMEOFFREASON_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "TIMEOFF":
            return TIMEOFF_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
        case "TIMEOFFREQUEST":
            return TIMEOFFREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES, nil
    }
    return 0, errors.New("Unknown WorkforceIntegrationSupportedEntities value: " + v)
}
func SerializeWorkforceIntegrationSupportedEntities(values []WorkforceIntegrationSupportedEntities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}