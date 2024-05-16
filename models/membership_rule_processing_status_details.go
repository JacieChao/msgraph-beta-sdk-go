package models
type MembershipRuleProcessingStatusDetails int

const (
    NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS MembershipRuleProcessingStatusDetails = iota
    RUNNING_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    FAILED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    SUCCEEDED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    UNSUPPORTEDFUTUREVALUE_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
)

func (i MembershipRuleProcessingStatusDetails) String() string {
    return []string{"NotStarted", "Running", "Failed", "Succeeded", "UnsupportedFutureValue"}[i]
}
func ParseMembershipRuleProcessingStatusDetails(v string) (any, error) {
    result := NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    switch v {
        case "NotStarted":
            result = NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "Running":
            result = RUNNING_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "Failed":
            result = FAILED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "Succeeded":
            result = SUCCEEDED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "UnsupportedFutureValue":
            result = UNSUPPORTEDFUTUREVALUE_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMembershipRuleProcessingStatusDetails(values []MembershipRuleProcessingStatusDetails) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MembershipRuleProcessingStatusDetails) isMultiValue() bool {
    return false
}
