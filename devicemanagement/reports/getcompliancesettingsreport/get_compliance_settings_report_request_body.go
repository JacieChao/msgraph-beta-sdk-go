package getcompliancesettingsreport

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetComplianceSettingsReportRequestBody 
type GetComplianceSettingsReportRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    filter *string;
    // 
    groupBy []string;
    // 
    name *string;
    // 
    orderBy []string;
    // 
    search *string;
    // 
    select_escaped []string;
    // 
    sessionId *string;
    // 
    skip *int32;
    // 
    top *int32;
}
// NewGetComplianceSettingsReportRequestBody instantiates a new getComplianceSettingsReportRequestBody and sets the default values.
func NewGetComplianceSettingsReportRequestBody()(*GetComplianceSettingsReportRequestBody) {
    m := &GetComplianceSettingsReportRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetComplianceSettingsReportRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFilter gets the filter property value. 
func (m *GetComplianceSettingsReportRequestBody) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// GetGroupBy gets the groupBy property value. 
func (m *GetComplianceSettingsReportRequestBody) GetGroupBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupBy
    }
}
// GetName gets the name property value. 
func (m *GetComplianceSettingsReportRequestBody) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOrderBy gets the orderBy property value. 
func (m *GetComplianceSettingsReportRequestBody) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
// GetSearch gets the search property value. 
func (m *GetComplianceSettingsReportRequestBody) GetSearch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.search
    }
}
// GetSelect gets the select property value. 
func (m *GetComplianceSettingsReportRequestBody) GetSelect()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
// GetSessionId gets the sessionId property value. 
func (m *GetComplianceSettingsReportRequestBody) GetSessionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionId
    }
}
// GetSkip gets the skip property value. 
func (m *GetComplianceSettingsReportRequestBody) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
// GetTop gets the top property value. 
func (m *GetComplianceSettingsReportRequestBody) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetComplianceSettingsReportRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["groupBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupBy(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["orderBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrderBy(res)
        }
        return nil
    }
    res["search"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val)
        }
        return nil
    }
    res["select"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelect(res)
        }
        return nil
    }
    res["sessionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionId(val)
        }
        return nil
    }
    res["skip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkip(val)
        }
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    return res
}
func (m *GetComplianceSettingsReportRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetComplianceSettingsReportRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    if m.GetGroupBy() != nil {
        err := writer.WriteCollectionOfStringValues("groupBy", m.GetGroupBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetOrderBy() != nil {
        err := writer.WriteCollectionOfStringValues("orderBy", m.GetOrderBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    if m.GetSelect() != nil {
        err := writer.WriteCollectionOfStringValues("select", m.GetSelect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionId", m.GetSessionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
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
func (m *GetComplianceSettingsReportRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFilter sets the filter property value. 
func (m *GetComplianceSettingsReportRequestBody) SetFilter(value *string)() {
    if m != nil {
        m.filter = value
    }
}
// SetGroupBy sets the groupBy property value. 
func (m *GetComplianceSettingsReportRequestBody) SetGroupBy(value []string)() {
    if m != nil {
        m.groupBy = value
    }
}
// SetName sets the name property value. 
func (m *GetComplianceSettingsReportRequestBody) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOrderBy sets the orderBy property value. 
func (m *GetComplianceSettingsReportRequestBody) SetOrderBy(value []string)() {
    if m != nil {
        m.orderBy = value
    }
}
// SetSearch sets the search property value. 
func (m *GetComplianceSettingsReportRequestBody) SetSearch(value *string)() {
    if m != nil {
        m.search = value
    }
}
// SetSelect sets the select property value. 
func (m *GetComplianceSettingsReportRequestBody) SetSelect(value []string)() {
    if m != nil {
        m.select_escaped = value
    }
}
// SetSessionId sets the sessionId property value. 
func (m *GetComplianceSettingsReportRequestBody) SetSessionId(value *string)() {
    if m != nil {
        m.sessionId = value
    }
}
// SetSkip sets the skip property value. 
func (m *GetComplianceSettingsReportRequestBody) SetSkip(value *int32)() {
    if m != nil {
        m.skip = value
    }
}
// SetTop sets the top property value. 
func (m *GetComplianceSettingsReportRequestBody) SetTop(value *int32)() {
    if m != nil {
        m.top = value
    }
}