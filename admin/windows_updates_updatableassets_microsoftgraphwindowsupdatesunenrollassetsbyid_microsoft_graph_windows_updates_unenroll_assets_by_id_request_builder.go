package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder provides operations to call the unenrollAssetsById method.
type WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderInternal instantiates a new WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder) {
    m := &WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatableAssets/microsoft.graph.windowsUpdates.unenrollAssetsById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder instantiates a new WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unenroll updatableAsset resources of the same type from update management by the deployment service. You can also use the method unenrollAssets to unenroll assets.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidUnenrollAssetsByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unenroll updatableAsset resources of the same type from update management by the deployment service. You can also use the method unenrollAssets to unenroll assets.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidUnenrollAssetsByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder when successful
func (m *WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder) {
    return NewWindowsUpdatesUpdatableassetsMicrosoftgraphwindowsupdatesunenrollassetsbyidMicrosoftGraphWindowsUpdatesUnenrollAssetsByIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
