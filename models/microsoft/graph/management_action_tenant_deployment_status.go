package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementActionTenantDeploymentStatus provides operations to manage the tenantRelationship singleton.
type ManagementActionTenantDeploymentStatus struct {
    Entity
    // The collection of deployment status for each instance of a management action. Optional.
    statuses []ManagementActionDeploymentStatusable;
    // The identifier for the tenant group that is associated with the management action. Required. Read-only.
    tenantGroupId *string;
    // The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
    tenantId *string;
}
// NewManagementActionTenantDeploymentStatus instantiates a new managementActionTenantDeploymentStatus and sets the default values.
func NewManagementActionTenantDeploymentStatus()(*ManagementActionTenantDeploymentStatus) {
    m := &ManagementActionTenantDeploymentStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagementActionTenantDeploymentStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementActionTenantDeploymentStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["statuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementActionDeploymentStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionDeploymentStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementActionDeploymentStatusable)
            }
            m.SetStatuses(res)
        }
        return nil
    }
    res["tenantGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantGroupId(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetStatuses gets the statuses property value. The collection of deployment status for each instance of a management action. Optional.
func (m *ManagementActionTenantDeploymentStatus) GetStatuses()([]ManagementActionDeploymentStatusable) {
    if m == nil {
        return nil
    } else {
        return m.statuses
    }
}
// GetTenantGroupId gets the tenantGroupId property value. The identifier for the tenant group that is associated with the management action. Required. Read-only.
func (m *ManagementActionTenantDeploymentStatus) GetTenantGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantGroupId
    }
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *ManagementActionTenantDeploymentStatus) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *ManagementActionTenantDeploymentStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementActionTenantDeploymentStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStatuses()))
        for i, v := range m.GetStatuses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("statuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantGroupId", m.GetTenantGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStatuses sets the statuses property value. The collection of deployment status for each instance of a management action. Optional.
func (m *ManagementActionTenantDeploymentStatus) SetStatuses(value []ManagementActionDeploymentStatusable)() {
    if m != nil {
        m.statuses = value
    }
}
// SetTenantGroupId sets the tenantGroupId property value. The identifier for the tenant group that is associated with the management action. Required. Read-only.
func (m *ManagementActionTenantDeploymentStatus) SetTenantGroupId(value *string)() {
    if m != nil {
        m.tenantGroupId = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *ManagementActionTenantDeploymentStatus) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
