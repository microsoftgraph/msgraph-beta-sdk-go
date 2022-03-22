package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcProvisioningPolicyAssignment 
type CloudPcProvisioningPolicyAssignment struct {
    Entity
    // The assignment target for the provisioning policy. Currently, the only target supported for this policy is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
    target CloudPcManagementAssignmentTargetable;
}
// NewCloudPcProvisioningPolicyAssignment instantiates a new cloudPcProvisioningPolicyAssignment and sets the default values.
func NewCloudPcProvisioningPolicyAssignment()(*CloudPcProvisioningPolicyAssignment) {
    m := &CloudPcProvisioningPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcProvisioningPolicyAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcProvisioningPolicyAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(CloudPcManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The assignment target for the provisioning policy. Currently, the only target supported for this policy is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
func (m *CloudPcProvisioningPolicyAssignment) GetTarget()(CloudPcManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Serialize serializes information the current object
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
// SetTarget sets the target property value. The assignment target for the provisioning policy. Currently, the only target supported for this policy is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
func (m *CloudPcProvisioningPolicyAssignment) SetTarget(value CloudPcManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
