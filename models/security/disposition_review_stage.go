package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DispositionReviewStage 
type DispositionReviewStage struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewDispositionReviewStage instantiates a new dispositionReviewStage and sets the default values.
func NewDispositionReviewStage()(*DispositionReviewStage) {
    m := &DispositionReviewStage{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDispositionReviewStageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDispositionReviewStageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDispositionReviewStage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DispositionReviewStage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["reviewersEmailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetReviewersEmailAddresses(res)
        }
        return nil
    }
    res["stageNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStageNumber(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name representing each stage within a collection.
func (m *DispositionReviewStage) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReviewersEmailAddresses gets the reviewersEmailAddresses property value. A collection of reviewers at each stage.
func (m *DispositionReviewStage) GetReviewersEmailAddresses()([]string) {
    val, err := m.GetBackingStore().Get("reviewersEmailAddresses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetStageNumber gets the stageNumber property value. The sequence number for each stage of the disposition review.
func (m *DispositionReviewStage) GetStageNumber()(*int32) {
    val, err := m.GetBackingStore().Get("stageNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DispositionReviewStage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetReviewersEmailAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("reviewersEmailAddresses", m.GetReviewersEmailAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("stageNumber", m.GetStageNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. Name representing each stage within a collection.
func (m *DispositionReviewStage) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewersEmailAddresses sets the reviewersEmailAddresses property value. A collection of reviewers at each stage.
func (m *DispositionReviewStage) SetReviewersEmailAddresses(value []string)() {
    err := m.GetBackingStore().Set("reviewersEmailAddresses", value)
    if err != nil {
        panic(err)
    }
}
// SetStageNumber sets the stageNumber property value. The sequence number for each stage of the disposition review.
func (m *DispositionReviewStage) SetStageNumber(value *int32)() {
    err := m.GetBackingStore().Set("stageNumber", value)
    if err != nil {
        panic(err)
    }
}
// DispositionReviewStageable 
type DispositionReviewStageable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetReviewersEmailAddresses()([]string)
    GetStageNumber()(*int32)
    SetName(value *string)()
    SetReviewersEmailAddresses(value []string)()
    SetStageNumber(value *int32)()
}
