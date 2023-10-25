package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegeEscalationAwsRoleFinding 
type PrivilegeEscalationAwsRoleFinding struct {
    PrivilegeEscalationFinding
}
// NewPrivilegeEscalationAwsRoleFinding instantiates a new privilegeEscalationAwsRoleFinding and sets the default values.
func NewPrivilegeEscalationAwsRoleFinding()(*PrivilegeEscalationAwsRoleFinding) {
    m := &PrivilegeEscalationAwsRoleFinding{
        PrivilegeEscalationFinding: *NewPrivilegeEscalationFinding(),
    }
    return m
}
// CreatePrivilegeEscalationAwsRoleFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegeEscalationAwsRoleFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegeEscalationAwsRoleFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegeEscalationAwsRoleFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegeEscalationFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PrivilegeEscalationAwsRoleFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegeEscalationFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// PrivilegeEscalationAwsRoleFindingable 
type PrivilegeEscalationAwsRoleFindingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegeEscalationFindingable
}
