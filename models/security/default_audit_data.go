package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultAuditData 
type DefaultAuditData struct {
    AuditData
}
// NewDefaultAuditData instantiates a new defaultAuditData and sets the default values.
func NewDefaultAuditData()(*DefaultAuditData) {
    m := &DefaultAuditData{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.defaultAuditData"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDefaultAuditDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultAuditDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultAuditData(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultAuditData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DefaultAuditData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// DefaultAuditDataable 
type DefaultAuditDataable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
