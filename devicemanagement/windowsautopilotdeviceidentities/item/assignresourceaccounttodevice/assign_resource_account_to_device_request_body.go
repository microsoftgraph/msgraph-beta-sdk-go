package assignresourceaccounttodevice

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignResourceAccountToDeviceRequestBody 
type AssignResourceAccountToDeviceRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    addressableUserName *string;
    // 
    resourceAccountName *string;
    // 
    userPrincipalName *string;
}
// NewAssignResourceAccountToDeviceRequestBody instantiates a new assignResourceAccountToDeviceRequestBody and sets the default values.
func NewAssignResourceAccountToDeviceRequestBody()(*AssignResourceAccountToDeviceRequestBody) {
    m := &AssignResourceAccountToDeviceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignResourceAccountToDeviceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddressableUserName gets the addressableUserName property value. 
func (m *AssignResourceAccountToDeviceRequestBody) GetAddressableUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.addressableUserName
    }
}
// GetResourceAccountName gets the resourceAccountName property value. 
func (m *AssignResourceAccountToDeviceRequestBody) GetResourceAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccountName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *AssignResourceAccountToDeviceRequestBody) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignResourceAccountToDeviceRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addressableUserName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressableUserName(val)
        }
        return nil
    }
    res["resourceAccountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAccountName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *AssignResourceAccountToDeviceRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignResourceAccountToDeviceRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddressableUserName sets the addressableUserName property value. 
func (m *AssignResourceAccountToDeviceRequestBody) SetAddressableUserName(value *string)() {
    if m != nil {
        m.addressableUserName = value
    }
}
// SetResourceAccountName sets the resourceAccountName property value. 
func (m *AssignResourceAccountToDeviceRequestBody) SetResourceAccountName(value *string)() {
    if m != nil {
        m.resourceAccountName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *AssignResourceAccountToDeviceRequestBody) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
