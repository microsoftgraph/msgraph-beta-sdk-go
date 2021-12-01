package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfficeClientConfigurationAssignment 
type OfficeClientConfigurationAssignment struct {
    Entity
    // The target assignment defined by the admin.
    target *OfficeConfigurationAssignmentTarget;
}
// NewOfficeClientConfigurationAssignment instantiates a new officeClientConfigurationAssignment and sets the default values.
func NewOfficeClientConfigurationAssignment()(*OfficeClientConfigurationAssignment) {
    m := &OfficeClientConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetTarget gets the target property value. The target assignment defined by the admin.
func (m *OfficeClientConfigurationAssignment) GetTarget()(*OfficeConfigurationAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeClientConfigurationAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeConfigurationAssignmentTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*OfficeConfigurationAssignmentTarget))
        }
        return nil
    }
    return res
}
func (m *OfficeClientConfigurationAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OfficeClientConfigurationAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTarget sets the target property value. The target assignment defined by the admin.
func (m *OfficeClientConfigurationAssignment) SetTarget(value *OfficeConfigurationAssignmentTarget)() {
    if m != nil {
        m.target = value
    }
}
