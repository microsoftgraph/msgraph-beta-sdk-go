package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemInsights 
type ItemInsights struct {
    OfficeGraphInsights
}
// NewItemInsights instantiates a new itemInsights and sets the default values.
func NewItemInsights()(*ItemInsights) {
    m := &ItemInsights{
        OfficeGraphInsights: *NewOfficeGraphInsights(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInsights) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OfficeGraphInsights.GetFieldDeserializers()
    return res
}
func (m *ItemInsights) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ItemInsights) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OfficeGraphInsights.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
