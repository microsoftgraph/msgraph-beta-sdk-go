package outlook

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedtimezoneswithtimezonestandard"
    i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups"
    i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders"
    ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/mastercategories"
    ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedtimezones"
    ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks"
    ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedlanguages"
    i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/mastercategories/item"
    i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item"
    ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item"
    id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item"
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
    Headers map[string]string
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
func (m *OutlookRequestBuilder) MasterCategories()(*ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296.MasterCategoriesRequestBuilder) {
    return ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296.NewMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MasterCategoriesById provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) MasterCategoriesById(id string)(*i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40.OutlookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory%2Did"] = id
    }
    return i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40.NewOutlookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SupportedLanguages provides operations to call the supportedLanguages method.
func (m *OutlookRequestBuilder) SupportedLanguages()(*ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be.SupportedLanguagesRequestBuilder) {
    return ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be.NewSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZones provides operations to call the supportedTimeZones method.
func (m *OutlookRequestBuilder) SupportedTimeZones()(*ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04.SupportedTimeZonesRequestBuilder) {
    return ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04.NewSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZonesWithTimeZoneStandard provides operations to call the supportedTimeZones method.
func (m *OutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a.SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a.NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard);
}
// TaskFolders provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TaskFolders()(*i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251.TaskFoldersRequestBuilder) {
    return i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251.NewTaskFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskFoldersById provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TaskFoldersById(id string)(*ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a.OutlookTaskFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskFolder%2Did"] = id
    }
    return ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a.NewOutlookTaskFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskGroups provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TaskGroups()(*i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21.TaskGroupsRequestBuilder) {
    return i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21.NewTaskGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskGroupsById provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TaskGroupsById(id string)(*i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb.OutlookTaskGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskGroup%2Did"] = id
    }
    return i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb.NewOutlookTaskGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) Tasks()(*ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00.TasksRequestBuilder) {
    return ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) TasksById(id string)(*id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00.OutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask%2Did"] = id
    }
    return id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00.NewOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
