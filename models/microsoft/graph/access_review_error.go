package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewError 
type AccessReviewError struct {
    GenericError
}
// NewAccessReviewError instantiates a new accessReviewError and sets the default values.
func NewAccessReviewError()(*AccessReviewError) {
    m := &AccessReviewError{
        GenericError: *NewGenericError(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.GenericError.GetFieldDeserializers()
    return res
}
func (m *AccessReviewError) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.GenericError.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
