package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedDeviceModelsAndManufacturers struct {
    additionalData map[string]interface{};
    deviceManufacturers []string;
    deviceModels []string;
}
func NewManagedDeviceModelsAndManufacturers()(*ManagedDeviceModelsAndManufacturers) {
    m := &ManagedDeviceModelsAndManufacturers{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ManagedDeviceModelsAndManufacturers) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ManagedDeviceModelsAndManufacturers) GetDeviceManufacturers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceManufacturers
    }
}
func (m *ManagedDeviceModelsAndManufacturers) GetDeviceModels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceModels
    }
}
func (m *ManagedDeviceModelsAndManufacturers) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceManufacturers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDeviceManufacturers(res)
        return nil
    }
    res["deviceModels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDeviceModels(res)
        return nil
    }
    return res
}
func (m *ManagedDeviceModelsAndManufacturers) IsNil()(bool) {
    return m == nil
}
func (m *ManagedDeviceModelsAndManufacturers) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("deviceManufacturers", m.GetDeviceManufacturers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("deviceModels", m.GetDeviceModels())
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
func (m *ManagedDeviceModelsAndManufacturers) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ManagedDeviceModelsAndManufacturers) SetDeviceManufacturers(value []string)() {
    m.deviceManufacturers = value
}
func (m *ManagedDeviceModelsAndManufacturers) SetDeviceModels(value []string)() {
    m.deviceModels = value
}
