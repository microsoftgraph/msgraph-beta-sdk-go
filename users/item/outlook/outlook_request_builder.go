package outlook

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0d34dad914e82d03ae5a89b298e4a19b8b8b6bf4eb26a5999888bd525b1fd9c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/supportedtimezoneswithtimezonestandard"
    i360efeb2914ada10b422929cbecf0f7f8fa695d7b1022878848afca7acfd0ddf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/mastercategories"
    i54cf645f94c231fabc20f9cde8f9bb429da86a32c740ad8c7fcbc3a398e8be24 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups"
    i6913acf90458035b6558bc17364451cf3efd2854e0918e1ebc3a438b11ba8b04 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/supportedtimezones"
    ia668b41f7ac335c1f54bbbbdde0f0ba9b8edb241acbb56fd81a1b8db3fb2ff10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/supportedlanguages"
    ib512bb7334ff30b38cfa6d102f049adfcb6a68103965b326b25efbd5874ed400 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks"
    idbf17811c520d1085f831150de278ee1bb194206c8b3ab498e0d0292a6511ace "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders"
    i5301c0330f5e20cf6591fab9dc299a0ee36c9818473c8c38663b5f5a43673c06 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item"
    i74db178f0907a8ff4bbfcfc53a94d5520a643e2563712703082afd071884dd1e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item"
    iacb0aac7a3842511fec05c93f217164eb10589c31806a96bf05179087b60e399 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/mastercategories/item"
    iaffe5863c3d1d6c8c2d7abc4fc3d594133e2a40279216bc540d6d0ad80f03059 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item"
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
// OutlookRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOutlookRequestBuilderInternal instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookRequestBuilder) {
    m := &OutlookRequestBuilder{
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
// NewOutlookRequestBuilder instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation selective Outlook services available to the user. Read-only. Nullable.
func (m *OutlookRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration selective Outlook services available to the user. Read-only. Nullable.
func (m *OutlookRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OutlookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property outlook in users
func (m *OutlookRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property outlook in users
func (m *OutlookRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, requestConfiguration *OutlookRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get selective Outlook services available to the user. Read-only. Nullable.
func (m *OutlookRequestBuilder) Get(ctx context.Context, requestConfiguration *OutlookRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// MasterCategories the masterCategories property
func (m *OutlookRequestBuilder) MasterCategories()(*i360efeb2914ada10b422929cbecf0f7f8fa695d7b1022878848afca7acfd0ddf.MasterCategoriesRequestBuilder) {
    return i360efeb2914ada10b422929cbecf0f7f8fa695d7b1022878848afca7acfd0ddf.NewMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MasterCategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.masterCategories.item collection
func (m *OutlookRequestBuilder) MasterCategoriesById(id string)(*iacb0aac7a3842511fec05c93f217164eb10589c31806a96bf05179087b60e399.OutlookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory%2Did"] = id
    }
    return iacb0aac7a3842511fec05c93f217164eb10589c31806a96bf05179087b60e399.NewOutlookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property outlook in users
func (m *OutlookRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, requestConfiguration *OutlookRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// SupportedLanguages provides operations to call the supportedLanguages method.
func (m *OutlookRequestBuilder) SupportedLanguages()(*ia668b41f7ac335c1f54bbbbdde0f0ba9b8edb241acbb56fd81a1b8db3fb2ff10.SupportedLanguagesRequestBuilder) {
    return ia668b41f7ac335c1f54bbbbdde0f0ba9b8edb241acbb56fd81a1b8db3fb2ff10.NewSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZones provides operations to call the supportedTimeZones method.
func (m *OutlookRequestBuilder) SupportedTimeZones()(*i6913acf90458035b6558bc17364451cf3efd2854e0918e1ebc3a438b11ba8b04.SupportedTimeZonesRequestBuilder) {
    return i6913acf90458035b6558bc17364451cf3efd2854e0918e1ebc3a438b11ba8b04.NewSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZonesWithTimeZoneStandard provides operations to call the supportedTimeZones method.
func (m *OutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*i0d34dad914e82d03ae5a89b298e4a19b8b8b6bf4eb26a5999888bd525b1fd9c5.SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return i0d34dad914e82d03ae5a89b298e4a19b8b8b6bf4eb26a5999888bd525b1fd9c5.NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard);
}
// TaskFolders the taskFolders property
func (m *OutlookRequestBuilder) TaskFolders()(*idbf17811c520d1085f831150de278ee1bb194206c8b3ab498e0d0292a6511ace.TaskFoldersRequestBuilder) {
    return idbf17811c520d1085f831150de278ee1bb194206c8b3ab498e0d0292a6511ace.NewTaskFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskFolders.item collection
func (m *OutlookRequestBuilder) TaskFoldersById(id string)(*i74db178f0907a8ff4bbfcfc53a94d5520a643e2563712703082afd071884dd1e.OutlookTaskFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskFolder%2Did"] = id
    }
    return i74db178f0907a8ff4bbfcfc53a94d5520a643e2563712703082afd071884dd1e.NewOutlookTaskFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskGroups the taskGroups property
func (m *OutlookRequestBuilder) TaskGroups()(*i54cf645f94c231fabc20f9cde8f9bb429da86a32c740ad8c7fcbc3a398e8be24.TaskGroupsRequestBuilder) {
    return i54cf645f94c231fabc20f9cde8f9bb429da86a32c740ad8c7fcbc3a398e8be24.NewTaskGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item collection
func (m *OutlookRequestBuilder) TaskGroupsById(id string)(*iaffe5863c3d1d6c8c2d7abc4fc3d594133e2a40279216bc540d6d0ad80f03059.OutlookTaskGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskGroup%2Did"] = id
    }
    return iaffe5863c3d1d6c8c2d7abc4fc3d594133e2a40279216bc540d6d0ad80f03059.NewOutlookTaskGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks the tasks property
func (m *OutlookRequestBuilder) Tasks()(*ib512bb7334ff30b38cfa6d102f049adfcb6a68103965b326b25efbd5874ed400.TasksRequestBuilder) {
    return ib512bb7334ff30b38cfa6d102f049adfcb6a68103965b326b25efbd5874ed400.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.tasks.item collection
func (m *OutlookRequestBuilder) TasksById(id string)(*i5301c0330f5e20cf6591fab9dc299a0ee36c9818473c8c38663b5f5a43673c06.OutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask%2Did"] = id
    }
    return i5301c0330f5e20cf6591fab9dc299a0ee36c9818473c8c38663b5f5a43673c06.NewOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
