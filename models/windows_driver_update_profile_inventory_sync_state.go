package models
// Windows DnF update inventory sync state.
type WindowsDriverUpdateProfileInventorySyncState int

const (
    // Pending sync.
    PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE WindowsDriverUpdateProfileInventorySyncState = iota
    // Successful sync.
    SUCCESS_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
    // Failed sync.
    FAILURE_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
)

func (i WindowsDriverUpdateProfileInventorySyncState) String() string {
    return []string{"pending", "success", "failure"}[i]
}
func ParseWindowsDriverUpdateProfileInventorySyncState(v string) (any, error) {
    result := PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
    switch v {
        case "pending":
            result = PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        case "success":
            result = SUCCESS_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        case "failure":
            result = FAILURE_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsDriverUpdateProfileInventorySyncState(values []WindowsDriverUpdateProfileInventorySyncState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsDriverUpdateProfileInventorySyncState) isMultiValue() bool {
    return false
}
