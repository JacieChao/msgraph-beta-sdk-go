package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type WindowsAutopilotEnrollmentType int

const (
    UNKNOWN_WINDOWSAUTOPILOTENROLLMENTTYPE WindowsAutopilotEnrollmentType = iota
    AZUREADJOINEDWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINED_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
)

func (i WindowsAutopilotEnrollmentType) String() string {
    return []string{"UNKNOWN", "AZUREADJOINEDWITHAUTOPILOTPROFILE", "OFFLINEDOMAINJOINED", "AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE", "AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE", "AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE", "AZUREADJOINEDWITHWHITEGLOVE", "OFFLINEDOMAINJOINEDWITHWHITEGLOVE", "OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE"}[i]
}
func ParseWindowsAutopilotEnrollmentType(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTENROLLMENTTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "AZUREADJOINEDWITHAUTOPILOTPROFILE":
            result = AZUREADJOINEDWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "OFFLINEDOMAINJOINED":
            result = OFFLINEDOMAINJOINED_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE":
            result = AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE":
            result = AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE":
            result = AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "AZUREADJOINEDWITHWHITEGLOVE":
            result = AZUREADJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "OFFLINEDOMAINJOINEDWITHWHITEGLOVE":
            result = OFFLINEDOMAINJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE":
            result = OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        default:
            return 0, errors.New("Unknown WindowsAutopilotEnrollmentType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotEnrollmentType(values []WindowsAutopilotEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}