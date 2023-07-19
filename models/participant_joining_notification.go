package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParticipantJoiningNotification 
type ParticipantJoiningNotification struct {
    Entity
}
// NewParticipantJoiningNotification instantiates a new participantJoiningNotification and sets the default values.
func NewParticipantJoiningNotification()(*ParticipantJoiningNotification) {
    m := &ParticipantJoiningNotification{
        Entity: *NewEntity(),
    }
    return m
}
// CreateParticipantJoiningNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParticipantJoiningNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipantJoiningNotification(), nil
}
// GetCall gets the call property value. The call property
func (m *ParticipantJoiningNotification) GetCall()(Callable) {
    val, err := m.GetBackingStore().Get("call")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Callable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParticipantJoiningNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["call"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCallFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCall(val.(Callable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ParticipantJoiningNotification) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ParticipantJoiningNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("call", m.GetCall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCall sets the call property value. The call property
func (m *ParticipantJoiningNotification) SetCall(value Callable)() {
    err := m.GetBackingStore().Set("call", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ParticipantJoiningNotification) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// ParticipantJoiningNotificationable 
type ParticipantJoiningNotificationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCall()(Callable)
    GetOdataType()(*string)
    SetCall(value Callable)()
    SetOdataType(value *string)()
}
