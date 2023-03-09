package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecureSignInSessionControl 
type SecureSignInSessionControl struct {
    ConditionalAccessSessionControl
}
// NewSecureSignInSessionControl instantiates a new SecureSignInSessionControl and sets the default values.
func NewSecureSignInSessionControl()(*SecureSignInSessionControl) {
    m := &SecureSignInSessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    odataTypeValue := "#microsoft.graph.secureSignInSessionControl"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSecureSignInSessionControlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecureSignInSessionControlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecureSignInSessionControl(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecureSignInSessionControl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecureSignInSessionControl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SecureSignInSessionControlable 
type SecureSignInSessionControlable interface {
    ConditionalAccessSessionControlable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
