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
    runSchedule DeviceHealthScriptRunScheduleable;
    // The Azure Active Directory group we are targeting the script to
    target DeviceAndAppManagementAssignmentTargetable;
}
// NewDeviceHealthScriptAssignment instantiates a new deviceHealthScriptAssignment and sets the default values.
func NewDeviceHealthScriptAssignment()(*DeviceHealthScriptAssignment) {
    m := &DeviceHealthScriptAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceHealthScriptAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceHealthScriptAssignment(), nil
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
        val, err := n.GetObjectValue(CreateDeviceHealthScriptRunScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSchedule(val.(DeviceHealthScriptRunScheduleable))
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
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
func (m *DeviceHealthScriptAssignment) GetRunSchedule()(DeviceHealthScriptRunScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.runSchedule
    }
}
// GetTarget gets the target property value. The Azure Active Directory group we are targeting the script to
func (m *DeviceHealthScriptAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
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
func (m *DeviceHealthScriptAssignment) SetRunSchedule(value DeviceHealthScriptRunScheduleable)() {
    if m != nil {
        m.runSchedule = value
    }
}
// SetTarget sets the target property value. The Azure Active Directory group we are targeting the script to
func (m *DeviceHealthScriptAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
