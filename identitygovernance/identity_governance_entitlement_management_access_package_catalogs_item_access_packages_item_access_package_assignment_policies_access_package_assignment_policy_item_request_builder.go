package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) AccessPackage()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesItemAccessPackageRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesItemAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageCatalog provides operations to manage the accessPackageCatalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) AccessPackageCatalog()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesItemAccessPackageCatalogRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesItemAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal instantiates a new AccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder instantiates a new AccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentPolicies for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable. Supports $expand.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageAssignmentPolicies in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CustomExtensionHandlers provides operations to manage the customExtensionHandlers property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CustomExtensionHandlers()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomExtensionHandlersById provides operations to manage the customExtensionHandlers property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CustomExtensionHandlersById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCustomExtensionHandlerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customExtensionHandler%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCustomExtensionHandlerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property accessPackageAssignmentPolicies for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Nullable. Supports $expand.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable), nil
}
// Patch update the navigation property accessPackageAssignmentPolicies in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable), nil
}
