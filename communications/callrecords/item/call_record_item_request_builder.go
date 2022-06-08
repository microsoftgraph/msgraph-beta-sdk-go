package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
    i8da91d80f834afa257263f71267b10122fb002ee361fdcad0ffef83b9e5cf724 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/callrecords/item/sessions"
    ic5dfb8f4b7db4be39da7edb542b78c201cb5a192410b1ee4562017341bf3fda3 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/callrecords/item/sessions/item"
)

// CallRecordItemRequestBuilder provides operations to manage the callRecords property of the microsoft.graph.cloudCommunications entity.
type CallRecordItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallRecordItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallRecordItemRequestBuilderGetQueryParameters get callRecords from communications
type CallRecordItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CallRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallRecordItemRequestBuilderGetQueryParameters
}
// CallRecordItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallRecordItemRequestBuilderInternal instantiates a new CallRecordItemRequestBuilder and sets the default values.
func NewCallRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordItemRequestBuilder) {
    m := &CallRecordItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/callRecords/{callRecord%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCallRecordItemRequestBuilder instantiates a new CallRecordItemRequestBuilder and sets the default values.
func NewCallRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property callRecords for communications
func (m *CallRecordItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property callRecords for communications
func (m *CallRecordItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *CallRecordItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get callRecords from communications
func (m *CallRecordItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get callRecords from communications
func (m *CallRecordItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CallRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property callRecords in communications
func (m *CallRecordItemRequestBuilder) CreatePatchRequestInformation(body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property callRecords in communications
func (m *CallRecordItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable, requestConfiguration *CallRecordItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property callRecords for communications
func (m *CallRecordItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property callRecords for communications
func (m *CallRecordItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *CallRecordItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get get callRecords from communications
func (m *CallRecordItemRequestBuilder) Get()(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get callRecords from communications
func (m *CallRecordItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *CallRecordItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateCallRecordFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable), nil
}
// Patch update the navigation property callRecords in communications
func (m *CallRecordItemRequestBuilder) Patch(body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property callRecords in communications
func (m *CallRecordItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CallRecordable, requestConfiguration *CallRecordItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Sessions the sessions property
func (m *CallRecordItemRequestBuilder) Sessions()(*i8da91d80f834afa257263f71267b10122fb002ee361fdcad0ffef83b9e5cf724.SessionsRequestBuilder) {
    return i8da91d80f834afa257263f71267b10122fb002ee361fdcad0ffef83b9e5cf724.NewSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SessionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.callRecords.item.sessions.item collection
func (m *CallRecordItemRequestBuilder) SessionsById(id string)(*ic5dfb8f4b7db4be39da7edb542b78c201cb5a192410b1ee4562017341bf3fda3.SessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["session%2Did"] = id
    }
    return ic5dfb8f4b7db4be39da7edb542b78c201cb5a192410b1ee4562017341bf3fda3.NewSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
