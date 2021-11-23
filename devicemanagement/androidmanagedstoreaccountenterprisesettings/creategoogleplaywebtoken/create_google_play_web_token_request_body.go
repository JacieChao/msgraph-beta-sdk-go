package creategoogleplaywebtoken

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// createGooglePlayWebTokenRequestBody 
type CreateGooglePlayWebTokenRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    parentUri *string;
}
// NewCreateGooglePlayWebTokenRequestBody instantiates a new createGooglePlayWebTokenRequestBody and sets the default values.
func NewCreateGooglePlayWebTokenRequestBody()(*CreateGooglePlayWebTokenRequestBody) {
    m := &CreateGooglePlayWebTokenRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateGooglePlayWebTokenRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetParentUri gets the parentUri property value. 
func (m *CreateGooglePlayWebTokenRequestBody) GetParentUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentUri
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateGooglePlayWebTokenRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["parentUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentUri(val)
        }
        return nil
    }
    return res
}
func (m *CreateGooglePlayWebTokenRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateGooglePlayWebTokenRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("parentUri", m.GetParentUri())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateGooglePlayWebTokenRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetParentUri sets the parentUri property value. 
func (m *CreateGooglePlayWebTokenRequestBody) SetParentUri(value *string)() {
    m.parentUri = value
}
