package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder provides operations to call the getSharePointSiteUsageStorage method.
type GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal instantiates a new GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder and sets the default values.
func NewGetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    m := &GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSharePointSiteUsageStorage(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder instantiates a new GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder and sets the default values.
func NewGetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getSharePointSiteUsageStorage
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getSharePointSiteUsageStorage
// returns a *RequestInformation when successful
func (m *GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder when successful
func (m *GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return NewGetsharepointsiteusagestoragewithperiodGetSharePointSiteUsageStorageWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
