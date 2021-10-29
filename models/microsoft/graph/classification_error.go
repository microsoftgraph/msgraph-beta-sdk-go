package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ClassificationError struct {
    ClassifcationErrorBase
    // 
    details []ClassifcationErrorBase;
}
// Instantiates a new classificationError and sets the default values.
func NewClassificationError()(*ClassificationError) {
    m := &ClassificationError{
        ClassifcationErrorBase: *NewClassifcationErrorBase(),
    }
    return m
}
// Gets the details property value. 
func (m *ClassificationError) GetDetails()([]ClassifcationErrorBase) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// The deserialization information for the current model
func (m *ClassificationError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ClassifcationErrorBase.GetFieldDeserializers()
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassifcationErrorBase() })
        if err != nil {
            return err
        }
        res := make([]ClassifcationErrorBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*ClassifcationErrorBase))
        }
        m.SetDetails(res)
        return nil
    }
    return res
}
func (m *ClassificationError) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ClassificationError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ClassifcationErrorBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("details", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the details property value. 
// Parameters:
//  - value : Value to set for the details property.
func (m *ClassificationError) SetDetails(value []ClassifcationErrorBase)() {
    m.details = value
}
