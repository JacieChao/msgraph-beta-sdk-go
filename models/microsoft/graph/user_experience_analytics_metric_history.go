package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// userExperienceAnalyticsMetricHistory 
type UserExperienceAnalyticsMetricHistory struct {
    Entity
    // The user experience analytics device id.
    deviceId *string;
    // The user experience analytics metric date time.
    metricDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The user experience analytics metric type.
    metricType *string;
    // User experience analytics metric.
    userExperienceAnalyticsMetric *UserExperienceAnalyticsMetric;
}
// NewUserExperienceAnalyticsMetricHistory instantiates a new userExperienceAnalyticsMetricHistory and sets the default values.
func NewUserExperienceAnalyticsMetricHistory()(*UserExperienceAnalyticsMetricHistory) {
    m := &UserExperienceAnalyticsMetricHistory{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeviceId gets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsMetricHistory) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetMetricDateTime gets the metricDateTime property value. The user experience analytics metric date time.
func (m *UserExperienceAnalyticsMetricHistory) GetMetricDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.metricDateTime
    }
}
// GetMetricType gets the metricType property value. The user experience analytics metric type.
func (m *UserExperienceAnalyticsMetricHistory) GetMetricType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.metricType
    }
}
// GetUserExperienceAnalyticsMetric gets the userExperienceAnalyticsMetric property value. User experience analytics metric.
func (m *UserExperienceAnalyticsMetricHistory) GetUserExperienceAnalyticsMetric()(*UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsMetric
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsMetricHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["metricDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetricDateTime(val)
        }
        return nil
    }
    res["metricType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetricType(val)
        }
        return nil
    }
    res["userExperienceAnalyticsMetric"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsMetric(val.(*UserExperienceAnalyticsMetric))
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsMetricHistory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsMetricHistory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("metricDateTime", m.GetMetricDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metricType", m.GetMetricType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsMetric", m.GetUserExperienceAnalyticsMetric())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceId sets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsMetricHistory) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetMetricDateTime sets the metricDateTime property value. The user experience analytics metric date time.
func (m *UserExperienceAnalyticsMetricHistory) SetMetricDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.metricDateTime = value
}
// SetMetricType sets the metricType property value. The user experience analytics metric type.
func (m *UserExperienceAnalyticsMetricHistory) SetMetricType(value *string)() {
    m.metricType = value
}
// SetUserExperienceAnalyticsMetric sets the userExperienceAnalyticsMetric property value. User experience analytics metric.
func (m *UserExperienceAnalyticsMetricHistory) SetUserExperienceAnalyticsMetric(value *UserExperienceAnalyticsMetric)() {
    m.userExperienceAnalyticsMetric = value
}
