package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ResultTemplateDictionary struct {
    Dictionary
}
func NewResultTemplateDictionary()(*ResultTemplateDictionary) {
    m := &ResultTemplateDictionary{
        Dictionary: *NewDictionary(),
    }
    return m
}
func (m *ResultTemplateDictionary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Dictionary.GetFieldDeserializers()
    return res
}
func (m *ResultTemplateDictionary) IsNil()(bool) {
    return m == nil
}
func (m *ResultTemplateDictionary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Dictionary.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
