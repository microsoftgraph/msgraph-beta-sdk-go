package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedAppleDeviceUser 
type SharedAppleDeviceUser struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Data quota
    dataQuota *int64;
    // Data to sync
    dataToSync *bool;
    // Data quota
    dataUsed *int64;
    // User name
    userPrincipalName *string;
}
// NewSharedAppleDeviceUser instantiates a new sharedAppleDeviceUser and sets the default values.
func NewSharedAppleDeviceUser()(*SharedAppleDeviceUser) {
    m := &SharedAppleDeviceUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharedAppleDeviceUser) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDataQuota gets the dataQuota property value. Data quota
func (m *SharedAppleDeviceUser) GetDataQuota()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dataQuota
    }
}
// GetDataToSync gets the dataToSync property value. Data to sync
func (m *SharedAppleDeviceUser) GetDataToSync()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataToSync
    }
}
// GetDataUsed gets the dataUsed property value. Data quota
func (m *SharedAppleDeviceUser) GetDataUsed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dataUsed
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User name
func (m *SharedAppleDeviceUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedAppleDeviceUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataQuota"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataQuota(val)
        }
        return nil
    }
    res["dataToSync"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataToSync(val)
        }
        return nil
    }
    res["dataUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataUsed(val)
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
func (m *SharedAppleDeviceUser) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharedAppleDeviceUser) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDataQuota sets the dataQuota property value. Data quota
func (m *SharedAppleDeviceUser) SetDataQuota(value *int64)() {
    if m != nil {
        m.dataQuota = value
    }
}
// SetDataToSync sets the dataToSync property value. Data to sync
func (m *SharedAppleDeviceUser) SetDataToSync(value *bool)() {
    if m != nil {
        m.dataToSync = value
    }
}
// SetDataUsed sets the dataUsed property value. Data quota
func (m *SharedAppleDeviceUser) SetDataUsed(value *int64)() {
    if m != nil {
        m.dataUsed = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User name
func (m *SharedAppleDeviceUser) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
