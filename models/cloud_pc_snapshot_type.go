package models
type CloudPcSnapshotType int

const (
    AUTOMATIC_CLOUDPCSNAPSHOTTYPE CloudPcSnapshotType = iota
    MANUAL_CLOUDPCSNAPSHOTTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCSNAPSHOTTYPE
)

func (i CloudPcSnapshotType) String() string {
    return []string{"automatic", "manual", "unknownFutureValue"}[i]
}
func ParseCloudPcSnapshotType(v string) (any, error) {
    result := AUTOMATIC_CLOUDPCSNAPSHOTTYPE
    switch v {
        case "automatic":
            result = AUTOMATIC_CLOUDPCSNAPSHOTTYPE
        case "manual":
            result = MANUAL_CLOUDPCSNAPSHOTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCSNAPSHOTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcSnapshotType(values []CloudPcSnapshotType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcSnapshotType) isMultiValue() bool {
    return false
}
