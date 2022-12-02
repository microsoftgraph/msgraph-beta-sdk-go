package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\incompatibleGroups\{group-id}
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/incompatibleGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder) Ref()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsItemRefRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
