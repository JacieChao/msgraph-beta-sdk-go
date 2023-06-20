package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody 
type WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody()(*WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIds(res)
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
    return res
}
// GetIds gets the ids property value. The ids property
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) GetIds()([]string) {
    val, err := m.GetBackingStore().Get("ids")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) GetMemberEntityType()(*string) {
    val, err := m.GetBackingStore().Get("memberEntityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIds sets the ids property value. The ids property
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) SetIds(value []string)() {
    err := m.GetBackingStore().Set("ids", value)
    if err != nil {
        panic(err)
    }
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBody) SetMemberEntityType(value *string)() {
    err := m.GetBackingStore().Set("memberEntityType", value)
    if err != nil {
        panic(err)
    }
}
// WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBodyable 
type WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIds()([]string)
    GetMemberEntityType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIds(value []string)()
    SetMemberEntityType(value *string)()
}
