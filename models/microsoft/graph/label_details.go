package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LabelDetails provides operations to call the extractLabel method.
type LabelDetails struct {
    ParentLabelDetails
}
// NewLabelDetails instantiates a new labelDetails and sets the default values.
func NewLabelDetails()(*LabelDetails) {
    m := &LabelDetails{
        ParentLabelDetails: *NewParentLabelDetails(),
    }
    return m
}
// CreateLabelDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelDetailsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewLabelDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ParentLabelDetails.GetFieldDeserializers()
    return res
}
func (m *LabelDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LabelDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ParentLabelDetails.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
