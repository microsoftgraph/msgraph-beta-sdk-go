package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityCounts method.
type Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderGetQueryParameters invoke function getOffice365GroupsActivityCounts
type Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderGetQueryParameters struct {
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
// Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderGetQueryParameters
}
// NewGetoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal instantiates a new Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    m := &Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityCounts(period='{period}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder instantiates a new Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityCounts
// Deprecated: This method is obsolete. Use GetAsGetOffice365GroupsActivityCountsWithPeriodGetResponse instead.
// returns a Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)(Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodResponseable), nil
}
// GetAsGetOffice365GroupsActivityCountsWithPeriodGetResponse invoke function getOffice365GroupsActivityCounts
// returns a Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) GetAsGetOffice365GroupsActivityCountsWithPeriodGetResponse(ctx context.Context, requestConfiguration *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)(Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365GroupsActivityCounts
// returns a *RequestInformation when successful
func (m *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder when successful
func (m *Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return NewGetoffice365groupsactivitycountswithperiodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
