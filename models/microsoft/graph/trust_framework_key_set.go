package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TrustFrameworkKeySet struct {
    Entity
    keys []TrustFrameworkKey;
}
func NewTrustFrameworkKeySet()(*TrustFrameworkKeySet) {
    m := &TrustFrameworkKeySet{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TrustFrameworkKeySet) GetKeys()([]TrustFrameworkKey) {
    if m == nil {
        return nil
    } else {
        return m.keys
    }
}
func (m *TrustFrameworkKeySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["keys"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTrustFrameworkKey() })
        if err != nil {
            return err
        }
        res := make([]TrustFrameworkKey, len(val))
        for i, v := range val {
            res[i] = *(v.(*TrustFrameworkKey))
        }
        m.SetKeys(res)
        return nil
    }
    return res
}
func (m *TrustFrameworkKeySet) IsNil()(bool) {
    return m == nil
}
func (m *TrustFrameworkKeySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetKeys()))
        for i, v := range m.GetKeys() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("keys", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TrustFrameworkKeySet) SetKeys(value []TrustFrameworkKey)() {
    m.keys = value
}
