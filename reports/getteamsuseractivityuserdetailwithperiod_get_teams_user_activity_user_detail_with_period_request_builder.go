package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder provides operations to call the getTeamsUserActivityUserDetail method.
type GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal instantiates a new GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    m := &GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getTeamsUserActivityUserDetail(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder instantiates a new GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getTeamsUserActivityUserDetail
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getTeamsUserActivityUserDetail
// returns a *RequestInformation when successful
func (m *GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder when successful
func (m *GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewGetteamsuseractivityuserdetailwithperiodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
