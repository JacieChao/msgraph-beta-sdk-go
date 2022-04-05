package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationEnforcedRestrictionsSessionControl 
type ApplicationEnforcedRestrictionsSessionControl struct {
    ConditionalAccessSessionControl
}
// NewApplicationEnforcedRestrictionsSessionControl instantiates a new applicationEnforcedRestrictionsSessionControl and sets the default values.
func NewApplicationEnforcedRestrictionsSessionControl()(*ApplicationEnforcedRestrictionsSessionControl) {
    m := &ApplicationEnforcedRestrictionsSessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
// CreateApplicationEnforcedRestrictionsSessionControlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationEnforcedRestrictionsSessionControlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationEnforcedRestrictionsSessionControl(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationEnforcedRestrictionsSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ApplicationEnforcedRestrictionsSessionControl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
