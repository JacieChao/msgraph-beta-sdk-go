package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BlockAccessAction 
type BlockAccessAction struct {
    DlpActionInfo
}
// NewBlockAccessAction instantiates a new BlockAccessAction and sets the default values.
func NewBlockAccessAction()(*BlockAccessAction) {
    m := &BlockAccessAction{
        DlpActionInfo: *NewDlpActionInfo(),
    }
    return m
}
// CreateBlockAccessActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBlockAccessActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlockAccessAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BlockAccessAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DlpActionInfo.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *BlockAccessAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DlpActionInfo.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// BlockAccessActionable 
type BlockAccessActionable interface {
    DlpActionInfoable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
