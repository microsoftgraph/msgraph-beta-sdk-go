package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.outlookTaskFolder entity.
type ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters the tasks in this task folder. Read-only. Nullable.
type ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters
}
// ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.outlookTask entity.
// returns a *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsRequestBuilder when successful
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Attachments()(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsRequestBuilder) {
    return NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Complete provides operations to call the complete method.
// returns a *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder when successful
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Complete()(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) {
    return NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal instantiates a new ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder and sets the default values.
func NewItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) {
    m := &ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook/taskGroups/{outlookTaskGroup%2Did}/taskFolders/{outlookTaskFolder%2Did}/tasks/{outlookTask%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder instantiates a new ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder and sets the default values.
func NewItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tasks for users
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the tasks in this task folder. Read-only. Nullable.
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
// returns a OutlookTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable), nil
}
// Patch update the navigation property tasks in users
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
// returns a OutlookTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable), nil
}
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemPermanentDeleteRequestBuilder when successful
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) PermanentDelete()(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemPermanentDeleteRequestBuilder) {
    return NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property tasks for users
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
// returns a *RequestInformation when successful
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the tasks in this task folder. Read-only. Nullable.
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
// returns a *RequestInformation when successful
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tasks in users
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
// returns a *RequestInformation when successful
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
// returns a *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder when successful
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) WithUrl(rawUrl string)(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) {
    return NewItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
