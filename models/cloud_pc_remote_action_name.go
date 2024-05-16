package models
type CloudPcRemoteActionName int

const (
    UNKNOWN_CLOUDPCREMOTEACTIONNAME CloudPcRemoteActionName = iota
    RESTART_CLOUDPCREMOTEACTIONNAME
    RENAME_CLOUDPCREMOTEACTIONNAME
    RESIZE_CLOUDPCREMOTEACTIONNAME
    RESTORE_CLOUDPCREMOTEACTIONNAME
    REPROVISION_CLOUDPCREMOTEACTIONNAME
    CHANGEUSERACCOUNTTYPE_CLOUDPCREMOTEACTIONNAME
    TROUBLESHOOT_CLOUDPCREMOTEACTIONNAME
    PLACEUNDERREVIEW_CLOUDPCREMOTEACTIONNAME
    UNKNOWNFUTUREVALUE_CLOUDPCREMOTEACTIONNAME
    CREATESNAPSHOT_CLOUDPCREMOTEACTIONNAME
    POWERON_CLOUDPCREMOTEACTIONNAME
    POWEROFF_CLOUDPCREMOTEACTIONNAME
    MOVEREGION_CLOUDPCREMOTEACTIONNAME
)

func (i CloudPcRemoteActionName) String() string {
    return []string{"unknown", "restart", "rename", "resize", "restore", "reprovision", "changeUserAccountType", "troubleshoot", "placeUnderReview", "unknownFutureValue", "createSnapshot", "powerOn", "powerOff", "moveRegion"}[i]
}
func ParseCloudPcRemoteActionName(v string) (any, error) {
    result := UNKNOWN_CLOUDPCREMOTEACTIONNAME
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDPCREMOTEACTIONNAME
        case "restart":
            result = RESTART_CLOUDPCREMOTEACTIONNAME
        case "rename":
            result = RENAME_CLOUDPCREMOTEACTIONNAME
        case "resize":
            result = RESIZE_CLOUDPCREMOTEACTIONNAME
        case "restore":
            result = RESTORE_CLOUDPCREMOTEACTIONNAME
        case "reprovision":
            result = REPROVISION_CLOUDPCREMOTEACTIONNAME
        case "changeUserAccountType":
            result = CHANGEUSERACCOUNTTYPE_CLOUDPCREMOTEACTIONNAME
        case "troubleshoot":
            result = TROUBLESHOOT_CLOUDPCREMOTEACTIONNAME
        case "placeUnderReview":
            result = PLACEUNDERREVIEW_CLOUDPCREMOTEACTIONNAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCREMOTEACTIONNAME
        case "createSnapshot":
            result = CREATESNAPSHOT_CLOUDPCREMOTEACTIONNAME
        case "powerOn":
            result = POWERON_CLOUDPCREMOTEACTIONNAME
        case "powerOff":
            result = POWEROFF_CLOUDPCREMOTEACTIONNAME
        case "moveRegion":
            result = MOVEREGION_CLOUDPCREMOTEACTIONNAME
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcRemoteActionName(values []CloudPcRemoteActionName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcRemoteActionName) isMultiValue() bool {
    return false
}
