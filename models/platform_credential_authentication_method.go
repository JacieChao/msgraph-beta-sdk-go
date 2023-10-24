package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlatformCredentialAuthenticationMethod 
type PlatformCredentialAuthenticationMethod struct {
    AuthenticationMethod
}
// NewPlatformCredentialAuthenticationMethod instantiates a new platformCredentialAuthenticationMethod and sets the default values.
func NewPlatformCredentialAuthenticationMethod()(*PlatformCredentialAuthenticationMethod) {
    m := &PlatformCredentialAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    odataTypeValue := "#microsoft.graph.platformCredentialAuthenticationMethod"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePlatformCredentialAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlatformCredentialAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlatformCredentialAuthenticationMethod(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *PlatformCredentialAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDevice gets the device property value. The device property
func (m *PlatformCredentialAuthenticationMethod) GetDevice()(Deviceable) {
    val, err := m.GetBackingStore().Get("device")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Deviceable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PlatformCredentialAuthenticationMethod) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlatformCredentialAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(Deviceable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["keyStrength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodKeyStrength)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyStrength(val.(*AuthenticationMethodKeyStrength))
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*AuthenticationMethodPlatform))
        }
        return nil
    }
    return res
}
// GetKeyStrength gets the keyStrength property value. The keyStrength property
func (m *PlatformCredentialAuthenticationMethod) GetKeyStrength()(*AuthenticationMethodKeyStrength) {
    val, err := m.GetBackingStore().Get("keyStrength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationMethodKeyStrength)
    }
    return nil
}
// GetPlatform gets the platform property value. The platform property
func (m *PlatformCredentialAuthenticationMethod) GetPlatform()(*AuthenticationMethodPlatform) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationMethodPlatform)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlatformCredentialAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetKeyStrength() != nil {
        cast := (*m.GetKeyStrength()).String()
        err = writer.WriteStringValue("keyStrength", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *PlatformCredentialAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDevice sets the device property value. The device property
func (m *PlatformCredentialAuthenticationMethod) SetDevice(value Deviceable)() {
    err := m.GetBackingStore().Set("device", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PlatformCredentialAuthenticationMethod) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyStrength sets the keyStrength property value. The keyStrength property
func (m *PlatformCredentialAuthenticationMethod) SetKeyStrength(value *AuthenticationMethodKeyStrength)() {
    err := m.GetBackingStore().Set("keyStrength", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. The platform property
func (m *PlatformCredentialAuthenticationMethod) SetPlatform(value *AuthenticationMethodPlatform)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// PlatformCredentialAuthenticationMethodable 
type PlatformCredentialAuthenticationMethodable interface {
    AuthenticationMethodable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDevice()(Deviceable)
    GetDisplayName()(*string)
    GetKeyStrength()(*AuthenticationMethodKeyStrength)
    GetPlatform()(*AuthenticationMethodPlatform)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDevice(value Deviceable)()
    SetDisplayName(value *string)()
    SetKeyStrength(value *AuthenticationMethodKeyStrength)()
    SetPlatform(value *AuthenticationMethodPlatform)()
}
