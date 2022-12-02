package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetQueryParameters list the updatableAsset resources that are members of a deploymentAudience.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetQueryParameters
}
// AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) {
    m := &AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/members{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) Count()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersCountRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation list the updatableAsset resources that are members of a deploymentAudience.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to members for admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) CreatePostRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// EnrollAssets provides operations to call the enrollAssets method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) EnrollAssets()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersEnrollAssetsRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersEnrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollAssetsById provides operations to call the enrollAssetsById method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) EnrollAssetsById()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersEnrollAssetsByIdRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersEnrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get list the updatableAsset resources that are members of a deploymentAudience.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable), nil
}
// Post create new navigation property to members for admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) Post(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// UnenrollAssets provides operations to call the unenrollAssets method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) UnenrollAssets()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersUnenrollAssetsRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersUnenrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnenrollAssetsById provides operations to call the unenrollAssetsById method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) UnenrollAssetsById()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersUnenrollAssetsByIdRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersUnenrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
