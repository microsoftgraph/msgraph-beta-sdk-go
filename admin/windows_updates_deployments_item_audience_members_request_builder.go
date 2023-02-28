package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetQueryParameters list the updatableAsset resources that are members of a deploymentAudience.
type WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetQueryParameters
}
// WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/members{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) Count()(*WindowsUpdatesDeploymentsItemAudienceMembersCountRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get list the updatableAsset resources that are members of a deploymentAudience.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/windowsupdates-deploymentaudience-list-members?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable), nil
}
// Post create new navigation property to members for admin
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) Post(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// ToGetRequestInformation list the updatableAsset resources that are members of a deploymentAudience.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to members for admin
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// WindowsUpdatesEnrollAssets provides operations to call the enrollAssets method.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) WindowsUpdatesEnrollAssets()(*WindowsUpdatesDeploymentsItemAudienceMembersWindowsUpdatesEnrollAssetsRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersWindowsUpdatesEnrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsUpdatesEnrollAssetsById provides operations to call the enrollAssetsById method.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) WindowsUpdatesEnrollAssetsById()(*WindowsUpdatesDeploymentsItemAudienceMembersWindowsUpdatesEnrollAssetsByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersWindowsUpdatesEnrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsUpdatesUnenrollAssets provides operations to call the unenrollAssets method.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) WindowsUpdatesUnenrollAssets()(*WindowsUpdatesDeploymentsItemAudienceMembersWindowsUpdatesUnenrollAssetsRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersWindowsUpdatesUnenrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsUpdatesUnenrollAssetsById provides operations to call the unenrollAssetsById method.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersRequestBuilder) WindowsUpdatesUnenrollAssetsById()(*WindowsUpdatesDeploymentsItemAudienceMembersWindowsUpdatesUnenrollAssetsByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersWindowsUpdatesUnenrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
