package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters specifies the assets to include in the audience.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/members/{updatableAsset%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder instantiates a new WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property members for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get specifies the assets to include in the audience.
// returns a UpdatableAssetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
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
// returns a *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) MicrosoftGraphWindowsUpdatesAddMembers()(*WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesAddMembersById provides operations to call the addMembersById method.
// returns a *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) MicrosoftGraphWindowsUpdatesAddMembersById()(*WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesRemoveMembers provides operations to call the removeMembers method.
// returns a *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) MicrosoftGraphWindowsUpdatesRemoveMembers()(*WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesRemoveMembersById provides operations to call the removeMembersById method.
// returns a *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) MicrosoftGraphWindowsUpdatesRemoveMembersById()(*WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property members in admin
// returns a UpdatableAssetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
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
// ToDeleteRequestInformation delete navigation property members for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation specifies the assets to include in the audience.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property members in admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
