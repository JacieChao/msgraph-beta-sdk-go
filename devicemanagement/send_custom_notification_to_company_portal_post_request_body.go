package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// SendCustomNotificationToCompanyPortalPostRequestBody 
type SendCustomNotificationToCompanyPortalPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSendCustomNotificationToCompanyPortalPostRequestBody instantiates a new SendCustomNotificationToCompanyPortalPostRequestBody and sets the default values.
func NewSendCustomNotificationToCompanyPortalPostRequestBody()(*SendCustomNotificationToCompanyPortalPostRequestBody) {
    m := &SendCustomNotificationToCompanyPortalPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSendCustomNotificationToCompanyPortalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSendCustomNotificationToCompanyPortalPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupsToNotify"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetGroupsToNotify(res)
        }
        return nil
    }
    res["notificationBody"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationBody(val)
        }
        return nil
    }
    res["notificationTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTitle(val)
        }
        return nil
    }
    return res
}
// GetGroupsToNotify gets the groupsToNotify property value. The groupsToNotify property
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) GetGroupsToNotify()([]string) {
    val, err := m.GetBackingStore().Get("groupsToNotify")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetNotificationBody gets the notificationBody property value. The notificationBody property
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationBody()(*string) {
    val, err := m.GetBackingStore().Get("notificationBody")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNotificationTitle gets the notificationTitle property value. The notificationTitle property
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) GetNotificationTitle()(*string) {
    val, err := m.GetBackingStore().Get("notificationTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroupsToNotify() != nil {
        err := writer.WriteCollectionOfStringValues("groupsToNotify", m.GetGroupsToNotify())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationBody", m.GetNotificationBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTitle", m.GetNotificationTitle())
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
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetGroupsToNotify sets the groupsToNotify property value. The groupsToNotify property
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) SetGroupsToNotify(value []string)() {
    err := m.GetBackingStore().Set("groupsToNotify", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationBody sets the notificationBody property value. The notificationBody property
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationBody(value *string)() {
    err := m.GetBackingStore().Set("notificationBody", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationTitle sets the notificationTitle property value. The notificationTitle property
func (m *SendCustomNotificationToCompanyPortalPostRequestBody) SetNotificationTitle(value *string)() {
    err := m.GetBackingStore().Set("notificationTitle", value)
    if err != nil {
        panic(err)
    }
}
// SendCustomNotificationToCompanyPortalPostRequestBodyable 
type SendCustomNotificationToCompanyPortalPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetGroupsToNotify()([]string)
    GetNotificationBody()(*string)
    GetNotificationTitle()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetGroupsToNotify(value []string)()
    SetNotificationBody(value *string)()
    SetNotificationTitle(value *string)()
}
