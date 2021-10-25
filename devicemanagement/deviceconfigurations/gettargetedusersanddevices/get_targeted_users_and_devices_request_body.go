package gettargetedusersanddevices

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GetTargetedUsersAndDevicesRequestBody struct {
    additionalData map[string]interface{};
    deviceConfigurationIds []string;
}
func NewGetTargetedUsersAndDevicesRequestBody()(*GetTargetedUsersAndDevicesRequestBody) {
    m := &GetTargetedUsersAndDevicesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GetTargetedUsersAndDevicesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GetTargetedUsersAndDevicesRequestBody) GetDeviceConfigurationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationIds
    }
}
func (m *GetTargetedUsersAndDevicesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceConfigurationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDeviceConfigurationIds(res)
        return nil
    }
    return res
}
func (m *GetTargetedUsersAndDevicesRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *GetTargetedUsersAndDevicesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("deviceConfigurationIds", m.GetDeviceConfigurationIds())
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
func (m *GetTargetedUsersAndDevicesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GetTargetedUsersAndDevicesRequestBody) SetDeviceConfigurationIds(value []string)() {
    m.deviceConfigurationIds = value
}
