package getoffice365groupsactivitygroupcountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetOffice365GroupsActivityGroupCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    active *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    total *int64;
}
func NewGetOffice365GroupsActivityGroupCountsWithPeriod()(*GetOffice365GroupsActivityGroupCountsWithPeriod) {
    m := &GetOffice365GroupsActivityGroupCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.active
    }
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetTotal()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["active"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetActive(val)
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportDate(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotal(val)
        return nil
    }
    return res
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetActive(value *int64)() {
    m.active = value
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetTotal(value *int64)() {
    m.total = value
}