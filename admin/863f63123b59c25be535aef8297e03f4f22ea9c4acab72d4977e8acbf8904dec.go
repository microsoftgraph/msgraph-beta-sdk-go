// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder provides operations to call the enrollAssets method.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/members/microsoft.graph.windowsUpdates.enrollAssets", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action enrollAssets
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action enrollAssets
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
