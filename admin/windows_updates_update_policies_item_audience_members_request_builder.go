package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderGetQueryParameters list the updatableAsset resources that are members of a deploymentAudience.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderGetQueryParameters
}
// WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUpdatableAssetId provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceMembersUpdatableAssetItemRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) ByUpdatableAssetId(updatableAssetId string)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if updatableAssetId != "" {
        urlTplParams["updatableAsset%2Did"] = updatableAssetId
    }
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderInternal instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/members{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceMembersCountRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) Count()(*WindowsUpdatesUpdatePoliciesItemAudienceMembersCountRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the updatableAsset resources that are members of a deploymentAudience.
// returns a UpdatableAssetCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-deploymentaudience-list-members?view=graph-rest-1.0
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable, error) {
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
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) MicrosoftGraphWindowsUpdatesEnrollAssets()(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesEnrollAssetsById provides operations to call the enrollAssetsById method.
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) MicrosoftGraphWindowsUpdatesEnrollAssetsById()(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesUnenrollAssets provides operations to call the unenrollAssets method.
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) MicrosoftGraphWindowsUpdatesUnenrollAssets()(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesUnenrollAssetsById provides operations to call the unenrollAssetsById method.
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) MicrosoftGraphWindowsUpdatesUnenrollAssetsById()(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to members for admin
// returns a UpdatableAssetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) Post(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
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
// ToGetRequestInformation list the updatableAsset resources that are members of a deploymentAudience.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to members for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
