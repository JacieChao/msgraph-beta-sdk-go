package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchQueryString struct {
    additionalData map[string]interface{};
    query *string;
}
func NewSearchQueryString()(*SearchQueryString) {
    m := &SearchQueryString{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchQueryString) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchQueryString) GetQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
func (m *SearchQueryString) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["query"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQuery(val)
        return nil
    }
    return res
}
func (m *SearchQueryString) IsNil()(bool) {
    return m == nil
}
func (m *SearchQueryString) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("query", m.GetQuery())
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
func (m *SearchQueryString) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchQueryString) SetQuery(value *string)() {
    m.query = value
}