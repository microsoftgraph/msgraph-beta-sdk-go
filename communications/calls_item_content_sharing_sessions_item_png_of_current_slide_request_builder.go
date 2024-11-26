package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder provides operations to manage the media for the cloudCommunications entity.
type CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderInternal instantiates a new CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder and sets the default values.
func NewCallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) {
    m := &CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/contentSharingSessions/{contentSharingSession%2Did}/pngOfCurrentSlide", pathParameters),
    }
    return m
}
// NewCallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder instantiates a new CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder and sets the default values.
func NewCallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete pngOfCurrentSlide for the navigation property contentSharingSessions in communications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) Delete(ctx context.Context, requestConfiguration *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get get pngOfCurrentSlide for the navigation property contentSharingSessions from communications
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) Get(ctx context.Context, requestConfiguration *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put update pngOfCurrentSlide for the navigation property contentSharingSessions in communications
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToDeleteRequestInformation delete pngOfCurrentSlide for the navigation property contentSharingSessions in communications
// returns a *RequestInformation when successful
func (m *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get pngOfCurrentSlide for the navigation property contentSharingSessions from communications
// returns a *RequestInformation when successful
func (m *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update pngOfCurrentSlide for the navigation property contentSharingSessions in communications
// returns a *RequestInformation when successful
func (m *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder when successful
func (m *CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) WithUrl(rawUrl string)(*CallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder) {
    return NewCallsItemContentSharingSessionsItemPngOfCurrentSlideRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
