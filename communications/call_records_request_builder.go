package communications

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
)

// CallRecordsRequestBuilder provides operations to manage the callRecords property of the microsoft.graph.cloudCommunications entity.
type CallRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallRecordsRequestBuilderGetQueryParameters get the list of callRecord objects and their properties. The results can be optionally filtered using the $filter query parameter on the startDateTime and participant id properties. Note that the listed call records don't include expandable relationships such as sessions and participants_v2. You can expand these relationships using Get callRecord for a specific record.
type CallRecordsRequestBuilderGetQueryParameters struct {
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
// CallRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallRecordsRequestBuilderGetQueryParameters
}
// CallRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCallRecordId provides operations to manage the callRecords property of the microsoft.graph.cloudCommunications entity.
// returns a *CallRecordsCallRecordItemRequestBuilder when successful
func (m *CallRecordsRequestBuilder) ByCallRecordId(callRecordId string)(*CallRecordsCallRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if callRecordId != "" {
        urlTplParams["callRecord%2Did"] = callRecordId
    }
    return NewCallRecordsCallRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCallRecordsRequestBuilderInternal instantiates a new CallRecordsRequestBuilder and sets the default values.
func NewCallRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordsRequestBuilder) {
    m := &CallRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCallRecordsRequestBuilder instantiates a new CallRecordsRequestBuilder and sets the default values.
func NewCallRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CallRecordsCountRequestBuilder when successful
func (m *CallRecordsRequestBuilder) Count()(*CallRecordsCountRequestBuilder) {
    return NewCallRecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the list of callRecord objects and their properties. The results can be optionally filtered using the $filter query parameter on the startDateTime and participant id properties. Note that the listed call records don't include expandable relationships such as sessions and participants_v2. You can expand these relationships using Get callRecord for a specific record.
// returns a CallRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/callrecords-cloudcommunications-list-callrecords?view=graph-rest-1.0
func (m *CallRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *CallRecordsRequestBuilderGetRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateCallRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordCollectionResponseable), nil
}
// MicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTime provides operations to call the getDirectRoutingCalls method.
// returns a *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallRecordsRequestBuilder) MicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTime(fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fromDateTime, toDateTime)
}
// MicrosoftGraphCallRecordsGetPstnBlockedUsersLogWithFromDateTimeWithToDateTime provides operations to call the getPstnBlockedUsersLog method.
// returns a *CallRecordsMicrosoftGraphCallRecordsGetPstnBlockedUsersLogWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallRecordsRequestBuilder) MicrosoftGraphCallRecordsGetPstnBlockedUsersLogWithFromDateTimeWithToDateTime(fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallRecordsMicrosoftGraphCallRecordsGetPstnBlockedUsersLogWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallRecordsMicrosoftGraphCallRecordsGetPstnBlockedUsersLogWithFromDateTimeWithToDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fromDateTime, toDateTime)
}
// MicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTime provides operations to call the getPstnCalls method.
// returns a *CallRecordsMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallRecordsRequestBuilder) MicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTime(fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallRecordsMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallRecordsMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fromDateTime, toDateTime)
}
// MicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTime provides operations to call the getPstnOnlineMeetingDialoutReport method.
// returns a *CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallRecordsRequestBuilder) MicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTime(fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fromDateTime, toDateTime)
}
// MicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTime provides operations to call the getSmsLog method.
// returns a *CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallRecordsRequestBuilder) MicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTime(fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallRecordsMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fromDateTime, toDateTime)
}
// Post create new navigation property to callRecords for communications
// returns a CallRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallRecordsRequestBuilder) Post(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable, requestConfiguration *CallRecordsRequestBuilderPostRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateCallRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable), nil
}
// ToGetRequestInformation get the list of callRecord objects and their properties. The results can be optionally filtered using the $filter query parameter on the startDateTime and participant id properties. Note that the listed call records don't include expandable relationships such as sessions and participants_v2. You can expand these relationships using Get callRecord for a specific record.
// returns a *RequestInformation when successful
func (m *CallRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to callRecords for communications
// returns a *RequestInformation when successful
func (m *CallRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable, requestConfiguration *CallRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/communications/callRecords", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CallRecordsRequestBuilder when successful
func (m *CallRecordsRequestBuilder) WithUrl(rawUrl string)(*CallRecordsRequestBuilder) {
    return NewCallRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
