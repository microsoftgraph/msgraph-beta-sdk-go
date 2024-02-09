package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderGetQueryParameters list the updatableAsset resources that are excluded from a deploymentAudience.
type WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderGetQueryParameters
}
// WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUpdatableAssetId provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) ByUpdatableAssetId(updatableAssetId string)(*WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if updatableAssetId != "" {
        urlTplParams["updatableAsset%2Did"] = updatableAssetId
    }
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/exclusions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder instantiates a new WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsCountRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) Count()(*WindowsUpdatesDeploymentsItemAudienceExclusionsCountRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the updatableAsset resources that are excluded from a deploymentAudience.
// returns a UpdatableAssetCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-deploymentaudience-list-exclusions?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable), nil
}
// MicrosoftGraphWindowsUpdatesEnrollAssets provides operations to call the enrollAssets method.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) MicrosoftGraphWindowsUpdatesEnrollAssets()(*WindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesEnrollAssetsById provides operations to call the enrollAssetsById method.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) MicrosoftGraphWindowsUpdatesEnrollAssetsById()(*WindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesUnenrollAssets provides operations to call the unenrollAssets method.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) MicrosoftGraphWindowsUpdatesUnenrollAssets()(*WindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesUnenrollAssetsById provides operations to call the unenrollAssetsById method.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) MicrosoftGraphWindowsUpdatesUnenrollAssetsById()(*WindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to exclusions for admin
// returns a UpdatableAssetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) Post(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// ToGetRequestInformation list the updatableAsset resources that are excluded from a deploymentAudience.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to exclusions for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/exclusions", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
