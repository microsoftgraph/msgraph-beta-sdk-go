package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Ediscoveryroot 
type Ediscoveryroot struct {
    Entity
    // 
    cases []Case_escaped;
}
// NewEdiscoveryroot instantiates a new ediscoveryroot and sets the default values.
func NewEdiscoveryroot()(*Ediscoveryroot) {
    m := &Ediscoveryroot{
        Entity: *NewEntity(),
    }
    return m
}
// GetCases gets the cases property value. 
func (m *Ediscoveryroot) GetCases()([]Case_escaped) {
    if m == nil {
        return nil
    } else {
        return m.cases
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Ediscoveryroot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cases"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCase_escaped() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Case_escaped, len(val))
            for i, v := range val {
                res[i] = *(v.(*Case_escaped))
            }
            m.SetCases(res)
        }
        return nil
    }
    return res
}
func (m *Ediscoveryroot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Ediscoveryroot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCases() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCases()))
        for i, v := range m.GetCases() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cases", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCases sets the cases property value. 
func (m *Ediscoveryroot) SetCases(value []Case_escaped)() {
    if m != nil {
        m.cases = value
    }
}
