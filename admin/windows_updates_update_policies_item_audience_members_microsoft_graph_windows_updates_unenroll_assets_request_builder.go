package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder provides operations to call the unenrollAssets method.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal instantiates a new MicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/members/microsoft.graph.windowsUpdates.unenrollAssets", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder instantiates a new MicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unenrollAssets
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsUnenrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action unenrollAssets
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsUnenrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
