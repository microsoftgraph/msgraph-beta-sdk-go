package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignment entity.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetQueryParameters read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentPolicies provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageAssignmentPolicies()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackageAssignmentPoliciesById provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AccessPackageCatalog provides operations to manage the accessPackageCatalog property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageCatalog()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageCatalogRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackageResourceRoleScopes provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageResourceRoleScopes()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackageResourceRoleScopesById provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope%2Did"] = id
    }
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AccessPackagesIncompatibleWith provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackagesIncompatibleWith()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackagesIncompatibleWithRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackagesIncompatibleWithById provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) AccessPackagesIncompatibleWithById(id string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackagesIncompatibleWithAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did"] = id
    }
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackagesIncompatibleWithAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderInternal instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackage for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable), nil
}
// IncompatibleAccessPackages provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) IncompatibleAccessPackages()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncompatibleAccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.incompatibleAccessPackages.item collection
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) IncompatibleAccessPackagesById(id string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did"] = id
    }
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// IncompatibleGroups provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) IncompatibleGroups()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncompatibleGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.incompatibleGroups.item collection
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) IncompatibleGroupsById(id string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleGroupsGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphGetApplicablePolicyRequirements provides operations to call the getApplicablePolicyRequirements method.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) MicrosoftGraphGetApplicablePolicyRequirements()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMoveToCatalog provides operations to call the moveToCatalog method.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) MicrosoftGraphMoveToCatalog()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property accessPackage in identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackage for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property accessPackage in identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
