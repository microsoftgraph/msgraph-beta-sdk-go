package assignresourceaccounttodevice

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignResourceAccountToDeviceRequestBody struct {
    additionalData map[string]interface{};
    addressableUserName *string;
    resourceAccountName *string;
    userPrincipalName *string;
}
func NewAssignResourceAccountToDeviceRequestBody()(*AssignResourceAccountToDeviceRequestBody) {
    m := &AssignResourceAccountToDeviceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssignResourceAccountToDeviceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssignResourceAccountToDeviceRequestBody) GetAddressableUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.addressableUserName
    }
}
func (m *AssignResourceAccountToDeviceRequestBody) GetResourceAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccountName
    }
}
func (m *AssignResourceAccountToDeviceRequestBody) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *AssignResourceAccountToDeviceRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addressableUserName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAddressableUserName(val)
        return nil
    }
    res["resourceAccountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceAccountName(val)
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
func (m *AssignResourceAccountToDeviceRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *AssignResourceAccountToDeviceRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceAccountName", m.GetResourceAccountName())
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
func (m *AssignResourceAccountToDeviceRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssignResourceAccountToDeviceRequestBody) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
func (m *AssignResourceAccountToDeviceRequestBody) SetResourceAccountName(value *string)() {
    m.resourceAccountName = value
}
func (m *AssignResourceAccountToDeviceRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
