package outlook

import (
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OutlookRequestBuilderDeleteOptions options for Delete
type OutlookRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OutlookRequestBuilderGetOptions options for Get
type OutlookRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *OutlookRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OutlookRequestBuilderGetQueryParameters read-only.
type OutlookRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// OutlookRequestBuilderPatchOptions options for Patch
type OutlookRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewOutlookRequestBuilderInternal instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookRequestBuilder) {
    m := &OutlookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook{?select}";
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
// CreateDeleteRequestInformation delete navigation property outlook for users
func (m *OutlookRequestBuilder) CreateDeleteRequestInformation(options *OutlookRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation read-only.
func (m *OutlookRequestBuilder) CreateGetRequestInformation(options *OutlookRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property outlook in users
func (m *OutlookRequestBuilder) CreatePatchRequestInformation(options *OutlookRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property outlook for users
func (m *OutlookRequestBuilder) Delete(options *OutlookRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only.
func (m *OutlookRequestBuilder) Get(options *OutlookRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
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
        urlTplParams["outlookCategory_id"] = id
    }
    return iacb0aac7a3842511fec05c93f217164eb10589c31806a96bf05179087b60e399.NewOutlookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property outlook in users
func (m *OutlookRequestBuilder) Patch(options *OutlookRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
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
        urlTplParams["outlookTaskFolder_id"] = id
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
        urlTplParams["outlookTaskGroup_id"] = id
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
        urlTplParams["outlookTask_id"] = id
    }
    return i5301c0330f5e20cf6591fab9dc299a0ee36c9818473c8c38663b5f5a43673c06.NewOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
