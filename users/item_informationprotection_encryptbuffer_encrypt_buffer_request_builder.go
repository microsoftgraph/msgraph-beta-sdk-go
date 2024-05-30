package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder provides operations to call the encryptBuffer method.
type ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationprotectionEncryptbufferEncryptBufferRequestBuilderInternal instantiates a new ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder and sets the default values.
func NewItemInformationprotectionEncryptbufferEncryptBufferRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder) {
    m := &ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/informationProtection/encryptBuffer", pathParameters),
    }
    return m
}
// NewItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder instantiates a new ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder and sets the default values.
func NewItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationprotectionEncryptbufferEncryptBufferRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action encryptBuffer
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a BufferEncryptionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder) Post(ctx context.Context, body ItemInformationprotectionEncryptbufferEncryptBufferPostRequestBodyable, requestConfiguration *ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BufferEncryptionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBufferEncryptionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BufferEncryptionResultable), nil
}
// ToPostRequestInformation invoke action encryptBuffer
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemInformationprotectionEncryptbufferEncryptBufferPostRequestBodyable, requestConfiguration *ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder when successful
func (m *ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder) WithUrl(rawUrl string)(*ItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder) {
    return NewItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
