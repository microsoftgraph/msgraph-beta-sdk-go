package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignment entity.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetQueryParameters read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentPolicies provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageAssignmentPolicies()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageCatalog provides operations to manage the accessPackageCatalog property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageCatalog()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageCatalogRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopes provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageResourceRoleScopes()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackagesIncompatibleWith provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackagesIncompatibleWith()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackagesIncompatibleWithRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesIncompatibleWithById provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackagesIncompatibleWithById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackagesIncompatibleWithAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackagesIncompatibleWithAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderInternal instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackage for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property accessPackage in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property accessPackage for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable), nil
}
// GetApplicablePolicyRequirements provides operations to call the getApplicablePolicyRequirements method.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) GetApplicablePolicyRequirements()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleAccessPackages provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) IncompatibleAccessPackages()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleAccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.incompatibleAccessPackages.item collection
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) IncompatibleAccessPackagesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IncompatibleGroups provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) IncompatibleGroups()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.incompatibleGroups.item collection
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) IncompatibleGroupsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MoveToCatalog provides operations to call the moveToCatalog method.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) MoveToCatalog()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property accessPackage in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable), nil
}
