package enrollassetsbyid

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
)

type EnrollAssetsByIdRequestBody struct {
    additionalData map[string]interface{};
    ids []string;
    memberEntityType *string;
    updateCategory *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory;
}
func NewEnrollAssetsByIdRequestBody()(*EnrollAssetsByIdRequestBody) {
    m := &EnrollAssetsByIdRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EnrollAssetsByIdRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EnrollAssetsByIdRequestBody) GetIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.ids
    }
}
func (m *EnrollAssetsByIdRequestBody) GetMemberEntityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberEntityType
    }
}
func (m *EnrollAssetsByIdRequestBody) GetUpdateCategory()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
func (m *EnrollAssetsByIdRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ids"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIds(res)
        return nil
    }
    res["memberEntityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMemberEntityType(val)
        return nil
    }
    res["updateCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseUpdateCategory)
        if err != nil {
            return err
        }
        cast := val.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory)
        m.SetUpdateCategory(&cast)
        return nil
    }
    return res
}
func (m *EnrollAssetsByIdRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *EnrollAssetsByIdRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memberEntityType", m.GetMemberEntityType())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := m.GetUpdateCategory().String()
        err := writer.WriteStringValue("updateCategory", &cast)
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
func (m *EnrollAssetsByIdRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EnrollAssetsByIdRequestBody) SetIds(value []string)() {
    m.ids = value
}
func (m *EnrollAssetsByIdRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
func (m *EnrollAssetsByIdRequestBody) SetUpdateCategory(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory)() {
    m.updateCategory = value
}