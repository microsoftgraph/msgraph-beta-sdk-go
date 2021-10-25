package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedDeviceReportedApp struct {
    additionalData map[string]interface{};
    appId *string;
}
func NewManagedDeviceReportedApp()(*ManagedDeviceReportedApp) {
    m := &ManagedDeviceReportedApp{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ManagedDeviceReportedApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ManagedDeviceReportedApp) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
func (m *ManagedDeviceReportedApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    return res
}
func (m *ManagedDeviceReportedApp) IsNil()(bool) {
    return m == nil
}
func (m *ManagedDeviceReportedApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
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
func (m *ManagedDeviceReportedApp) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ManagedDeviceReportedApp) SetAppId(value *string)() {
    m.appId = value
}
