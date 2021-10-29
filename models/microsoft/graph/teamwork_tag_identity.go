package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamworkTagIdentity struct {
    Identity
}
func NewTeamworkTagIdentity()(*TeamworkTagIdentity) {
    m := &TeamworkTagIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
func (m *TeamworkTagIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    return res
}
func (m *TeamworkTagIdentity) IsNil()(bool) {
    return m == nil
}
func (m *TeamworkTagIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}