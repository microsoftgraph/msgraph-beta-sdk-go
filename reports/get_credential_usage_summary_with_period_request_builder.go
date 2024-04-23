package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetCredentialUsageSummaryWithPeriodRequestBuilder provides operations to call the getCredentialUsageSummary method.
type GetCredentialUsageSummaryWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetCredentialUsageSummaryWithPeriodRequestBuilderGetQueryParameters report the current state of how many users in your organization used the self-service password reset capabilities.
type GetCredentialUsageSummaryWithPeriodRequestBuilderGetQueryParameters struct {
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
// GetCredentialUsageSummaryWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetCredentialUsageSummaryWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetCredentialUsageSummaryWithPeriodRequestBuilderGetQueryParameters
}
// NewGetCredentialUsageSummaryWithPeriodRequestBuilderInternal instantiates a new GetCredentialUsageSummaryWithPeriodRequestBuilder and sets the default values.
func NewGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetCredentialUsageSummaryWithPeriodRequestBuilder) {
    m := &GetCredentialUsageSummaryWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getCredentialUsageSummary(period='{period}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetCredentialUsageSummaryWithPeriodRequestBuilder instantiates a new GetCredentialUsageSummaryWithPeriodRequestBuilder and sets the default values.
func NewGetCredentialUsageSummaryWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetCredentialUsageSummaryWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get report the current state of how many users in your organization used the self-service password reset capabilities.
// Deprecated: This method is obsolete. Use GetAsGetCredentialUsageSummaryWithPeriodGetResponse instead.
// returns a GetCredentialUsageSummaryWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getcredentialusagesummary?view=graph-rest-beta
func (m *GetCredentialUsageSummaryWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetCredentialUsageSummaryWithPeriodRequestBuilderGetRequestConfiguration)(GetCredentialUsageSummaryWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetCredentialUsageSummaryWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetCredentialUsageSummaryWithPeriodResponseable), nil
}
// GetAsGetCredentialUsageSummaryWithPeriodGetResponse report the current state of how many users in your organization used the self-service password reset capabilities.
// returns a GetCredentialUsageSummaryWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getcredentialusagesummary?view=graph-rest-beta
func (m *GetCredentialUsageSummaryWithPeriodRequestBuilder) GetAsGetCredentialUsageSummaryWithPeriodGetResponse(ctx context.Context, requestConfiguration *GetCredentialUsageSummaryWithPeriodRequestBuilderGetRequestConfiguration)(GetCredentialUsageSummaryWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetCredentialUsageSummaryWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetCredentialUsageSummaryWithPeriodGetResponseable), nil
}
// ToGetRequestInformation report the current state of how many users in your organization used the self-service password reset capabilities.
// returns a *RequestInformation when successful
func (m *GetCredentialUsageSummaryWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetCredentialUsageSummaryWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetCredentialUsageSummaryWithPeriodRequestBuilder when successful
func (m *GetCredentialUsageSummaryWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return NewGetCredentialUsageSummaryWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
