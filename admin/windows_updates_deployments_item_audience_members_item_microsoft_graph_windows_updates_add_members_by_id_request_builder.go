package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder provides operations to call the addMembersById method.
type WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal instantiates a new MicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/members/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.addMembersById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder instantiates a new MicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add members of the same type to an updatableAssetGroup. You can also use the method addMembers to add members. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-addmembersbyid?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation add members of the same type to an updatableAssetGroup. You can also use the method addMembers to add members. This API is available in the following national cloud deployments.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
