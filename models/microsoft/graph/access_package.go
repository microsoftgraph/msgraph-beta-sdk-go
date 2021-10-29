package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackage struct {
    Entity
    // Read-only. Nullable.
    accessPackageAssignmentPolicies []AccessPackageAssignmentPolicy;
    // Read-only. Nullable.
    accessPackageCatalog *AccessPackageCatalog;
    // Nullable.
    accessPackageResourceRoleScopes []AccessPackageResourceRoleScope;
    // The access packages that are incompatible with this package. Read-only.
    accessPackagesIncompatibleWith []AccessPackage;
    // ID of the access package catalog referencing this access package. Read-only.
    catalogId *string;
    // UPN of the user or identity of the subject who created this resource. Read-only.
    createdBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the access package.
    description *string;
    // The display name of the access package.
    displayName *string;
    // The  access packages whose assigned users are ineligible to be assigned this access package.
    incompatibleAccessPackages []AccessPackage;
    // The groups whose members are ineligible to be assigned this access package.
    incompatibleGroups []Group;
    // Whether the access package is hidden from the requestor.
    isHidden *bool;
    // Indicates whether role scopes are visible.
    isRoleScopesVisible *bool;
    // The UPN of the user who last modified this resource. Read-only.
    modifiedBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new accessPackage and sets the default values.
func NewAccessPackage()(*AccessPackage) {
    m := &AccessPackage{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackageAssignmentPolicies property value. Read-only. Nullable.
func (m *AccessPackage) GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicy) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentPolicies
    }
}
// Gets the accessPackageCatalog property value. Read-only. Nullable.
func (m *AccessPackage) GetAccessPackageCatalog()(*AccessPackageCatalog) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageCatalog
    }
}
// Gets the accessPackageResourceRoleScopes property value. Nullable.
func (m *AccessPackage) GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScope) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRoleScopes
    }
}
// Gets the accessPackagesIncompatibleWith property value. The access packages that are incompatible with this package. Read-only.
func (m *AccessPackage) GetAccessPackagesIncompatibleWith()([]AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackagesIncompatibleWith
    }
}
// Gets the catalogId property value. ID of the access package catalog referencing this access package. Read-only.
func (m *AccessPackage) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
// Gets the createdBy property value. UPN of the user or identity of the subject who created this resource. Read-only.
func (m *AccessPackage) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The description of the access package.
func (m *AccessPackage) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name of the access package.
func (m *AccessPackage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the incompatibleAccessPackages property value. The  access packages whose assigned users are ineligible to be assigned this access package.
func (m *AccessPackage) GetIncompatibleAccessPackages()([]AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.incompatibleAccessPackages
    }
}
// Gets the incompatibleGroups property value. The groups whose members are ineligible to be assigned this access package.
func (m *AccessPackage) GetIncompatibleGroups()([]Group) {
    if m == nil {
        return nil
    } else {
        return m.incompatibleGroups
    }
}
// Gets the isHidden property value. Whether the access package is hidden from the requestor.
func (m *AccessPackage) GetIsHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHidden
    }
}
// Gets the isRoleScopesVisible property value. Indicates whether role scopes are visible.
func (m *AccessPackage) GetIsRoleScopesVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRoleScopesVisible
    }
}
// Gets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
func (m *AccessPackage) GetModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modifiedBy
    }
}
// Gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// The deserialization information for the current model
func (m *AccessPackage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignmentPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentPolicy() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageAssignmentPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageAssignmentPolicy))
        }
        m.SetAccessPackageAssignmentPolicies(res)
        return nil
    }
    res["accessPackageCatalog"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageCatalog() })
        if err != nil {
            return err
        }
        m.SetAccessPackageCatalog(val.(*AccessPackageCatalog))
        return nil
    }
    res["accessPackageResourceRoleScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceRoleScope() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageResourceRoleScope, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageResourceRoleScope))
        }
        m.SetAccessPackageResourceRoleScopes(res)
        return nil
    }
    res["accessPackagesIncompatibleWith"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        res := make([]AccessPackage, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackage))
        }
        m.SetAccessPackagesIncompatibleWith(res)
        return nil
    }
    res["catalogId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCatalogId(val)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreatedBy(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["incompatibleAccessPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        res := make([]AccessPackage, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackage))
        }
        m.SetIncompatibleAccessPackages(res)
        return nil
    }
    res["incompatibleGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        res := make([]Group, len(val))
        for i, v := range val {
            res[i] = *(v.(*Group))
        }
        m.SetIncompatibleGroups(res)
        return nil
    }
    res["isHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsHidden(val)
        return nil
    }
    res["isRoleScopesVisible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRoleScopesVisible(val)
        return nil
    }
    res["modifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModifiedBy(val)
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetModifiedDateTime(val)
        return nil
    }
    return res
}
func (m *AccessPackage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
// Sets the accessPackageAssignmentPolicies property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackageAssignmentPolicies property.
func (m *AccessPackage) SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicy)() {
    m.accessPackageAssignmentPolicies = value
}
// Sets the accessPackageCatalog property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackageCatalog property.
func (m *AccessPackage) SetAccessPackageCatalog(value *AccessPackageCatalog)() {
    m.accessPackageCatalog = value
}
// Sets the accessPackageResourceRoleScopes property value. Nullable.
// Parameters:
//  - value : Value to set for the accessPackageResourceRoleScopes property.
func (m *AccessPackage) SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScope)() {
    m.accessPackageResourceRoleScopes = value
}
// Sets the accessPackagesIncompatibleWith property value. The access packages that are incompatible with this package. Read-only.
// Parameters:
//  - value : Value to set for the accessPackagesIncompatibleWith property.
func (m *AccessPackage) SetAccessPackagesIncompatibleWith(value []AccessPackage)() {
    m.accessPackagesIncompatibleWith = value
}
// Sets the catalogId property value. ID of the access package catalog referencing this access package. Read-only.
// Parameters:
//  - value : Value to set for the catalogId property.
func (m *AccessPackage) SetCatalogId(value *string)() {
    m.catalogId = value
}
// Sets the createdBy property value. UPN of the user or identity of the subject who created this resource. Read-only.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *AccessPackage) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AccessPackage) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The description of the access package.
// Parameters:
//  - value : Value to set for the description property.
func (m *AccessPackage) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name of the access package.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AccessPackage) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the incompatibleAccessPackages property value. The  access packages whose assigned users are ineligible to be assigned this access package.
// Parameters:
//  - value : Value to set for the incompatibleAccessPackages property.
func (m *AccessPackage) SetIncompatibleAccessPackages(value []AccessPackage)() {
    m.incompatibleAccessPackages = value
}
// Sets the incompatibleGroups property value. The groups whose members are ineligible to be assigned this access package.
// Parameters:
//  - value : Value to set for the incompatibleGroups property.
func (m *AccessPackage) SetIncompatibleGroups(value []Group)() {
    m.incompatibleGroups = value
}
// Sets the isHidden property value. Whether the access package is hidden from the requestor.
// Parameters:
//  - value : Value to set for the isHidden property.
func (m *AccessPackage) SetIsHidden(value *bool)() {
    m.isHidden = value
}
// Sets the isRoleScopesVisible property value. Indicates whether role scopes are visible.
// Parameters:
//  - value : Value to set for the isRoleScopesVisible property.
func (m *AccessPackage) SetIsRoleScopesVisible(value *bool)() {
    m.isRoleScopesVisible = value
}
// Sets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
// Parameters:
//  - value : Value to set for the modifiedBy property.
func (m *AccessPackage) SetModifiedBy(value *string)() {
    m.modifiedBy = value
}
// Sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *AccessPackage) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
