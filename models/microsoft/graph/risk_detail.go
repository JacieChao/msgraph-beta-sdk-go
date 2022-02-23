package graph
import (
    "strings"
    "errors"
)
// 
type RiskDetail int

const (
    NONE_RISKDETAIL RiskDetail = iota
    ADMINGENERATEDTEMPORARYPASSWORD_RISKDETAIL
    USERPERFORMEDSECUREDPASSWORDCHANGE_RISKDETAIL
    USERPERFORMEDSECUREDPASSWORDRESET_RISKDETAIL
    ADMINCONFIRMEDSIGNINSAFE_RISKDETAIL
    AICONFIRMEDSIGNINSAFE_RISKDETAIL
    USERPASSEDMFADRIVENBYRISKBASEDPOLICY_RISKDETAIL
    ADMINDISMISSEDALLRISKFORUSER_RISKDETAIL
    ADMINCONFIRMEDSIGNINCOMPROMISED_RISKDETAIL
    HIDDEN_RISKDETAIL
    ADMINCONFIRMEDUSERCOMPROMISED_RISKDETAIL
    UNKNOWNFUTUREVALUE_RISKDETAIL
    ADMINCONFIRMEDSERVICEPRINCIPALCOMPROMISED_RISKDETAIL
    ADMINDISMISSEDALLRISKFORSERVICEPRINCIPAL_RISKDETAIL
)

func (i RiskDetail) String() string {
    return []string{"NONE", "ADMINGENERATEDTEMPORARYPASSWORD", "USERPERFORMEDSECUREDPASSWORDCHANGE", "USERPERFORMEDSECUREDPASSWORDRESET", "ADMINCONFIRMEDSIGNINSAFE", "AICONFIRMEDSIGNINSAFE", "USERPASSEDMFADRIVENBYRISKBASEDPOLICY", "ADMINDISMISSEDALLRISKFORUSER", "ADMINCONFIRMEDSIGNINCOMPROMISED", "HIDDEN", "ADMINCONFIRMEDUSERCOMPROMISED", "UNKNOWNFUTUREVALUE", "ADMINCONFIRMEDSERVICEPRINCIPALCOMPROMISED", "ADMINDISMISSEDALLRISKFORSERVICEPRINCIPAL"}[i]
}
func ParseRiskDetail(v string) (interface{}, error) {
    result := NONE_RISKDETAIL
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_RISKDETAIL
        case "ADMINGENERATEDTEMPORARYPASSWORD":
            result = ADMINGENERATEDTEMPORARYPASSWORD_RISKDETAIL
        case "USERPERFORMEDSECUREDPASSWORDCHANGE":
            result = USERPERFORMEDSECUREDPASSWORDCHANGE_RISKDETAIL
        case "USERPERFORMEDSECUREDPASSWORDRESET":
            result = USERPERFORMEDSECUREDPASSWORDRESET_RISKDETAIL
        case "ADMINCONFIRMEDSIGNINSAFE":
            result = ADMINCONFIRMEDSIGNINSAFE_RISKDETAIL
        case "AICONFIRMEDSIGNINSAFE":
            result = AICONFIRMEDSIGNINSAFE_RISKDETAIL
        case "USERPASSEDMFADRIVENBYRISKBASEDPOLICY":
            result = USERPASSEDMFADRIVENBYRISKBASEDPOLICY_RISKDETAIL
        case "ADMINDISMISSEDALLRISKFORUSER":
            result = ADMINDISMISSEDALLRISKFORUSER_RISKDETAIL
        case "ADMINCONFIRMEDSIGNINCOMPROMISED":
            result = ADMINCONFIRMEDSIGNINCOMPROMISED_RISKDETAIL
        case "HIDDEN":
            result = HIDDEN_RISKDETAIL
        case "ADMINCONFIRMEDUSERCOMPROMISED":
            result = ADMINCONFIRMEDUSERCOMPROMISED_RISKDETAIL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RISKDETAIL
        case "ADMINCONFIRMEDSERVICEPRINCIPALCOMPROMISED":
            result = ADMINCONFIRMEDSERVICEPRINCIPALCOMPROMISED_RISKDETAIL
        case "ADMINDISMISSEDALLRISKFORSERVICEPRINCIPAL":
            result = ADMINDISMISSEDALLRISKFORSERVICEPRINCIPAL_RISKDETAIL
        default:
            return 0, errors.New("Unknown RiskDetail value: " + v)
    }
    return &result, nil
}
func SerializeRiskDetail(values []RiskDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
