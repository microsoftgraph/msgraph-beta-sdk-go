package communications

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder provides operations to call the getSmsLog method.
type CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters invoke function getSmsLog
type CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters struct {
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
// CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters
}
// NewCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderInternal instantiates a new MicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) {
    m := &CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/microsoft.graph.callRecords.getSmsLog(fromDateTime={fromDateTime},toDateTime={toDateTime}){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    if fromDateTime != nil {
        m.BaseRequestBuilder.PathParameters["fromDateTime"] = (*fromDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if toDateTime != nil {
        m.BaseRequestBuilder.PathParameters["toDateTime"] = (*toDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder instantiates a new MicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getSmsLog
// Deprecated: This method is obsolete. Use GetAsGetSmsLogWithFromDateTimeWithToDateTimeGetResponse instead.
func (m *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeGetSmsLogWithFromDateTimeWithToDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeGetSmsLogWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeGetSmsLogWithFromDateTimeWithToDateTimeResponseable), nil
}
// GetAsGetSmsLogWithFromDateTimeWithToDateTimeGetResponse invoke function getSmsLog
func (m *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) GetAsGetSmsLogWithFromDateTimeWithToDateTimeGetResponse(ctx context.Context, requestConfiguration *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeGetSmsLogWithFromDateTimeWithToDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeGetSmsLogWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeGetSmsLogWithFromDateTimeWithToDateTimeGetResponseable), nil
}
// ToGetRequestInformation invoke function getSmsLog
func (m *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) WithUrl(rawUrl string)(*CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
