package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingsGetAllTranscriptsRequestBuilder provides operations to call the getAllTranscripts method.
type OnlineMeetingsGetAllTranscriptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingsGetAllTranscriptsRequestBuilderGetQueryParameters get all transcripts from scheduled onlineMeeting instances for which the specified user is the organizer. This API currently doesn't support getting call transcripts from channel meetings. You can apply the delta function on getAllTranscripts to synchronize and get callTranscript resources as they are added for onlineMeeting instances organized by the specified user. Delta query supports both full synchronization that gets all the transcripts for online meetings organized by the user, and incremental synchronization that gets transcripts that have been added since the last synchronization. Typically, you would do an initial full synchronization, and then get incremental changes to that transcript view periodically. Find more information in the delta query documentation. For additional examples, see callTranscript: delta. To learn more about using the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
type OnlineMeetingsGetAllTranscriptsRequestBuilderGetQueryParameters struct {
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
// OnlineMeetingsGetAllTranscriptsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsGetAllTranscriptsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlineMeetingsGetAllTranscriptsRequestBuilderGetQueryParameters
}
// NewOnlineMeetingsGetAllTranscriptsRequestBuilderInternal instantiates a new OnlineMeetingsGetAllTranscriptsRequestBuilder and sets the default values.
func NewOnlineMeetingsGetAllTranscriptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsGetAllTranscriptsRequestBuilder) {
    m := &OnlineMeetingsGetAllTranscriptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/getAllTranscripts(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOnlineMeetingsGetAllTranscriptsRequestBuilder instantiates a new OnlineMeetingsGetAllTranscriptsRequestBuilder and sets the default values.
func NewOnlineMeetingsGetAllTranscriptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsGetAllTranscriptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsGetAllTranscriptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all transcripts from scheduled onlineMeeting instances for which the specified user is the organizer. This API currently doesn't support getting call transcripts from channel meetings. You can apply the delta function on getAllTranscripts to synchronize and get callTranscript resources as they are added for onlineMeeting instances organized by the specified user. Delta query supports both full synchronization that gets all the transcripts for online meetings organized by the user, and incremental synchronization that gets transcripts that have been added since the last synchronization. Typically, you would do an initial full synchronization, and then get incremental changes to that transcript view periodically. Find more information in the delta query documentation. For additional examples, see callTranscript: delta. To learn more about using the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// Deprecated: This method is obsolete. Use GetAsGetAllTranscriptsGetResponse instead.
// returns a OnlineMeetingsGetAllTranscriptsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onlinemeeting-getalltranscripts?view=graph-rest-beta
func (m *OnlineMeetingsGetAllTranscriptsRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingsGetAllTranscriptsRequestBuilderGetRequestConfiguration)(OnlineMeetingsGetAllTranscriptsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOnlineMeetingsGetAllTranscriptsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OnlineMeetingsGetAllTranscriptsResponseable), nil
}
// GetAsGetAllTranscriptsGetResponse get all transcripts from scheduled onlineMeeting instances for which the specified user is the organizer. This API currently doesn't support getting call transcripts from channel meetings. You can apply the delta function on getAllTranscripts to synchronize and get callTranscript resources as they are added for onlineMeeting instances organized by the specified user. Delta query supports both full synchronization that gets all the transcripts for online meetings organized by the user, and incremental synchronization that gets transcripts that have been added since the last synchronization. Typically, you would do an initial full synchronization, and then get incremental changes to that transcript view periodically. Find more information in the delta query documentation. For additional examples, see callTranscript: delta. To learn more about using the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// returns a OnlineMeetingsGetAllTranscriptsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onlinemeeting-getalltranscripts?view=graph-rest-beta
func (m *OnlineMeetingsGetAllTranscriptsRequestBuilder) GetAsGetAllTranscriptsGetResponse(ctx context.Context, requestConfiguration *OnlineMeetingsGetAllTranscriptsRequestBuilderGetRequestConfiguration)(OnlineMeetingsGetAllTranscriptsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOnlineMeetingsGetAllTranscriptsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OnlineMeetingsGetAllTranscriptsGetResponseable), nil
}
// ToGetRequestInformation get all transcripts from scheduled onlineMeeting instances for which the specified user is the organizer. This API currently doesn't support getting call transcripts from channel meetings. You can apply the delta function on getAllTranscripts to synchronize and get callTranscript resources as they are added for onlineMeeting instances organized by the specified user. Delta query supports both full synchronization that gets all the transcripts for online meetings organized by the user, and incremental synchronization that gets transcripts that have been added since the last synchronization. Typically, you would do an initial full synchronization, and then get incremental changes to that transcript view periodically. Find more information in the delta query documentation. For additional examples, see callTranscript: delta. To learn more about using the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// returns a *RequestInformation when successful
func (m *OnlineMeetingsGetAllTranscriptsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsGetAllTranscriptsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlineMeetingsGetAllTranscriptsRequestBuilder when successful
func (m *OnlineMeetingsGetAllTranscriptsRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingsGetAllTranscriptsRequestBuilder) {
    return NewOnlineMeetingsGetAllTranscriptsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
