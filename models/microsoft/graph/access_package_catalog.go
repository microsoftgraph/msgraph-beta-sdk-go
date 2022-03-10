package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageCatalog provides operations to manage the identityGovernance singleton.
type AccessPackageCatalog struct {
    Entity
    // The roles in each resource in a catalog. Read-only.
    accessPackageResourceRoles []AccessPackageResourceRoleable;
    // Read-only. Nullable.
    accessPackageResources []AccessPackageResourceable;
    // Read-only.
    accessPackageResourceScopes []AccessPackageResourceScopeable;
    // The access packages in this catalog. Read-only. Nullable.
    accessPackages []AccessPackageable;
    // Has the value Published if the access packages are available for management.
    catalogStatus *string;
    // Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
    catalogType *string;
    // UPN of the user who created this resource. Read-only.
    createdBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    customAccessPackageWorkflowExtensions []CustomAccessPackageWorkflowExtensionable;
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
// NewAccessPackageCatalog instantiates a new accessPackageCatalog and sets the default values.
func NewAccessPackageCatalog()(*AccessPackageCatalog) {
    m := &AccessPackageCatalog{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageCatalogFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessPackageCatalog(), nil
}
// GetAccessPackageResourceRoles gets the accessPackageResourceRoles property value. The roles in each resource in a catalog. Read-only.
func (m *AccessPackageCatalog) GetAccessPackageResourceRoles()([]AccessPackageResourceRoleable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRoles
    }
}
// GetAccessPackageResources gets the accessPackageResources property value. Read-only. Nullable.
func (m *AccessPackageCatalog) GetAccessPackageResources()([]AccessPackageResourceable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResources
    }
}
// GetAccessPackageResourceScopes gets the accessPackageResourceScopes property value. Read-only.
func (m *AccessPackageCatalog) GetAccessPackageResourceScopes()([]AccessPackageResourceScopeable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceScopes
    }
}
// GetAccessPackages gets the accessPackages property value. The access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageCatalog) GetAccessPackages()([]AccessPackageable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackages
    }
}
// GetCatalogStatus gets the catalogStatus property value. Has the value Published if the access packages are available for management.
func (m *AccessPackageCatalog) GetCatalogStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogStatus
    }
}
// GetCatalogType gets the catalogType property value. Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
func (m *AccessPackageCatalog) GetCatalogType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogType
    }
}
// GetCreatedBy gets the createdBy property value. UPN of the user who created this resource. Read-only.
func (m *AccessPackageCatalog) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCustomAccessPackageWorkflowExtensions gets the customAccessPackageWorkflowExtensions property value. 
func (m *AccessPackageCatalog) GetCustomAccessPackageWorkflowExtensions()([]CustomAccessPackageWorkflowExtensionable) {
    if m == nil {
        return nil
    } else {
        return m.customAccessPackageWorkflowExtensions
    }
}
// GetDescription gets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the access package catalog.
func (m *AccessPackageCatalog) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageCatalog) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResourceRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRoleable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageResourceRoleable)
            }
            m.SetAccessPackageResourceRoles(res)
        }
        return nil
    }
    res["accessPackageResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageResourceable)
            }
            m.SetAccessPackageResources(res)
        }
        return nil
    }
    res["accessPackageResourceScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceScopeable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageResourceScopeable)
            }
            m.SetAccessPackageResourceScopes(res)
        }
        return nil
    }
    res["accessPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageable)
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
    res["customAccessPackageWorkflowExtensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomAccessPackageWorkflowExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomAccessPackageWorkflowExtensionable, len(val))
            for i, v := range val {
                res[i] = v.(CustomAccessPackageWorkflowExtensionable)
            }
            m.SetCustomAccessPackageWorkflowExtensions(res)
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
// GetIsExternallyVisible gets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) GetIsExternallyVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isExternallyVisible
    }
}
// GetModifiedBy gets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
func (m *AccessPackageCatalog) GetModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modifiedBy
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
func (m *AccessPackageCatalog) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessPackageCatalog) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageResourceRoles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceRoles()))
        for i, v := range m.GetAccessPackageResourceRoles() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResources()))
        for i, v := range m.GetAccessPackageResources() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceScopes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceScopes()))
        for i, v := range m.GetAccessPackageResourceScopes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomAccessPackageWorkflowExtensions()))
        for i, v := range m.GetCustomAccessPackageWorkflowExtensions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAccessPackageResourceRoles sets the accessPackageResourceRoles property value. The roles in each resource in a catalog. Read-only.
func (m *AccessPackageCatalog) SetAccessPackageResourceRoles(value []AccessPackageResourceRoleable)() {
    if m != nil {
        m.accessPackageResourceRoles = value
    }
}
// SetAccessPackageResources sets the accessPackageResources property value. Read-only. Nullable.
func (m *AccessPackageCatalog) SetAccessPackageResources(value []AccessPackageResourceable)() {
    if m != nil {
        m.accessPackageResources = value
    }
}
// SetAccessPackageResourceScopes sets the accessPackageResourceScopes property value. Read-only.
func (m *AccessPackageCatalog) SetAccessPackageResourceScopes(value []AccessPackageResourceScopeable)() {
    if m != nil {
        m.accessPackageResourceScopes = value
    }
}
// SetAccessPackages sets the accessPackages property value. The access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageCatalog) SetAccessPackages(value []AccessPackageable)() {
    if m != nil {
        m.accessPackages = value
    }
}
// SetCatalogStatus sets the catalogStatus property value. Has the value Published if the access packages are available for management.
func (m *AccessPackageCatalog) SetCatalogStatus(value *string)() {
    if m != nil {
        m.catalogStatus = value
    }
}
// SetCatalogType sets the catalogType property value. Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
func (m *AccessPackageCatalog) SetCatalogType(value *string)() {
    if m != nil {
        m.catalogType = value
    }
}
// SetCreatedBy sets the createdBy property value. UPN of the user who created this resource. Read-only.
func (m *AccessPackageCatalog) SetCreatedBy(value *string)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCustomAccessPackageWorkflowExtensions sets the customAccessPackageWorkflowExtensions property value. 
func (m *AccessPackageCatalog) SetCustomAccessPackageWorkflowExtensions(value []CustomAccessPackageWorkflowExtensionable)() {
    if m != nil {
        m.customAccessPackageWorkflowExtensions = value
    }
}
// SetDescription sets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the access package catalog.
func (m *AccessPackageCatalog) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsExternallyVisible sets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) SetIsExternallyVisible(value *bool)() {
    if m != nil {
        m.isExternallyVisible = value
    }
}
// SetModifiedBy sets the modifiedBy property value. The UPN of the user who last modified this resource. Read-only.
func (m *AccessPackageCatalog) SetModifiedBy(value *string)() {
    if m != nil {
        m.modifiedBy = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
