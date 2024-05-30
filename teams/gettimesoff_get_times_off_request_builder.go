package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GettimesoffGetTimesOffRequestBuilder provides operations to call the getTimesOff method.
type GettimesoffGetTimesOffRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GettimesoffGetTimesOffRequestBuilderGetQueryParameters get all time off objects across all teams to which a user is a direct member.
type GettimesoffGetTimesOffRequestBuilderGetQueryParameters struct {
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
// GettimesoffGetTimesOffRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GettimesoffGetTimesOffRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GettimesoffGetTimesOffRequestBuilderGetQueryParameters
}
// NewGettimesoffGetTimesOffRequestBuilderInternal instantiates a new GettimesoffGetTimesOffRequestBuilder and sets the default values.
func NewGettimesoffGetTimesOffRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GettimesoffGetTimesOffRequestBuilder) {
    m := &GettimesoffGetTimesOffRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/getTimesOff(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGettimesoffGetTimesOffRequestBuilder instantiates a new GettimesoffGetTimesOffRequestBuilder and sets the default values.
func NewGettimesoffGetTimesOffRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GettimesoffGetTimesOffRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGettimesoffGetTimesOffRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all time off objects across all teams to which a user is a direct member.
// Deprecated: This method is obsolete. Use GetAsGetTimesOffGetResponse instead.
// returns a GettimesoffGetTimesOffResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-gettimesoff?view=graph-rest-beta
func (m *GettimesoffGetTimesOffRequestBuilder) Get(ctx context.Context, requestConfiguration *GettimesoffGetTimesOffRequestBuilderGetRequestConfiguration)(GettimesoffGetTimesOffResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGettimesoffGetTimesOffResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GettimesoffGetTimesOffResponseable), nil
}
// GetAsGetTimesOffGetResponse get all time off objects across all teams to which a user is a direct member.
// returns a GettimesoffGetTimesOffGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-gettimesoff?view=graph-rest-beta
func (m *GettimesoffGetTimesOffRequestBuilder) GetAsGetTimesOffGetResponse(ctx context.Context, requestConfiguration *GettimesoffGetTimesOffRequestBuilderGetRequestConfiguration)(GettimesoffGetTimesOffGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGettimesoffGetTimesOffGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GettimesoffGetTimesOffGetResponseable), nil
}
// ToGetRequestInformation get all time off objects across all teams to which a user is a direct member.
// returns a *RequestInformation when successful
func (m *GettimesoffGetTimesOffRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GettimesoffGetTimesOffRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GettimesoffGetTimesOffRequestBuilder when successful
func (m *GettimesoffGetTimesOffRequestBuilder) WithUrl(rawUrl string)(*GettimesoffGetTimesOffRequestBuilder) {
    return NewGettimesoffGetTimesOffRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
