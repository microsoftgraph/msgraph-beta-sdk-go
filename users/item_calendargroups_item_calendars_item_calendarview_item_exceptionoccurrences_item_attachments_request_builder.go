package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder provides operations to manage the attachments property of the microsoft.graph.event entity.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderGetQueryParameters the collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderGetQueryParameters struct {
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
// ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderGetQueryParameters
}
// ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAttachmentId provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsAttachmentItemRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) ByAttachmentId(attachmentId string)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attachmentId != "" {
        urlTplParams["attachment%2Did"] = attachmentId
    }
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/attachments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsCountRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) Count()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsCountRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) CreateUploadSession()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
// returns a AttachmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentCollectionResponseable), nil
}
// Post create new navigation property to attachments for users
// returns a Attachmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attachments for users
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemExceptionoccurrencesItemAttachmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
