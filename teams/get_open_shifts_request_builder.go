package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetOpenShiftsRequestBuilder provides operations to call the getOpenShifts method.
type GetOpenShiftsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetOpenShiftsRequestBuilderGetQueryParameters get all openShift objects across all teams a user is a direct member of.
type GetOpenShiftsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// GetOpenShiftsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetOpenShiftsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetOpenShiftsRequestBuilderGetQueryParameters
}
// NewGetOpenShiftsRequestBuilderInternal instantiates a new GetOpenShiftsRequestBuilder and sets the default values.
func NewGetOpenShiftsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOpenShiftsRequestBuilder) {
    m := &GetOpenShiftsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/getOpenShifts(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGetOpenShiftsRequestBuilder instantiates a new GetOpenShiftsRequestBuilder and sets the default values.
func NewGetOpenShiftsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOpenShiftsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOpenShiftsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all openShift objects across all teams a user is a direct member of.
// Deprecated: This method is obsolete. Use GetAsGetOpenShiftsGetResponse instead.
// returns a GetOpenShiftsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-getopenshifts?view=graph-rest-beta
func (m *GetOpenShiftsRequestBuilder) Get(ctx context.Context, requestConfiguration *GetOpenShiftsRequestBuilderGetRequestConfiguration)(GetOpenShiftsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetOpenShiftsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetOpenShiftsResponseable), nil
}
// GetAsGetOpenShiftsGetResponse get all openShift objects across all teams a user is a direct member of.
// returns a GetOpenShiftsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-getopenshifts?view=graph-rest-beta
func (m *GetOpenShiftsRequestBuilder) GetAsGetOpenShiftsGetResponse(ctx context.Context, requestConfiguration *GetOpenShiftsRequestBuilderGetRequestConfiguration)(GetOpenShiftsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetOpenShiftsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetOpenShiftsGetResponseable), nil
}
// ToGetRequestInformation get all openShift objects across all teams a user is a direct member of.
// returns a *RequestInformation when successful
func (m *GetOpenShiftsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetOpenShiftsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetOpenShiftsRequestBuilder when successful
func (m *GetOpenShiftsRequestBuilder) WithUrl(rawUrl string)(*GetOpenShiftsRequestBuilder) {
    return NewGetOpenShiftsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
