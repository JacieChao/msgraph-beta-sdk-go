package models
import (
    "errors"
    "math"
    "strings"
)
type ConditionalAccessTransferMethods int

const (
    NONE_CONDITIONALACCESSTRANSFERMETHODS = 1
    DEVICECODEFLOW_CONDITIONALACCESSTRANSFERMETHODS = 2
    AUTHENTICATIONTRANSFER_CONDITIONALACCESSTRANSFERMETHODS = 4
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSTRANSFERMETHODS = 8
)

func (i ConditionalAccessTransferMethods) String() string {
    var values []string
    options := []string{"none", "deviceCodeFlow", "authenticationTransfer", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := ConditionalAccessTransferMethods(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseConditionalAccessTransferMethods(v string) (any, error) {
    var result ConditionalAccessTransferMethods
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CONDITIONALACCESSTRANSFERMETHODS
            case "deviceCodeFlow":
                result |= DEVICECODEFLOW_CONDITIONALACCESSTRANSFERMETHODS
            case "authenticationTransfer":
                result |= AUTHENTICATIONTRANSFER_CONDITIONALACCESSTRANSFERMETHODS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CONDITIONALACCESSTRANSFERMETHODS
            default:
                return 0, errors.New("Unknown ConditionalAccessTransferMethods value: " + v)
        }
    }
    return &result, nil
}
func SerializeConditionalAccessTransferMethods(values []ConditionalAccessTransferMethods) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConditionalAccessTransferMethods) isMultiValue() bool {
    return true
}
