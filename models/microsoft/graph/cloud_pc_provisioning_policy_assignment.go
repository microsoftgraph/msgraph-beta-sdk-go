package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CloudPcProvisioningPolicyAssignment struct {
    Entity
    target *CloudPcManagementAssignmentTarget;
}
func NewCloudPcProvisioningPolicyAssignment()(*CloudPcProvisioningPolicyAssignment) {
    m := &CloudPcProvisioningPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *CloudPcProvisioningPolicyAssignment) GetTarget()(*CloudPcManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *CloudPcProvisioningPolicyAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        m.SetTarget(val.(*CloudPcManagementAssignmentTarget))
        return nil
    }
    return res
}
func (m *CloudPcProvisioningPolicyAssignment) IsNil()(bool) {
    return m == nil
}
func (m *CloudPcProvisioningPolicyAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *CloudPcProvisioningPolicyAssignment) SetTarget(value *CloudPcManagementAssignmentTarget)() {
    m.target = value
}
