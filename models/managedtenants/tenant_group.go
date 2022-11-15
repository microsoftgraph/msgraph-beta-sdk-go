package managedtenants

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantGroup provides operations to manage the collection of accessReviewDecision entities.
type TenantGroup struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
    allTenantsIncluded *bool
    // The display name for the tenant group. Optional. Read-only.
    displayName *string
    // The collection of management action associated with the tenant group. Optional. Read-only.
    managementActions []ManagementActionInfoable
    // The collection of management intents associated with the tenant group. Optional. Read-only.
    managementIntents []ManagementIntentInfoable
    // The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
    tenantIds []string
}
// NewTenantGroup instantiates a new tenantGroup and sets the default values.
func NewTenantGroup()(*TenantGroup) {
    m := &TenantGroup{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.tenantGroup";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTenantGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantGroup(), nil
}
// GetAllTenantsIncluded gets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
func (m *TenantGroup) GetAllTenantsIncluded()(*bool) {
    return m.allTenantsIncluded
}
// GetDisplayName gets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allTenantsIncluded"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllTenantsIncluded)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["managementActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementActionInfoFromDiscriminatorValue , m.SetManagementActions)
    res["managementIntents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementIntentInfoFromDiscriminatorValue , m.SetManagementIntents)
    res["tenantIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTenantIds)
    return res
}
// GetManagementActions gets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementActions()([]ManagementActionInfoable) {
    return m.managementActions
}
// GetManagementIntents gets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementIntents()([]ManagementIntentInfoable) {
    return m.managementIntents
}
// GetTenantIds gets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) GetTenantIds()([]string) {
    return m.tenantIds
}
// Serialize serializes information the current object
func (m *TenantGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allTenantsIncluded", m.GetAllTenantsIncluded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetManagementActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementActions())
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementIntents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementIntents())
        err = writer.WriteCollectionOfObjectValues("managementIntents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantIds() != nil {
        err = writer.WriteCollectionOfStringValues("tenantIds", m.GetTenantIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllTenantsIncluded sets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
func (m *TenantGroup) SetAllTenantsIncluded(value *bool)() {
    m.allTenantsIncluded = value
}
// SetDisplayName sets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetManagementActions sets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) SetManagementActions(value []ManagementActionInfoable)() {
    m.managementActions = value
}
// SetManagementIntents sets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) SetManagementIntents(value []ManagementIntentInfoable)() {
    m.managementIntents = value
}
// SetTenantIds sets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) SetTenantIds(value []string)() {
    m.tenantIds = value
}
