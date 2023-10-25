package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OverprovisionedServerlessFunctionFinding 
type OverprovisionedServerlessFunctionFinding struct {
    IdentityFinding
}
// NewOverprovisionedServerlessFunctionFinding instantiates a new overprovisionedServerlessFunctionFinding and sets the default values.
func NewOverprovisionedServerlessFunctionFinding()(*OverprovisionedServerlessFunctionFinding) {
    m := &OverprovisionedServerlessFunctionFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateOverprovisionedServerlessFunctionFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOverprovisionedServerlessFunctionFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOverprovisionedServerlessFunctionFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OverprovisionedServerlessFunctionFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OverprovisionedServerlessFunctionFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OverprovisionedServerlessFunctionFindingable 
type OverprovisionedServerlessFunctionFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
