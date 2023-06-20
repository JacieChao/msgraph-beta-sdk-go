package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppIntentAndState 
type MobileAppIntentAndState struct {
    Entity
}
// NewMobileAppIntentAndState instantiates a new MobileAppIntentAndState and sets the default values.
func NewMobileAppIntentAndState()(*MobileAppIntentAndState) {
    m := &MobileAppIntentAndState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppIntentAndStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppIntentAndStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppIntentAndState(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppIntentAndState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["managedDeviceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceIdentifier(val)
        }
        return nil
    }
    res["mobileAppList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppIntentAndStateDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppIntentAndStateDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileAppIntentAndStateDetailable)
                }
            }
            m.SetMobileAppList(res)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceIdentifier gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppIntentAndState) GetManagedDeviceIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMobileAppList gets the mobileAppList property value. The list of payload intents and states for the tenant.
func (m *MobileAppIntentAndState) GetMobileAppList()([]MobileAppIntentAndStateDetailable) {
    val, err := m.GetBackingStore().Get("mobileAppList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobileAppIntentAndStateDetailable)
    }
    return nil
}
// GetUserId gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppIntentAndState) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppIntentAndState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("managedDeviceIdentifier", m.GetManagedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppList()))
        for i, v := range m.GetMobileAppList() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManagedDeviceIdentifier sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppIntentAndState) SetManagedDeviceIdentifier(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileAppList sets the mobileAppList property value. The list of payload intents and states for the tenant.
func (m *MobileAppIntentAndState) SetMobileAppList(value []MobileAppIntentAndStateDetailable)() {
    err := m.GetBackingStore().Set("mobileAppList", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppIntentAndState) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// MobileAppIntentAndStateable 
type MobileAppIntentAndStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManagedDeviceIdentifier()(*string)
    GetMobileAppList()([]MobileAppIntentAndStateDetailable)
    GetUserId()(*string)
    SetManagedDeviceIdentifier(value *string)()
    SetMobileAppList(value []MobileAppIntentAndStateDetailable)()
    SetUserId(value *string)()
}
