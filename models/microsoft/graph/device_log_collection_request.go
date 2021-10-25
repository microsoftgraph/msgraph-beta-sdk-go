package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceLogCollectionRequest struct {
    additionalData map[string]interface{};
    id *string;
    templateType *DeviceLogCollectionTemplateType;
}
func NewDeviceLogCollectionRequest()(*DeviceLogCollectionRequest) {
    m := &DeviceLogCollectionRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceLogCollectionRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceLogCollectionRequest) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *DeviceLogCollectionRequest) GetTemplateType()(*DeviceLogCollectionTemplateType) {
    if m == nil {
        return nil
    } else {
        return m.templateType
    }
}
func (m *DeviceLogCollectionRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["templateType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceLogCollectionTemplateType)
        if err != nil {
            return err
        }
        cast := val.(DeviceLogCollectionTemplateType)
        m.SetTemplateType(&cast)
        return nil
    }
    return res
}
func (m *DeviceLogCollectionRequest) IsNil()(bool) {
    return m == nil
}
func (m *DeviceLogCollectionRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateType() != nil {
        cast := m.GetTemplateType().String()
        err := writer.WriteStringValue("templateType", &cast)
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
func (m *DeviceLogCollectionRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceLogCollectionRequest) SetId(value *string)() {
    m.id = value
}
func (m *DeviceLogCollectionRequest) SetTemplateType(value *DeviceLogCollectionTemplateType)() {
    m.templateType = value
}
