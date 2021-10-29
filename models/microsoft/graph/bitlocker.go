package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Bitlocker struct {
    Entity
    recoveryKeys []BitlockerRecoveryKey;
}
func NewBitlocker()(*Bitlocker) {
    m := &Bitlocker{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Bitlocker) GetRecoveryKeys()([]BitlockerRecoveryKey) {
    if m == nil {
        return nil
    } else {
        return m.recoveryKeys
    }
}
func (m *Bitlocker) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["recoveryKeys"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBitlockerRecoveryKey() })
        if err != nil {
            return err
        }
        res := make([]BitlockerRecoveryKey, len(val))
        for i, v := range val {
            res[i] = *(v.(*BitlockerRecoveryKey))
        }
        m.SetRecoveryKeys(res)
        return nil
    }
    return res
}
func (m *Bitlocker) IsNil()(bool) {
    return m == nil
}
func (m *Bitlocker) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecoveryKeys()))
        for i, v := range m.GetRecoveryKeys() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("recoveryKeys", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Bitlocker) SetRecoveryKeys(value []BitlockerRecoveryKey)() {
    m.recoveryKeys = value
}