package models
type ApprovalState int

const (
    PENDING_APPROVALSTATE ApprovalState = iota
    APPROVED_APPROVALSTATE
    DENIED_APPROVALSTATE
    ABORTED_APPROVALSTATE
    CANCELED_APPROVALSTATE
)

func (i ApprovalState) String() string {
    return []string{"pending", "approved", "denied", "aborted", "canceled"}[i]
}
func ParseApprovalState(v string) (any, error) {
    result := PENDING_APPROVALSTATE
    switch v {
        case "pending":
            result = PENDING_APPROVALSTATE
        case "approved":
            result = APPROVED_APPROVALSTATE
        case "denied":
            result = DENIED_APPROVALSTATE
        case "aborted":
            result = ABORTED_APPROVALSTATE
        case "canceled":
            result = CANCELED_APPROVALSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeApprovalState(values []ApprovalState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApprovalState) isMultiValue() bool {
    return false
}
