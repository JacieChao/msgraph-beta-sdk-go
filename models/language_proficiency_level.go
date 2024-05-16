package models
type LanguageProficiencyLevel int

const (
    ELEMENTARY_LANGUAGEPROFICIENCYLEVEL LanguageProficiencyLevel = iota
    CONVERSATIONAL_LANGUAGEPROFICIENCYLEVEL
    LIMITEDWORKING_LANGUAGEPROFICIENCYLEVEL
    PROFESSIONALWORKING_LANGUAGEPROFICIENCYLEVEL
    FULLPROFESSIONAL_LANGUAGEPROFICIENCYLEVEL
    NATIVEORBILINGUAL_LANGUAGEPROFICIENCYLEVEL
    UNKNOWNFUTUREVALUE_LANGUAGEPROFICIENCYLEVEL
)

func (i LanguageProficiencyLevel) String() string {
    return []string{"elementary", "conversational", "limitedWorking", "professionalWorking", "fullProfessional", "nativeOrBilingual", "unknownFutureValue"}[i]
}
func ParseLanguageProficiencyLevel(v string) (any, error) {
    result := ELEMENTARY_LANGUAGEPROFICIENCYLEVEL
    switch v {
        case "elementary":
            result = ELEMENTARY_LANGUAGEPROFICIENCYLEVEL
        case "conversational":
            result = CONVERSATIONAL_LANGUAGEPROFICIENCYLEVEL
        case "limitedWorking":
            result = LIMITEDWORKING_LANGUAGEPROFICIENCYLEVEL
        case "professionalWorking":
            result = PROFESSIONALWORKING_LANGUAGEPROFICIENCYLEVEL
        case "fullProfessional":
            result = FULLPROFESSIONAL_LANGUAGEPROFICIENCYLEVEL
        case "nativeOrBilingual":
            result = NATIVEORBILINGUAL_LANGUAGEPROFICIENCYLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LANGUAGEPROFICIENCYLEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLanguageProficiencyLevel(values []LanguageProficiencyLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LanguageProficiencyLevel) isMultiValue() bool {
    return false
}
