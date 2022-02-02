package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackage 
type AccessPackage struct {
    Entity
    // Read-only. Nullable. Supports $expand.
    accessPackageAssignmentPolicies []AccessPackageAssignmentPolicy;
    // Read-only. Nullable.
    accessPackageCatalog *AccessPackageCatalog;
    // Nullable.
    accessPackageResourceRoleScopes []AccessPackageResourceRoleScope;
    // The access packages that are incompatible with this package. Read-only.
    accessPackagesIncompatibleWith []AccessPackage;
    // Identifier of the access package catalog referencing this access package. Read-only.
    catalogId *string;
    // The userPrincipalName of the user or identity of the subject who created this resource. Read-only.
    createdBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the access package.
    description *string;
    // The display name of the access package. Supports $filter (eq, contains).
    displayName *string;
    // The  access packages whose assigned users are ineligible to be assigned this access package.
    incompatibleAccessPackages []AccessPackage;
    // The groups whose members are ineligible to be assigned this access package.
    incompatibleGroups []Group;
    // Whether the access package is hidden from the requestor.
    isHidden *bool;
    // Indicates whether role scopes are visible.
    isRoleScopesVisible *bool;
    // The userPrincipalName of the user who last modified this resource. Read-only.
    modifiedBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewAccessPackage instantiates a new accessPackage and sets the default values.
func NewAccessPackage()(*AccessPackage) {
    m := &AccessPackage{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackageAssignmentPolicies gets the accessPackageAssignmentPolicies property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackage) GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicy) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentPolicies
    }
}
// GetAccessPackageCatalog gets the accessPackageCatalog property value. Read-only. Nullable.
func (m *AccessPackage) GetAccessPackageCatalog()(*AccessPackageCatalog) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageCatalog
    }
}
// GetAccessPackageResourceRoleScopes gets the accessPackageResourceRoleScopes property value. Nullable.
func (m *AccessPackage) GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScope) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRoleScopes
    }
}
// GetAccessPackagesIncompatibleWith gets the accessPackagesIncompatibleWith property value. The access packages that are incompatible with this package. Read-only.
func (m *AccessPackage) GetAccessPackagesIncompatibleWith()([]AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackagesIncompatibleWith
    }
}
// GetCatalogId gets the catalogId property value. Identifier of the access package catalog referencing this access package. Read-only.
func (m *AccessPackage) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
// GetCreatedBy gets the createdBy property value. The userPrincipalName of the user or identity of the subject who created this resource. Read-only.
func (m *AccessPackage) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description of the access package.
func (m *AccessPackage) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the access package. Supports $filter (eq, contains).
func (m *AccessPackage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIncompatibleAccessPackages gets the incompatibleAccessPackages property value. The  access packages whose assigned users are ineligible to be assigned this access package.
func (m *AccessPackage) GetIncompatibleAccessPackages()([]AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.incompatibleAccessPackages
    }
}
// GetIncompatibleGroups gets the incompatibleGroups property value. The groups whose members are ineligible to be assigned this access package.
func (m *AccessPackage) GetIncompatibleGroups()([]Group) {
    if m == nil {
        return nil
    } else {
        return m.incompatibleGroups
    }
}
// GetIsHidden gets the isHidden property value. Whether the access package is hidden from the requestor.
func (m *AccessPackage) GetIsHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHidden
    }
}
// GetIsRoleScopesVisible gets the isRoleScopesVisible property value. Indicates whether role scopes are visible.
func (m *AccessPackage) GetIsRoleScopesVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRoleScopesVisible
    }
}
// GetModifiedBy gets the modifiedBy property value. The userPrincipalName of the user who last modified this resource. Read-only.
func (m *AccessPackage) GetModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modifiedBy
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignmentPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageAssignmentPolicy))
            }
            m.SetAccessPackageAssignmentPolicies(res)
        }
        return nil
    }
    res["accessPackageCatalog"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageCatalog() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageCatalog(val.(*AccessPackageCatalog))
        }
        return nil
    }
    res["accessPackageResourceRoleScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceRoleScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRoleScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageResourceRoleScope))
            }
            m.SetAccessPackageResourceRoleScopes(res)
        }
        return nil
    }
    res["accessPackagesIncompatibleWith"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackage, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackage))
            }
            m.SetAccessPackagesIncompatibleWith(res)
        }
        return nil
    }
    res["catalogId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["incompatibleAccessPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackage, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackage))
            }
            m.SetIncompatibleAccessPackages(res)
        }
        return nil
    }
    res["incompatibleGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Group, len(val))
            for i, v := range val {
                res[i] = *(v.(*Group))
            }
            m.SetIncompatibleGroups(res)
        }
        return nil
    }
    res["isHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHidden(val)
        }
        return nil
    }
    res["isRoleScopesVisible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRoleScopesVisible(val)
        }
        return nil
    }
    res["modifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedBy(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *AccessPackage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessPackage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageAssignmentPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignmentPolicies()))
        for i, v := range m.GetAccessPackageAssignmentPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceRoleScopes()))
        for i, v := range m.GetAccessPackageResourceRoleScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoleScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackagesIncompatibleWith() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackagesIncompatibleWith()))
        for i, v := range m.GetAccessPackagesIncompatibleWith() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncompatibleAccessPackages()))
        for i, v := range m.GetIncompatibleAccessPackages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("incompatibleAccessPackages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncompatibleGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncompatibleGroups()))
        for i, v := range m.GetIncompatibleGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *AccessPackage) SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicy)() {
    if m != nil {
        m.accessPackageAssignmentPolicies = value
    }
}
// SetAccessPackageCatalog sets the accessPackageCatalog property value. Read-only. Nullable.
func (m *AccessPackage) SetAccessPackageCatalog(value *AccessPackageCatalog)() {
    if m != nil {
        m.accessPackageCatalog = value
    }
}
// SetAccessPackageResourceRoleScopes sets the accessPackageResourceRoleScopes property value. Nullable.
func (m *AccessPackage) SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScope)() {
    if m != nil {
        m.accessPackageResourceRoleScopes = value
    }
}
// SetAccessPackagesIncompatibleWith sets the accessPackagesIncompatibleWith property value. The access packages that are incompatible with this package. Read-only.
func (m *AccessPackage) SetAccessPackagesIncompatibleWith(value []AccessPackage)() {
    if m != nil {
        m.accessPackagesIncompatibleWith = value
    }
}
// SetCatalogId sets the catalogId property value. Identifier of the access package catalog referencing this access package. Read-only.
func (m *AccessPackage) SetCatalogId(value *string)() {
    if m != nil {
        m.catalogId = value
    }
}
// SetCreatedBy sets the createdBy property value. The userPrincipalName of the user or identity of the subject who created this resource. Read-only.
func (m *AccessPackage) SetCreatedBy(value *string)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description of the access package.
func (m *AccessPackage) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the access package. Supports $filter (eq, contains).
func (m *AccessPackage) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIncompatibleAccessPackages sets the incompatibleAccessPackages property value. The  access packages whose assigned users are ineligible to be assigned this access package.
func (m *AccessPackage) SetIncompatibleAccessPackages(value []AccessPackage)() {
    if m != nil {
        m.incompatibleAccessPackages = value
    }
}
// SetIncompatibleGroups sets the incompatibleGroups property value. The groups whose members are ineligible to be assigned this access package.
func (m *AccessPackage) SetIncompatibleGroups(value []Group)() {
    if m != nil {
        m.incompatibleGroups = value
    }
}
// SetIsHidden sets the isHidden property value. Whether the access package is hidden from the requestor.
func (m *AccessPackage) SetIsHidden(value *bool)() {
    if m != nil {
        m.isHidden = value
    }
}
// SetIsRoleScopesVisible sets the isRoleScopesVisible property value. Indicates whether role scopes are visible.
func (m *AccessPackage) SetIsRoleScopesVisible(value *bool)() {
    if m != nil {
        m.isRoleScopesVisible = value
    }
}
// SetModifiedBy sets the modifiedBy property value. The userPrincipalName of the user who last modified this resource. Read-only.
func (m *AccessPackage) SetModifiedBy(value *string)() {
    if m != nil {
        m.modifiedBy = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
