package models
type ActionState int

const (
    NONE_ACTIONSTATE ActionState = iota
    PENDING_ACTIONSTATE
    CANCELED_ACTIONSTATE
    ACTIVE_ACTIONSTATE
    DONE_ACTIONSTATE
    FAILED_ACTIONSTATE
    NOTSUPPORTED_ACTIONSTATE
)

func (i ActionState) String() string {
    return []string{"none", "pending", "canceled", "active", "done", "failed", "notSupported"}[i]
}
func ParseActionState(v string) (any, error) {
    result := NONE_ACTIONSTATE
    switch v {
        case "none":
            result = NONE_ACTIONSTATE
        case "pending":
            result = PENDING_ACTIONSTATE
        case "canceled":
            result = CANCELED_ACTIONSTATE
        case "active":
            result = ACTIVE_ACTIONSTATE
        case "done":
            result = DONE_ACTIONSTATE
        case "failed":
            result = FAILED_ACTIONSTATE
        case "notSupported":
            result = NOTSUPPORTED_ACTIONSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeActionState(values []ActionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ActionState) isMultiValue() bool {
    return false
}
