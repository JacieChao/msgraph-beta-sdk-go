package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExactMatchDataStore struct {
    ExactMatchDataStoreBase
    sessions []ExactMatchSession;
}
func NewExactMatchDataStore()(*ExactMatchDataStore) {
    m := &ExactMatchDataStore{
        ExactMatchDataStoreBase: *NewExactMatchDataStoreBase(),
    }
    return m
}
func (m *ExactMatchDataStore) GetSessions()([]ExactMatchSession) {
    if m == nil {
        return nil
    } else {
        return m.sessions
    }
}
func (m *ExactMatchDataStore) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ExactMatchDataStoreBase.GetFieldDeserializers()
    res["sessions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchSession() })
        if err != nil {
            return err
        }
        res := make([]ExactMatchSession, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExactMatchSession))
        }
        m.SetSessions(res)
        return nil
    }
    return res
}
func (m *ExactMatchDataStore) IsNil()(bool) {
    return m == nil
}
func (m *ExactMatchDataStore) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ExactMatchDataStoreBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ExactMatchDataStore) SetSessions(value []ExactMatchSession)() {
    m.sessions = value
}