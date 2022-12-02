package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder provides operations to manage the accessPackageResourceEnvironments property of the microsoft.graph.entitlementManagement entity.
type IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderGetQueryParameters a reference to the geolocation environment in which a resource is located.
type IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResources provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageResourceEnvironment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) AccessPackageResources()(*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourcesById provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageResourceEnvironment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) AccessPackageResourcesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesAccessPackageResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesAccessPackageResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderInternal instantiates a new AccessPackageResourceEnvironmentItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/{accessPackageResourceEnvironment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder instantiates a new AccessPackageResourceEnvironmentItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageResourceEnvironments for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a reference to the geolocation environment in which a resource is located.
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageResourceEnvironments in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackageResourceEnvironments for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a reference to the geolocation environment in which a resource is located.
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable), nil
}
// Patch update the navigation property accessPackageResourceEnvironments in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable), nil
}
