package models
// An enum representing the Android Play Integrity API evaluation types.
type AndroidSafetyNetEvaluationType int

const (
    // Default value. Typical measurements and reference data were used.
    BASIC_ANDROIDSAFETYNETEVALUATIONTYPE AndroidSafetyNetEvaluationType = iota
    // Strong Integrity checks (such as a hardware-backed proof of boot integrity) were used.
    HARDWAREBACKED_ANDROIDSAFETYNETEVALUATIONTYPE
)

func (i AndroidSafetyNetEvaluationType) String() string {
    return []string{"basic", "hardwareBacked"}[i]
}
func ParseAndroidSafetyNetEvaluationType(v string) (any, error) {
    result := BASIC_ANDROIDSAFETYNETEVALUATIONTYPE
    switch v {
        case "basic":
            result = BASIC_ANDROIDSAFETYNETEVALUATIONTYPE
        case "hardwareBacked":
            result = HARDWAREBACKED_ANDROIDSAFETYNETEVALUATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidSafetyNetEvaluationType(values []AndroidSafetyNetEvaluationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidSafetyNetEvaluationType) isMultiValue() bool {
    return false
}
