package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder provides operations to call the getOffice365ActiveUserDetail method.
type Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetQueryParameters invoke function getOffice365ActiveUserDetail
type Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetQueryParameters struct {
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
// Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetQueryParameters
}
// NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal instantiates a new Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    m := &Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365ActiveUserDetail(period='{period}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder instantiates a new Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365ActiveUserDetail
// Deprecated: This method is obsolete. Use GetAsGetOffice365ActiveUserDetailWithPeriodGetResponse instead.
// returns a Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodResponseable), nil
}
// GetAsGetOffice365ActiveUserDetailWithPeriodGetResponse invoke function getOffice365ActiveUserDetail
// returns a Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) GetAsGetOffice365ActiveUserDetailWithPeriodGetResponse(ctx context.Context, requestConfiguration *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365ActiveUserDetail
// returns a *RequestInformation when successful
func (m *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder when successful
func (m *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
