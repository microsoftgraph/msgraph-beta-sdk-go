package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ElevationRequestsItemGetAllElevationRequestsRequestBuilder provides operations to call the getAllElevationRequests method.
type ElevationRequestsItemGetAllElevationRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ElevationRequestsItemGetAllElevationRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ElevationRequestsItemGetAllElevationRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewElevationRequestsItemGetAllElevationRequestsRequestBuilderInternal instantiates a new GetAllElevationRequestsRequestBuilder and sets the default values.
func NewElevationRequestsItemGetAllElevationRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ElevationRequestsItemGetAllElevationRequestsRequestBuilder) {
    m := &ElevationRequestsItemGetAllElevationRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/elevationRequests/{privilegeManagementElevationRequest%2Did}/getAllElevationRequests", pathParameters),
    }
    return m
}
// NewElevationRequestsItemGetAllElevationRequestsRequestBuilder instantiates a new GetAllElevationRequestsRequestBuilder and sets the default values.
func NewElevationRequestsItemGetAllElevationRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ElevationRequestsItemGetAllElevationRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewElevationRequestsItemGetAllElevationRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getAllElevationRequests
// Deprecated: This method is obsolete. Use PostAsGetAllElevationRequestsPostResponse instead.
func (m *ElevationRequestsItemGetAllElevationRequestsRequestBuilder) Post(ctx context.Context, requestConfiguration *ElevationRequestsItemGetAllElevationRequestsRequestBuilderPostRequestConfiguration)(ElevationRequestsItemGetAllElevationRequestsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateElevationRequestsItemGetAllElevationRequestsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ElevationRequestsItemGetAllElevationRequestsResponseable), nil
}
// PostAsGetAllElevationRequestsPostResponse invoke action getAllElevationRequests
func (m *ElevationRequestsItemGetAllElevationRequestsRequestBuilder) PostAsGetAllElevationRequestsPostResponse(ctx context.Context, requestConfiguration *ElevationRequestsItemGetAllElevationRequestsRequestBuilderPostRequestConfiguration)(ElevationRequestsItemGetAllElevationRequestsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateElevationRequestsItemGetAllElevationRequestsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ElevationRequestsItemGetAllElevationRequestsPostResponseable), nil
}
// ToPostRequestInformation invoke action getAllElevationRequests
func (m *ElevationRequestsItemGetAllElevationRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ElevationRequestsItemGetAllElevationRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ElevationRequestsItemGetAllElevationRequestsRequestBuilder) WithUrl(rawUrl string)(*ElevationRequestsItemGetAllElevationRequestsRequestBuilder) {
    return NewElevationRequestsItemGetAllElevationRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
