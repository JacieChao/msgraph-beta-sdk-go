package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody 
type WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody instantiates a new WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody()(*WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) {
    m := &WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody(), nil
}
// GetAddExclusions gets the addExclusions property value. The addExclusions property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetAddExclusions()([]string) {
    val, err := m.GetBackingStore().Get("addExclusions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAddMembers gets the addMembers property value. The addMembers property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetAddMembers()([]string) {
    val, err := m.GetBackingStore().Get("addMembers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addExclusions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAddExclusions(res)
        }
        return nil
    }
    res["addMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAddMembers(res)
        }
        return nil
    }
    res["memberEntityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEntityType(val)
        }
        return nil
    }
    res["removeExclusions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveExclusions(res)
        }
        return nil
    }
    res["removeMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveMembers(res)
        }
        return nil
    }
    return res
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetMemberEntityType()(*string) {
    val, err := m.GetBackingStore().Get("memberEntityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemoveExclusions gets the removeExclusions property value. The removeExclusions property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetRemoveExclusions()([]string) {
    val, err := m.GetBackingStore().Get("removeExclusions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRemoveMembers gets the removeMembers property value. The removeMembers property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetRemoveMembers()([]string) {
    val, err := m.GetBackingStore().Get("removeMembers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddExclusions() != nil {
        err := writer.WriteCollectionOfStringValues("addExclusions", m.GetAddExclusions())
        if err != nil {
            return err
        }
    }
    if m.GetAddMembers() != nil {
        err := writer.WriteCollectionOfStringValues("addMembers", m.GetAddMembers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memberEntityType", m.GetMemberEntityType())
        if err != nil {
            return err
        }
    }
    if m.GetRemoveExclusions() != nil {
        err := writer.WriteCollectionOfStringValues("removeExclusions", m.GetRemoveExclusions())
        if err != nil {
            return err
        }
    }
    if m.GetRemoveMembers() != nil {
        err := writer.WriteCollectionOfStringValues("removeMembers", m.GetRemoveMembers())
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
// SetAddExclusions sets the addExclusions property value. The addExclusions property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetAddExclusions(value []string)() {
    err := m.GetBackingStore().Set("addExclusions", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddMembers sets the addMembers property value. The addMembers property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetAddMembers(value []string)() {
    err := m.GetBackingStore().Set("addMembers", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetMemberEntityType(value *string)() {
    err := m.GetBackingStore().Set("memberEntityType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoveExclusions sets the removeExclusions property value. The removeExclusions property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetRemoveExclusions(value []string)() {
    err := m.GetBackingStore().Set("removeExclusions", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoveMembers sets the removeMembers property value. The removeMembers property
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetRemoveMembers(value []string)() {
    err := m.GetBackingStore().Set("removeMembers", value)
    if err != nil {
        panic(err)
    }
}
// WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyable 
type WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddExclusions()([]string)
    GetAddMembers()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetMemberEntityType()(*string)
    GetRemoveExclusions()([]string)
    GetRemoveMembers()([]string)
    SetAddExclusions(value []string)()
    SetAddMembers(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetMemberEntityType(value *string)()
    SetRemoveExclusions(value []string)()
    SetRemoveMembers(value []string)()
}
