package deletetiindicatorsbyexternalid

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeleteTiIndicatorsByExternalIdRequestBody struct {
    additionalData map[string]interface{};
    value []string;
}
func NewDeleteTiIndicatorsByExternalIdRequestBody()(*DeleteTiIndicatorsByExternalIdRequestBody) {
    m := &DeleteTiIndicatorsByExternalIdRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeleteTiIndicatorsByExternalIdRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeleteTiIndicatorsByExternalIdRequestBody) GetValue()([]string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *DeleteTiIndicatorsByExternalIdRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetValue(res)
        return nil
    }
    return res
}
func (m *DeleteTiIndicatorsByExternalIdRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *DeleteTiIndicatorsByExternalIdRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeleteTiIndicatorsByExternalIdRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeleteTiIndicatorsByExternalIdRequestBody) SetValue(value []string)() {
    m.value = value
}
