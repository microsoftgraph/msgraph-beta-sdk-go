package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder provides operations to call the removeMembersById method.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderInternal instantiates a new MicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/members/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.removeMembersById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder instantiates a new MicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove members of the same type from an updatableAssetGroup. You can also use the method removeMembers to remove members.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-removemembersbyid?view=graph-rest-1.0
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRemoveMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove members of the same type from an updatableAssetGroup. You can also use the method removeMembers to remove members.
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRemoveMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersItemMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
