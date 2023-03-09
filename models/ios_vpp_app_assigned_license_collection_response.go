package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppAssignedLicenseCollectionResponse 
type IosVppAppAssignedLicenseCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewIosVppAppAssignedLicenseCollectionResponse instantiates a new IosVppAppAssignedLicenseCollectionResponse and sets the default values.
func NewIosVppAppAssignedLicenseCollectionResponse()(*IosVppAppAssignedLicenseCollectionResponse) {
    m := &IosVppAppAssignedLicenseCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateIosVppAppAssignedLicenseCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppAppAssignedLicenseCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVppAppAssignedLicenseCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVppAppAssignedLicenseCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosVppAppAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosVppAppAssignedLicenseable, len(val))
            for i, v := range val {
                res[i] = v.(IosVppAppAssignedLicenseable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *IosVppAppAssignedLicenseCollectionResponse) GetValue()([]IosVppAppAssignedLicenseable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosVppAppAssignedLicenseable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosVppAppAssignedLicenseCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *IosVppAppAssignedLicenseCollectionResponse) SetValue(value []IosVppAppAssignedLicenseable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// IosVppAppAssignedLicenseCollectionResponseable 
type IosVppAppAssignedLicenseCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]IosVppAppAssignedLicenseable)
    SetValue(value []IosVppAppAssignedLicenseable)()
}
