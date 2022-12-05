package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder provides operations to manage the taskFolders property of the microsoft.graph.outlookTaskGroup entity.
type UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderGetQueryParameters the collection of task folders in the task group. Read-only. Nullable.
type UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderGetQueryParameters
}
// UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderInternal instantiates a new OutlookTaskFolderItemRequestBuilder and sets the default values.
func NewUsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) {
    m := &UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/outlook/taskGroups/{outlookTaskGroup%2Did}/taskFolders/{outlookTaskFolder%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder instantiates a new OutlookTaskFolderItemRequestBuilder and sets the default values.
func NewUsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property taskFolders for users
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of task folders in the task group. Read-only. Nullable.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property taskFolders in users
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property taskFolders for users
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of task folders in the task group. Read-only. Nullable.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable), nil
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.outlookTaskFolder entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemOutlookTaskGroupsItemTaskFoldersItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.outlookTaskFolder entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemOutlookTaskGroupsItemTaskFoldersItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property taskFolders in users
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable, requestConfiguration *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.outlookTaskFolder entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemOutlookTaskGroupsItemTaskFoldersItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.outlookTaskFolder entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemOutlookTaskGroupsItemTaskFoldersItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.outlookTaskFolder entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) Tasks()(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksRequestBuilder) {
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.outlookTaskFolder entity.
func (m *UsersItemOutlookTaskGroupsItemTaskFoldersOutlookTaskFolderItemRequestBuilder) TasksById(id string)(*UsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask%2Did"] = id
    }
    return NewUsersItemOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
