package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemOutlookRequestBuilder provides operations to manage the outlook property of the microsoft.graph.user entity.
type UsersItemOutlookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemOutlookRequestBuilderGetQueryParameters selective Outlook services available to the user. Read-only. Nullable.
type UsersItemOutlookRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemOutlookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemOutlookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemOutlookRequestBuilderGetQueryParameters
}
// NewUsersItemOutlookRequestBuilderInternal instantiates a new OutlookRequestBuilder and sets the default values.
func NewUsersItemOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemOutlookRequestBuilder) {
    m := &UsersItemOutlookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/outlook{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemOutlookRequestBuilder instantiates a new OutlookRequestBuilder and sets the default values.
func NewUsersItemOutlookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemOutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation selective Outlook services available to the user. Read-only. Nullable.
func (m *UsersItemOutlookRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemOutlookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get selective Outlook services available to the user. Read-only. Nullable.
func (m *UsersItemOutlookRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemOutlookRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable), nil
}
// MasterCategories provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
func (m *UsersItemOutlookRequestBuilder) MasterCategories()(*UsersItemOutlookMasterCategoriesRequestBuilder) {
    return NewUsersItemOutlookMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MasterCategoriesById provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
func (m *UsersItemOutlookRequestBuilder) MasterCategoriesById(id string)(*UsersItemOutlookMasterCategoriesOutlookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory%2Did"] = id
    }
    return NewUsersItemOutlookMasterCategoriesOutlookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SupportedLanguages provides operations to call the supportedLanguages method.
func (m *UsersItemOutlookRequestBuilder) SupportedLanguages()(*UsersItemOutlookSupportedLanguagesRequestBuilder) {
    return NewUsersItemOutlookSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZones provides operations to call the supportedTimeZones method.
func (m *UsersItemOutlookRequestBuilder) SupportedTimeZones()(*UsersItemOutlookSupportedTimeZonesRequestBuilder) {
    return NewUsersItemOutlookSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZonesWithTimeZoneStandard provides operations to call the supportedTimeZones method.
func (m *UsersItemOutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*UsersItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return NewUsersItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard);
}
// TaskFolders provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
func (m *UsersItemOutlookRequestBuilder) TaskFolders()(*UsersItemOutlookTaskFoldersRequestBuilder) {
    return NewUsersItemOutlookTaskFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskFoldersById provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
func (m *UsersItemOutlookRequestBuilder) TaskFoldersById(id string)(*UsersItemOutlookTaskFoldersOutlookTaskFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskFolder%2Did"] = id
    }
    return NewUsersItemOutlookTaskFoldersOutlookTaskFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskGroups provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
func (m *UsersItemOutlookRequestBuilder) TaskGroups()(*UsersItemOutlookTaskGroupsRequestBuilder) {
    return NewUsersItemOutlookTaskGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskGroupsById provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
func (m *UsersItemOutlookRequestBuilder) TaskGroupsById(id string)(*UsersItemOutlookTaskGroupsOutlookTaskGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskGroup%2Did"] = id
    }
    return NewUsersItemOutlookTaskGroupsOutlookTaskGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
func (m *UsersItemOutlookRequestBuilder) Tasks()(*UsersItemOutlookTasksRequestBuilder) {
    return NewUsersItemOutlookTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
func (m *UsersItemOutlookRequestBuilder) TasksById(id string)(*UsersItemOutlookTasksOutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask%2Did"] = id
    }
    return NewUsersItemOutlookTasksOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
