package dismiss

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DismissRequestBody provides operations to call the dismiss method.
type DismissRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The dismissReason property
    dismissReason *string;
}
// NewDismissRequestBody instantiates a new dismissRequestBody and sets the default values.
func NewDismissRequestBody()(*DismissRequestBody) {
    m := &DismissRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDismissRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDismissRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDismissRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DismissRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDismissReason gets the dismissReason property value. The dismissReason property
func (m *DismissRequestBody) GetDismissReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dismissReason
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DismissRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dismissReason"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissReason(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DismissRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dismissReason", m.GetDismissReason())
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
func (m *DismissRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDismissReason sets the dismissReason property value. The dismissReason property
func (m *DismissRequestBody) SetDismissReason(value *string)() {
    if m != nil {
        m.dismissReason = value
    }
}
