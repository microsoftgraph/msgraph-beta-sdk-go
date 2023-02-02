package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesRequestBuilder provides operations to manage the updates property of the microsoft.graph.adminWindows entity.
type WindowsUpdatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsUpdatesRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesRequestBuilderGetQueryParameters get updates from admin
type WindowsUpdatesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesRequestBuilderGetQueryParameters
}
// WindowsUpdatesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) Catalog()(*WindowsUpdatesCatalogRequestBuilder) {
    return NewWindowsUpdatesCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewWindowsUpdatesRequestBuilderInternal instantiates a new UpdatesRequestBuilder and sets the default values.
func NewWindowsUpdatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesRequestBuilder) {
    m := &WindowsUpdatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWindowsUpdatesRequestBuilder instantiates a new UpdatesRequestBuilder and sets the default values.
func NewWindowsUpdatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property updates for admin
func (m *WindowsUpdatesRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentAudiences provides operations to manage the deploymentAudiences property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) DeploymentAudiences()(*WindowsUpdatesDeploymentAudiencesRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeploymentAudiencesById provides operations to manage the deploymentAudiences property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) DeploymentAudiencesById(id string)(*WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deploymentAudience%2Did"] = id
    }
    return NewWindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Deployments provides operations to manage the deployments property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) Deployments()(*WindowsUpdatesDeploymentsRequestBuilder) {
    return NewWindowsUpdatesDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeploymentsById provides operations to manage the deployments property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) DeploymentsById(id string)(*WindowsUpdatesDeploymentsDeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deployment%2Did"] = id
    }
    return NewWindowsUpdatesDeploymentsDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get updates from admin
func (m *WindowsUpdatesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdminWindowsUpdatesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdminWindowsUpdatesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdminWindowsUpdatesable), nil
}
// Patch update the navigation property updates in admin
func (m *WindowsUpdatesRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdminWindowsUpdatesable, requestConfiguration *WindowsUpdatesRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdminWindowsUpdatesable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdminWindowsUpdatesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdminWindowsUpdatesable), nil
}
// ResourceConnections provides operations to manage the resourceConnections property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) ResourceConnections()(*WindowsUpdatesResourceConnectionsRequestBuilder) {
    return NewWindowsUpdatesResourceConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourceConnectionsById provides operations to manage the resourceConnections property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) ResourceConnectionsById(id string)(*WindowsUpdatesResourceConnectionsResourceConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceConnection%2Did"] = id
    }
    return NewWindowsUpdatesResourceConnectionsResourceConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property updates for admin
func (m *WindowsUpdatesRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get updates from admin
func (m *WindowsUpdatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property updates in admin
func (m *WindowsUpdatesRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdminWindowsUpdatesable, requestConfiguration *WindowsUpdatesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdatableAssets provides operations to manage the updatableAssets property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) UpdatableAssets()(*WindowsUpdatesUpdatableAssetsRequestBuilder) {
    return NewWindowsUpdatesUpdatableAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UpdatableAssetsById provides operations to manage the updatableAssets property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) UpdatableAssetsById(id string)(*WindowsUpdatesUpdatableAssetsUpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset%2Did"] = id
    }
    return NewWindowsUpdatesUpdatableAssetsUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// UpdatePolicies provides operations to manage the updatePolicies property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) UpdatePolicies()(*WindowsUpdatesUpdatePoliciesRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UpdatePoliciesById provides operations to manage the updatePolicies property of the microsoft.graph.adminWindowsUpdates entity.
func (m *WindowsUpdatesRequestBuilder) UpdatePoliciesById(id string)(*WindowsUpdatesUpdatePoliciesUpdatePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatePolicy%2Did"] = id
    }
    return NewWindowsUpdatesUpdatePoliciesUpdatePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
