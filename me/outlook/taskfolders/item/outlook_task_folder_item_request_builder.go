package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks"
    i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/singlevalueextendedproperties"
    i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/multivalueextendedproperties"
    i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item"
    i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/multivalueextendedproperties/item"
    i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/singlevalueextendedproperties/item"
)

// OutlookTaskFolderItemRequestBuilder provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
type OutlookTaskFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OutlookTaskFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookTaskFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OutlookTaskFolderItemRequestBuilderGetQueryParameters get taskFolders from me
type OutlookTaskFolderItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OutlookTaskFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookTaskFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OutlookTaskFolderItemRequestBuilderGetQueryParameters
}
// OutlookTaskFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookTaskFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOutlookTaskFolderItemRequestBuilderInternal instantiates a new OutlookTaskFolderItemRequestBuilder and sets the default values.
func NewOutlookTaskFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookTaskFolderItemRequestBuilder) {
    m := &OutlookTaskFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook/taskFolders/{outlookTaskFolder%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutlookTaskFolderItemRequestBuilder instantiates a new OutlookTaskFolderItemRequestBuilder and sets the default values.
func NewOutlookTaskFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookTaskFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property taskFolders for me
func (m *OutlookTaskFolderItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property taskFolders for me
func (m *OutlookTaskFolderItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *OutlookTaskFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get taskFolders from me
func (m *OutlookTaskFolderItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get taskFolders from me
func (m *OutlookTaskFolderItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OutlookTaskFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property taskFolders in me
func (m *OutlookTaskFolderItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property taskFolders in me
func (m *OutlookTaskFolderItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable, requestConfiguration *OutlookTaskFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property taskFolders for me
func (m *OutlookTaskFolderItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property taskFolders for me
func (m *OutlookTaskFolderItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *OutlookTaskFolderItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get taskFolders from me
func (m *OutlookTaskFolderItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get taskFolders from me
func (m *OutlookTaskFolderItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OutlookTaskFolderItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFolderFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *OutlookTaskFolderItemRequestBuilder) MultiValueExtendedProperties()(*i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c.MultiValueExtendedPropertiesRequestBuilder) {
    return i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property taskFolders in me
func (m *OutlookTaskFolderItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property taskFolders in me
func (m *OutlookTaskFolderItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskFolderable, requestConfiguration *OutlookTaskFolderItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *OutlookTaskFolderItemRequestBuilder) SingleValueExtendedProperties()(*i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a.SingleValueExtendedPropertiesRequestBuilder) {
    return i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks the tasks property
func (m *OutlookTaskFolderItemRequestBuilder) Tasks()(*i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a.TasksRequestBuilder) {
    return i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.tasks.item collection
func (m *OutlookTaskFolderItemRequestBuilder) TasksById(id string)(*i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9.OutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask%2Did"] = id
    }
    return i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9.NewOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
