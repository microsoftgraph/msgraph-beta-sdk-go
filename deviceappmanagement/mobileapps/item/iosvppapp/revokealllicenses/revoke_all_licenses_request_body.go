package revokealllicenses

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RevokeAllLicensesRequestBody struct {
    additionalData map[string]interface{};
    notifyManagedDevices *bool;
}
func NewRevokeAllLicensesRequestBody()(*RevokeAllLicensesRequestBody) {
    m := &RevokeAllLicensesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RevokeAllLicensesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RevokeAllLicensesRequestBody) GetNotifyManagedDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyManagedDevices
    }
}
func (m *RevokeAllLicensesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["notifyManagedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetNotifyManagedDevices(val)
        return nil
    }
    return res
}
func (m *RevokeAllLicensesRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *RevokeAllLicensesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("notifyManagedDevices", m.GetNotifyManagedDevices())
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
func (m *RevokeAllLicensesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RevokeAllLicensesRequestBody) SetNotifyManagedDevices(value *bool)() {
    m.notifyManagedDevices = value
}
