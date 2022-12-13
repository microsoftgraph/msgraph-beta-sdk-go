package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OutlookRequestBuilder provides operations to manage the outlook property of the microsoft.graph.user entity.
type OutlookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OutlookRequestBuilderGetQueryParameters selective Outlook services available to the user. Read-only. Nullable.
type OutlookRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OutlookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OutlookRequestBuilderGetQueryParameters
}
// NewOutlookRequestBuilderInternal instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookRequestBuilder) {
    m := &OutlookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutlookRequestBuilder instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation selective Outlook services available to the user. Read-only. Nullable.
func (m *OutlookRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OutlookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get selective Outlook services available to the user. Read-only. Nullable.
func (m *OutlookRequestBuilder) Get(ctx context.Context, requestConfiguration *OutlookRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, error) {
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
func (m *OutlookRequestBuilder) MasterCategories()(*OutlookMasterCategoriesRequestBuilder) {
    return NewOutlookMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MasterCategoriesById provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) MasterCategoriesById(id string)(*OutlookMasterCategoriesOutlookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory%2Did"] = id
    }
    return NewOutlookMasterCategoriesOutlookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SupportedLanguages provides operations to call the supportedLanguages method.
func (m *OutlookRequestBuilder) SupportedLanguages()(*OutlookSupportedLanguagesRequestBuilder) {
    return NewOutlookSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZones provides operations to call the supportedTimeZones method.
func (m *OutlookRequestBuilder) SupportedTimeZones()(*OutlookSupportedTimeZonesRequestBuilder) {
    return NewOutlookSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZonesWithTimeZoneStandard provides operations to call the supportedTimeZones method.
func (m *OutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*OutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return NewOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard);
}
// TaskFolders provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TaskFolders()(*OutlookTaskFoldersRequestBuilder) {
    return NewOutlookTaskFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskFoldersById provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TaskFoldersById(id string)(*OutlookTaskFoldersOutlookTaskFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskFolder%2Did"] = id
    }
    return NewOutlookTaskFoldersOutlookTaskFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskGroups provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TaskGroups()(*OutlookTaskGroupsRequestBuilder) {
    return NewOutlookTaskGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskGroupsById provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TaskGroupsById(id string)(*OutlookTaskGroupsOutlookTaskGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskGroup%2Did"] = id
    }
    return NewOutlookTaskGroupsOutlookTaskGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) Tasks()(*OutlookTasksRequestBuilder) {
    return NewOutlookTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TasksById(id string)(*OutlookTasksOutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask%2Did"] = id
    }
    return NewOutlookTasksOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
