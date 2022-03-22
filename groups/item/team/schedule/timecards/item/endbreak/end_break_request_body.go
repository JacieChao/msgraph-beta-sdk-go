package endbreak

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// EndBreakRequestBody provides operations to call the endBreak method.
type EndBreakRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    atApprovedLocation *bool;
    // 
    notes i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBodyable;
}
// NewEndBreakRequestBody instantiates a new endBreakRequestBody and sets the default values.
func NewEndBreakRequestBody()(*EndBreakRequestBody) {
    m := &EndBreakRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEndBreakRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEndBreakRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEndBreakRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EndBreakRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAtApprovedLocation gets the atApprovedLocation property value. 
func (m *EndBreakRequestBody) GetAtApprovedLocation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.atApprovedLocation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EndBreakRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["atApprovedLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtApprovedLocation(val)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBodyable))
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. 
func (m *EndBreakRequestBody) GetNotes()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Serialize serializes information the current object
func (m *EndBreakRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("atApprovedLocation", m.GetAtApprovedLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notes", m.GetNotes())
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
func (m *EndBreakRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAtApprovedLocation sets the atApprovedLocation property value. 
func (m *EndBreakRequestBody) SetAtApprovedLocation(value *bool)() {
    if m != nil {
        m.atApprovedLocation = value
    }
}
// SetNotes sets the notes property value. 
func (m *EndBreakRequestBody) SetNotes(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBodyable)() {
    if m != nil {
        m.notes = value
    }
}