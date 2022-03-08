package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementScriptUserState provides operations to manage the deviceManagement singleton.
type DeviceManagementScriptUserState struct {
    Entity
    // List of run states for this script across all devices of specific user.
    deviceRunStates []DeviceManagementScriptDeviceStateable;
    // Error device count for specific user.
    errorDeviceCount *int32;
    // Success device count for specific user.
    successDeviceCount *int32;
    // User principle name of specific user.
    userPrincipalName *string;
}
// NewDeviceManagementScriptUserState instantiates a new deviceManagementScriptUserState and sets the default values.
func NewDeviceManagementScriptUserState()(*DeviceManagementScriptUserState) {
    m := &DeviceManagementScriptUserState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementScriptUserStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptUserStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementScriptUserState(), nil
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for this script across all devices of specific user.
func (m *DeviceManagementScriptUserState) GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Error device count for specific user.
func (m *DeviceManagementScriptUserState) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptUserState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptDeviceStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementScriptDeviceStateable)
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
// GetSuccessDeviceCount gets the successDeviceCount property value. Success device count for specific user.
func (m *DeviceManagementScriptUserState) GetSuccessDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successDeviceCount
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User principle name of specific user.
func (m *DeviceManagementScriptUserState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *DeviceManagementScriptUserState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptUserState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeviceRunStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for this script across all devices of specific user.
func (m *DeviceManagementScriptUserState) SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)() {
    if m != nil {
        m.deviceRunStates = value
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Error device count for specific user.
func (m *DeviceManagementScriptUserState) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetSuccessDeviceCount sets the successDeviceCount property value. Success device count for specific user.
func (m *DeviceManagementScriptUserState) SetSuccessDeviceCount(value *int32)() {
    if m != nil {
        m.successDeviceCount = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principle name of specific user.
func (m *DeviceManagementScriptUserState) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
