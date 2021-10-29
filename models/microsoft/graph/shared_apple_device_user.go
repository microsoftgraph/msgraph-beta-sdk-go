package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new sharedAppleDeviceUser and sets the default values.
func NewSharedAppleDeviceUser()(*SharedAppleDeviceUser) {
    m := &SharedAppleDeviceUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharedAppleDeviceUser) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the dataQuota property value. Data quota
func (m *SharedAppleDeviceUser) GetDataQuota()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dataQuota
    }
}
// Gets the dataToSync property value. Data to sync
func (m *SharedAppleDeviceUser) GetDataToSync()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataToSync
    }
}
// Gets the dataUsed property value. Data quota
func (m *SharedAppleDeviceUser) GetDataUsed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dataUsed
    }
}
// Gets the userPrincipalName property value. User name
func (m *SharedAppleDeviceUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SharedAppleDeviceUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the dataQuota property value. Data quota
// Parameters:
//  - value : Value to set for the dataQuota property.
func (m *SharedAppleDeviceUser) SetDataQuota(value *int64)() {
    m.dataQuota = value
}
// Sets the dataToSync property value. Data to sync
// Parameters:
//  - value : Value to set for the dataToSync property.
func (m *SharedAppleDeviceUser) SetDataToSync(value *bool)() {
    m.dataToSync = value
}
// Sets the dataUsed property value. Data quota
// Parameters:
//  - value : Value to set for the dataUsed property.
func (m *SharedAppleDeviceUser) SetDataUsed(value *int64)() {
    m.dataUsed = value
}
// Sets the userPrincipalName property value. User name
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *SharedAppleDeviceUser) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
