package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody 
type AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody instantiates a new AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody and sets the default values.
func NewAccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody()(*AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) {
    m := &AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDecision gets the decision property value. The decision property
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) GetDecision()(*string) {
    val, err := m.GetBackingStore().Get("decision")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecision(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
// GetJustification gets the justification property value. The justification property
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) GetJustification()(*string) {
    val, err := m.GetBackingStore().Get("justification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) GetPrincipalId()(*string) {
    val, err := m.GetBackingStore().Get("principalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceId gets the resourceId property value. The resourceId property
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) GetResourceId()(*string) {
    val, err := m.GetBackingStore().Get("resourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("decision", m.GetDecision())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDecision sets the decision property value. The decision property
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) SetDecision(value *string)() {
    err := m.GetBackingStore().Set("decision", value)
    if err != nil {
        panic(err)
    }
}
// SetJustification sets the justification property value. The justification property
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) SetJustification(value *string)() {
    err := m.GetBackingStore().Set("justification", value)
    if err != nil {
        panic(err)
    }
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) SetPrincipalId(value *string)() {
    err := m.GetBackingStore().Set("principalId", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceId sets the resourceId property value. The resourceId property
func (m *AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBody) SetResourceId(value *string)() {
    err := m.GetBackingStore().Set("resourceId", value)
    if err != nil {
        panic(err)
    }
}
// AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBodyable 
type AccessReviewsDecisionsItemInstanceStagesItemDecisionsRecordAllDecisionsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDecision()(*string)
    GetJustification()(*string)
    GetPrincipalId()(*string)
    GetResourceId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDecision(value *string)()
    SetJustification(value *string)()
    SetPrincipalId(value *string)()
    SetResourceId(value *string)()
}
