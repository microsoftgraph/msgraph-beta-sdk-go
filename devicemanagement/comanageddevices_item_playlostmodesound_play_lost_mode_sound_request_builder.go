package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder provides operations to call the playLostModeSound method.
type ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderInternal instantiates a new ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder and sets the default values.
func NewComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) {
    m := &ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/playLostModeSound", pathParameters),
    }
    return m
}
// NewComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder instantiates a new ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder and sets the default values.
func NewComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderInternal(urlParams, requestAdapter)
}
// Post play lost mode sound
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) Post(ctx context.Context, body ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundPostRequestBodyable, requestConfiguration *ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation play lost mode sound
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundPostRequestBodyable, requestConfiguration *ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder when successful
func (m *ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) {
    return NewComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
