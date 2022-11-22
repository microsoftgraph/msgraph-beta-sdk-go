package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageCatalog 
type AccessPackageCatalog struct {
    Entity
    // The roles in each resource in a catalog. Read-only.
    accessPackageResourceRoles []AccessPackageResourceRoleable
    // The accessPackageResources property
    accessPackageResources []AccessPackageResourceable
    // The accessPackageResourceScopes property
    accessPackageResourceScopes []AccessPackageResourceScopeable
    // The access packages in this catalog. Read-only. Nullable. Supports $expand.
    accessPackages []AccessPackageable
    // Has the value Published if the access packages are available for management.
    catalogStatus *string
    // One of UserManaged or ServiceDefault.
    catalogType *string
    // UPN of the user who created this resource. Read-only.
    createdBy *string
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customAccessPackageWorkflowExtensions property
    customAccessPackageWorkflowExtensions []CustomAccessPackageWorkflowExtensionable
    // The description of the access package catalog.
    description *string
    // The display name of the access package catalog. Supports $filter (eq, contains).
    displayName *string
    // Whether the access packages in this catalog can be requested by users outside of the tenant.
    isExternallyVisible *bool
    // The UPN of the user who last modified this resource. Read-only.
    modifiedBy *string
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
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
// GetAccessPackageResourceRoles gets the accessPackageResourceRoles property value. The roles in each resource in a catalog. Read-only.
func (m *AccessPackageCatalog) GetAccessPackageResourceRoles()([]AccessPackageResourceRoleable) {
    return m.accessPackageResourceRoles
}
// GetAccessPackageResources gets the accessPackageResources property value. The accessPackageResources property
func (m *AccessPackageCatalog) GetAccessPackageResources()([]AccessPackageResourceable) {
    return m.accessPackageResources
}
// GetAccessPackageResourceScopes gets the accessPackageResourceScopes property value. The accessPackageResourceScopes property
func (m *AccessPackageCatalog) GetAccessPackageResourceScopes()([]AccessPackageResourceScopeable) {
    return m.accessPackageResourceScopes
}
// GetAccessPackages gets the accessPackages property value. The access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *AccessPackageCatalog) GetAccessPackages()([]AccessPackageable) {
    return m.accessPackages
}
// GetCatalogStatus gets the catalogStatus property value. Has the value Published if the access packages are available for management.
func (m *AccessPackageCatalog) GetCatalogStatus()(*string) {
    return m.catalogStatus
}
// GetCatalogType gets the catalogType property value. One of UserManaged or ServiceDefault.
func (m *AccessPackageCatalog) GetCatalogType()(*string) {
    return m.catalogType
}
// GetCreatedBy gets the createdBy property value. UPN of the user who created this resource. Read-only.
func (m *AccessPackageCatalog) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomAccessPackageWorkflowExtensions gets the customAccessPackageWorkflowExtensions property value. The customAccessPackageWorkflowExtensions property
func (m *AccessPackageCatalog) GetCustomAccessPackageWorkflowExtensions()([]CustomAccessPackageWorkflowExtensionable) {
    return m.customAccessPackageWorkflowExtensions
}
// GetDescription gets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the access package catalog. Supports $filter (eq, contains).
func (m *AccessPackageCatalog) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageCatalog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResourceRoles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageResourceRoleFromDiscriminatorValue , m.SetAccessPackageResourceRoles)
    res["accessPackageResources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageResourceFromDiscriminatorValue , m.SetAccessPackageResources)
    res["accessPackageResourceScopes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageResourceScopeFromDiscriminatorValue , m.SetAccessPackageResourceScopes)
    res["accessPackages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageFromDiscriminatorValue , m.SetAccessPackages)
    res["catalogStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCatalogStatus)
    res["catalogType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCatalogType)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["customAccessPackageWorkflowExtensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomAccessPackageWorkflowExtensionFromDiscriminatorValue , m.SetCustomAccessPackageWorkflowExtensions)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isExternallyVisible"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsExternallyVisible)
    res["modifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModifiedBy)
    res["modifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetModifiedDateTime)
    return res
}
// GetIsExternallyVisible gets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) GetIsExternallyVisible()(*bool) {
    return m.isExternallyVisible
}
// GetModifiedBy gets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
func (m *AccessPackageCatalog) GetModifiedBy()(*string) {
    return m.modifiedBy
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// Serialize serializes information the current object
func (m *AccessPackageCatalog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageResourceRoles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackageResourceRoles())
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackageResources())
        err = writer.WriteCollectionOfObjectValues("accessPackageResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceScopes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackageResourceScopes())
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackages())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomAccessPackageWorkflowExtensions())
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
// SetAccessPackageResourceRoles sets the accessPackageResourceRoles property value. The roles in each resource in a catalog. Read-only.
func (m *AccessPackageCatalog) SetAccessPackageResourceRoles(value []AccessPackageResourceRoleable)() {
    m.accessPackageResourceRoles = value
}
// SetAccessPackageResources sets the accessPackageResources property value. The accessPackageResources property
func (m *AccessPackageCatalog) SetAccessPackageResources(value []AccessPackageResourceable)() {
    m.accessPackageResources = value
}
// SetAccessPackageResourceScopes sets the accessPackageResourceScopes property value. The accessPackageResourceScopes property
func (m *AccessPackageCatalog) SetAccessPackageResourceScopes(value []AccessPackageResourceScopeable)() {
    m.accessPackageResourceScopes = value
}
// SetAccessPackages sets the accessPackages property value. The access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *AccessPackageCatalog) SetAccessPackages(value []AccessPackageable)() {
    m.accessPackages = value
}
// SetCatalogStatus sets the catalogStatus property value. Has the value Published if the access packages are available for management.
func (m *AccessPackageCatalog) SetCatalogStatus(value *string)() {
    m.catalogStatus = value
}
// SetCatalogType sets the catalogType property value. One of UserManaged or ServiceDefault.
func (m *AccessPackageCatalog) SetCatalogType(value *string)() {
    m.catalogType = value
}
// SetCreatedBy sets the createdBy property value. UPN of the user who created this resource. Read-only.
func (m *AccessPackageCatalog) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomAccessPackageWorkflowExtensions sets the customAccessPackageWorkflowExtensions property value. The customAccessPackageWorkflowExtensions property
func (m *AccessPackageCatalog) SetCustomAccessPackageWorkflowExtensions(value []CustomAccessPackageWorkflowExtensionable)() {
    m.customAccessPackageWorkflowExtensions = value
}
// SetDescription sets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the access package catalog. Supports $filter (eq, contains).
func (m *AccessPackageCatalog) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsExternallyVisible sets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) SetIsExternallyVisible(value *bool)() {
    m.isExternallyVisible = value
}
// SetModifiedBy sets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
func (m *AccessPackageCatalog) SetModifiedBy(value *string)() {
    m.modifiedBy = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
