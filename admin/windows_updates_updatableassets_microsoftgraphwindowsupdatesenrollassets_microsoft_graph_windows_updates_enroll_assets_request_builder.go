package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder provides operations to call the enrollAssets method.
type WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal instantiates a new WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    m := &WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatableAssets/microsoft.graph.windowsUpdates.enrollAssets", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder instantiates a new WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enroll updatableAsset resources in update management by the deployment service. You can enroll an azureADDevice resource in update management, but may not enroll an updatableAssetGroup in update management. Enrolling a Microsoft Entra device in update management automatically creates an azureADDevice object if it does not already exist. You can also use the method enrollAssetsById to enroll assets.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsEnrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation enroll updatableAsset resources in update management by the deployment service. You can enroll an azureADDevice resource in update management, but may not enroll an updatableAssetGroup in update management. Enrolling a Microsoft Entra device in update management automatically creates an azureADDevice object if it does not already exist. You can also use the method enrollAssetsById to enroll assets.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsEnrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    return NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesenrollassetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
