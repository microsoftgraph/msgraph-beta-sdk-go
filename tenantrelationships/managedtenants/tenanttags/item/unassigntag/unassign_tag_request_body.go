package unassigntag

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnassignTagRequestBody struct {
    additionalData map[string]interface{};
    tenantIds []string;
}
func NewUnassignTagRequestBody()(*UnassignTagRequestBody) {
    m := &UnassignTagRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UnassignTagRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UnassignTagRequestBody) GetTenantIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tenantIds
    }
}
func (m *UnassignTagRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["tenantIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTenantIds(res)
        return nil
    }
    return res
}
func (m *UnassignTagRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *UnassignTagRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("tenantIds", m.GetTenantIds())
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
func (m *UnassignTagRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UnassignTagRequestBody) SetTenantIds(value []string)() {
    m.tenantIds = value
}
