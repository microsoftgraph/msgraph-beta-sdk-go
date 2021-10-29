package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EntitlementManagement struct {
    Entity
    // 
    accessPackageAssignmentApprovals []Approval;
    // 
    accessPackageAssignmentPolicies []AccessPackageAssignmentPolicy;
    // 
    accessPackageAssignmentRequests []AccessPackageAssignmentRequest;
    // 
    accessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRole;
    // 
    accessPackageAssignments []AccessPackageAssignment;
    // 
    accessPackageCatalogs []AccessPackageCatalog;
    // 
    accessPackageResourceEnvironments []AccessPackageResourceEnvironment;
    // 
    accessPackageResourceRequests []AccessPackageResourceRequest;
    // 
    accessPackageResourceRoleScopes []AccessPackageResourceRoleScope;
    // 
    accessPackageResources []AccessPackageResource;
    // 
    accessPackages []AccessPackage;
    // 
    connectedOrganizations []ConnectedOrganization;
    // 
    settings *EntitlementManagementSettings;
}
// Instantiates a new entitlementManagement and sets the default values.
func NewEntitlementManagement()(*EntitlementManagement) {
    m := &EntitlementManagement{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackageAssignmentApprovals property value. 
func (m *EntitlementManagement) GetAccessPackageAssignmentApprovals()([]Approval) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentApprovals
    }
}
// Gets the accessPackageAssignmentPolicies property value. 
func (m *EntitlementManagement) GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicy) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentPolicies
    }
}
// Gets the accessPackageAssignmentRequests property value. 
func (m *EntitlementManagement) GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentRequests
    }
}
// Gets the accessPackageAssignmentResourceRoles property value. 
func (m *EntitlementManagement) GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRole) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentResourceRoles
    }
}
// Gets the accessPackageAssignments property value. 
func (m *EntitlementManagement) GetAccessPackageAssignments()([]AccessPackageAssignment) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignments
    }
}
// Gets the accessPackageCatalogs property value. 
func (m *EntitlementManagement) GetAccessPackageCatalogs()([]AccessPackageCatalog) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageCatalogs
    }
}
// Gets the accessPackageResourceEnvironments property value. 
func (m *EntitlementManagement) GetAccessPackageResourceEnvironments()([]AccessPackageResourceEnvironment) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceEnvironments
    }
}
// Gets the accessPackageResourceRequests property value. 
func (m *EntitlementManagement) GetAccessPackageResourceRequests()([]AccessPackageResourceRequest) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRequests
    }
}
// Gets the accessPackageResourceRoleScopes property value. 
func (m *EntitlementManagement) GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScope) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResourceRoleScopes
    }
}
// Gets the accessPackageResources property value. 
func (m *EntitlementManagement) GetAccessPackageResources()([]AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResources
    }
}
// Gets the accessPackages property value. 
func (m *EntitlementManagement) GetAccessPackages()([]AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackages
    }
}
// Gets the connectedOrganizations property value. 
func (m *EntitlementManagement) GetConnectedOrganizations()([]ConnectedOrganization) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganizations
    }
}
// Gets the settings property value. 
func (m *EntitlementManagement) GetSettings()(*EntitlementManagementSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// The deserialization information for the current model
func (m *EntitlementManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignmentApprovals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApproval() })
        if err != nil {
            return err
        }
        res := make([]Approval, len(val))
        for i, v := range val {
            res[i] = *(v.(*Approval))
        }
        m.SetAccessPackageAssignmentApprovals(res)
        return nil
    }
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
    res["accessPackageAssignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentRequest() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageAssignmentRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageAssignmentRequest))
        }
        m.SetAccessPackageAssignmentRequests(res)
        return nil
    }
    res["accessPackageAssignmentResourceRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentResourceRole() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageAssignmentResourceRole, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageAssignmentResourceRole))
        }
        m.SetAccessPackageAssignmentResourceRoles(res)
        return nil
    }
    res["accessPackageAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignment() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageAssignment))
        }
        m.SetAccessPackageAssignments(res)
        return nil
    }
    res["accessPackageCatalogs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageCatalog() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageCatalog, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageCatalog))
        }
        m.SetAccessPackageCatalogs(res)
        return nil
    }
    res["accessPackageResourceEnvironments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceEnvironment() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageResourceEnvironment, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageResourceEnvironment))
        }
        m.SetAccessPackageResourceEnvironments(res)
        return nil
    }
    res["accessPackageResourceRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceRequest() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageResourceRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageResourceRequest))
        }
        m.SetAccessPackageResourceRequests(res)
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
    res["accessPackageResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResource() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageResource))
        }
        m.SetAccessPackageResources(res)
        return nil
    }
    res["accessPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        res := make([]AccessPackage, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackage))
        }
        m.SetAccessPackages(res)
        return nil
    }
    res["connectedOrganizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectedOrganization() })
        if err != nil {
            return err
        }
        res := make([]ConnectedOrganization, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConnectedOrganization))
        }
        m.SetConnectedOrganizations(res)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntitlementManagementSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*EntitlementManagementSettings))
        return nil
    }
    return res
}
func (m *EntitlementManagement) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EntitlementManagement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignmentApprovals()))
        for i, v := range m.GetAccessPackageAssignmentApprovals() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentApprovals", cast)
        if err != nil {
            return err
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignmentRequests()))
        for i, v := range m.GetAccessPackageAssignmentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignmentResourceRoles()))
        for i, v := range m.GetAccessPackageAssignmentResourceRoles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentResourceRoles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignments()))
        for i, v := range m.GetAccessPackageAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageCatalogs()))
        for i, v := range m.GetAccessPackageCatalogs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageCatalogs", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceEnvironments()))
        for i, v := range m.GetAccessPackageResourceEnvironments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceEnvironments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResourceRequests()))
        for i, v := range m.GetAccessPackageResourceRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRequests", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConnectedOrganizations()))
        for i, v := range m.GetConnectedOrganizations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("connectedOrganizations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accessPackageAssignmentApprovals property value. 
// Parameters:
//  - value : Value to set for the accessPackageAssignmentApprovals property.
func (m *EntitlementManagement) SetAccessPackageAssignmentApprovals(value []Approval)() {
    m.accessPackageAssignmentApprovals = value
}
// Sets the accessPackageAssignmentPolicies property value. 
// Parameters:
//  - value : Value to set for the accessPackageAssignmentPolicies property.
func (m *EntitlementManagement) SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicy)() {
    m.accessPackageAssignmentPolicies = value
}
// Sets the accessPackageAssignmentRequests property value. 
// Parameters:
//  - value : Value to set for the accessPackageAssignmentRequests property.
func (m *EntitlementManagement) SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequest)() {
    m.accessPackageAssignmentRequests = value
}
// Sets the accessPackageAssignmentResourceRoles property value. 
// Parameters:
//  - value : Value to set for the accessPackageAssignmentResourceRoles property.
func (m *EntitlementManagement) SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRole)() {
    m.accessPackageAssignmentResourceRoles = value
}
// Sets the accessPackageAssignments property value. 
// Parameters:
//  - value : Value to set for the accessPackageAssignments property.
func (m *EntitlementManagement) SetAccessPackageAssignments(value []AccessPackageAssignment)() {
    m.accessPackageAssignments = value
}
// Sets the accessPackageCatalogs property value. 
// Parameters:
//  - value : Value to set for the accessPackageCatalogs property.
func (m *EntitlementManagement) SetAccessPackageCatalogs(value []AccessPackageCatalog)() {
    m.accessPackageCatalogs = value
}
// Sets the accessPackageResourceEnvironments property value. 
// Parameters:
//  - value : Value to set for the accessPackageResourceEnvironments property.
func (m *EntitlementManagement) SetAccessPackageResourceEnvironments(value []AccessPackageResourceEnvironment)() {
    m.accessPackageResourceEnvironments = value
}
// Sets the accessPackageResourceRequests property value. 
// Parameters:
//  - value : Value to set for the accessPackageResourceRequests property.
func (m *EntitlementManagement) SetAccessPackageResourceRequests(value []AccessPackageResourceRequest)() {
    m.accessPackageResourceRequests = value
}
// Sets the accessPackageResourceRoleScopes property value. 
// Parameters:
//  - value : Value to set for the accessPackageResourceRoleScopes property.
func (m *EntitlementManagement) SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScope)() {
    m.accessPackageResourceRoleScopes = value
}
// Sets the accessPackageResources property value. 
// Parameters:
//  - value : Value to set for the accessPackageResources property.
func (m *EntitlementManagement) SetAccessPackageResources(value []AccessPackageResource)() {
    m.accessPackageResources = value
}
// Sets the accessPackages property value. 
// Parameters:
//  - value : Value to set for the accessPackages property.
func (m *EntitlementManagement) SetAccessPackages(value []AccessPackage)() {
    m.accessPackages = value
}
// Sets the connectedOrganizations property value. 
// Parameters:
//  - value : Value to set for the connectedOrganizations property.
func (m *EntitlementManagement) SetConnectedOrganizations(value []ConnectedOrganization)() {
    m.connectedOrganizations = value
}
// Sets the settings property value. 
// Parameters:
//  - value : Value to set for the settings property.
func (m *EntitlementManagement) SetSettings(value *EntitlementManagementSettings)() {
    m.settings = value
}
