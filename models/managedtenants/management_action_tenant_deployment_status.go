package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementActionTenantDeploymentStatus provides operations to manage the collection of activityStatistics entities.
type ManagementActionTenantDeploymentStatus struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The collection of deployment status for each instance of a management action. Optional.
    statuses []ManagementActionDeploymentStatusable
    // The identifier for the tenant group that is associated with the management action. Required. Read-only.
    tenantGroupId *string
    // The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
    tenantId *string
}
// NewManagementActionTenantDeploymentStatus instantiates a new managementActionTenantDeploymentStatus and sets the default values.
func NewManagementActionTenantDeploymentStatus()(*ManagementActionTenantDeploymentStatus) {
    m := &ManagementActionTenantDeploymentStatus{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.managementActionTenantDeploymentStatus";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementActionTenantDeploymentStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementActionTenantDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["statuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["tenantGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantGroupId(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *ManagementActionTenantDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStatuses()))
        for i, v := range m.GetStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
