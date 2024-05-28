package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder provides operations to call the generateDownloadUrl method.
type MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderInternal instantiates a new MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder and sets the default values.
func NewMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder) {
    m := &MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelServerLogCollectionResponses/{microsoftTunnelServerLogCollectionResponse%2Did}/generateDownloadUrl", pathParameters),
    }
    return m
}
// NewMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder instantiates a new MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder and sets the default values.
func NewMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action generateDownloadUrl
// Deprecated: This method is obsolete. Use PostAsGenerateDownloadUrlPostResponse instead.
// returns a MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder) Post(ctx context.Context, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderPostRequestConfiguration)(MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlResponseable), nil
}
// PostAsGenerateDownloadUrlPostResponse invoke action generateDownloadUrl
// returns a MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder) PostAsGenerateDownloadUrlPostResponse(ctx context.Context, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderPostRequestConfiguration)(MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlPostResponseable), nil
}
// ToPostRequestInformation invoke action generateDownloadUrl
// returns a *RequestInformation when successful
func (m *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder when successful
func (m *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder) {
    return NewMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
