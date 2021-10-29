package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/termstore"
)

type Relation struct {
    Entity
    fromTerm *Term;
    relationship *i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.RelationType;
    set *Set;
    toTerm *Term;
}
func NewRelation()(*Relation) {
    m := &Relation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Relation) GetFromTerm()(*Term) {
    if m == nil {
        return nil
    } else {
        return m.fromTerm
    }
}
func (m *Relation) GetRelationship()(*i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.RelationType) {
    if m == nil {
        return nil
    } else {
        return m.relationship
    }
}
func (m *Relation) GetSet()(*Set) {
    if m == nil {
        return nil
    } else {
        return m.set
    }
}
func (m *Relation) GetToTerm()(*Term) {
    if m == nil {
        return nil
    } else {
        return m.toTerm
    }
}
func (m *Relation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fromTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        m.SetFromTerm(val.(*Term))
        return nil
    }
    res["relationship"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.ParseRelationType)
        if err != nil {
            return err
        }
        cast := val.(i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.RelationType)
        m.SetRelationship(&cast)
        return nil
    }
    res["set"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSet() })
        if err != nil {
            return err
        }
        m.SetSet(val.(*Set))
        return nil
    }
    res["toTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        m.SetToTerm(val.(*Term))
        return nil
    }
    return res
}
func (m *Relation) IsNil()(bool) {
    return m == nil
}
func (m *Relation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fromTerm", m.GetFromTerm())
        if err != nil {
            return err
        }
    }
    if m.GetRelationship() != nil {
        cast := m.GetRelationship().String()
        err = writer.WriteStringValue("relationship", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("set", m.GetSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("toTerm", m.GetToTerm())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Relation) SetFromTerm(value *Term)() {
    m.fromTerm = value
}
func (m *Relation) SetRelationship(value *i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.RelationType)() {
    m.relationship = value
}
func (m *Relation) SetSet(value *Set)() {
    m.set = value
}
func (m *Relation) SetToTerm(value *Term)() {
    m.toTerm = value
}