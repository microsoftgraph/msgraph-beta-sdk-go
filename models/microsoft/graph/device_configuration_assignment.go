package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceConfigurationAssignment 
type DeviceConfigurationAssignment struct {
    Entity
    // The admin intent to apply or remove the profile. Possible values are: apply, remove.
    intent *DeviceConfigAssignmentIntent;
    // The assignment source for the device configuration, direct or parcel/policySet. This property is read-only. Possible values are: direct, policySets.
    source *DeviceAndAppManagementAssignmentSource;
    // The identifier of the source of the assignment. This property is read-only.
    sourceId *string;
    // The assignment target for the device configuration.
    target *DeviceAndAppManagementAssignmentTarget;
}
// NewDeviceConfigurationAssignment instantiates a new deviceConfigurationAssignment and sets the default values.
func NewDeviceConfigurationAssignment()(*DeviceConfigurationAssignment) {
    m := &DeviceConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetIntent gets the intent property value. The admin intent to apply or remove the profile. Possible values are: apply, remove.
func (m *DeviceConfigurationAssignment) GetIntent()(*DeviceConfigAssignmentIntent) {
    if m == nil {
        return nil
    } else {
        return m.intent
    }
}
// GetSource gets the source property value. The assignment source for the device configuration, direct or parcel/policySet. This property is read-only. Possible values are: direct, policySets.
func (m *DeviceConfigurationAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment. This property is read-only.
func (m *DeviceConfigurationAssignment) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTarget gets the target property value. The assignment target for the device configuration.
func (m *DeviceConfigurationAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceConfigAssignmentIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(*DeviceConfigAssignmentIntent))
        }
        return nil
    }
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
func (m *DeviceConfigurationAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceConfigurationAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetIntent sets the intent property value. The admin intent to apply or remove the profile. Possible values are: apply, remove.
func (m *DeviceConfigurationAssignment) SetIntent(value *DeviceConfigAssignmentIntent)() {
    if m != nil {
        m.intent = value
    }
}
// SetSource sets the source property value. The assignment source for the device configuration, direct or parcel/policySet. This property is read-only. Possible values are: direct, policySets.
func (m *DeviceConfigurationAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    if m != nil {
        m.source = value
    }
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment. This property is read-only.
func (m *DeviceConfigurationAssignment) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetTarget sets the target property value. The assignment target for the device configuration.
func (m *DeviceConfigurationAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    if m != nil {
        m.target = value
    }
}
