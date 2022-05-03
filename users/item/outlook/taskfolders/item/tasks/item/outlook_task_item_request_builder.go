package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07134841e1432d33390c1388213ee45e6595d43d6e0eb485c3d75c5572f84b26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/singlevalueextendedproperties"
    i30d7001f23621738f5c0ebf4de402aa0940898dabde9c286e7bde7c8fe3a247a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/attachments"
    ibeff8c7ce77c8ec152f96297bac9743c951b7257a49b9309cd240df9aa287826 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/multivalueextendedproperties"
    ieeee0fbd5dbbdd405ef962f5cf1518d5b11621a62b9593c79e43559e54f18acc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/complete"
    i42d4fa154ff58194c41d7a068d2ef0ea649f669575467c50bac921280cece148 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/singlevalueextendedproperties/item"
    i8c19f39e583cd9a1b9f36e6c8fe3732bbb945cb45363ddd749683c9976564002 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/multivalueextendedproperties/item"
    ic6491e7b1ea237fd1a5fc05790050c780572b7f37e967ddffd4225e1dc62ce51 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/attachments/item"
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
func (m *OutlookTaskItemRequestBuilder) Attachments()(*i30d7001f23621738f5c0ebf4de402aa0940898dabde9c286e7bde7c8fe3a247a.AttachmentsRequestBuilder) {
    return i30d7001f23621738f5c0ebf4de402aa0940898dabde9c286e7bde7c8fe3a247a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskFolders.item.tasks.item.attachments.item collection
func (m *OutlookTaskItemRequestBuilder) AttachmentsById(id string)(*ic6491e7b1ea237fd1a5fc05790050c780572b7f37e967ddffd4225e1dc62ce51.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ic6491e7b1ea237fd1a5fc05790050c780572b7f37e967ddffd4225e1dc62ce51.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Complete the complete property
func (m *OutlookTaskItemRequestBuilder) Complete()(*ieeee0fbd5dbbdd405ef962f5cf1518d5b11621a62b9593c79e43559e54f18acc.CompleteRequestBuilder) {
    return ieeee0fbd5dbbdd405ef962f5cf1518d5b11621a62b9593c79e43559e54f18acc.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOutlookTaskItemRequestBuilderInternal instantiates a new OutlookTaskItemRequestBuilder and sets the default values.
func NewOutlookTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookTaskItemRequestBuilder) {
    m := &OutlookTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/outlook/taskFolders/{outlookTaskFolder%2Did}/tasks/{outlookTask%2Did}{?%24select}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property tasks for users
func (m *OutlookTaskItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OutlookTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property tasks in users
func (m *OutlookTaskItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property tasks for users
func (m *OutlookTaskItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *OutlookTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property tasks for users
func (m *OutlookTaskItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *OutlookTaskItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskItemRequestBuilder) GetWithResponseHandler(requestConfiguration *OutlookTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskItemRequestBuilder) GetWithResponseHandler(requestConfiguration *OutlookTaskItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
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
func (m *OutlookTaskItemRequestBuilder) MultiValueExtendedProperties()(*ibeff8c7ce77c8ec152f96297bac9743c951b7257a49b9309cd240df9aa287826.MultiValueExtendedPropertiesRequestBuilder) {
    return ibeff8c7ce77c8ec152f96297bac9743c951b7257a49b9309cd240df9aa287826.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskFolders.item.tasks.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i8c19f39e583cd9a1b9f36e6c8fe3732bbb945cb45363ddd749683c9976564002.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i8c19f39e583cd9a1b9f36e6c8fe3732bbb945cb45363ddd749683c9976564002.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property tasks in users
func (m *OutlookTaskItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *OutlookTaskItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property tasks in users
func (m *OutlookTaskItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *OutlookTaskItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *OutlookTaskItemRequestBuilder) SingleValueExtendedProperties()(*i07134841e1432d33390c1388213ee45e6595d43d6e0eb485c3d75c5572f84b26.SingleValueExtendedPropertiesRequestBuilder) {
    return i07134841e1432d33390c1388213ee45e6595d43d6e0eb485c3d75c5572f84b26.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskFolders.item.tasks.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i42d4fa154ff58194c41d7a068d2ef0ea649f669575467c50bac921280cece148.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i42d4fa154ff58194c41d7a068d2ef0ea649f669575467c50bac921280cece148.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
