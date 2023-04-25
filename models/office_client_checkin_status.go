package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OfficeClientCheckinStatus 
type OfficeClientCheckinStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOfficeClientCheckinStatus instantiates a new officeClientCheckinStatus and sets the default values.
func NewOfficeClientCheckinStatus()(*OfficeClientCheckinStatus) {
    m := &OfficeClientCheckinStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOfficeClientCheckinStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeClientCheckinStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeClientCheckinStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeClientCheckinStatus) GetAdditionalData()(map[string]any) {
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
// GetAppliedPolicies gets the appliedPolicies property value. The appliedPolicies property
func (m *OfficeClientCheckinStatus) GetAppliedPolicies()([]string) {
    val, err := m.GetBackingStore().Get("appliedPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *OfficeClientCheckinStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCheckinDateTime gets the checkinDateTime property value. The checkinDateTime property
func (m *OfficeClientCheckinStatus) GetCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("checkinDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The deviceName property
func (m *OfficeClientCheckinStatus) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDevicePlatform gets the devicePlatform property value. The devicePlatform property
func (m *OfficeClientCheckinStatus) GetDevicePlatform()(*string) {
    val, err := m.GetBackingStore().Get("devicePlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDevicePlatformVersion gets the devicePlatformVersion property value. The devicePlatformVersion property
func (m *OfficeClientCheckinStatus) GetDevicePlatformVersion()(*string) {
    val, err := m.GetBackingStore().Get("devicePlatformVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetErrorMessage gets the errorMessage property value. The errorMessage property
func (m *OfficeClientCheckinStatus) GetErrorMessage()(*string) {
    val, err := m.GetBackingStore().Get("errorMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeClientCheckinStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appliedPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAppliedPolicies(res)
        }
        return nil
    }
    res["checkinDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckinDateTime(val)
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
    res["devicePlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatform(val)
        }
        return nil
    }
    res["devicePlatformVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatformVersion(val)
        }
        return nil
    }
    res["errorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["wasSuccessful"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWasSuccessful(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OfficeClientCheckinStatus) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The userId property
func (m *OfficeClientCheckinStatus) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *OfficeClientCheckinStatus) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWasSuccessful gets the wasSuccessful property value. The wasSuccessful property
func (m *OfficeClientCheckinStatus) GetWasSuccessful()(*bool) {
    val, err := m.GetBackingStore().Get("wasSuccessful")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OfficeClientCheckinStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppliedPolicies() != nil {
        err := writer.WriteCollectionOfStringValues("appliedPolicies", m.GetAppliedPolicies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("checkinDateTime", m.GetCheckinDateTime())
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
        err := writer.WriteStringValue("devicePlatform", m.GetDevicePlatform())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("devicePlatformVersion", m.GetDevicePlatformVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorMessage", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wasSuccessful", m.GetWasSuccessful())
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
func (m *OfficeClientCheckinStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliedPolicies sets the appliedPolicies property value. The appliedPolicies property
func (m *OfficeClientCheckinStatus) SetAppliedPolicies(value []string)() {
    err := m.GetBackingStore().Set("appliedPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *OfficeClientCheckinStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCheckinDateTime sets the checkinDateTime property value. The checkinDateTime property
func (m *OfficeClientCheckinStatus) SetCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("checkinDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *OfficeClientCheckinStatus) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDevicePlatform sets the devicePlatform property value. The devicePlatform property
func (m *OfficeClientCheckinStatus) SetDevicePlatform(value *string)() {
    err := m.GetBackingStore().Set("devicePlatform", value)
    if err != nil {
        panic(err)
    }
}
// SetDevicePlatformVersion sets the devicePlatformVersion property value. The devicePlatformVersion property
func (m *OfficeClientCheckinStatus) SetDevicePlatformVersion(value *string)() {
    err := m.GetBackingStore().Set("devicePlatformVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorMessage sets the errorMessage property value. The errorMessage property
func (m *OfficeClientCheckinStatus) SetErrorMessage(value *string)() {
    err := m.GetBackingStore().Set("errorMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OfficeClientCheckinStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The userId property
func (m *OfficeClientCheckinStatus) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *OfficeClientCheckinStatus) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetWasSuccessful sets the wasSuccessful property value. The wasSuccessful property
func (m *OfficeClientCheckinStatus) SetWasSuccessful(value *bool)() {
    err := m.GetBackingStore().Set("wasSuccessful", value)
    if err != nil {
        panic(err)
    }
}
// OfficeClientCheckinStatusable 
type OfficeClientCheckinStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppliedPolicies()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceName()(*string)
    GetDevicePlatform()(*string)
    GetDevicePlatformVersion()(*string)
    GetErrorMessage()(*string)
    GetOdataType()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetWasSuccessful()(*bool)
    SetAppliedPolicies(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceName(value *string)()
    SetDevicePlatform(value *string)()
    SetDevicePlatformVersion(value *string)()
    SetErrorMessage(value *string)()
    SetOdataType(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetWasSuccessful(value *bool)()
}
