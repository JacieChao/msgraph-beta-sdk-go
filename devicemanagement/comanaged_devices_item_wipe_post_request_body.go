package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ComanagedDevicesItemWipePostRequestBody 
type ComanagedDevicesItemWipePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewComanagedDevicesItemWipePostRequestBody instantiates a new ComanagedDevicesItemWipePostRequestBody and sets the default values.
func NewComanagedDevicesItemWipePostRequestBody()(*ComanagedDevicesItemWipePostRequestBody) {
    m := &ComanagedDevicesItemWipePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemWipePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemWipePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemWipePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemWipePostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ComanagedDevicesItemWipePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemWipePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["macOsUnlockCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOsUnlockCode(val)
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
    res["useProtectedWipe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseProtectedWipe(val)
        }
        return nil
    }
    return res
}
// GetKeepEnrollmentData gets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *ComanagedDevicesItemWipePostRequestBody) GetKeepEnrollmentData()(*bool) {
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
func (m *ComanagedDevicesItemWipePostRequestBody) GetKeepUserData()(*bool) {
    val, err := m.GetBackingStore().Get("keepUserData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMacOsUnlockCode gets the macOsUnlockCode property value. The macOsUnlockCode property
func (m *ComanagedDevicesItemWipePostRequestBody) GetMacOsUnlockCode()(*string) {
    val, err := m.GetBackingStore().Get("macOsUnlockCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *ComanagedDevicesItemWipePostRequestBody) GetPersistEsimDataPlan()(*bool) {
    val, err := m.GetBackingStore().Get("persistEsimDataPlan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUseProtectedWipe gets the useProtectedWipe property value. The useProtectedWipe property
func (m *ComanagedDevicesItemWipePostRequestBody) GetUseProtectedWipe()(*bool) {
    val, err := m.GetBackingStore().Get("useProtectedWipe")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemWipePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("macOsUnlockCode", m.GetMacOsUnlockCode())
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
        err := writer.WriteBoolValue("useProtectedWipe", m.GetUseProtectedWipe())
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
func (m *ComanagedDevicesItemWipePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ComanagedDevicesItemWipePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *ComanagedDevicesItemWipePostRequestBody) SetKeepEnrollmentData(value *bool)() {
    err := m.GetBackingStore().Set("keepEnrollmentData", value)
    if err != nil {
        panic(err)
    }
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *ComanagedDevicesItemWipePostRequestBody) SetKeepUserData(value *bool)() {
    err := m.GetBackingStore().Set("keepUserData", value)
    if err != nil {
        panic(err)
    }
}
// SetMacOsUnlockCode sets the macOsUnlockCode property value. The macOsUnlockCode property
func (m *ComanagedDevicesItemWipePostRequestBody) SetMacOsUnlockCode(value *string)() {
    err := m.GetBackingStore().Set("macOsUnlockCode", value)
    if err != nil {
        panic(err)
    }
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *ComanagedDevicesItemWipePostRequestBody) SetPersistEsimDataPlan(value *bool)() {
    err := m.GetBackingStore().Set("persistEsimDataPlan", value)
    if err != nil {
        panic(err)
    }
}
// SetUseProtectedWipe sets the useProtectedWipe property value. The useProtectedWipe property
func (m *ComanagedDevicesItemWipePostRequestBody) SetUseProtectedWipe(value *bool)() {
    err := m.GetBackingStore().Set("useProtectedWipe", value)
    if err != nil {
        panic(err)
    }
}
// ComanagedDevicesItemWipePostRequestBodyable 
type ComanagedDevicesItemWipePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetKeepEnrollmentData()(*bool)
    GetKeepUserData()(*bool)
    GetMacOsUnlockCode()(*string)
    GetPersistEsimDataPlan()(*bool)
    GetUseProtectedWipe()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetKeepEnrollmentData(value *bool)()
    SetKeepUserData(value *bool)()
    SetMacOsUnlockCode(value *string)()
    SetPersistEsimDataPlan(value *bool)()
    SetUseProtectedWipe(value *bool)()
}
