package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody 
type ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The updateWindowsDeviceAccountActionParameter property
    updateWindowsDeviceAccountActionParameter ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable
}
// NewItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody instantiates a new ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody and sets the default values.
func NewItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody()(*ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) {
    m := &ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["updateWindowsDeviceAccountActionParameter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUpdateWindowsDeviceAccountActionParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindowsDeviceAccountActionParameter(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable))
        }
        return nil
    }
    return res
}
// GetUpdateWindowsDeviceAccountActionParameter gets the updateWindowsDeviceAccountActionParameter property value. The updateWindowsDeviceAccountActionParameter property
func (m *ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) GetUpdateWindowsDeviceAccountActionParameter()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable) {
    return m.updateWindowsDeviceAccountActionParameter
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("updateWindowsDeviceAccountActionParameter", m.GetUpdateWindowsDeviceAccountActionParameter())
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
func (m *ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUpdateWindowsDeviceAccountActionParameter sets the updateWindowsDeviceAccountActionParameter property value. The updateWindowsDeviceAccountActionParameter property
func (m *ItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) SetUpdateWindowsDeviceAccountActionParameter(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable)() {
    m.updateWindowsDeviceAccountActionParameter = value
}
