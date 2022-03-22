package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TrustFrameworkKeySet 
type TrustFrameworkKeySet struct {
    Entity
    // A collection of the keys.
    keys []TrustFrameworkKeyable;
}
// NewTrustFrameworkKeySet instantiates a new trustFrameworkKeySet and sets the default values.
func NewTrustFrameworkKeySet()(*TrustFrameworkKeySet) {
    m := &TrustFrameworkKeySet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrustFrameworkKeySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkKeySetFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTrustFrameworkKeySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkKeySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["keys"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkKeyable, len(val))
            for i, v := range val {
                res[i] = v.(TrustFrameworkKeyable)
            }
            m.SetKeys(res)
        }
        return nil
    }
    return res
}
// GetKeys gets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) GetKeys()([]TrustFrameworkKeyable) {
    if m == nil {
        return nil
    } else {
        return m.keys
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("keys", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeys sets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) SetKeys(value []TrustFrameworkKeyable)() {
    if m != nil {
        m.keys = value
    }
}
