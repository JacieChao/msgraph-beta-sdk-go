package models
type PrinterFeedOrientation int

const (
    LONGEDGEFIRST_PRINTERFEEDORIENTATION PrinterFeedOrientation = iota
    SHORTEDGEFIRST_PRINTERFEEDORIENTATION
)

func (i PrinterFeedOrientation) String() string {
    return []string{"longEdgeFirst", "shortEdgeFirst"}[i]
}
func ParsePrinterFeedOrientation(v string) (any, error) {
    result := LONGEDGEFIRST_PRINTERFEEDORIENTATION
    switch v {
        case "longEdgeFirst":
            result = LONGEDGEFIRST_PRINTERFEEDORIENTATION
        case "shortEdgeFirst":
            result = SHORTEDGEFIRST_PRINTERFEEDORIENTATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePrinterFeedOrientation(values []PrinterFeedOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PrinterFeedOrientation) isMultiValue() bool {
    return false
}
