package monitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
)

// AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse 
type AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse instantiates a new AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse and sets the default values.
func NewAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse()(*AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse) {
    m := &AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.CreatePortalNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse) GetValue()([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponse) SetValue(value []i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponseable 
type AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)
    SetValue(value []i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)()
}
