package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DisplayTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The text identifier for the display template; for example, contosoTickets.
    id *string;
    // 
    layout *Json;
    // Defines the priority of a display template. A display template with priority 1 is evaluated before a template with priority 4. Gaps in priority values are supported.
    priority *int32;
    // Specifies additional rules for selecting this display template based on the item schema. Optional.
    rules []PropertyRule;
}
// Instantiates a new displayTemplate and sets the default values.
func NewDisplayTemplate()(*DisplayTemplate) {
    m := &DisplayTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DisplayTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the id property value. The text identifier for the display template; for example, contosoTickets.
func (m *DisplayTemplate) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the layout property value. 
func (m *DisplayTemplate) GetLayout()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.layout
    }
}
// Gets the priority property value. Defines the priority of a display template. A display template with priority 1 is evaluated before a template with priority 4. Gaps in priority values are supported.
func (m *DisplayTemplate) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// Gets the rules property value. Specifies additional rules for selecting this display template based on the item schema. Optional.
func (m *DisplayTemplate) GetRules()([]PropertyRule) {
    if m == nil {
        return nil
    } else {
        return m.rules
    }
}
// The deserialization information for the current model
func (m *DisplayTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["layout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetLayout(val.(*Json))
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPriority(val)
        return nil
    }
    res["rules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPropertyRule() })
        if err != nil {
            return err
        }
        res := make([]PropertyRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*PropertyRule))
        }
        m.SetRules(res)
        return nil
    }
    return res
}
func (m *DisplayTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DisplayTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("layout", m.GetLayout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DisplayTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the id property value. The text identifier for the display template; for example, contosoTickets.
// Parameters:
//  - value : Value to set for the id property.
func (m *DisplayTemplate) SetId(value *string)() {
    m.id = value
}
// Sets the layout property value. 
// Parameters:
//  - value : Value to set for the layout property.
func (m *DisplayTemplate) SetLayout(value *Json)() {
    m.layout = value
}
// Sets the priority property value. Defines the priority of a display template. A display template with priority 1 is evaluated before a template with priority 4. Gaps in priority values are supported.
// Parameters:
//  - value : Value to set for the priority property.
func (m *DisplayTemplate) SetPriority(value *int32)() {
    m.priority = value
}
// Sets the rules property value. Specifies additional rules for selecting this display template based on the item schema. Optional.
// Parameters:
//  - value : Value to set for the rules property.
func (m *DisplayTemplate) SetRules(value []PropertyRule)() {
    m.rules = value
}