package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityFileCounts method.
type Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetQueryParameters invoke function getOffice365GroupsActivityFileCounts
type Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetQueryParameters struct {
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
// Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetQueryParameters
}
// NewGetoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal instantiates a new Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    m := &Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityFileCounts(period='{period}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder instantiates a new Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityFileCounts
// Deprecated: This method is obsolete. Use GetAsGetOffice365GroupsActivityFileCountsWithPeriodGetResponse instead.
// returns a Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration)(Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodResponseable), nil
}
// GetAsGetOffice365GroupsActivityFileCountsWithPeriodGetResponse invoke function getOffice365GroupsActivityFileCounts
// returns a Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) GetAsGetOffice365GroupsActivityFileCountsWithPeriodGetResponse(ctx context.Context, requestConfiguration *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration)(Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365GroupsActivityFileCounts
// returns a *RequestInformation when successful
func (m *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder when successful
func (m *Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return NewGetoffice365groupsactivityfilecountswithperiodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
