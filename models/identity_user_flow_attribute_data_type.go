package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type IdentityUserFlowAttributeDataType int

const (
    STRING_IDENTITYUSERFLOWATTRIBUTEDATATYPE IdentityUserFlowAttributeDataType = iota
    BOOLEAN_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    INT64_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    STRINGCOLLECTION_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    DATETIME_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    UNKNOWNFUTUREVALUE_IDENTITYUSERFLOWATTRIBUTEDATATYPE
)

func (i IdentityUserFlowAttributeDataType) String() string {
    return []string{"string", "boolean", "int64", "stringCollection", "dateTime", "unknownFutureValue"}[i]
}
func ParseIdentityUserFlowAttributeDataType(v string) (interface{}, error) {
    result := STRING_IDENTITYUSERFLOWATTRIBUTEDATATYPE
    switch v {
        case "string":
            result = STRING_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "boolean":
            result = BOOLEAN_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "int64":
            result = INT64_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "stringCollection":
            result = STRINGCOLLECTION_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "dateTime":
            result = DATETIME_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IDENTITYUSERFLOWATTRIBUTEDATATYPE
        default:
            return 0, errors.New("Unknown IdentityUserFlowAttributeDataType value: " + v)
    }
    return &result, nil
}
func SerializeIdentityUserFlowAttributeDataType(values []IdentityUserFlowAttributeDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
