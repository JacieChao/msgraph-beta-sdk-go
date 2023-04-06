package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileGeneralDeviceConfiguration 
type AndroidWorkProfileGeneralDeviceConfiguration struct {
    DeviceConfiguration
}
// NewAndroidWorkProfileGeneralDeviceConfiguration instantiates a new AndroidWorkProfileGeneralDeviceConfiguration and sets the default values.
func NewAndroidWorkProfileGeneralDeviceConfiguration()(*AndroidWorkProfileGeneralDeviceConfiguration) {
    m := &AndroidWorkProfileGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileGeneralDeviceConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidWorkProfileGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfileGeneralDeviceConfiguration(), nil
}
// GetAllowedGoogleAccountDomains gets the allowedGoogleAccountDomains property value. Determine domains allow-list for accounts that can be added to work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetAllowedGoogleAccountDomains()([]string) {
    val, err := m.GetBackingStore().Get("allowedGoogleAccountDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["allowedGoogleAccountDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedGoogleAccountDomains(res)
        }
        return nil
    }
    res["passwordBlockFaceUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockFaceUnlock(val)
        }
        return nil
    }
    res["passwordBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["passwordBlockIrisUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockIrisUnlock(val)
        }
        return nil
    }
    res["passwordBlockTrustAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockTrustAgents(val)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passwordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidWorkProfileRequiredPasswordType))
        }
        return nil
    }
    res["passwordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["requiredPasswordComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidRequiredPasswordComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredPasswordComplexity(val.(*AndroidRequiredPasswordComplexity))
        }
        return nil
    }
    res["securityRequireVerifyApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireVerifyApps(val)
        }
        return nil
    }
    res["vpnAlwaysOnPackageIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnAlwaysOnPackageIdentifier(val)
        }
        return nil
    }
    res["vpnEnableAlwaysOnLockdownMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnEnableAlwaysOnLockdownMode(val)
        }
        return nil
    }
    res["workProfileAccountUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileAccountUse)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileAccountUse(val.(*AndroidWorkProfileAccountUse))
        }
        return nil
    }
    res["workProfileAllowAppInstallsFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileAllowAppInstallsFromUnknownSources(val)
        }
        return nil
    }
    res["workProfileAllowWidgets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileAllowWidgets(val)
        }
        return nil
    }
    res["workProfileBlockAddingAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockAddingAccounts(val)
        }
        return nil
    }
    res["workProfileBlockCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCamera(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileCallerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileCallerId(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileContactsSearch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileContactsSearch(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileCopyPaste"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileCopyPaste(val)
        }
        return nil
    }
    res["workProfileBlockNotificationsWhileDeviceLocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockNotificationsWhileDeviceLocked(val)
        }
        return nil
    }
    res["workProfileBlockPersonalAppInstallsFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockPersonalAppInstallsFromUnknownSources(val)
        }
        return nil
    }
    res["workProfileBlockScreenCapture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockScreenCapture(val)
        }
        return nil
    }
    res["workProfileBluetoothEnableContactSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBluetoothEnableContactSharing(val)
        }
        return nil
    }
    res["workProfileDataSharingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileCrossProfileDataSharingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileDataSharingType(val.(*AndroidWorkProfileCrossProfileDataSharingType))
        }
        return nil
    }
    res["workProfileDefaultAppPermissionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileDefaultAppPermissionPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileDefaultAppPermissionPolicy(val.(*AndroidWorkProfileDefaultAppPermissionPolicyType))
        }
        return nil
    }
    res["workProfilePasswordBlockFaceUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockFaceUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockIrisUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockIrisUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockTrustAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockTrustAgents(val)
        }
        return nil
    }
    res["workProfilePasswordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordExpirationDays(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLength(val)
        }
        return nil
    }
    res["workProfilePasswordMinLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinLowerCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinNonLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinNumericCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinSymbolCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinUpperCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["workProfilePasswordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["workProfilePasswordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordRequiredType(val.(*AndroidWorkProfileRequiredPasswordType))
        }
        return nil
    }
    res["workProfilePasswordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["workProfileRequiredPasswordComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidRequiredPasswordComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileRequiredPasswordComplexity(val.(*AndroidRequiredPasswordComplexity))
        }
        return nil
    }
    res["workProfileRequirePassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileRequirePassword(val)
        }
        return nil
    }
    return res
}
// GetPasswordBlockFaceUnlock gets the passwordBlockFaceUnlock property value. Indicates whether or not to block face unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockFaceUnlock()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockFaceUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockFingerprintUnlock gets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockFingerprintUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockIrisUnlock gets the passwordBlockIrisUnlock property value. Indicates whether or not to block iris unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockIrisUnlock()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockIrisUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockTrustAgents gets the passwordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockTrustAgents()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockTrustAgents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordExpirationDays()(*int32) {
    val, err := m.GetBackingStore().Get("passwordExpirationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum length of passwords. Valid values 4 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesOfInactivityBeforeScreenTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("passwordPreviousPasswordBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Android Work Profile required password type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordRequiredType()(*AndroidWorkProfileRequiredPasswordType) {
    val, err := m.GetBackingStore().Get("passwordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidWorkProfileRequiredPasswordType)
    }
    return nil
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    val, err := m.GetBackingStore().Get("passwordSignInFailureCountBeforeFactoryReset")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRequiredPasswordComplexity gets the requiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity) {
    val, err := m.GetBackingStore().Get("requiredPasswordComplexity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidRequiredPasswordComplexity)
    }
    return nil
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetSecurityRequireVerifyApps()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireVerifyApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetVpnAlwaysOnPackageIdentifier gets the vpnAlwaysOnPackageIdentifier property value. Enable lockdown mode for always-on VPN.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetVpnAlwaysOnPackageIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("vpnAlwaysOnPackageIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVpnEnableAlwaysOnLockdownMode gets the vpnEnableAlwaysOnLockdownMode property value. Enable lockdown mode for always-on VPN.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetVpnEnableAlwaysOnLockdownMode()(*bool) {
    val, err := m.GetBackingStore().Get("vpnEnableAlwaysOnLockdownMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileAccountUse gets the workProfileAccountUse property value. An enum representing possible values for account use in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileAccountUse()(*AndroidWorkProfileAccountUse) {
    val, err := m.GetBackingStore().Get("workProfileAccountUse")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidWorkProfileAccountUse)
    }
    return nil
}
// GetWorkProfileAllowAppInstallsFromUnknownSources gets the workProfileAllowAppInstallsFromUnknownSources property value. Indicates whether to allow installation of apps from unknown sources.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileAllowAppInstallsFromUnknownSources()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileAllowAppInstallsFromUnknownSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileAllowWidgets gets the workProfileAllowWidgets property value. Allow widgets from work profile apps.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileAllowWidgets()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileAllowWidgets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBlockAddingAccounts gets the workProfileBlockAddingAccounts property value. Block users from adding/removing accounts in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockAddingAccounts()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBlockAddingAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBlockCamera gets the workProfileBlockCamera property value. Block work profile camera.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCamera()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBlockCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBlockCrossProfileCallerId gets the workProfileBlockCrossProfileCallerId property value. Block display work profile caller ID in personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCallerId()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBlockCrossProfileCallerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBlockCrossProfileContactsSearch gets the workProfileBlockCrossProfileContactsSearch property value. Block work profile contacts availability in personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileContactsSearch()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBlockCrossProfileContactsSearch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBlockCrossProfileCopyPaste gets the workProfileBlockCrossProfileCopyPaste property value. Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCopyPaste()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBlockCrossProfileCopyPaste")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBlockNotificationsWhileDeviceLocked gets the workProfileBlockNotificationsWhileDeviceLocked property value. Indicates whether or not to block notifications while device locked.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockNotificationsWhileDeviceLocked()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBlockNotificationsWhileDeviceLocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBlockPersonalAppInstallsFromUnknownSources gets the workProfileBlockPersonalAppInstallsFromUnknownSources property value. Prevent app installations from unknown sources in the personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockPersonalAppInstallsFromUnknownSources()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBlockPersonalAppInstallsFromUnknownSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBlockScreenCapture gets the workProfileBlockScreenCapture property value. Block screen capture in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockScreenCapture()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBlockScreenCapture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileBluetoothEnableContactSharing gets the workProfileBluetoothEnableContactSharing property value. Allow bluetooth devices to access enterprise contacts.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBluetoothEnableContactSharing()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileBluetoothEnableContactSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfileDataSharingType gets the workProfileDataSharingType property value. Android Work Profile cross profile data sharing type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDataSharingType()(*AndroidWorkProfileCrossProfileDataSharingType) {
    val, err := m.GetBackingStore().Get("workProfileDataSharingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidWorkProfileCrossProfileDataSharingType)
    }
    return nil
}
// GetWorkProfileDefaultAppPermissionPolicy gets the workProfileDefaultAppPermissionPolicy property value. Android Work Profile default app permission policy type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDefaultAppPermissionPolicy()(*AndroidWorkProfileDefaultAppPermissionPolicyType) {
    val, err := m.GetBackingStore().Get("workProfileDefaultAppPermissionPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidWorkProfileDefaultAppPermissionPolicyType)
    }
    return nil
}
// GetWorkProfilePasswordBlockFaceUnlock gets the workProfilePasswordBlockFaceUnlock property value. Indicates whether or not to block face unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFaceUnlock()(*bool) {
    val, err := m.GetBackingStore().Get("workProfilePasswordBlockFaceUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfilePasswordBlockFingerprintUnlock gets the workProfilePasswordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFingerprintUnlock()(*bool) {
    val, err := m.GetBackingStore().Get("workProfilePasswordBlockFingerprintUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfilePasswordBlockIrisUnlock gets the workProfilePasswordBlockIrisUnlock property value. Indicates whether or not to block iris unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockIrisUnlock()(*bool) {
    val, err := m.GetBackingStore().Get("workProfilePasswordBlockIrisUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfilePasswordBlockTrustAgents gets the workProfilePasswordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockTrustAgents()(*bool) {
    val, err := m.GetBackingStore().Get("workProfilePasswordBlockTrustAgents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfilePasswordExpirationDays gets the workProfilePasswordExpirationDays property value. Number of days before the work profile password expires. Valid values 1 to 365
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDays()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordExpirationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinimumLength gets the workProfilePasswordMinimumLength property value. Minimum length of work profile password. Valid values 4 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinLetterCharacters gets the workProfilePasswordMinLetterCharacters property value. Minimum # of letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLetterCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinLetterCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinLowerCaseCharacters gets the workProfilePasswordMinLowerCaseCharacters property value. Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLowerCaseCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinLowerCaseCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinNonLetterCharacters gets the workProfilePasswordMinNonLetterCharacters property value. Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNonLetterCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinNonLetterCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinNumericCharacters gets the workProfilePasswordMinNumericCharacters property value. Minimum # of numeric characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNumericCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinNumericCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinSymbolCharacters gets the workProfilePasswordMinSymbolCharacters property value. Minimum # of symbols required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinSymbolCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinSymbolCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinUpperCaseCharacters gets the workProfilePasswordMinUpperCaseCharacters property value. Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinUpperCaseCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinUpperCaseCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout gets the workProfilePasswordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinutesOfInactivityBeforeScreenTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordPreviousPasswordBlockCount gets the workProfilePasswordPreviousPasswordBlockCount property value. Number of previous work profile passwords to block. Valid values 0 to 24
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordPreviousPasswordBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordRequiredType gets the workProfilePasswordRequiredType property value. Android Work Profile required password type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredType()(*AndroidWorkProfileRequiredPasswordType) {
    val, err := m.GetBackingStore().Get("workProfilePasswordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidWorkProfileRequiredPasswordType)
    }
    return nil
}
// GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset gets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordSignInFailureCountBeforeFactoryReset")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfileRequiredPasswordComplexity gets the workProfileRequiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity) {
    val, err := m.GetBackingStore().Get("workProfileRequiredPasswordComplexity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidRequiredPasswordComplexity)
    }
    return nil
}
// GetWorkProfileRequirePassword gets the workProfileRequirePassword property value. Password is required or not for work profile
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileRequirePassword()(*bool) {
    val, err := m.GetBackingStore().Get("workProfileRequirePassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidWorkProfileGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedGoogleAccountDomains() != nil {
        err = writer.WriteCollectionOfStringValues("allowedGoogleAccountDomains", m.GetAllowedGoogleAccountDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockFaceUnlock", m.GetPasswordBlockFaceUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockFingerprintUnlock", m.GetPasswordBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockIrisUnlock", m.GetPasswordBlockIrisUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockTrustAgents", m.GetPasswordBlockTrustAgents())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeScreenTimeout", m.GetPasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordBlockCount", m.GetPasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredPasswordComplexity() != nil {
        cast := (*m.GetRequiredPasswordComplexity()).String()
        err = writer.WriteStringValue("requiredPasswordComplexity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireVerifyApps", m.GetSecurityRequireVerifyApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vpnAlwaysOnPackageIdentifier", m.GetVpnAlwaysOnPackageIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("vpnEnableAlwaysOnLockdownMode", m.GetVpnEnableAlwaysOnLockdownMode())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileAccountUse() != nil {
        cast := (*m.GetWorkProfileAccountUse()).String()
        err = writer.WriteStringValue("workProfileAccountUse", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileAllowAppInstallsFromUnknownSources", m.GetWorkProfileAllowAppInstallsFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileAllowWidgets", m.GetWorkProfileAllowWidgets())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockAddingAccounts", m.GetWorkProfileBlockAddingAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCamera", m.GetWorkProfileBlockCamera())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileCallerId", m.GetWorkProfileBlockCrossProfileCallerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileContactsSearch", m.GetWorkProfileBlockCrossProfileContactsSearch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileCopyPaste", m.GetWorkProfileBlockCrossProfileCopyPaste())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockNotificationsWhileDeviceLocked", m.GetWorkProfileBlockNotificationsWhileDeviceLocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockPersonalAppInstallsFromUnknownSources", m.GetWorkProfileBlockPersonalAppInstallsFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockScreenCapture", m.GetWorkProfileBlockScreenCapture())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBluetoothEnableContactSharing", m.GetWorkProfileBluetoothEnableContactSharing())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileDataSharingType() != nil {
        cast := (*m.GetWorkProfileDataSharingType()).String()
        err = writer.WriteStringValue("workProfileDataSharingType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileDefaultAppPermissionPolicy() != nil {
        cast := (*m.GetWorkProfileDefaultAppPermissionPolicy()).String()
        err = writer.WriteStringValue("workProfileDefaultAppPermissionPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockFaceUnlock", m.GetWorkProfilePasswordBlockFaceUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockFingerprintUnlock", m.GetWorkProfilePasswordBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockIrisUnlock", m.GetWorkProfilePasswordBlockIrisUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockTrustAgents", m.GetWorkProfilePasswordBlockTrustAgents())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordExpirationDays", m.GetWorkProfilePasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLength", m.GetWorkProfilePasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinLetterCharacters", m.GetWorkProfilePasswordMinLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinLowerCaseCharacters", m.GetWorkProfilePasswordMinLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinNonLetterCharacters", m.GetWorkProfilePasswordMinNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinNumericCharacters", m.GetWorkProfilePasswordMinNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinSymbolCharacters", m.GetWorkProfilePasswordMinSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinUpperCaseCharacters", m.GetWorkProfilePasswordMinUpperCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinutesOfInactivityBeforeScreenTimeout", m.GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordPreviousPasswordBlockCount", m.GetWorkProfilePasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfilePasswordRequiredType() != nil {
        cast := (*m.GetWorkProfilePasswordRequiredType()).String()
        err = writer.WriteStringValue("workProfilePasswordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordSignInFailureCountBeforeFactoryReset", m.GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileRequiredPasswordComplexity() != nil {
        cast := (*m.GetWorkProfileRequiredPasswordComplexity()).String()
        err = writer.WriteStringValue("workProfileRequiredPasswordComplexity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileRequirePassword", m.GetWorkProfileRequirePassword())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedGoogleAccountDomains sets the allowedGoogleAccountDomains property value. Determine domains allow-list for accounts that can be added to work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetAllowedGoogleAccountDomains(value []string)() {
    err := m.GetBackingStore().Set("allowedGoogleAccountDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockFaceUnlock sets the passwordBlockFaceUnlock property value. Indicates whether or not to block face unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockFaceUnlock(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockFaceUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockFingerprintUnlock sets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockFingerprintUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockIrisUnlock sets the passwordBlockIrisUnlock property value. Indicates whether or not to block iris unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockIrisUnlock(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockIrisUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockTrustAgents sets the passwordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockTrustAgents(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockTrustAgents", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum length of passwords. Valid values 4 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesOfInactivityBeforeScreenTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("passwordPreviousPasswordBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Android Work Profile required password type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordRequiredType(value *AndroidWorkProfileRequiredPasswordType)() {
    err := m.GetBackingStore().Set("passwordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    err := m.GetBackingStore().Set("passwordSignInFailureCountBeforeFactoryReset", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredPasswordComplexity sets the requiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)() {
    err := m.GetBackingStore().Set("requiredPasswordComplexity", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireVerifyApps", value)
    if err != nil {
        panic(err)
    }
}
// SetVpnAlwaysOnPackageIdentifier sets the vpnAlwaysOnPackageIdentifier property value. Enable lockdown mode for always-on VPN.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetVpnAlwaysOnPackageIdentifier(value *string)() {
    err := m.GetBackingStore().Set("vpnAlwaysOnPackageIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetVpnEnableAlwaysOnLockdownMode sets the vpnEnableAlwaysOnLockdownMode property value. Enable lockdown mode for always-on VPN.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetVpnEnableAlwaysOnLockdownMode(value *bool)() {
    err := m.GetBackingStore().Set("vpnEnableAlwaysOnLockdownMode", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileAccountUse sets the workProfileAccountUse property value. An enum representing possible values for account use in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileAccountUse(value *AndroidWorkProfileAccountUse)() {
    err := m.GetBackingStore().Set("workProfileAccountUse", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileAllowAppInstallsFromUnknownSources sets the workProfileAllowAppInstallsFromUnknownSources property value. Indicates whether to allow installation of apps from unknown sources.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileAllowAppInstallsFromUnknownSources(value *bool)() {
    err := m.GetBackingStore().Set("workProfileAllowAppInstallsFromUnknownSources", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileAllowWidgets sets the workProfileAllowWidgets property value. Allow widgets from work profile apps.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileAllowWidgets(value *bool)() {
    err := m.GetBackingStore().Set("workProfileAllowWidgets", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBlockAddingAccounts sets the workProfileBlockAddingAccounts property value. Block users from adding/removing accounts in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockAddingAccounts(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBlockAddingAccounts", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBlockCamera sets the workProfileBlockCamera property value. Block work profile camera.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCamera(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBlockCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBlockCrossProfileCallerId sets the workProfileBlockCrossProfileCallerId property value. Block display work profile caller ID in personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCallerId(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBlockCrossProfileCallerId", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBlockCrossProfileContactsSearch sets the workProfileBlockCrossProfileContactsSearch property value. Block work profile contacts availability in personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileContactsSearch(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBlockCrossProfileContactsSearch", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBlockCrossProfileCopyPaste sets the workProfileBlockCrossProfileCopyPaste property value. Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCopyPaste(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBlockCrossProfileCopyPaste", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBlockNotificationsWhileDeviceLocked sets the workProfileBlockNotificationsWhileDeviceLocked property value. Indicates whether or not to block notifications while device locked.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockNotificationsWhileDeviceLocked(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBlockNotificationsWhileDeviceLocked", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBlockPersonalAppInstallsFromUnknownSources sets the workProfileBlockPersonalAppInstallsFromUnknownSources property value. Prevent app installations from unknown sources in the personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockPersonalAppInstallsFromUnknownSources(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBlockPersonalAppInstallsFromUnknownSources", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBlockScreenCapture sets the workProfileBlockScreenCapture property value. Block screen capture in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockScreenCapture(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBlockScreenCapture", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileBluetoothEnableContactSharing sets the workProfileBluetoothEnableContactSharing property value. Allow bluetooth devices to access enterprise contacts.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBluetoothEnableContactSharing(value *bool)() {
    err := m.GetBackingStore().Set("workProfileBluetoothEnableContactSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileDataSharingType sets the workProfileDataSharingType property value. Android Work Profile cross profile data sharing type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileDataSharingType(value *AndroidWorkProfileCrossProfileDataSharingType)() {
    err := m.GetBackingStore().Set("workProfileDataSharingType", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileDefaultAppPermissionPolicy sets the workProfileDefaultAppPermissionPolicy property value. Android Work Profile default app permission policy type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileDefaultAppPermissionPolicy(value *AndroidWorkProfileDefaultAppPermissionPolicyType)() {
    err := m.GetBackingStore().Set("workProfileDefaultAppPermissionPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordBlockFaceUnlock sets the workProfilePasswordBlockFaceUnlock property value. Indicates whether or not to block face unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockFaceUnlock(value *bool)() {
    err := m.GetBackingStore().Set("workProfilePasswordBlockFaceUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordBlockFingerprintUnlock sets the workProfilePasswordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockFingerprintUnlock(value *bool)() {
    err := m.GetBackingStore().Set("workProfilePasswordBlockFingerprintUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordBlockIrisUnlock sets the workProfilePasswordBlockIrisUnlock property value. Indicates whether or not to block iris unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockIrisUnlock(value *bool)() {
    err := m.GetBackingStore().Set("workProfilePasswordBlockIrisUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordBlockTrustAgents sets the workProfilePasswordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockTrustAgents(value *bool)() {
    err := m.GetBackingStore().Set("workProfilePasswordBlockTrustAgents", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordExpirationDays sets the workProfilePasswordExpirationDays property value. Number of days before the work profile password expires. Valid values 1 to 365
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinimumLength sets the workProfilePasswordMinimumLength property value. Minimum length of work profile password. Valid values 4 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinLetterCharacters sets the workProfilePasswordMinLetterCharacters property value. Minimum # of letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLetterCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinLetterCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinLowerCaseCharacters sets the workProfilePasswordMinLowerCaseCharacters property value. Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLowerCaseCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinLowerCaseCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinNonLetterCharacters sets the workProfilePasswordMinNonLetterCharacters property value. Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNonLetterCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinNonLetterCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinNumericCharacters sets the workProfilePasswordMinNumericCharacters property value. Minimum # of numeric characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNumericCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinNumericCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinSymbolCharacters sets the workProfilePasswordMinSymbolCharacters property value. Minimum # of symbols required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinSymbolCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinSymbolCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinUpperCaseCharacters sets the workProfilePasswordMinUpperCaseCharacters property value. Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinUpperCaseCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinUpperCaseCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout sets the workProfilePasswordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinutesOfInactivityBeforeScreenTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordPreviousPasswordBlockCount sets the workProfilePasswordPreviousPasswordBlockCount property value. Number of previous work profile passwords to block. Valid values 0 to 24
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordPreviousPasswordBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordRequiredType sets the workProfilePasswordRequiredType property value. Android Work Profile required password type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordRequiredType(value *AndroidWorkProfileRequiredPasswordType)() {
    err := m.GetBackingStore().Set("workProfilePasswordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset sets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordSignInFailureCountBeforeFactoryReset", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileRequiredPasswordComplexity sets the workProfileRequiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)() {
    err := m.GetBackingStore().Set("workProfileRequiredPasswordComplexity", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfileRequirePassword sets the workProfileRequirePassword property value. Password is required or not for work profile
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileRequirePassword(value *bool)() {
    err := m.GetBackingStore().Set("workProfileRequirePassword", value)
    if err != nil {
        panic(err)
    }
}
// AndroidWorkProfileGeneralDeviceConfigurationable 
type AndroidWorkProfileGeneralDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedGoogleAccountDomains()([]string)
    GetPasswordBlockFaceUnlock()(*bool)
    GetPasswordBlockFingerprintUnlock()(*bool)
    GetPasswordBlockIrisUnlock()(*bool)
    GetPasswordBlockTrustAgents()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequiredType()(*AndroidWorkProfileRequiredPasswordType)
    GetPasswordSignInFailureCountBeforeFactoryReset()(*int32)
    GetRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity)
    GetSecurityRequireVerifyApps()(*bool)
    GetVpnAlwaysOnPackageIdentifier()(*string)
    GetVpnEnableAlwaysOnLockdownMode()(*bool)
    GetWorkProfileAccountUse()(*AndroidWorkProfileAccountUse)
    GetWorkProfileAllowAppInstallsFromUnknownSources()(*bool)
    GetWorkProfileAllowWidgets()(*bool)
    GetWorkProfileBlockAddingAccounts()(*bool)
    GetWorkProfileBlockCamera()(*bool)
    GetWorkProfileBlockCrossProfileCallerId()(*bool)
    GetWorkProfileBlockCrossProfileContactsSearch()(*bool)
    GetWorkProfileBlockCrossProfileCopyPaste()(*bool)
    GetWorkProfileBlockNotificationsWhileDeviceLocked()(*bool)
    GetWorkProfileBlockPersonalAppInstallsFromUnknownSources()(*bool)
    GetWorkProfileBlockScreenCapture()(*bool)
    GetWorkProfileBluetoothEnableContactSharing()(*bool)
    GetWorkProfileDataSharingType()(*AndroidWorkProfileCrossProfileDataSharingType)
    GetWorkProfileDefaultAppPermissionPolicy()(*AndroidWorkProfileDefaultAppPermissionPolicyType)
    GetWorkProfilePasswordBlockFaceUnlock()(*bool)
    GetWorkProfilePasswordBlockFingerprintUnlock()(*bool)
    GetWorkProfilePasswordBlockIrisUnlock()(*bool)
    GetWorkProfilePasswordBlockTrustAgents()(*bool)
    GetWorkProfilePasswordExpirationDays()(*int32)
    GetWorkProfilePasswordMinimumLength()(*int32)
    GetWorkProfilePasswordMinLetterCharacters()(*int32)
    GetWorkProfilePasswordMinLowerCaseCharacters()(*int32)
    GetWorkProfilePasswordMinNonLetterCharacters()(*int32)
    GetWorkProfilePasswordMinNumericCharacters()(*int32)
    GetWorkProfilePasswordMinSymbolCharacters()(*int32)
    GetWorkProfilePasswordMinUpperCaseCharacters()(*int32)
    GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetWorkProfilePasswordPreviousPasswordBlockCount()(*int32)
    GetWorkProfilePasswordRequiredType()(*AndroidWorkProfileRequiredPasswordType)
    GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset()(*int32)
    GetWorkProfileRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity)
    GetWorkProfileRequirePassword()(*bool)
    SetAllowedGoogleAccountDomains(value []string)()
    SetPasswordBlockFaceUnlock(value *bool)()
    SetPasswordBlockFingerprintUnlock(value *bool)()
    SetPasswordBlockIrisUnlock(value *bool)()
    SetPasswordBlockTrustAgents(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequiredType(value *AndroidWorkProfileRequiredPasswordType)()
    SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)()
    SetRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)()
    SetSecurityRequireVerifyApps(value *bool)()
    SetVpnAlwaysOnPackageIdentifier(value *string)()
    SetVpnEnableAlwaysOnLockdownMode(value *bool)()
    SetWorkProfileAccountUse(value *AndroidWorkProfileAccountUse)()
    SetWorkProfileAllowAppInstallsFromUnknownSources(value *bool)()
    SetWorkProfileAllowWidgets(value *bool)()
    SetWorkProfileBlockAddingAccounts(value *bool)()
    SetWorkProfileBlockCamera(value *bool)()
    SetWorkProfileBlockCrossProfileCallerId(value *bool)()
    SetWorkProfileBlockCrossProfileContactsSearch(value *bool)()
    SetWorkProfileBlockCrossProfileCopyPaste(value *bool)()
    SetWorkProfileBlockNotificationsWhileDeviceLocked(value *bool)()
    SetWorkProfileBlockPersonalAppInstallsFromUnknownSources(value *bool)()
    SetWorkProfileBlockScreenCapture(value *bool)()
    SetWorkProfileBluetoothEnableContactSharing(value *bool)()
    SetWorkProfileDataSharingType(value *AndroidWorkProfileCrossProfileDataSharingType)()
    SetWorkProfileDefaultAppPermissionPolicy(value *AndroidWorkProfileDefaultAppPermissionPolicyType)()
    SetWorkProfilePasswordBlockFaceUnlock(value *bool)()
    SetWorkProfilePasswordBlockFingerprintUnlock(value *bool)()
    SetWorkProfilePasswordBlockIrisUnlock(value *bool)()
    SetWorkProfilePasswordBlockTrustAgents(value *bool)()
    SetWorkProfilePasswordExpirationDays(value *int32)()
    SetWorkProfilePasswordMinimumLength(value *int32)()
    SetWorkProfilePasswordMinLetterCharacters(value *int32)()
    SetWorkProfilePasswordMinLowerCaseCharacters(value *int32)()
    SetWorkProfilePasswordMinNonLetterCharacters(value *int32)()
    SetWorkProfilePasswordMinNumericCharacters(value *int32)()
    SetWorkProfilePasswordMinSymbolCharacters(value *int32)()
    SetWorkProfilePasswordMinUpperCaseCharacters(value *int32)()
    SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetWorkProfilePasswordPreviousPasswordBlockCount(value *int32)()
    SetWorkProfilePasswordRequiredType(value *AndroidWorkProfileRequiredPasswordType)()
    SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(value *int32)()
    SetWorkProfileRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)()
    SetWorkProfileRequirePassword(value *bool)()
}
