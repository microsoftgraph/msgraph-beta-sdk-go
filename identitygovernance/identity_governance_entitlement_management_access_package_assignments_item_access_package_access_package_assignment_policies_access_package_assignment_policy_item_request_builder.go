package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) AccessPackage()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemAccessPackageRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageCatalog provides operations to manage the accessPackageCatalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) AccessPackageCatalog()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemAccessPackageCatalogRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal instantiates a new AccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder instantiates a new AccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentPolicies for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CustomExtensionHandlers()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomExtensionHandlersById provides operations to manage the customExtensionHandlers property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) CustomExtensionHandlersById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCustomExtensionHandlerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customExtensionHandler%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCustomExtensionHandlerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property accessPackageAssignmentPolicies for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, error) {
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
