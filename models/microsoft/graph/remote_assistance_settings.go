package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RemoteAssistanceSettings struct {
    Entity
    // Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
    allowSessionsToUnenrolledDevices *bool;
    // The current state of remote assistance for the account. Possible values are: notConfigured, disabled, enabled. This setting is configurable by the admin. Remote assistance settings that have not yet been configured by the admin have a notConfigured state. Returned by default. Possible values are: notConfigured, disabled, enabled.
    remoteAssistanceState *RemoteAssistanceState;
}
// Instantiates a new remoteAssistanceSettings and sets the default values.
func NewRemoteAssistanceSettings()(*RemoteAssistanceSettings) {
    m := &RemoteAssistanceSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allowSessionsToUnenrolledDevices property value. Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) GetAllowSessionsToUnenrolledDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowSessionsToUnenrolledDevices
    }
}
// Gets the remoteAssistanceState property value. The current state of remote assistance for the account. Possible values are: notConfigured, disabled, enabled. This setting is configurable by the admin. Remote assistance settings that have not yet been configured by the admin have a notConfigured state. Returned by default. Possible values are: notConfigured, disabled, enabled.
func (m *RemoteAssistanceSettings) GetRemoteAssistanceState()(*RemoteAssistanceState) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceState
    }
}
// The deserialization information for the current model
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
            cast := val.(RemoteAssistanceState)
            m.SetRemoteAssistanceState(&cast)
        }
        return nil
    }
    return res
}
func (m *RemoteAssistanceSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetRemoteAssistanceState().String()
        err = writer.WriteStringValue("remoteAssistanceState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowSessionsToUnenrolledDevices property value. Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
// Parameters:
//  - value : Value to set for the allowSessionsToUnenrolledDevices property.
func (m *RemoteAssistanceSettings) SetAllowSessionsToUnenrolledDevices(value *bool)() {
    m.allowSessionsToUnenrolledDevices = value
}
// Sets the remoteAssistanceState property value. The current state of remote assistance for the account. Possible values are: notConfigured, disabled, enabled. This setting is configurable by the admin. Remote assistance settings that have not yet been configured by the admin have a notConfigured state. Returned by default. Possible values are: notConfigured, disabled, enabled.
// Parameters:
//  - value : Value to set for the remoteAssistanceState property.
func (m *RemoteAssistanceSettings) SetRemoteAssistanceState(value *RemoteAssistanceState)() {
    m.remoteAssistanceState = value
}
