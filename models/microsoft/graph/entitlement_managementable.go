package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EntitlementManagementable 
type EntitlementManagementable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackageAssignmentApprovals()([]Approvalable)
    GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicyable)
    GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequestable)
    GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRoleable)
    GetAccessPackageAssignments()([]AccessPackageAssignmentable)
    GetAccessPackageCatalogs()([]AccessPackageCatalogable)
    GetAccessPackageResourceEnvironments()([]AccessPackageResourceEnvironmentable)
    GetAccessPackageResourceRequests()([]AccessPackageResourceRequestable)
    GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScopeable)
    GetAccessPackageResources()([]AccessPackageResourceable)
    GetAccessPackages()([]AccessPackageable)
    GetConnectedOrganizations()([]ConnectedOrganizationable)
    GetSettings()(EntitlementManagementSettingsable)
    SetAccessPackageAssignmentApprovals(value []Approvalable)()
    SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicyable)()
    SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequestable)()
    SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRoleable)()
    SetAccessPackageAssignments(value []AccessPackageAssignmentable)()
    SetAccessPackageCatalogs(value []AccessPackageCatalogable)()
    SetAccessPackageResourceEnvironments(value []AccessPackageResourceEnvironmentable)()
    SetAccessPackageResourceRequests(value []AccessPackageResourceRequestable)()
    SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScopeable)()
    SetAccessPackageResources(value []AccessPackageResourceable)()
    SetAccessPackages(value []AccessPackageable)()
    SetConnectedOrganizations(value []ConnectedOrganizationable)()
    SetSettings(value EntitlementManagementSettingsable)()
}
