package synchronization

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates"
    i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/ping"
    i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs"
    ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/acquireaccesstoken"
    i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item"
    ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item"
)

// SynchronizationRequestBuilder provides operations to manage the synchronization property of the microsoft.graph.application entity.
type SynchronizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SynchronizationRequestBuilderDeleteOptions options for Delete
type SynchronizationRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SynchronizationRequestBuilderGetOptions options for Get
type SynchronizationRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SynchronizationRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SynchronizationRequestBuilderGetQueryParameters get synchronization from applications
type SynchronizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SynchronizationRequestBuilderPatchOptions options for Patch
type SynchronizationRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AcquireAccessToken the acquireAccessToken property
func (m *SynchronizationRequestBuilder) AcquireAccessToken()(*ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050.AcquireAccessTokenRequestBuilder) {
    return ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050.NewAcquireAccessTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSynchronizationRequestBuilderInternal instantiates a new SynchronizationRequestBuilder and sets the default values.
func NewSynchronizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationRequestBuilder) {
    m := &SynchronizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/synchronization{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSynchronizationRequestBuilder instantiates a new SynchronizationRequestBuilder and sets the default values.
func NewSynchronizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property synchronization for applications
func (m *SynchronizationRequestBuilder) CreateDeleteRequestInformation(options *SynchronizationRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get synchronization from applications
func (m *SynchronizationRequestBuilder) CreateGetRequestInformation(options *SynchronizationRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property synchronization in applications
func (m *SynchronizationRequestBuilder) CreatePatchRequestInformation(options *SynchronizationRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property synchronization for applications
func (m *SynchronizationRequestBuilder) Delete(options *SynchronizationRequestBuilderDeleteOptions)(error) {
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
// Get get synchronization from applications
func (m *SynchronizationRequestBuilder) Get(options *SynchronizationRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable), nil
}
// Jobs the jobs property
func (m *SynchronizationRequestBuilder) Jobs()(*i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373.JobsRequestBuilder) {
    return i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373.NewJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.synchronization.jobs.item collection
func (m *SynchronizationRequestBuilder) JobsById(id string)(*ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2.SynchronizationJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationJob%2Did"] = id
    }
    return ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2.NewSynchronizationJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property synchronization in applications
func (m *SynchronizationRequestBuilder) Patch(options *SynchronizationRequestBuilderPatchOptions)(error) {
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
// Ping provides operations to call the Ping method.
func (m *SynchronizationRequestBuilder) Ping()(*i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0.PingRequestBuilder) {
    return i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0.NewPingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Templates the templates property
func (m *SynchronizationRequestBuilder) Templates()(*i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1.TemplatesRequestBuilder) {
    return i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1.NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.synchronization.templates.item collection
func (m *SynchronizationRequestBuilder) TemplatesById(id string)(*i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a.SynchronizationTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationTemplate%2Did"] = id
    }
    return i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a.NewSynchronizationTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
