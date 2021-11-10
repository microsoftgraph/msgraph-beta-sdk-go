package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageCatalog struct {
    Entity
    // 
    accessPackageResourceRoles []AccessPackageResourceRole;
    // Read-only. Nullable.
    accessPackageResources []AccessPackageResource;
    // 
    accessPackageResourceScopes []AccessPackageResourceScope;
    // The access packages in this catalog. Read-only. Nullable.
    accessPackages []AccessPackage;
    // Has the value Published if the access packages are available for management.
    catalogStatus *string;
    // One of UserManaged or ServiceDefault.
    catalogType *string;
    // UPN of the user who created this resource. Read-only.
    createdBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the access package catalog.
    description *string;
    // The display name of the access package catalog.
    displayName *string;
    // Whether the access packages in this catalog can be requested by users outside of the tenant.
    isExternallyVisible *bool;
    // The UPN of the user who last modified this resource. Read-only.
    modifiedBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new accessPackageCatalog and sets the default values.
func NewAccessPackageCatalog()(*AccessPackageCatalog) {
    m := &AccessPackageCatalog{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackageResourceRoles property value. 
func (m *AccessPackageCatalog) GetAccessPackageResourceRoles()([]AccessPackageResourceRole) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRoles
    }
}
// Gets the accessPackageResources property value. Read-only. Nullable.
func (m *AccessPackageCatalog) GetAccessPackageResources()([]AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResources
    }
}
// Gets the accessPackageResourceScopes property value. 
func (m *AccessPackageCatalog) GetAccessPackageResourceScopes()([]AccessPackageResourceScope) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceScopes
    }
}
// Gets the accessPackages property value. The access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageCatalog) GetAccessPackages()([]AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackages
    }
}
// Gets the catalogStatus property value. Has the value Published if the access packages are available for management.
func (m *AccessPackageCatalog) GetCatalogStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogStatus
    }
}
// Gets the catalogType property value. One of UserManaged or ServiceDefault.
func (m *AccessPackageCatalog) GetCatalogType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogType
    }
}
// Gets the createdBy property value. UPN of the user who created this resource. Read-only.
func (m *AccessPackageCatalog) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name of the access package catalog.
func (m *AccessPackageCatalog) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) GetIsExternallyVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isExternallyVisible
    }
}
// Gets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
func (m *AccessPackageCatalog) GetModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modifiedBy
    }
}
// Gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// The deserialization information for the current model
func (m *AccessPackageCatalog) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResourceRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceRole() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRole, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageResourceRole))
            }
            m.SetAccessPackageResourceRoles(res)
        }
        return nil
    }
    res["accessPackageResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageResource))
            }
            m.SetAccessPackageResources(res)
        }
        return nil
    }
    res["accessPackageResourceScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageResourceScope))
            }
            m.SetAccessPackageResourceScopes(res)
        }
        return nil
    }
    res["accessPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackage, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackage))
            }
            m.SetAccessPackages(res)
        }
        return nil
    }
    res["catalogStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogStatus(val)
        }
        return nil
    }
    res["catalogType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogType(val)
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
    res["isExternallyVisible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternallyVisible(val)
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
func (m *AccessPackageCatalog) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageCatalog) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceRoles()))
        for i, v := range m.GetAccessPackageResourceRoles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResources()))
        for i, v := range m.GetAccessPackageResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceScopes()))
        for i, v := range m.GetAccessPackageResourceScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceScopes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the accessPackageResourceRoles property value. 
// Parameters:
//  - value : Value to set for the accessPackageResourceRoles property.
func (m *AccessPackageCatalog) SetAccessPackageResourceRoles(value []AccessPackageResourceRole)() {
    m.accessPackageResourceRoles = value
}
// Sets the accessPackageResources property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackageResources property.
func (m *AccessPackageCatalog) SetAccessPackageResources(value []AccessPackageResource)() {
    m.accessPackageResources = value
}
// Sets the accessPackageResourceScopes property value. 
// Parameters:
//  - value : Value to set for the accessPackageResourceScopes property.
func (m *AccessPackageCatalog) SetAccessPackageResourceScopes(value []AccessPackageResourceScope)() {
    m.accessPackageResourceScopes = value
}
// Sets the accessPackages property value. The access packages in this catalog. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackages property.
func (m *AccessPackageCatalog) SetAccessPackages(value []AccessPackage)() {
    m.accessPackages = value
}
// Sets the catalogStatus property value. Has the value Published if the access packages are available for management.
// Parameters:
//  - value : Value to set for the catalogStatus property.
func (m *AccessPackageCatalog) SetCatalogStatus(value *string)() {
    m.catalogStatus = value
}
// Sets the catalogType property value. One of UserManaged or ServiceDefault.
// Parameters:
//  - value : Value to set for the catalogType property.
func (m *AccessPackageCatalog) SetCatalogType(value *string)() {
    m.catalogType = value
}
// Sets the createdBy property value. UPN of the user who created this resource. Read-only.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *AccessPackageCatalog) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AccessPackageCatalog) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The description of the access package catalog.
// Parameters:
//  - value : Value to set for the description property.
func (m *AccessPackageCatalog) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name of the access package catalog.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AccessPackageCatalog) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
// Parameters:
//  - value : Value to set for the isExternallyVisible property.
func (m *AccessPackageCatalog) SetIsExternallyVisible(value *bool)() {
    m.isExternallyVisible = value
}
// Sets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
// Parameters:
//  - value : Value to set for the modifiedBy property.
func (m *AccessPackageCatalog) SetModifiedBy(value *string)() {
    m.modifiedBy = value
}
// Sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *AccessPackageCatalog) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
