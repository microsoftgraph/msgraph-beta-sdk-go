package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityFileCounts method.
type PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetQueryParameters invoke function getOffice365GroupsActivityFileCounts
type PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetQueryParameters struct {
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
// PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetQueryParameters
}
// NewPrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal instantiates a new GetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder and sets the default values.
func NewPrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    m := &PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/microsoft.graph.getOffice365GroupsActivityFileCounts(period='{period}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder instantiates a new GetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder and sets the default values.
func NewPrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getOffice365GroupsActivityFileCounts
func (m *PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getOffice365GroupsActivityFileCounts
func (m *PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintReportsGetOffice365GroupsActivityFileCountsWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodResponseable), nil
}
