package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder provides operations to call the addMembersById method.
type WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal instantiates a new MicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/exclusions/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.addMembersById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder instantiates a new MicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add members of the same type to an updatableAssetGroup. You can also use the method addMembers to add members. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-addmembersbyid?view=graph-rest-1.0
func (m *WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration)(error) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdAddMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
