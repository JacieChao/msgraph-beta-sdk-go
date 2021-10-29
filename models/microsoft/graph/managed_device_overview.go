package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedDeviceOverview struct {
    Entity
    deviceExchangeAccessStateSummary *DeviceExchangeAccessStateSummary;
    deviceOperatingSystemSummary *DeviceOperatingSystemSummary;
    dualEnrolledDeviceCount *int32;
    enrolledDeviceCount *int32;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    managedDeviceModelsAndManufacturers *ManagedDeviceModelsAndManufacturers;
    mdmEnrolledCount *int32;
}
func NewManagedDeviceOverview()(*ManagedDeviceOverview) {
    m := &ManagedDeviceOverview{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ManagedDeviceOverview) GetDeviceExchangeAccessStateSummary()(*DeviceExchangeAccessStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceExchangeAccessStateSummary
    }
}
func (m *ManagedDeviceOverview) GetDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceOperatingSystemSummary
    }
}
func (m *ManagedDeviceOverview) GetDualEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dualEnrolledDeviceCount
    }
}
func (m *ManagedDeviceOverview) GetEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDeviceCount
    }
}
func (m *ManagedDeviceOverview) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *ManagedDeviceOverview) GetManagedDeviceModelsAndManufacturers()(*ManagedDeviceModelsAndManufacturers) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceModelsAndManufacturers
    }
}
func (m *ManagedDeviceOverview) GetMdmEnrolledCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mdmEnrolledCount
    }
}
func (m *ManagedDeviceOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceExchangeAccessStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceExchangeAccessStateSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceExchangeAccessStateSummary(val.(*DeviceExchangeAccessStateSummary))
        return nil
    }
    res["deviceOperatingSystemSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceOperatingSystemSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceOperatingSystemSummary(val.(*DeviceOperatingSystemSummary))
        return nil
    }
    res["dualEnrolledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDualEnrolledDeviceCount(val)
        return nil
    }
    res["enrolledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEnrolledDeviceCount(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["managedDeviceModelsAndManufacturers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceModelsAndManufacturers() })
        if err != nil {
            return err
        }
        m.SetManagedDeviceModelsAndManufacturers(val.(*ManagedDeviceModelsAndManufacturers))
        return nil
    }
    res["mdmEnrolledCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMdmEnrolledCount(val)
        return nil
    }
    return res
}
func (m *ManagedDeviceOverview) IsNil()(bool) {
    return m == nil
}
func (m *ManagedDeviceOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceExchangeAccessStateSummary", m.GetDeviceExchangeAccessStateSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceOperatingSystemSummary", m.GetDeviceOperatingSystemSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("dualEnrolledDeviceCount", m.GetDualEnrolledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("enrolledDeviceCount", m.GetEnrolledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDeviceModelsAndManufacturers", m.GetManagedDeviceModelsAndManufacturers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mdmEnrolledCount", m.GetMdmEnrolledCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ManagedDeviceOverview) SetDeviceExchangeAccessStateSummary(value *DeviceExchangeAccessStateSummary)() {
    m.deviceExchangeAccessStateSummary = value
}
func (m *ManagedDeviceOverview) SetDeviceOperatingSystemSummary(value *DeviceOperatingSystemSummary)() {
    m.deviceOperatingSystemSummary = value
}
func (m *ManagedDeviceOverview) SetDualEnrolledDeviceCount(value *int32)() {
    m.dualEnrolledDeviceCount = value
}
func (m *ManagedDeviceOverview) SetEnrolledDeviceCount(value *int32)() {
    m.enrolledDeviceCount = value
}
func (m *ManagedDeviceOverview) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *ManagedDeviceOverview) SetManagedDeviceModelsAndManufacturers(value *ManagedDeviceModelsAndManufacturers)() {
    m.managedDeviceModelsAndManufacturers = value
}
func (m *ManagedDeviceOverview) SetMdmEnrolledCount(value *int32)() {
    m.mdmEnrolledCount = value
}