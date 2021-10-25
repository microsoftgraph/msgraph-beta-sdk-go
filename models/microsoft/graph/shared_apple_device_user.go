package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SharedAppleDeviceUser struct {
    additionalData map[string]interface{};
    dataQuota *int64;
    dataToSync *bool;
    dataUsed *int64;
    userPrincipalName *string;
}
func NewSharedAppleDeviceUser()(*SharedAppleDeviceUser) {
    m := &SharedAppleDeviceUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SharedAppleDeviceUser) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SharedAppleDeviceUser) GetDataQuota()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dataQuota
    }
}
func (m *SharedAppleDeviceUser) GetDataToSync()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataToSync
    }
}
func (m *SharedAppleDeviceUser) GetDataUsed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dataUsed
    }
}
func (m *SharedAppleDeviceUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *SharedAppleDeviceUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataQuota"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDataQuota(val)
        return nil
    }
    res["dataToSync"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDataToSync(val)
        return nil
    }
    res["dataUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDataUsed(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *SharedAppleDeviceUser) IsNil()(bool) {
    return m == nil
}
func (m *SharedAppleDeviceUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("dataQuota", m.GetDataQuota())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dataToSync", m.GetDataToSync())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("dataUsed", m.GetDataUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *SharedAppleDeviceUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SharedAppleDeviceUser) SetDataQuota(value *int64)() {
    m.dataQuota = value
}
func (m *SharedAppleDeviceUser) SetDataToSync(value *bool)() {
    m.dataToSync = value
}
func (m *SharedAppleDeviceUser) SetDataUsed(value *int64)() {
    m.dataUsed = value
}
func (m *SharedAppleDeviceUser) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
