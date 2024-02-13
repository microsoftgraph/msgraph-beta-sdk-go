package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder provides operations to manage the attachments property of the microsoft.graph.event entity.
type ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters read the properties, relationships, or raw contents of an attachment that is attached to a user event, message, Outlook task, or group post. An attachment can be one of the following types: All these types of attachments are derived from the attachment resource.
type ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters
}
// NewItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal instantiates a new ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    m := &ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/attachments/{attachment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder instantiates a new ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an attachment from a user calendar event, message, Outlook task, or post.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/attachment-delete?view=graph-rest-1.0
func (m *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties, relationships, or raw contents of an attachment that is attached to a user event, message, Outlook task, or group post. An attachment can be one of the following types: All these types of attachments are derived from the attachment resource.
// returns a Attachmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/attachment-get?view=graph-rest-1.0
func (m *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable), nil
}
// ToDeleteRequestInformation delete an attachment from a user calendar event, message, Outlook task, or post.
// returns a *RequestInformation when successful
func (m *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/attachments/{attachment%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties, relationships, or raw contents of an attachment that is attached to a user event, message, Outlook task, or group post. An attachment can be one of the following types: All these types of attachments are derived from the attachment resource.
// returns a *RequestInformation when successful
func (m *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder when successful
func (m *ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    return NewItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
