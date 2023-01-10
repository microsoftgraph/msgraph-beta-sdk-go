package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOutlookRequestBuilder provides operations to manage the outlook property of the microsoft.graph.user entity.
type ItemOutlookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOutlookRequestBuilderGetQueryParameters selective Outlook services available to the user. Read-only. Nullable.
type ItemOutlookRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOutlookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOutlookRequestBuilderGetQueryParameters
}
// NewItemOutlookRequestBuilderInternal instantiates a new OutlookRequestBuilder and sets the default values.
func NewItemOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookRequestBuilder) {
    m := &ItemOutlookRequestBuilder{
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
// NewItemOutlookRequestBuilder instantiates a new OutlookRequestBuilder and sets the default values.
func NewItemOutlookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
// Get selective Outlook services available to the user. Read-only. Nullable.
func (m *ItemOutlookRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutlookRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
func (m *ItemOutlookRequestBuilder) MasterCategories()(*ItemOutlookMasterCategoriesRequestBuilder) {
    return NewItemOutlookMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MasterCategoriesById provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
func (m *ItemOutlookRequestBuilder) MasterCategoriesById(id string)(*ItemOutlookMasterCategoriesOutlookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory%2Did"] = id
    }
    return NewItemOutlookMasterCategoriesOutlookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SupportedLanguages provides operations to call the supportedLanguages method.
func (m *ItemOutlookRequestBuilder) SupportedLanguages()(*ItemOutlookSupportedLanguagesRequestBuilder) {
    return NewItemOutlookSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZones provides operations to call the supportedTimeZones method.
func (m *ItemOutlookRequestBuilder) SupportedTimeZones()(*ItemOutlookSupportedTimeZonesRequestBuilder) {
    return NewItemOutlookSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZonesWithTimeZoneStandard provides operations to call the supportedTimeZones method.
func (m *ItemOutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*ItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return NewItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard);
}
// TaskFolders provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
func (m *ItemOutlookRequestBuilder) TaskFolders()(*ItemOutlookTaskFoldersRequestBuilder) {
    return NewItemOutlookTaskFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskFoldersById provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
func (m *ItemOutlookRequestBuilder) TaskFoldersById(id string)(*ItemOutlookTaskFoldersOutlookTaskFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskFolder%2Did"] = id
    }
    return NewItemOutlookTaskFoldersOutlookTaskFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskGroups provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
func (m *ItemOutlookRequestBuilder) TaskGroups()(*ItemOutlookTaskGroupsRequestBuilder) {
    return NewItemOutlookTaskGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskGroupsById provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
func (m *ItemOutlookRequestBuilder) TaskGroupsById(id string)(*ItemOutlookTaskGroupsOutlookTaskGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskGroup%2Did"] = id
    }
    return NewItemOutlookTaskGroupsOutlookTaskGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
func (m *ItemOutlookRequestBuilder) Tasks()(*ItemOutlookTasksRequestBuilder) {
    return NewItemOutlookTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
func (m *ItemOutlookRequestBuilder) TasksById(id string)(*ItemOutlookTasksOutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask%2Did"] = id
    }
    return NewItemOutlookTasksOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToGetRequestInformation selective Outlook services available to the user. Read-only. Nullable.
func (m *ItemOutlookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
