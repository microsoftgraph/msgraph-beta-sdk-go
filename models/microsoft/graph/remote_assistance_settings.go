package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RemoteAssistanceSettings struct {
    Entity
    allowSessionsToUnenrolledDevices *bool;
    remoteAssistanceState *RemoteAssistanceState;
}
func NewRemoteAssistanceSettings()(*RemoteAssistanceSettings) {
    m := &RemoteAssistanceSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *RemoteAssistanceSettings) GetAllowSessionsToUnenrolledDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowSessionsToUnenrolledDevices
    }
}
func (m *RemoteAssistanceSettings) GetRemoteAssistanceState()(*RemoteAssistanceState) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceState
    }
}
func (m *RemoteAssistanceSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowSessionsToUnenrolledDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowSessionsToUnenrolledDevices(val)
        return nil
    }
    res["remoteAssistanceState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteAssistanceState)
        if err != nil {
            return err
        }
        cast := val.(RemoteAssistanceState)
        m.SetRemoteAssistanceState(&cast)
        return nil
    }
    return res
}
func (m *RemoteAssistanceSettings) IsNil()(bool) {
    return m == nil
}
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
func (m *RemoteAssistanceSettings) SetAllowSessionsToUnenrolledDevices(value *bool)() {
    m.allowSessionsToUnenrolledDevices = value
}
func (m *RemoteAssistanceSettings) SetRemoteAssistanceState(value *RemoteAssistanceState)() {
    m.remoteAssistanceState = value
}
