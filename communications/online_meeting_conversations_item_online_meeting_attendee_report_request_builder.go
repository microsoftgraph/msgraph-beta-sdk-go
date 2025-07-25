// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder provides operations to manage the media for the cloudCommunications entity.
type OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderInternal instantiates a new OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder and sets the default values.
func NewOnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) {
    m := &OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetingConversations/{onlineMeetingEngagementConversation%2Did}/onlineMeeting/attendeeReport", pathParameters),
    }
    return m
}
// NewOnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder instantiates a new OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder and sets the default values.
func NewOnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the content stream of the attendee report of a Teams live event. Read-only.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the content stream of the attendee report of a Teams live event. Read-only.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put the content stream of the attendee report of a Teams live event. Read-only.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation the content stream of the attendee report of a Teams live event. Read-only.
// returns a *RequestInformation when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the content stream of the attendee report of a Teams live event. Read-only.
// returns a *RequestInformation when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the content stream of the attendee report of a Teams live event. Read-only.
// returns a *RequestInformation when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder when successful
func (m *OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder) {
    return NewOnlineMeetingConversationsItemOnlineMeetingAttendeeReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
