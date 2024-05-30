package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder provides operations to call the getOffice365ActiveUserCounts method.
type Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetQueryParameters invoke function getOffice365ActiveUserCounts
type Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetQueryParameters struct {
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
// Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetQueryParameters
}
// NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal instantiates a new Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    m := &Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365ActiveUserCounts(period='{period}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder instantiates a new Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365ActiveUserCounts
// Deprecated: This method is obsolete. Use GetAsGetOffice365ActiveUserCountsWithPeriodGetResponse instead.
// returns a Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodResponseable), nil
}
// GetAsGetOffice365ActiveUserCountsWithPeriodGetResponse invoke function getOffice365ActiveUserCounts
// returns a Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) GetAsGetOffice365ActiveUserCountsWithPeriodGetResponse(ctx context.Context, requestConfiguration *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365ActiveUserCounts
// returns a *RequestInformation when successful
func (m *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder when successful
func (m *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
