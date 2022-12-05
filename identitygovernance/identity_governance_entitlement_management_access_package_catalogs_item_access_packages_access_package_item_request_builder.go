package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters the access packages in this catalog. Read-only. Nullable. Supports $expand.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentPolicies provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackageAssignmentPolicies()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageCatalog provides operations to manage the accessPackageCatalog property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackageCatalog()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageCatalogRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopes provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackageResourceRoleScopes()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackagesIncompatibleWith provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackagesIncompatibleWithRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesIncompatibleWithById provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackagesIncompatibleWithById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackagesIncompatibleWithAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackagesIncompatibleWithAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackages for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackages in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackages for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) GetApplicablePolicyRequirements()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleAccessPackages provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleAccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.incompatibleAccessPackages.item collection
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) IncompatibleAccessPackagesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IncompatibleGroups provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) IncompatibleGroups()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.incompatibleGroups.item collection
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) IncompatibleGroupsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MoveToCatalog provides operations to call the moveToCatalog method.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) MoveToCatalog()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property accessPackages in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
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
