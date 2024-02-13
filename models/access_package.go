package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackage struct {
    Entity
}
// NewAccessPackage instantiates a new AccessPackage and sets the default values.
func NewAccessPackage()(*AccessPackage) {
    m := &AccessPackage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackage(), nil
}
// GetAccessPackageAssignmentPolicies gets the accessPackageAssignmentPolicies property value. Read-only. Nullable. Supports $expand.
// returns a []AccessPackageAssignmentPolicyable when successful
func (m *AccessPackage) GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicyable) {
    val, err := m.GetBackingStore().Get("accessPackageAssignmentPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageAssignmentPolicyable)
    }
    return nil
}
// GetAccessPackageCatalog gets the accessPackageCatalog property value. The accessPackageCatalog property
// returns a AccessPackageCatalogable when successful
func (m *AccessPackage) GetAccessPackageCatalog()(AccessPackageCatalogable) {
    val, err := m.GetBackingStore().Get("accessPackageCatalog")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageCatalogable)
    }
    return nil
}
// GetAccessPackageResourceRoleScopes gets the accessPackageResourceRoleScopes property value. The accessPackageResourceRoleScopes property
// returns a []AccessPackageResourceRoleScopeable when successful
func (m *AccessPackage) GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScopeable) {
    val, err := m.GetBackingStore().Get("accessPackageResourceRoleScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageResourceRoleScopeable)
    }
    return nil
}
// GetAccessPackagesIncompatibleWith gets the accessPackagesIncompatibleWith property value. The access packages that are incompatible with this package. Read-only.
// returns a []AccessPackageable when successful
func (m *AccessPackage) GetAccessPackagesIncompatibleWith()([]AccessPackageable) {
    val, err := m.GetBackingStore().Get("accessPackagesIncompatibleWith")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageable)
    }
    return nil
}
// GetCatalogId gets the catalogId property value. Identifier of the access package catalog referencing this access package. Read-only.
// returns a *string when successful
func (m *AccessPackage) GetCatalogId()(*string) {
    val, err := m.GetBackingStore().Get("catalogId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The userPrincipalName of the user or identity of the subject who created this resource. Read-only.
// returns a *string when successful
func (m *AccessPackage) GetCreatedBy()(*string) {
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
// returns a *Time when successful
func (m *AccessPackage) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description of the access package.
// returns a *string when successful
func (m *AccessPackage) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the access package. Supports $filter (eq, contains).
// returns a *string when successful
func (m *AccessPackage) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignmentPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageAssignmentPolicyable)
                }
            }
            m.SetAccessPackageAssignmentPolicies(res)
        }
        return nil
    }
    res["accessPackageCatalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageCatalog(val.(AccessPackageCatalogable))
        }
        return nil
    }
    res["accessPackageResourceRoleScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRoleScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRoleScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceRoleScopeable)
                }
            }
            m.SetAccessPackageResourceRoleScopes(res)
        }
        return nil
    }
    res["accessPackagesIncompatibleWith"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAccessPackagesIncompatibleWith(res)
        }
        return nil
    }
    res["catalogId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
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
    res["incompatibleAccessPackages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIncompatibleAccessPackages(res)
        }
        return nil
    }
    res["incompatibleGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Groupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Groupable)
                }
            }
            m.SetIncompatibleGroups(res)
        }
        return nil
    }
    res["isHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHidden(val)
        }
        return nil
    }
    res["isRoleScopesVisible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRoleScopesVisible(val)
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
// GetIncompatibleAccessPackages gets the incompatibleAccessPackages property value. The  access packages whose assigned users are ineligible to be assigned this access package.
// returns a []AccessPackageable when successful
func (m *AccessPackage) GetIncompatibleAccessPackages()([]AccessPackageable) {
    val, err := m.GetBackingStore().Get("incompatibleAccessPackages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageable)
    }
    return nil
}
// GetIncompatibleGroups gets the incompatibleGroups property value. The groups whose members are ineligible to be assigned this access package.
// returns a []Groupable when successful
func (m *AccessPackage) GetIncompatibleGroups()([]Groupable) {
    val, err := m.GetBackingStore().Get("incompatibleGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Groupable)
    }
    return nil
}
// GetIsHidden gets the isHidden property value. Whether the access package is hidden from the requestor.
// returns a *bool when successful
func (m *AccessPackage) GetIsHidden()(*bool) {
    val, err := m.GetBackingStore().Get("isHidden")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsRoleScopesVisible gets the isRoleScopesVisible property value. Indicates whether role scopes are visible.
// returns a *bool when successful
func (m *AccessPackage) GetIsRoleScopesVisible()(*bool) {
    val, err := m.GetBackingStore().Get("isRoleScopesVisible")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetModifiedBy gets the modifiedBy property value. The userPrincipalName of the user who last modified this resource. Read-only.
// returns a *string when successful
func (m *AccessPackage) GetModifiedBy()(*string) {
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
// returns a *Time when successful
func (m *AccessPackage) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *AccessPackage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageAssignmentPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignmentPolicies()))
        for i, v := range m.GetAccessPackageAssignmentPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageCatalog", m.GetAccessPackageCatalog())
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceRoleScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageResourceRoleScopes()))
        for i, v := range m.GetAccessPackageResourceRoleScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoleScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackagesIncompatibleWith() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackagesIncompatibleWith()))
        for i, v := range m.GetAccessPackagesIncompatibleWith() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackagesIncompatibleWith", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("catalogId", m.GetCatalogId())
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
    if m.GetIncompatibleAccessPackages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncompatibleAccessPackages()))
        for i, v := range m.GetIncompatibleAccessPackages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("incompatibleAccessPackages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncompatibleGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncompatibleGroups()))
        for i, v := range m.GetIncompatibleGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("incompatibleGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHidden", m.GetIsHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRoleScopesVisible", m.GetIsRoleScopesVisible())
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
// SetAccessPackageAssignmentPolicies sets the accessPackageAssignmentPolicies property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackage) SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicyable)() {
    err := m.GetBackingStore().Set("accessPackageAssignmentPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageCatalog sets the accessPackageCatalog property value. The accessPackageCatalog property
func (m *AccessPackage) SetAccessPackageCatalog(value AccessPackageCatalogable)() {
    err := m.GetBackingStore().Set("accessPackageCatalog", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageResourceRoleScopes sets the accessPackageResourceRoleScopes property value. The accessPackageResourceRoleScopes property
func (m *AccessPackage) SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScopeable)() {
    err := m.GetBackingStore().Set("accessPackageResourceRoleScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackagesIncompatibleWith sets the accessPackagesIncompatibleWith property value. The access packages that are incompatible with this package. Read-only.
func (m *AccessPackage) SetAccessPackagesIncompatibleWith(value []AccessPackageable)() {
    err := m.GetBackingStore().Set("accessPackagesIncompatibleWith", value)
    if err != nil {
        panic(err)
    }
}
// SetCatalogId sets the catalogId property value. Identifier of the access package catalog referencing this access package. Read-only.
func (m *AccessPackage) SetCatalogId(value *string)() {
    err := m.GetBackingStore().Set("catalogId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The userPrincipalName of the user or identity of the subject who created this resource. Read-only.
func (m *AccessPackage) SetCreatedBy(value *string)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the access package.
func (m *AccessPackage) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the access package. Supports $filter (eq, contains).
func (m *AccessPackage) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIncompatibleAccessPackages sets the incompatibleAccessPackages property value. The  access packages whose assigned users are ineligible to be assigned this access package.
func (m *AccessPackage) SetIncompatibleAccessPackages(value []AccessPackageable)() {
    err := m.GetBackingStore().Set("incompatibleAccessPackages", value)
    if err != nil {
        panic(err)
    }
}
// SetIncompatibleGroups sets the incompatibleGroups property value. The groups whose members are ineligible to be assigned this access package.
func (m *AccessPackage) SetIncompatibleGroups(value []Groupable)() {
    err := m.GetBackingStore().Set("incompatibleGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHidden sets the isHidden property value. Whether the access package is hidden from the requestor.
func (m *AccessPackage) SetIsHidden(value *bool)() {
    err := m.GetBackingStore().Set("isHidden", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRoleScopesVisible sets the isRoleScopesVisible property value. Indicates whether role scopes are visible.
func (m *AccessPackage) SetIsRoleScopesVisible(value *bool)() {
    err := m.GetBackingStore().Set("isRoleScopesVisible", value)
    if err != nil {
        panic(err)
    }
}
// SetModifiedBy sets the modifiedBy property value. The userPrincipalName of the user who last modified this resource. Read-only.
func (m *AccessPackage) SetModifiedBy(value *string)() {
    err := m.GetBackingStore().Set("modifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("modifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
type AccessPackageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicyable)
    GetAccessPackageCatalog()(AccessPackageCatalogable)
    GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScopeable)
    GetAccessPackagesIncompatibleWith()([]AccessPackageable)
    GetCatalogId()(*string)
    GetCreatedBy()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIncompatibleAccessPackages()([]AccessPackageable)
    GetIncompatibleGroups()([]Groupable)
    GetIsHidden()(*bool)
    GetIsRoleScopesVisible()(*bool)
    GetModifiedBy()(*string)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicyable)()
    SetAccessPackageCatalog(value AccessPackageCatalogable)()
    SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScopeable)()
    SetAccessPackagesIncompatibleWith(value []AccessPackageable)()
    SetCatalogId(value *string)()
    SetCreatedBy(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIncompatibleAccessPackages(value []AccessPackageable)()
    SetIncompatibleGroups(value []Groupable)()
    SetIsHidden(value *bool)()
    SetIsRoleScopesVisible(value *bool)()
    SetModifiedBy(value *string)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
