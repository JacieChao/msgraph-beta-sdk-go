package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type PagePromotionType int

const (
    MICROSOFTRESERVED_PAGEPROMOTIONTYPE PagePromotionType = iota
    PAGE_PAGEPROMOTIONTYPE
    NEWSPOST_PAGEPROMOTIONTYPE
    UNKNOWNFUTUREVALUE_PAGEPROMOTIONTYPE
)

func (i PagePromotionType) String() string {
    return []string{"microsoftReserved", "page", "newsPost", "unknownFutureValue"}[i]
}
func ParsePagePromotionType(v string) (interface{}, error) {
    result := MICROSOFTRESERVED_PAGEPROMOTIONTYPE
    switch v {
        case "microsoftReserved":
            result = MICROSOFTRESERVED_PAGEPROMOTIONTYPE
        case "page":
            result = PAGE_PAGEPROMOTIONTYPE
        case "newsPost":
            result = NEWSPOST_PAGEPROMOTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PAGEPROMOTIONTYPE
        default:
            return 0, errors.New("Unknown PagePromotionType value: " + v)
    }
    return &result, nil
}
func SerializePagePromotionType(values []PagePromotionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
