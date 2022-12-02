package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the getUserArchivedPrintJobs method.
type ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters invoke function getUserArchivedPrintJobs
type ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getUserArchivedPrintJobs(userId='{userId}',startDateTime={startDateTime},endDateTime={endDateTime}){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if endDateTime != nil {
        urlTplParams["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        urlTplParams["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if userId != nil {
        urlTplParams["userId"] = *userId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// CreateGetRequestInformation invoke function getUserArchivedPrintJobs
func (m *ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getUserArchivedPrintJobs
func (m *ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseable), nil
}
