package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityDetail method.
type ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetQueryParameters invoke function getOffice365GroupsActivityDetail
type ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetQueryParameters
}
// NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal instantiates a new GetOffice365GroupsActivityDetailWithPeriodRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    m := &ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/microsoft.graph.getOffice365GroupsActivityDetail(period='{period}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder instantiates a new GetOffice365GroupsActivityDetailWithPeriodRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityDetail
func (m *ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration)(ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365GroupsActivityDetail
func (m *ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
