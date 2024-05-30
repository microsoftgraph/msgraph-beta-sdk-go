package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
type GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal instantiates a new GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    m := &GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSkypeForBusinessDeviceUsageDistributionUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder instantiates a new GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getSkypeForBusinessDeviceUsageDistributionUserCounts
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getSkypeForBusinessDeviceUsageDistributionUserCounts
// returns a *RequestInformation when successful
func (m *GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder when successful
func (m *GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewGetskypeforbusinessdeviceusagedistributionusercountswithperiodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
