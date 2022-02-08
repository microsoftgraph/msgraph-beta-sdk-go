package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RemoteAssistanceSettings 
type RemoteAssistanceSettings struct {
    Entity
    // Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
    allowSessionsToUnenrolledDevices *bool;
    // The current state of remote assistance for the account. Possible values are: disabled, enabled. This setting is configurable by the admin. Remote assistance settings that have not yet been configured by the admin have a disabled state. Returned by default. Possible values are: disabled, enabled.
    remoteAssistanceState *RemoteAssistanceState;
}
// NewRemoteAssistanceSettings instantiates a new remoteAssistanceSettings and sets the default values.
func NewRemoteAssistanceSettings()(*RemoteAssistanceSettings) {
    m := &RemoteAssistanceSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllowSessionsToUnenrolledDevices gets the allowSessionsToUnenrolledDevices property value. Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) GetAllowSessionsToUnenrolledDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowSessionsToUnenrolledDevices
    }
}
// GetRemoteAssistanceState gets the remoteAssistanceState property value. The current state of remote assistance for the account. Possible values are: disabled, enabled. This setting is configurable by the admin. Remote assistance settings that have not yet been configured by the admin have a disabled state. Returned by default. Possible values are: disabled, enabled.
func (m *RemoteAssistanceSettings) GetRemoteAssistanceState()(*RemoteAssistanceState) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoteAssistanceSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowSessionsToUnenrolledDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSessionsToUnenrolledDevices(val)
        }
        return nil
    }
    res["remoteAssistanceState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteAssistanceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceState(val.(*RemoteAssistanceState))
        }
        return nil
    }
    return res
}
func (m *RemoteAssistanceSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RemoteAssistanceSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowSessionsToUnenrolledDevices", m.GetAllowSessionsToUnenrolledDevices())
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAssistanceState() != nil {
        cast := (*m.GetRemoteAssistanceState()).String()
        err = writer.WriteStringValue("remoteAssistanceState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowSessionsToUnenrolledDevices sets the allowSessionsToUnenrolledDevices property value. Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) SetAllowSessionsToUnenrolledDevices(value *bool)() {
    if m != nil {
        m.allowSessionsToUnenrolledDevices = value
    }
}
// SetRemoteAssistanceState sets the remoteAssistanceState property value. The current state of remote assistance for the account. Possible values are: disabled, enabled. This setting is configurable by the admin. Remote assistance settings that have not yet been configured by the admin have a disabled state. Returned by default. Possible values are: disabled, enabled.
func (m *RemoteAssistanceSettings) SetRemoteAssistanceState(value *RemoteAssistanceState)() {
    if m != nil {
        m.remoteAssistanceState = value
    }
}
