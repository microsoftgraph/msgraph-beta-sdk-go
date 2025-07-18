// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingConversationsItemOnlineMeetingRequestBuilder provides operations to manage the onlineMeeting property of the microsoft.graph.onlineMeetingEngagementConversation entity.
type OnlineMeetingConversationsItemOnlineMeetingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingConversationsItemOnlineMeetingRequestBuilderGetQueryParameters the online meeting associated with the conversation.
type OnlineMeetingConversationsItemOnlineMeetingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnlineMeetingConversationsItemOnlineMeetingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingConversationsItemOnlineMeetingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlineMeetingConversationsItemOnlineMeetingRequestBuilderGetQueryParameters
}
// AlternativeRecording provides operations to manage the media for the cloudCommunications entity.
// returns a *OnlineMeetingConversationsItemOnlineMeetingAlternativeRecordingRequestBuilder when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) AlternativeRecording()(*OnlineMeetingConversationsItemOnlineMeetingAlternativeRecordingRequestBuilder) {
    return NewOnlineMeetingConversationsItemOnlineMeetingAlternativeRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttendeeReport provides operations to manage the media for the cloudCommunications entity.
// returns a *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) AttendeeReport()(*OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) {
    return NewOnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BroadcastRecording provides operations to manage the media for the cloudCommunications entity.
// returns a *OnlineMeetingConversationsItemOnlineMeetingBroadcastRecordingRequestBuilder when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) BroadcastRecording()(*OnlineMeetingConversationsItemOnlineMeetingBroadcastRecordingRequestBuilder) {
    return NewOnlineMeetingConversationsItemOnlineMeetingBroadcastRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewOnlineMeetingConversationsItemOnlineMeetingRequestBuilderInternal instantiates a new OnlineMeetingConversationsItemOnlineMeetingRequestBuilder and sets the default values.
func NewOnlineMeetingConversationsItemOnlineMeetingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) {
    m := &OnlineMeetingConversationsItemOnlineMeetingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetingConversations/{onlineMeetingEngagementConversation%2Did}/onlineMeeting{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOnlineMeetingConversationsItemOnlineMeetingRequestBuilder instantiates a new OnlineMeetingConversationsItemOnlineMeetingRequestBuilder and sets the default values.
func NewOnlineMeetingConversationsItemOnlineMeetingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingConversationsItemOnlineMeetingRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the online meeting associated with the conversation.
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingConversationsItemOnlineMeetingRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// Recording provides operations to manage the media for the cloudCommunications entity.
// returns a *OnlineMeetingConversationsItemOnlineMeetingRecordingRequestBuilder when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) Recording()(*OnlineMeetingConversationsItemOnlineMeetingRecordingRequestBuilder) {
    return NewOnlineMeetingConversationsItemOnlineMeetingRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the online meeting associated with the conversation.
// returns a *RequestInformation when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingConversationsItemOnlineMeetingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlineMeetingConversationsItemOnlineMeetingRequestBuilder when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingConversationsItemOnlineMeetingRequestBuilder) {
    return NewOnlineMeetingConversationsItemOnlineMeetingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
