package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OfficeClientConfigurationAssignment struct {
    Entity
    target *OfficeConfigurationAssignmentTarget;
}
func NewOfficeClientConfigurationAssignment()(*OfficeClientConfigurationAssignment) {
    m := &OfficeClientConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OfficeClientConfigurationAssignment) GetTarget()(*OfficeConfigurationAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *OfficeClientConfigurationAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeConfigurationAssignmentTarget() })
        if err != nil {
            return err
        }
        m.SetTarget(val.(*OfficeConfigurationAssignmentTarget))
        return nil
    }
    return res
}
func (m *OfficeClientConfigurationAssignment) IsNil()(bool) {
    return m == nil
}
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
func (m *OfficeClientConfigurationAssignment) SetTarget(value *OfficeConfigurationAssignmentTarget)() {
    m.target = value
}
