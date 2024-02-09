package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder provides operations to call the removeMembers method.
type WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder) {
    m := &WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/exclusions/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.removeMembers", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder instantiates a new WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove members from an updatableAssetGroup. You can also use the method removeMembersById to remove members.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-removemembers?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRemoveMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove members from an updatableAssetGroup. You can also use the method removeMembersById to remove members.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRemoveMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder when successful
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesRemoveMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
