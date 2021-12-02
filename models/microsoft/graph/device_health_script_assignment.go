package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScriptAssignment 
type DeviceHealthScriptAssignment struct {
    Entity
    // Determine whether we want to run detection script only or run both detection script and remediation script
    runRemediationScript *bool;
    // Script run schedule for the target group
    runSchedule *DeviceHealthScriptRunSchedule;
    // The Azure Active Directory group we are targeting the script to
    target *DeviceAndAppManagementAssignmentTarget;
}
// NewDeviceHealthScriptAssignment instantiates a new deviceHealthScriptAssignment and sets the default values.
func NewDeviceHealthScriptAssignment()(*DeviceHealthScriptAssignment) {
    m := &DeviceHealthScriptAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetRunRemediationScript gets the runRemediationScript property value. Determine whether we want to run detection script only or run both detection script and remediation script
func (m *DeviceHealthScriptAssignment) GetRunRemediationScript()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.runRemediationScript
    }
}
// GetRunSchedule gets the runSchedule property value. Script run schedule for the target group
func (m *DeviceHealthScriptAssignment) GetRunSchedule()(*DeviceHealthScriptRunSchedule) {
    if m == nil {
        return nil
    } else {
        return m.runSchedule
    }
}
// GetTarget gets the target property value. The Azure Active Directory group we are targeting the script to
func (m *DeviceHealthScriptAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["runRemediationScript"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunRemediationScript(val)
        }
        return nil
    }
    res["runSchedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptRunSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSchedule(val.(*DeviceHealthScriptRunSchedule))
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*DeviceAndAppManagementAssignmentTarget))
        }
        return nil
    }
    return res
}
func (m *DeviceHealthScriptAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetRunRemediationScript sets the runRemediationScript property value. Determine whether we want to run detection script only or run both detection script and remediation script
func (m *DeviceHealthScriptAssignment) SetRunRemediationScript(value *bool)() {
    if m != nil {
        m.runRemediationScript = value
    }
}
// SetRunSchedule sets the runSchedule property value. Script run schedule for the target group
func (m *DeviceHealthScriptAssignment) SetRunSchedule(value *DeviceHealthScriptRunSchedule)() {
    if m != nil {
        m.runSchedule = value
    }
}
// SetTarget sets the target property value. The Azure Active Directory group we are targeting the script to
func (m *DeviceHealthScriptAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    if m != nil {
        m.target = value
    }
}
