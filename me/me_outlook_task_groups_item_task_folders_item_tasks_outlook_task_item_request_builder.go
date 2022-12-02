package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.outlookTaskFolder entity.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters the tasks in this task folder. Read-only. Nullable.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetQueryParameters
}
// MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.outlookTask entity.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Attachments()(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsRequestBuilder) {
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.outlookTask entity.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) AttachmentsById(id string)(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Complete provides operations to call the complete method.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Complete()(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) {
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal instantiates a new OutlookTaskItemRequestBuilder and sets the default values.
func NewMeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) {
    m := &MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook/taskGroups/{outlookTaskGroup%2Did}/taskFolders/{outlookTaskFolder%2Did}/tasks/{outlookTask%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder instantiates a new OutlookTaskItemRequestBuilder and sets the default values.
func NewMeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property tasks for me
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property tasks in me
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property tasks for me
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the tasks in this task folder. Read-only. Nullable.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable), nil
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.outlookTask entity.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) MultiValueExtendedProperties()(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.outlookTask entity.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property tasks in me
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, requestConfiguration *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookTaskable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.outlookTask entity.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) SingleValueExtendedProperties()(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.outlookTask entity.
func (m *MeOutlookTaskGroupsItemTaskFoldersItemTasksOutlookTaskItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*MeOutlookTaskGroupsItemTaskFoldersItemTasksItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeOutlookTaskGroupsItemTaskFoldersItemTasksItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
