package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder provides operations to call the enrollAssets method.
type WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal instantiates a new WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    m := &WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatableAssets/microsoft.graph.windowsUpdates.enrollAssets", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder instantiates a new WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enroll updatableAsset resources in update management by Windows Autopatch. You can enroll an azureADDevice resource in update management, but may not enroll an updatableAssetGroup in update management. Enrolling a Microsoft Entra device in update management automatically creates an azureADDevice object if it does not already exist. You can also use the method enrollAssetsById to enroll assets.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation enroll updatableAsset resources in update management by Windows Autopatch. You can enroll an azureADDevice resource in update management, but may not enroll an updatableAssetGroup in update management. Enrolling a Microsoft Entra device in update management automatically creates an azureADDevice object if it does not already exist. You can also use the method enrollAssetsById to enroll assets.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder) {
    return NewWindowsUpdatesUpdatableAssetsMicrosoftGraphWindowsUpdatesEnrollAssetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
