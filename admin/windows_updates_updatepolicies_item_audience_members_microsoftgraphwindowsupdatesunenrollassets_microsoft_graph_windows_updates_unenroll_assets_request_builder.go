package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder provides operations to call the unenrollAssets method.
type WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    m := &WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/members/microsoft.graph.windowsUpdates.unenrollAssets", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unenrollAssets
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsUnenrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action unenrollAssets
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsUnenrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemAudienceMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
