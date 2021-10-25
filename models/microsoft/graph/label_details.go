package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LabelDetails struct {
    ParentLabelDetails
}
func NewLabelDetails()(*LabelDetails) {
    m := &LabelDetails{
        ParentLabelDetails: *NewParentLabelDetails(),
    }
    return m
}
func (m *LabelDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ParentLabelDetails.GetFieldDeserializers()
    return res
}
func (m *LabelDetails) IsNil()(bool) {
    return m == nil
}
func (m *LabelDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ParentLabelDetails.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
