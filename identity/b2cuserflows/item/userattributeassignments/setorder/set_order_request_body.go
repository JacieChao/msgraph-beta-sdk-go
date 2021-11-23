package setorder

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// setOrderRequestBody 
type SetOrderRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    newAssignmentOrder *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentOrder;
}
// NewSetOrderRequestBody instantiates a new setOrderRequestBody and sets the default values.
func NewSetOrderRequestBody()(*SetOrderRequestBody) {
    m := &SetOrderRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetOrderRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetNewAssignmentOrder gets the newAssignmentOrder property value. 
func (m *SetOrderRequestBody) GetNewAssignmentOrder()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentOrder) {
    if m == nil {
        return nil
    } else {
        return m.newAssignmentOrder
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetOrderRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newAssignmentOrder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAssignmentOrder() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewAssignmentOrder(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentOrder))
        }
        return nil
    }
    return res
}
func (m *SetOrderRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SetOrderRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newAssignmentOrder", m.GetNewAssignmentOrder())
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
func (m *SetOrderRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewAssignmentOrder sets the newAssignmentOrder property value. 
func (m *SetOrderRequestBody) SetNewAssignmentOrder(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentOrder)() {
    m.newAssignmentOrder = value
}
