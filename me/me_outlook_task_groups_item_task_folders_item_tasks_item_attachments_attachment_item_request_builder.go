package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder provides operations to manage the attachments property of the microsoft.graph.outlookTask entity.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters the collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters
}
// NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderInternal instantiates a new AttachmentItemRequestBuilder and sets the default values.
func NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder) {
    m := &MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook/taskGroups/{outlookTaskGroup%2Did}/taskFolders/{outlookTaskFolder%2Did}/tasks/{outlookTask%2Did}/attachments/{attachment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder instantiates a new AttachmentItemRequestBuilder and sets the default values.
func NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property attachments for me
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property attachments for me
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, error) {
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
