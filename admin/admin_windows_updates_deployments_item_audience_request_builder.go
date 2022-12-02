package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder provides operations to manage the audience property of the microsoft.graph.windowsUpdates.deployment entity.
type AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderGetQueryParameters specifies the audience to which content is deployed.
type AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderGetQueryParameters
}
// AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderInternal instantiates a new AudienceRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) {
    m := &AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder instantiates a new AudienceRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property audience for admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation specifies the audience to which content is deployed.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property audience in admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property audience for admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Exclusions provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) Exclusions()(*AdminWindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExclusionsById provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) ExclusionsById(id string)(*AdminWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset%2Did"] = id
    }
    return NewAdminWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get specifies the audience to which content is deployed.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateDeploymentAudienceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) Members()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) MembersById(id string)(*AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset%2Did"] = id
    }
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property audience in admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateDeploymentAudienceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable), nil
}
// UpdateAudience provides operations to call the updateAudience method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) UpdateAudience()(*AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateAudienceById provides operations to call the updateAudienceById method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceRequestBuilder) UpdateAudienceById()(*AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceByIdRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
