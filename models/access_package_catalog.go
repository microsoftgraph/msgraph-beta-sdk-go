package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageCatalog 
type AccessPackageCatalog struct {
    Entity
}
// NewAccessPackageCatalog instantiates a new accessPackageCatalog and sets the default values.
func NewAccessPackageCatalog()(*AccessPackageCatalog) {
    m := &AccessPackageCatalog{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageCatalogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageCatalog(), nil
}
// GetAccessPackageCustomWorkflowExtensions gets the accessPackageCustomWorkflowExtensions property value. The attributes of a logic app, which can be called at various stages of an access package request and assignment cycle.
func (m *AccessPackageCatalog) GetAccessPackageCustomWorkflowExtensions()([]CustomCalloutExtensionable) {
    val, err := m.GetBackingStore().Get("accessPackageCustomWorkflowExtensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomCalloutExtensionable)
    }
    return nil
}
// GetAccessPackageResourceRoles gets the accessPackageResourceRoles property value. The roles in each resource in a catalog. Read-only.
func (m *AccessPackageCatalog) GetAccessPackageResourceRoles()([]AccessPackageResourceRoleable) {
    val, err := m.GetBackingStore().Get("accessPackageResourceRoles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageResourceRoleable)
    }
    return nil
}
// GetAccessPackageResources gets the accessPackageResources property value. The accessPackageResources property
func (m *AccessPackageCatalog) GetAccessPackageResources()([]AccessPackageResourceable) {
    val, err := m.GetBackingStore().Get("accessPackageResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageResourceable)
    }
    return nil
}
// GetAccessPackageResourceScopes gets the accessPackageResourceScopes property value. The accessPackageResourceScopes property
func (m *AccessPackageCatalog) GetAccessPackageResourceScopes()([]AccessPackageResourceScopeable) {
    val, err := m.GetBackingStore().Get("accessPackageResourceScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageResourceScopeable)
    }
    return nil
}
// GetAccessPackages gets the accessPackages property value. The access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *AccessPackageCatalog) GetAccessPackages()([]AccessPackageable) {
    val, err := m.GetBackingStore().Get("accessPackages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageable)
    }
    return nil
}
// GetCatalogStatus gets the catalogStatus property value. Has the value Published if the access packages are available for management.
func (m *AccessPackageCatalog) GetCatalogStatus()(*string) {
    val, err := m.GetBackingStore().Get("catalogStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCatalogType gets the catalogType property value. One of UserManaged or ServiceDefault.
func (m *AccessPackageCatalog) GetCatalogType()(*string) {
    val, err := m.GetBackingStore().Get("catalogType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. UPN of the user who created this resource. Read-only.
func (m *AccessPackageCatalog) GetCreatedBy()(*string) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCustomAccessPackageWorkflowExtensions gets the customAccessPackageWorkflowExtensions property value. The customAccessPackageWorkflowExtensions property
func (m *AccessPackageCatalog) GetCustomAccessPackageWorkflowExtensions()([]CustomAccessPackageWorkflowExtensionable) {
    val, err := m.GetBackingStore().Get("customAccessPackageWorkflowExtensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomAccessPackageWorkflowExtensionable)
    }
    return nil
}
// GetDescription gets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the access package catalog. Supports $filter (eq, contains).
func (m *AccessPackageCatalog) GetDisplayName()(*string) {
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
func (m *AccessPackageCatalog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageCustomWorkflowExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomCalloutExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomCalloutExtensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomCalloutExtensionable)
                }
            }
            m.SetAccessPackageCustomWorkflowExtensions(res)
        }
        return nil
    }
    res["accessPackageResourceRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceRoleable)
                }
            }
            m.SetAccessPackageResourceRoles(res)
        }
        return nil
    }
    res["accessPackageResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceable)
                }
            }
            m.SetAccessPackageResources(res)
        }
        return nil
    }
    res["accessPackageResourceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceScopeable)
                }
            }
            m.SetAccessPackageResourceScopes(res)
        }
        return nil
    }
    res["accessPackages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageable)
                }
            }
            m.SetAccessPackages(res)
        }
        return nil
    }
    res["catalogStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogStatus(val)
        }
        return nil
    }
    res["catalogType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogType(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["customAccessPackageWorkflowExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomAccessPackageWorkflowExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomAccessPackageWorkflowExtensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomAccessPackageWorkflowExtensionable)
                }
            }
            m.SetCustomAccessPackageWorkflowExtensions(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["isExternallyVisible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternallyVisible(val)
        }
        return nil
    }
    res["modifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedBy(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetIsExternallyVisible gets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) GetIsExternallyVisible()(*bool) {
    val, err := m.GetBackingStore().Get("isExternallyVisible")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetModifiedBy gets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
func (m *AccessPackageCatalog) GetModifiedBy()(*string) {
    val, err := m.GetBackingStore().Get("modifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("modifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageCatalog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageCustomWorkflowExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageCustomWorkflowExtensions()))
        for i, v := range m.GetAccessPackageCustomWorkflowExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageCustomWorkflowExtensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageResourceRoles()))
        for i, v := range m.GetAccessPackageResourceRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageResources()))
        for i, v := range m.GetAccessPackageResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageResourceScopes()))
        for i, v := range m.GetAccessPackageResourceScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("catalogStatus", m.GetCatalogStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("catalogType", m.GetCatalogType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCustomAccessPackageWorkflowExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomAccessPackageWorkflowExtensions()))
        for i, v := range m.GetCustomAccessPackageWorkflowExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customAccessPackageWorkflowExtensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    {
        err = writer.WriteBoolValue("isExternallyVisible", m.GetIsExternallyVisible())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modifiedBy", m.GetModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageCustomWorkflowExtensions sets the accessPackageCustomWorkflowExtensions property value. The attributes of a logic app, which can be called at various stages of an access package request and assignment cycle.
func (m *AccessPackageCatalog) SetAccessPackageCustomWorkflowExtensions(value []CustomCalloutExtensionable)() {
    err := m.GetBackingStore().Set("accessPackageCustomWorkflowExtensions", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageResourceRoles sets the accessPackageResourceRoles property value. The roles in each resource in a catalog. Read-only.
func (m *AccessPackageCatalog) SetAccessPackageResourceRoles(value []AccessPackageResourceRoleable)() {
    err := m.GetBackingStore().Set("accessPackageResourceRoles", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageResources sets the accessPackageResources property value. The accessPackageResources property
func (m *AccessPackageCatalog) SetAccessPackageResources(value []AccessPackageResourceable)() {
    err := m.GetBackingStore().Set("accessPackageResources", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageResourceScopes sets the accessPackageResourceScopes property value. The accessPackageResourceScopes property
func (m *AccessPackageCatalog) SetAccessPackageResourceScopes(value []AccessPackageResourceScopeable)() {
    err := m.GetBackingStore().Set("accessPackageResourceScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackages sets the accessPackages property value. The access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *AccessPackageCatalog) SetAccessPackages(value []AccessPackageable)() {
    err := m.GetBackingStore().Set("accessPackages", value)
    if err != nil {
        panic(err)
    }
}
// SetCatalogStatus sets the catalogStatus property value. Has the value Published if the access packages are available for management.
func (m *AccessPackageCatalog) SetCatalogStatus(value *string)() {
    err := m.GetBackingStore().Set("catalogStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCatalogType sets the catalogType property value. One of UserManaged or ServiceDefault.
func (m *AccessPackageCatalog) SetCatalogType(value *string)() {
    err := m.GetBackingStore().Set("catalogType", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. UPN of the user who created this resource. Read-only.
func (m *AccessPackageCatalog) SetCreatedBy(value *string)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomAccessPackageWorkflowExtensions sets the customAccessPackageWorkflowExtensions property value. The customAccessPackageWorkflowExtensions property
func (m *AccessPackageCatalog) SetCustomAccessPackageWorkflowExtensions(value []CustomAccessPackageWorkflowExtensionable)() {
    err := m.GetBackingStore().Set("customAccessPackageWorkflowExtensions", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the access package catalog. Supports $filter (eq, contains).
func (m *AccessPackageCatalog) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExternallyVisible sets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) SetIsExternallyVisible(value *bool)() {
    err := m.GetBackingStore().Set("isExternallyVisible", value)
    if err != nil {
        panic(err)
    }
}
// SetModifiedBy sets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
func (m *AccessPackageCatalog) SetModifiedBy(value *string)() {
    err := m.GetBackingStore().Set("modifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("modifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageCatalogable 
type AccessPackageCatalogable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageCustomWorkflowExtensions()([]CustomCalloutExtensionable)
    GetAccessPackageResourceRoles()([]AccessPackageResourceRoleable)
    GetAccessPackageResources()([]AccessPackageResourceable)
    GetAccessPackageResourceScopes()([]AccessPackageResourceScopeable)
    GetAccessPackages()([]AccessPackageable)
    GetCatalogStatus()(*string)
    GetCatalogType()(*string)
    GetCreatedBy()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomAccessPackageWorkflowExtensions()([]CustomAccessPackageWorkflowExtensionable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsExternallyVisible()(*bool)
    GetModifiedBy()(*string)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessPackageCustomWorkflowExtensions(value []CustomCalloutExtensionable)()
    SetAccessPackageResourceRoles(value []AccessPackageResourceRoleable)()
    SetAccessPackageResources(value []AccessPackageResourceable)()
    SetAccessPackageResourceScopes(value []AccessPackageResourceScopeable)()
    SetAccessPackages(value []AccessPackageable)()
    SetCatalogStatus(value *string)()
    SetCatalogType(value *string)()
    SetCreatedBy(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomAccessPackageWorkflowExtensions(value []CustomAccessPackageWorkflowExtensionable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsExternallyVisible(value *bool)()
    SetModifiedBy(value *string)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
