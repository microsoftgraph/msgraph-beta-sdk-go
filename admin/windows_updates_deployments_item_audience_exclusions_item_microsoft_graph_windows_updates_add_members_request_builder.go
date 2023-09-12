package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder provides operations to call the addMembers method.
type WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal instantiates a new MicrosoftGraphWindowsUpdatesAddMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/exclusions/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.addMembers", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder instantiates a new MicrosoftGraphWindowsUpdatesAddMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add members to an updatableAssetGroup. You can add azureADDevice resources as members, but may not add updatableAssetGroup resources as members. Adding an Azure AD device as a member of an updatable asset group automatically creates an azureADDevice object, if it does not already exist. You can also use the method addMembersById to add members.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-addmembers?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation add members to an updatableAssetGroup. You can add azureADDevice resources as members, but may not add updatableAssetGroup resources as members. Adding an Azure AD device as a member of an updatable asset group automatically creates an azureADDevice object, if it does not already exist. You can also use the method addMembersById to add members.
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
