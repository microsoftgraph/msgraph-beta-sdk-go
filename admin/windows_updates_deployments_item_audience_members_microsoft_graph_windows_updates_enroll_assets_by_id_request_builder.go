package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder provides operations to call the enrollAssetsById method.
type WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderInternal instantiates a new MicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/members/microsoft.graph.windowsUpdates.enrollAssetsById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder instantiates a new MicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action enrollAssetsById
func (m *WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action enrollAssetsById
func (m *WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
