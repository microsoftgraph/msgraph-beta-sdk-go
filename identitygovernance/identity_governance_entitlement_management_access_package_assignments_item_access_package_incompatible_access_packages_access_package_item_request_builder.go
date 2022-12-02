package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\incompatibleAccessPackages\{accessPackage-id}
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/incompatibleAccessPackages/{accessPackage%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder) Ref()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
