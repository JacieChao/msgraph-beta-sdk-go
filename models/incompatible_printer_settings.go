package models
import (
    "errors"
)
// 
type IncompatiblePrinterSettings int

const (
    SHOW_INCOMPATIBLEPRINTERSETTINGS IncompatiblePrinterSettings = iota
    HIDE_INCOMPATIBLEPRINTERSETTINGS
    UNKNOWNFUTUREVALUE_INCOMPATIBLEPRINTERSETTINGS
)

func (i IncompatiblePrinterSettings) String() string {
    return []string{"show", "hide", "unknownFutureValue"}[i]
}
func ParseIncompatiblePrinterSettings(v string) (any, error) {
    result := SHOW_INCOMPATIBLEPRINTERSETTINGS
    switch v {
        case "show":
            result = SHOW_INCOMPATIBLEPRINTERSETTINGS
        case "hide":
            result = HIDE_INCOMPATIBLEPRINTERSETTINGS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INCOMPATIBLEPRINTERSETTINGS
        default:
            return 0, errors.New("Unknown IncompatiblePrinterSettings value: " + v)
    }
    return &result, nil
}
func SerializeIncompatiblePrinterSettings(values []IncompatiblePrinterSettings) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IncompatiblePrinterSettings) isMultiValue() bool {
    return false
}
