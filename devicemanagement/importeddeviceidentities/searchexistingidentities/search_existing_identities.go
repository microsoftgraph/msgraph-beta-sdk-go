package searchexistingidentities

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchExistingIdentities struct {
    additionalData map[string]interface{};
}
func NewSearchExistingIdentities()(*SearchExistingIdentities) {
    m := &SearchExistingIdentities{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchExistingIdentities) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchExistingIdentities) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *SearchExistingIdentities) IsNil()(bool) {
    return m == nil
}
func (m *SearchExistingIdentities) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SearchExistingIdentities) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
