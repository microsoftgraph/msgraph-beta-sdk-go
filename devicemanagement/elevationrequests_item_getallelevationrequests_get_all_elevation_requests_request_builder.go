package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder provides operations to call the getAllElevationRequests method.
type ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderInternal instantiates a new ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder and sets the default values.
func NewElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder) {
    m := &ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/elevationRequests/{privilegeManagementElevationRequest%2Did}/getAllElevationRequests", pathParameters),
    }
    return m
}
// NewElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder instantiates a new ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder and sets the default values.
func NewElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getAllElevationRequests
// Deprecated: This method is obsolete. Use PostAsGetAllElevationRequestsPostResponse instead.
// returns a ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder) Post(ctx context.Context, requestConfiguration *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderPostRequestConfiguration)(ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsResponseable), nil
}
// PostAsGetAllElevationRequestsPostResponse invoke action getAllElevationRequests
// returns a ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder) PostAsGetAllElevationRequestsPostResponse(ctx context.Context, requestConfiguration *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderPostRequestConfiguration)(ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsPostResponseable), nil
}
// ToPostRequestInformation invoke action getAllElevationRequests
// returns a *RequestInformation when successful
func (m *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder when successful
func (m *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder) WithUrl(rawUrl string)(*ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder) {
    return NewElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
