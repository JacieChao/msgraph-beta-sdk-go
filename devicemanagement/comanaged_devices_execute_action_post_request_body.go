package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ComanagedDevicesExecuteActionPostRequestBody 
type ComanagedDevicesExecuteActionPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewComanagedDevicesExecuteActionPostRequestBody instantiates a new ComanagedDevicesExecuteActionPostRequestBody and sets the default values.
func NewComanagedDevicesExecuteActionPostRequestBody()(*ComanagedDevicesExecuteActionPostRequestBody) {
    m := &ComanagedDevicesExecuteActionPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesExecuteActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesExecuteActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesExecuteActionPostRequestBody(), nil
}
// GetActionName gets the actionName property value. The actionName property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetActionName()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceRemoteAction) {
    val, err := m.GetBackingStore().Get("actionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceRemoteAction)
    }
    return nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCarrierUrl gets the carrierUrl property value. The carrierUrl property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetCarrierUrl()(*string) {
    val, err := m.GetBackingStore().Get("carrierUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeprovisionReason gets the deprovisionReason property value. The deprovisionReason property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetDeprovisionReason()(*string) {
    val, err := m.GetBackingStore().Get("deprovisionReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceIds gets the deviceIds property value. The deviceIds property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetDeviceIds()([]string) {
    val, err := m.GetBackingStore().Get("deviceIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The deviceName property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseManagedDeviceRemoteAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceRemoteAction))
        }
        return nil
    }
    res["carrierUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierUrl(val)
        }
        return nil
    }
    res["deprovisionReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeprovisionReason(val)
        }
        return nil
    }
    res["deviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDeviceIds(res)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["keepEnrollmentData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepEnrollmentData(val)
        }
        return nil
    }
    res["keepUserData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
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
    res["organizationalUnitPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnitPath(val)
        }
        return nil
    }
    res["persistEsimDataPlan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistEsimDataPlan(val)
        }
        return nil
    }
    return res
}
// GetKeepEnrollmentData gets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetKeepEnrollmentData()(*bool) {
    val, err := m.GetBackingStore().Get("keepEnrollmentData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKeepUserData gets the keepUserData property value. The keepUserData property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetKeepUserData()(*bool) {
    val, err := m.GetBackingStore().Get("keepUserData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetNotificationBody gets the notificationBody property value. The notificationBody property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetNotificationBody()(*string) {
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
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetNotificationTitle()(*string) {
    val, err := m.GetBackingStore().Get("notificationTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrganizationalUnitPath gets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetOrganizationalUnitPath()(*string) {
    val, err := m.GetBackingStore().Get("organizationalUnitPath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *ComanagedDevicesExecuteActionPostRequestBody) GetPersistEsimDataPlan()(*bool) {
    val, err := m.GetBackingStore().Get("persistEsimDataPlan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ComanagedDevicesExecuteActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActionName() != nil {
        cast := (*m.GetActionName()).String()
        err := writer.WriteStringValue("actionName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("carrierUrl", m.GetCarrierUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deprovisionReason", m.GetDeprovisionReason())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("deviceIds", m.GetDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("keepEnrollmentData", m.GetKeepEnrollmentData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
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
        err := writer.WriteStringValue("organizationalUnitPath", m.GetOrganizationalUnitPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("persistEsimDataPlan", m.GetPersistEsimDataPlan())
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
// SetActionName sets the actionName property value. The actionName property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetActionName(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceRemoteAction)() {
    err := m.GetBackingStore().Set("actionName", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCarrierUrl sets the carrierUrl property value. The carrierUrl property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetCarrierUrl(value *string)() {
    err := m.GetBackingStore().Set("carrierUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDeprovisionReason sets the deprovisionReason property value. The deprovisionReason property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetDeprovisionReason(value *string)() {
    err := m.GetBackingStore().Set("deprovisionReason", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceIds sets the deviceIds property value. The deviceIds property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetDeviceIds(value []string)() {
    err := m.GetBackingStore().Set("deviceIds", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetKeepEnrollmentData(value *bool)() {
    err := m.GetBackingStore().Set("keepEnrollmentData", value)
    if err != nil {
        panic(err)
    }
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetKeepUserData(value *bool)() {
    err := m.GetBackingStore().Set("keepUserData", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationBody sets the notificationBody property value. The notificationBody property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetNotificationBody(value *string)() {
    err := m.GetBackingStore().Set("notificationBody", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationTitle sets the notificationTitle property value. The notificationTitle property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetNotificationTitle(value *string)() {
    err := m.GetBackingStore().Set("notificationTitle", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationalUnitPath sets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetOrganizationalUnitPath(value *string)() {
    err := m.GetBackingStore().Set("organizationalUnitPath", value)
    if err != nil {
        panic(err)
    }
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *ComanagedDevicesExecuteActionPostRequestBody) SetPersistEsimDataPlan(value *bool)() {
    err := m.GetBackingStore().Set("persistEsimDataPlan", value)
    if err != nil {
        panic(err)
    }
}
// ComanagedDevicesExecuteActionPostRequestBodyable 
type ComanagedDevicesExecuteActionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionName()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceRemoteAction)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCarrierUrl()(*string)
    GetDeprovisionReason()(*string)
    GetDeviceIds()([]string)
    GetDeviceName()(*string)
    GetKeepEnrollmentData()(*bool)
    GetKeepUserData()(*bool)
    GetNotificationBody()(*string)
    GetNotificationTitle()(*string)
    GetOrganizationalUnitPath()(*string)
    GetPersistEsimDataPlan()(*bool)
    SetActionName(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceRemoteAction)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCarrierUrl(value *string)()
    SetDeprovisionReason(value *string)()
    SetDeviceIds(value []string)()
    SetDeviceName(value *string)()
    SetKeepEnrollmentData(value *bool)()
    SetKeepUserData(value *bool)()
    SetNotificationBody(value *string)()
    SetNotificationTitle(value *string)()
    SetOrganizationalUnitPath(value *string)()
    SetPersistEsimDataPlan(value *bool)()
}
