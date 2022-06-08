package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0718f2c7a8c775532496a98d185ae4cfcb0a71392db198afe0ebb5deddd3bfdd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/attachments"
    i0e3cd5b3643da0403d06ec109f55ebdd821409ca2b64e8929398963347543678 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/complete"
    i2488a196ac751025aa25b85acf5ef7ee5e6c2a37ea695ecd83a9367d0fbc5c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/singlevalueextendedproperties"
    iec0dcd96a66f87e472d58355850a2262b0d0a17eb7b7dc13ff86168db555b568 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/multivalueextendedproperties"
    i175900f9235a31a08f8f7870a3a233657ed75f27ea0f9c33a35029ae11dadd54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/attachments/item"
    ie3373c186191b55aac268824553eab67e16cbf405be52509be0da3cd61c6f601 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/multivalueextendedproperties/item"
    if8a609bdb585ca1f49d83f9d345d43bf56ed08b184a513e1f1cd22d928dad3ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/singlevalueextendedproperties/item"
)

// OutlookTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.outlookTaskFolder entity.
type OutlookTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OutlookTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OutlookTaskItemRequestBuilderGetQueryParameters the tasks in this task folder. Read-only. Nullable.
type OutlookTaskItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OutlookTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OutlookTaskItemRequestBuilderGetQueryParameters
}
// OutlookTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments the attachments property
func (m *OutlookTaskItemRequestBuilder) Attachments()(*i0718f2c7a8c775532496a98d185ae4cfcb0a71392db198afe0ebb5deddd3bfdd.AttachmentsRequestBuilder) {
    return i0718f2c7a8c775532496a98d185ae4cfcb0a71392db198afe0ebb5deddd3bfdd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.tasks.item.attachments.item collection
func (m *OutlookTaskItemRequestBuilder) AttachmentsById(id string)(*i175900f9235a31a08f8f7870a3a233657ed75f27ea0f9c33a35029ae11dadd54.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i175900f9235a31a08f8f7870a3a233657ed75f27ea0f9c33a35029ae11dadd54.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Complete the complete property
func (m *OutlookTaskItemRequestBuilder) Complete()(*i0e3cd5b3643da0403d06ec109f55ebdd821409ca2b64e8929398963347543678.CompleteRequestBuilder) {
    return i0e3cd5b3643da0403d06ec109f55ebdd821409ca2b64e8929398963347543678.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOutlookTaskItemRequestBuilderInternal instantiates a new OutlookTaskItemRequestBuilder and sets the default values.
func NewOutlookTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookTaskItemRequestBuilder) {
    m := &OutlookTaskItemRequestBuilder{
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
// NewOutlookTaskItemRequestBuilder instantiates a new OutlookTaskItemRequestBuilder and sets the default values.
func NewOutlookTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property tasks for users
func (m *OutlookTaskItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property tasks for users
func (m *OutlookTaskItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *OutlookTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OutlookTaskItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OutlookTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OutlookTaskItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property tasks in users
func (m *OutlookTaskItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *OutlookTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property tasks for users
func (m *OutlookTaskItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property tasks for users
func (m *OutlookTaskItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *OutlookTaskItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OutlookTaskItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *OutlookTaskItemRequestBuilder) MultiValueExtendedProperties()(*iec0dcd96a66f87e472d58355850a2262b0d0a17eb7b7dc13ff86168db555b568.MultiValueExtendedPropertiesRequestBuilder) {
    return iec0dcd96a66f87e472d58355850a2262b0d0a17eb7b7dc13ff86168db555b568.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.tasks.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie3373c186191b55aac268824553eab67e16cbf405be52509be0da3cd61c6f601.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ie3373c186191b55aac268824553eab67e16cbf405be52509be0da3cd61c6f601.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property tasks in users
func (m *OutlookTaskItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property tasks in users
func (m *OutlookTaskItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *OutlookTaskItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *OutlookTaskItemRequestBuilder) SingleValueExtendedProperties()(*i2488a196ac751025aa25b85acf5ef7ee5e6c2a37ea695ecd83a9367d0fbc5c85.SingleValueExtendedPropertiesRequestBuilder) {
    return i2488a196ac751025aa25b85acf5ef7ee5e6c2a37ea695ecd83a9367d0fbc5c85.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.tasks.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if8a609bdb585ca1f49d83f9d345d43bf56ed08b184a513e1f1cd22d928dad3ed.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return if8a609bdb585ca1f49d83f9d345d43bf56ed08b184a513e1f1cd22d928dad3ed.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
