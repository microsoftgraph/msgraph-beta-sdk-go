package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ClassificationError provides operations to manage the dataClassificationService singleton.
type ClassificationError struct {
    ClassifcationErrorBase
    // 
    details []ClassifcationErrorBaseable;
}
// NewClassificationError instantiates a new classificationError and sets the default values.
func NewClassificationError()(*ClassificationError) {
    m := &ClassificationError{
        ClassifcationErrorBase: *NewClassifcationErrorBase(),
    }
    return m
}
// CreateClassificationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationErrorFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewClassificationError(), nil
}
// GetDetails gets the details property value. 
func (m *ClassificationError) GetDetails()([]ClassifcationErrorBaseable) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ClassifcationErrorBase.GetFieldDeserializers()
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClassifcationErrorBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassifcationErrorBaseable, len(val))
            for i, v := range val {
                res[i] = v.(ClassifcationErrorBaseable)
            }
            m.SetDetails(res)
        }
        return nil
    }
    return res
}
func (m *ClassificationError) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ClassificationError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ClassifcationErrorBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("details", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetails sets the details property value. 
func (m *ClassificationError) SetDetails(value []ClassifcationErrorBaseable)() {
    if m != nil {
        m.details = value
    }
}
