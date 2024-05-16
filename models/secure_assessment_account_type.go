package models
// Type of accounts that are allowed for Windows10SecureAssessment ConfigurationAccount.
type SecureAssessmentAccountType int

const (
    // Indicates an Azure AD account in format of AzureAD\username@tenant.com.
    AZUREADACCOUNT_SECUREASSESSMENTACCOUNTTYPE SecureAssessmentAccountType = iota
    // Indicates a domain account in format of domain\user or user@domain.com.
    DOMAINACCOUNT_SECUREASSESSMENTACCOUNTTYPE
    // Indicates a local account in format of username.
    LOCALACCOUNT_SECUREASSESSMENTACCOUNTTYPE
    // Indicates a local guest account in format of test name.
    LOCALGUESTACCOUNT_SECUREASSESSMENTACCOUNTTYPE
)

func (i SecureAssessmentAccountType) String() string {
    return []string{"azureADAccount", "domainAccount", "localAccount", "localGuestAccount"}[i]
}
func ParseSecureAssessmentAccountType(v string) (any, error) {
    result := AZUREADACCOUNT_SECUREASSESSMENTACCOUNTTYPE
    switch v {
        case "azureADAccount":
            result = AZUREADACCOUNT_SECUREASSESSMENTACCOUNTTYPE
        case "domainAccount":
            result = DOMAINACCOUNT_SECUREASSESSMENTACCOUNTTYPE
        case "localAccount":
            result = LOCALACCOUNT_SECUREASSESSMENTACCOUNTTYPE
        case "localGuestAccount":
            result = LOCALGUESTACCOUNT_SECUREASSESSMENTACCOUNTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecureAssessmentAccountType(values []SecureAssessmentAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecureAssessmentAccountType) isMultiValue() bool {
    return false
}
