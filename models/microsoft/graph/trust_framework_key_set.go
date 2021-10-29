package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TrustFrameworkKeySet struct {
    Entity
    // A collection of the keys.
    keys []TrustFrameworkKey;
}
// Instantiates a new trustFrameworkKeySet and sets the default values.
func NewTrustFrameworkKeySet()(*TrustFrameworkKeySet) {
    m := &TrustFrameworkKeySet{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) GetKeys()([]TrustFrameworkKey) {
    if m == nil {
        return nil
    } else {
        return m.keys
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the keys property value. A collection of the keys.
// Parameters:
//  - value : Value to set for the keys property.
func (m *TrustFrameworkKeySet) SetKeys(value []TrustFrameworkKey)() {
    m.keys = value
}
