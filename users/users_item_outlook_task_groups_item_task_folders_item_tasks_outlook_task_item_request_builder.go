package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.outlookTaskFolder entity.
type UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters the tasks in this task folder. Read-only. Nullable.
type UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters
}
// UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.outlookTask entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Attachments()(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsRequestBuilder) {
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.outlookTask entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) AttachmentsById(id string)(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Complete provides operations to call the complete method.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Complete()(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) {
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal instantiates a new OutlookTaskItemRequestBuilder and sets the default values.
func NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) {
    m := &UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/outlook/taskGroups/{outlookTaskGroup%2Did}/taskFolders/{outlookTaskFolder%2Did}/tasks/{outlookTask%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder instantiates a new OutlookTaskItemRequestBuilder and sets the default values.
func NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property tasks for users
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the tasks in this task folder. Read-only. Nullable.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property tasks in users
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property tasks for users
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the tasks in this task folder. Read-only. Nullable.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable), nil
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.outlookTask entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.outlookTask entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property tasks in users
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.outlookTask entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.outlookTask entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
