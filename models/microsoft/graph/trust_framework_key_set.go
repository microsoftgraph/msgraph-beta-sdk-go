package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TrustFrameworkKeySet 
type TrustFrameworkKeySet struct {
    Entity
    // A collection of the keys.
    keys []TrustFrameworkKey;
}
// NewTrustFrameworkKeySet instantiates a new trustFrameworkKeySet and sets the default values.
func NewTrustFrameworkKeySet()(*TrustFrameworkKeySet) {
    m := &TrustFrameworkKeySet{
        Entity: *NewEntity(),
    }
    return m
}
// GetKeys gets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) GetKeys()([]TrustFrameworkKey) {
    if m == nil {
        return nil
    } else {
        return m.keys
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkKeySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["keys"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTrustFrameworkKey() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkKey, len(val))
            for i, v := range val {
                res[i] = *(v.(*TrustFrameworkKey))
            }
            m.SetKeys(res)
        }
        return nil
    }
    return res
}
func (m *TrustFrameworkKeySet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TrustFrameworkKeySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetKeys() != nil {
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
// SetKeys sets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) SetKeys(value []TrustFrameworkKey)() {
    if m != nil {
        m.keys = value
    }
}
