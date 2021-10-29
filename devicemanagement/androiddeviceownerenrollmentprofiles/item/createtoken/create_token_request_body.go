package createtoken

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CreateTokenRequestBody struct {
    additionalData map[string]interface{};
    tokenValidityInSeconds *int32;
}
func NewCreateTokenRequestBody()(*CreateTokenRequestBody) {
    m := &CreateTokenRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CreateTokenRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CreateTokenRequestBody) GetTokenValidityInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.tokenValidityInSeconds
    }
}
func (m *CreateTokenRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["tokenValidityInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTokenValidityInSeconds(val)
        return nil
    }
    return res
}
func (m *CreateTokenRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *CreateTokenRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("tokenValidityInSeconds", m.GetTokenValidityInSeconds())
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
func (m *CreateTokenRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CreateTokenRequestBody) SetTokenValidityInSeconds(value *int32)() {
    m.tokenValidityInSeconds = value
}