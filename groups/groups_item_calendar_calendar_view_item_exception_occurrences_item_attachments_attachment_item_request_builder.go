package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder provides operations to manage the attachments property of the microsoft.graph.event entity.
type GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters the collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
type GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters
}
// NewGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal instantiates a new AttachmentItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    m := &GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/attachments/{attachment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder instantiates a new AttachmentItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property attachments for groups
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property attachments for groups
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable), nil
}
