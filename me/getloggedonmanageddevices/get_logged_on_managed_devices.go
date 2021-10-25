package getloggedonmanageddevices

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GetLoggedOnManagedDevices struct {
    additionalData map[string]interface{};
}
func NewGetLoggedOnManagedDevices()(*GetLoggedOnManagedDevices) {
    m := &GetLoggedOnManagedDevices{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GetLoggedOnManagedDevices) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GetLoggedOnManagedDevices) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *GetLoggedOnManagedDevices) IsNil()(bool) {
    return m == nil
}
func (m *GetLoggedOnManagedDevices) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetLoggedOnManagedDevices) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
