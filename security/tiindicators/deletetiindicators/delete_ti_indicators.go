package deletetiindicators

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeleteTiIndicators struct {
    additionalData map[string]interface{};
}
func NewDeleteTiIndicators()(*DeleteTiIndicators) {
    m := &DeleteTiIndicators{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeleteTiIndicators) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeleteTiIndicators) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *DeleteTiIndicators) IsNil()(bool) {
    return m == nil
}
func (m *DeleteTiIndicators) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeleteTiIndicators) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
