package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InactiveServerlessFunctionFinding 
type InactiveServerlessFunctionFinding struct {
    IdentityFinding
}
// NewInactiveServerlessFunctionFinding instantiates a new inactiveServerlessFunctionFinding and sets the default values.
func NewInactiveServerlessFunctionFinding()(*InactiveServerlessFunctionFinding) {
    m := &InactiveServerlessFunctionFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateInactiveServerlessFunctionFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInactiveServerlessFunctionFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInactiveServerlessFunctionFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InactiveServerlessFunctionFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InactiveServerlessFunctionFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// InactiveServerlessFunctionFindingable 
type InactiveServerlessFunctionFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
