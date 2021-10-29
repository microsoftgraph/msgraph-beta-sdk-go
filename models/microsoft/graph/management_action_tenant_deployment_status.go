package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagementActionTenantDeploymentStatus struct {
    Entity
    // The collection of deployment status for each instance of a management action. Optional.
    statuses []ManagementActionDeploymentStatus;
    // The identifier for the tenant group that is associated with the management action. Required. Read-only.
    tenantGroupId *string;
    // The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
    tenantId *string;
}
// Instantiates a new managementActionTenantDeploymentStatus and sets the default values.
func NewManagementActionTenantDeploymentStatus()(*ManagementActionTenantDeploymentStatus) {
    m := &ManagementActionTenantDeploymentStatus{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the statuses property value. The collection of deployment status for each instance of a management action. Optional.
func (m *ManagementActionTenantDeploymentStatus) GetStatuses()([]ManagementActionDeploymentStatus) {
    if m == nil {
        return nil
    } else {
        return m.statuses
    }
}
// Gets the tenantGroupId property value. The identifier for the tenant group that is associated with the management action. Required. Read-only.
func (m *ManagementActionTenantDeploymentStatus) GetTenantGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantGroupId
    }
}
// Gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *ManagementActionTenantDeploymentStatus) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// The deserialization information for the current model
func (m *ManagementActionTenantDeploymentStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["statuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementActionDeploymentStatus() })
        if err != nil {
            return err
        }
        res := make([]ManagementActionDeploymentStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementActionDeploymentStatus))
        }
        m.SetStatuses(res)
        return nil
    }
    res["tenantGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantGroupId(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    return res
}
func (m *ManagementActionTenantDeploymentStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementActionTenantDeploymentStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStatuses()))
        for i, v := range m.GetStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the statuses property value. The collection of deployment status for each instance of a management action. Optional.
// Parameters:
//  - value : Value to set for the statuses property.
func (m *ManagementActionTenantDeploymentStatus) SetStatuses(value []ManagementActionDeploymentStatus)() {
    m.statuses = value
}
// Sets the tenantGroupId property value. The identifier for the tenant group that is associated with the management action. Required. Read-only.
// Parameters:
//  - value : Value to set for the tenantGroupId property.
func (m *ManagementActionTenantDeploymentStatus) SetTenantGroupId(value *string)() {
    m.tenantGroupId = value
}
// Sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *ManagementActionTenantDeploymentStatus) SetTenantId(value *string)() {
    m.tenantId = value
}
