package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IntuneBrandable 
type IntuneBrandable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompanyPortalBlockedActions()([]CompanyPortalBlockedActionable)
    GetContactITEmailAddress()(*string)
    GetContactITName()(*string)
    GetContactITNotes()(*string)
    GetContactITPhoneNumber()(*string)
    GetCustomCanSeePrivacyMessage()(*string)
    GetCustomCantSeePrivacyMessage()(*string)
    GetCustomPrivacyMessage()(*string)
    GetDarkBackgroundLogo()(MimeContentable)
    GetDisableClientTelemetry()(*bool)
    GetDisplayName()(*string)
    GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions)
    GetIsFactoryResetDisabled()(*bool)
    GetIsRemoveDeviceDisabled()(*bool)
    GetLandingPageCustomizedImage()(MimeContentable)
    GetLightBackgroundLogo()(MimeContentable)
    GetOnlineSupportSiteName()(*string)
    GetOnlineSupportSiteUrl()(*string)
    GetPrivacyUrl()(*string)
    GetRoleScopeTagIds()([]string)
    GetSendDeviceOwnershipChangePushNotification()(*bool)
    GetShowAzureADEnterpriseApps()(*bool)
    GetShowDisplayNameNextToLogo()(*bool)
    GetShowLogo()(*bool)
    GetShowNameNextToLogo()(*bool)
    GetShowOfficeWebApps()(*bool)
    GetThemeColor()(RgbColorable)
    SetCompanyPortalBlockedActions(value []CompanyPortalBlockedActionable)()
    SetContactITEmailAddress(value *string)()
    SetContactITName(value *string)()
    SetContactITNotes(value *string)()
    SetContactITPhoneNumber(value *string)()
    SetCustomCanSeePrivacyMessage(value *string)()
    SetCustomCantSeePrivacyMessage(value *string)()
    SetCustomPrivacyMessage(value *string)()
    SetDarkBackgroundLogo(value MimeContentable)()
    SetDisableClientTelemetry(value *bool)()
    SetDisplayName(value *string)()
    SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)()
    SetIsFactoryResetDisabled(value *bool)()
    SetIsRemoveDeviceDisabled(value *bool)()
    SetLandingPageCustomizedImage(value MimeContentable)()
    SetLightBackgroundLogo(value MimeContentable)()
    SetOnlineSupportSiteName(value *string)()
    SetOnlineSupportSiteUrl(value *string)()
    SetPrivacyUrl(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetSendDeviceOwnershipChangePushNotification(value *bool)()
    SetShowAzureADEnterpriseApps(value *bool)()
    SetShowDisplayNameNextToLogo(value *bool)()
    SetShowLogo(value *bool)()
    SetShowNameNextToLogo(value *bool)()
    SetShowOfficeWebApps(value *bool)()
    SetThemeColor(value RgbColorable)()
}