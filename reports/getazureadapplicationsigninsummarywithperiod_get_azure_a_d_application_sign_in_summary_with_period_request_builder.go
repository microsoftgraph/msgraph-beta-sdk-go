package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder provides operations to call the getAzureADApplicationSignInSummary method.
type GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetQueryParameters retrieve applicationSigninSummary objects within either the last seven or 30 days.
type GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetQueryParameters struct {
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
// GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetQueryParameters
}
// NewGetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal instantiates a new GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder and sets the default values.
func NewGetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    m := &GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getAzureADApplicationSignInSummary(period='{period}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder instantiates a new GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder and sets the default values.
func NewGetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get retrieve applicationSigninSummary objects within either the last seven or 30 days.
// Deprecated: This method is obsolete. Use GetAsGetAzureADApplicationSignInSummaryWithPeriodGetResponse instead.
// returns a GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getazureadapplicationsigninsummary?view=graph-rest-beta
func (m *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration)(GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodResponseable), nil
}
// GetAsGetAzureADApplicationSignInSummaryWithPeriodGetResponse retrieve applicationSigninSummary objects within either the last seven or 30 days.
// returns a GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getazureadapplicationsigninsummary?view=graph-rest-beta
func (m *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) GetAsGetAzureADApplicationSignInSummaryWithPeriodGetResponse(ctx context.Context, requestConfiguration *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration)(GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodGetResponseable), nil
}
// ToGetRequestInformation retrieve applicationSigninSummary objects within either the last seven or 30 days.
// returns a *RequestInformation when successful
func (m *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder when successful
func (m *GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return NewGetazureadapplicationsigninsummarywithperiodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
