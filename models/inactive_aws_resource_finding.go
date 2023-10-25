package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InactiveAwsResourceFinding 
type InactiveAwsResourceFinding struct {
    IdentityFinding
}
// NewInactiveAwsResourceFinding instantiates a new inactiveAwsResourceFinding and sets the default values.
func NewInactiveAwsResourceFinding()(*InactiveAwsResourceFinding) {
    m := &InactiveAwsResourceFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateInactiveAwsResourceFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInactiveAwsResourceFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInactiveAwsResourceFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InactiveAwsResourceFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InactiveAwsResourceFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// InactiveAwsResourceFindingable 
type InactiveAwsResourceFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
