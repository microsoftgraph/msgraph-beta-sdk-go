package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantGroup 
type TenantGroup struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewTenantGroup instantiates a new TenantGroup and sets the default values.
func NewTenantGroup()(*TenantGroup) {
    m := &TenantGroup{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateTenantGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantGroup(), nil
}
// GetAllTenantsIncluded gets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
func (m *TenantGroup) GetAllTenantsIncluded()(*bool) {
    val, err := m.GetBackingStore().Get("allTenantsIncluded")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allTenantsIncluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllTenantsIncluded(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["managementActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementActionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementActionInfoable)
            }
            m.SetManagementActions(res)
        }
        return nil
    }
    res["managementIntents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementIntentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementIntentInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementIntentInfoable)
            }
            m.SetManagementIntents(res)
        }
        return nil
    }
    res["tenantIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTenantIds(res)
        }
        return nil
    }
    return res
}
// GetManagementActions gets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementActions()([]ManagementActionInfoable) {
    val, err := m.GetBackingStore().Get("managementActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementActionInfoable)
    }
    return nil
}
// GetManagementIntents gets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementIntents()([]ManagementIntentInfoable) {
    val, err := m.GetBackingStore().Get("managementIntents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementIntentInfoable)
    }
    return nil
}
// GetTenantIds gets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) GetTenantIds()([]string) {
    val, err := m.GetBackingStore().Get("tenantIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementActions()))
        for i, v := range m.GetManagementActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementIntents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementIntents()))
        for i, v := range m.GetManagementIntents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("allTenantsIncluded", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementActions sets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) SetManagementActions(value []ManagementActionInfoable)() {
    err := m.GetBackingStore().Set("managementActions", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementIntents sets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) SetManagementIntents(value []ManagementIntentInfoable)() {
    err := m.GetBackingStore().Set("managementIntents", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantIds sets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) SetTenantIds(value []string)() {
    err := m.GetBackingStore().Set("tenantIds", value)
    if err != nil {
        panic(err)
    }
}
// TenantGroupable 
type TenantGroupable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllTenantsIncluded()(*bool)
    GetDisplayName()(*string)
    GetManagementActions()([]ManagementActionInfoable)
    GetManagementIntents()([]ManagementIntentInfoable)
    GetTenantIds()([]string)
    SetAllTenantsIncluded(value *bool)()
    SetDisplayName(value *string)()
    SetManagementActions(value []ManagementActionInfoable)()
    SetManagementIntents(value []ManagementIntentInfoable)()
    SetTenantIds(value []string)()
}
