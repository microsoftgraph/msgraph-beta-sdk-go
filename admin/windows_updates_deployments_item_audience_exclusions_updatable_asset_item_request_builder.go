package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderGetQueryParameters specifies the assets to exclude from the audience.
type WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/exclusions/{updatableAsset%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder instantiates a new WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exclusions for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get specifies the assets to exclude from the audience.
// returns a UpdatableAssetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// MicrosoftGraphWindowsUpdatesAddMembers provides operations to call the addMembers method.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) MicrosoftGraphWindowsUpdatesAddMembers()(*WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesAddMembersById provides operations to call the addMembersById method.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) MicrosoftGraphWindowsUpdatesAddMembersById()(*WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesRemoveMembers provides operations to call the removeMembers method.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesremovemembersMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) MicrosoftGraphWindowsUpdatesRemoveMembers()(*WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesremovemembersMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesremovemembersMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesRemoveMembersById provides operations to call the removeMembersById method.
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) MicrosoftGraphWindowsUpdatesRemoveMembersById()(*WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property exclusions in admin
// returns a UpdatableAssetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property exclusions for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation specifies the assets to exclude from the audience.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property exclusions in admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsUpdatableAssetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
