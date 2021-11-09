package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementScriptUserState struct {
    Entity
    // List of run states for this script across all devices of specific user.
    deviceRunStates []DeviceManagementScriptDeviceState;
    // Error device count for specific user.
    errorDeviceCount *int32;
    // Success device count for specific user.
    successDeviceCount *int32;
    // User principle name of specific user.
    userPrincipalName *string;
}
// Instantiates a new deviceManagementScriptUserState and sets the default values.
func NewDeviceManagementScriptUserState()(*DeviceManagementScriptUserState) {
    m := &DeviceManagementScriptUserState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceRunStates property value. List of run states for this script across all devices of specific user.
func (m *DeviceManagementScriptUserState) GetDeviceRunStates()([]DeviceManagementScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// Gets the errorDeviceCount property value. Error device count for specific user.
func (m *DeviceManagementScriptUserState) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// Gets the successDeviceCount property value. Success device count for specific user.
func (m *DeviceManagementScriptUserState) GetSuccessDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successDeviceCount
    }
}
// Gets the userPrincipalName property value. User principle name of specific user.
func (m *DeviceManagementScriptUserState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *DeviceManagementScriptUserState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptDeviceState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptDeviceState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementScriptDeviceState))
            }
            m.SetDeviceRunStates(res)
        }
        return nil
    }
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["successDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessDeviceCount(val)
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
func (m *DeviceManagementScriptUserState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementScriptUserState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceRunStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successDeviceCount", m.GetSuccessDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceRunStates property value. List of run states for this script across all devices of specific user.
// Parameters:
//  - value : Value to set for the deviceRunStates property.
func (m *DeviceManagementScriptUserState) SetDeviceRunStates(value []DeviceManagementScriptDeviceState)() {
    m.deviceRunStates = value
}
// Sets the errorDeviceCount property value. Error device count for specific user.
// Parameters:
//  - value : Value to set for the errorDeviceCount property.
func (m *DeviceManagementScriptUserState) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// Sets the successDeviceCount property value. Success device count for specific user.
// Parameters:
//  - value : Value to set for the successDeviceCount property.
func (m *DeviceManagementScriptUserState) SetSuccessDeviceCount(value *int32)() {
    m.successDeviceCount = value
}
// Sets the userPrincipalName property value. User principle name of specific user.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *DeviceManagementScriptUserState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
