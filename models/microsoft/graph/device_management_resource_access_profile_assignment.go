package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementResourceAccessProfileAssignment 
type DeviceManagementResourceAccessProfileAssignment struct {
    Entity
    // The assignment intent for the resource access profile. Possible values are: apply, remove.
    intent *DeviceManagementResourceAccessProfileIntent;
    // The identifier of the source of the assignment.
    sourceId *string;
    // Base type for assignment targets.
    target *DeviceAndAppManagementAssignmentTarget;
}
// NewDeviceManagementResourceAccessProfileAssignment instantiates a new deviceManagementResourceAccessProfileAssignment and sets the default values.
func NewDeviceManagementResourceAccessProfileAssignment()(*DeviceManagementResourceAccessProfileAssignment) {
    m := &DeviceManagementResourceAccessProfileAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetIntent gets the intent property value. The assignment intent for the resource access profile. Possible values are: apply, remove.
func (m *DeviceManagementResourceAccessProfileAssignment) GetIntent()(*DeviceManagementResourceAccessProfileIntent) {
    if m == nil {
        return nil
    } else {
        return m.intent
    }
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementResourceAccessProfileAssignment) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTarget gets the target property value. Base type for assignment targets.
func (m *DeviceManagementResourceAccessProfileAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementResourceAccessProfileAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementResourceAccessProfileIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(*DeviceManagementResourceAccessProfileIntent))
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
func (m *DeviceManagementResourceAccessProfileAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementResourceAccessProfileAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := (*m.GetIntent()).String()
        err = writer.WriteStringValue("intent", &cast)
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
// SetIntent sets the intent property value. The assignment intent for the resource access profile. Possible values are: apply, remove.
func (m *DeviceManagementResourceAccessProfileAssignment) SetIntent(value *DeviceManagementResourceAccessProfileIntent)() {
    if m != nil {
        m.intent = value
    }
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementResourceAccessProfileAssignment) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetTarget sets the target property value. Base type for assignment targets.
func (m *DeviceManagementResourceAccessProfileAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    if m != nil {
        m.target = value
    }
}
