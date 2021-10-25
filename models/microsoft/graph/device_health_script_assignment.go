package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthScriptAssignment struct {
    Entity
    runRemediationScript *bool;
    runSchedule *DeviceHealthScriptRunSchedule;
    target *DeviceAndAppManagementAssignmentTarget;
}
func NewDeviceHealthScriptAssignment()(*DeviceHealthScriptAssignment) {
    m := &DeviceHealthScriptAssignment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceHealthScriptAssignment) GetRunRemediationScript()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.runRemediationScript
    }
}
func (m *DeviceHealthScriptAssignment) GetRunSchedule()(*DeviceHealthScriptRunSchedule) {
    if m == nil {
        return nil
    } else {
        return m.runSchedule
    }
}
func (m *DeviceHealthScriptAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *DeviceHealthScriptAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["runRemediationScript"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRunRemediationScript(val)
        return nil
    }
    res["runSchedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptRunSchedule() })
        if err != nil {
            return err
        }
        m.SetRunSchedule(val.(*DeviceHealthScriptRunSchedule))
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        m.SetTarget(val.(*DeviceAndAppManagementAssignmentTarget))
        return nil
    }
    return res
}
func (m *DeviceHealthScriptAssignment) IsNil()(bool) {
    return m == nil
}
func (m *DeviceHealthScriptAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("runRemediationScript", m.GetRunRemediationScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("runSchedule", m.GetRunSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceHealthScriptAssignment) SetRunRemediationScript(value *bool)() {
    m.runRemediationScript = value
}
func (m *DeviceHealthScriptAssignment) SetRunSchedule(value *DeviceHealthScriptRunSchedule)() {
    m.runSchedule = value
}
func (m *DeviceHealthScriptAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    m.target = value
}
