package security
type SubmissionCategory int

const (
    NOTJUNK_SUBMISSIONCATEGORY SubmissionCategory = iota
    SPAM_SUBMISSIONCATEGORY
    PHISHING_SUBMISSIONCATEGORY
    MALWARE_SUBMISSIONCATEGORY
    UNKNOWNFUTUREVALUE_SUBMISSIONCATEGORY
)

func (i SubmissionCategory) String() string {
    return []string{"notJunk", "spam", "phishing", "malware", "unknownFutureValue"}[i]
}
func ParseSubmissionCategory(v string) (any, error) {
    result := NOTJUNK_SUBMISSIONCATEGORY
    switch v {
        case "notJunk":
            result = NOTJUNK_SUBMISSIONCATEGORY
        case "spam":
            result = SPAM_SUBMISSIONCATEGORY
        case "phishing":
            result = PHISHING_SUBMISSIONCATEGORY
        case "malware":
            result = MALWARE_SUBMISSIONCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SUBMISSIONCATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSubmissionCategory(values []SubmissionCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SubmissionCategory) isMultiValue() bool {
    return false
}
