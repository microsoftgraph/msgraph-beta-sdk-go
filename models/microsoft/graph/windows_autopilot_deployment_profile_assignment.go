package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsAutopilotDeploymentProfileAssignment 
type WindowsAutopilotDeploymentProfileAssignment struct {
    Entity
    // Type of resource used for deployment to a group, direct or parcel/policySet. Possible values are: direct, policySets.
    source *DeviceAndAppManagementAssignmentSource;
    // Identifier for resource used for deployment to a group
    sourceId *string;
    // The assignment target for the Windows Autopilot deployment profile.
    target *DeviceAndAppManagementAssignmentTarget;
}
// NewWindowsAutopilotDeploymentProfileAssignment instantiates a new windowsAutopilotDeploymentProfileAssignment and sets the default values.
func NewWindowsAutopilotDeploymentProfileAssignment()(*WindowsAutopilotDeploymentProfileAssignment) {
    m := &WindowsAutopilotDeploymentProfileAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetSource gets the source property value. Type of resource used for deployment to a group, direct or parcel/policySet. Possible values are: direct, policySets.
func (m *WindowsAutopilotDeploymentProfileAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetSourceId gets the sourceId property value. Identifier for resource used for deployment to a group
func (m *WindowsAutopilotDeploymentProfileAssignment) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTarget gets the target property value. The assignment target for the Windows Autopilot deployment profile.
func (m *WindowsAutopilotDeploymentProfileAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeploymentProfileAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*DeviceAndAppManagementAssignmentSource))
        }
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
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
func (m *WindowsAutopilotDeploymentProfileAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfileAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceId", m.GetSourceId())
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
// SetSource sets the source property value. Type of resource used for deployment to a group, direct or parcel/policySet. Possible values are: direct, policySets.
func (m *WindowsAutopilotDeploymentProfileAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    if m != nil {
        m.source = value
    }
}
// SetSourceId sets the sourceId property value. Identifier for resource used for deployment to a group
func (m *WindowsAutopilotDeploymentProfileAssignment) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetTarget sets the target property value. The assignment target for the Windows Autopilot deployment profile.
func (m *WindowsAutopilotDeploymentProfileAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    if m != nil {
        m.target = value
    }
}
