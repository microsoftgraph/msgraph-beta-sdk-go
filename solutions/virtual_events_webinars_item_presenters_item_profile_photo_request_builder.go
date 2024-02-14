package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder provides operations to manage the media for the solutionsRoot entity.
type VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderInternal instantiates a new VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder) {
    m := &VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/presenters/{virtualEventPresenter%2Did}/profilePhoto", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder instantiates a new VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get profilePhoto for the navigation property presenters from solutions
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put update profilePhoto for the navigation property presenters in solutions
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderPutRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get profilePhoto for the navigation property presenters from solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update profilePhoto for the navigation property presenters in solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder when successful
func (m *VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder) {
    return NewVirtualEventsWebinarsItemPresentersItemProfilePhotoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
