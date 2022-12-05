package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalsItemSynchronizationRequestBuilder provides operations to manage the synchronization property of the microsoft.graph.servicePrincipal entity.
type ServicePrincipalsItemSynchronizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalsItemSynchronizationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalsItemSynchronizationRequestBuilderGetQueryParameters get synchronization from servicePrincipals
type ServicePrincipalsItemSynchronizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalsItemSynchronizationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalsItemSynchronizationRequestBuilderGetQueryParameters
}
// ServicePrincipalsItemSynchronizationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcquireAccessToken provides operations to call the acquireAccessToken method.
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) AcquireAccessToken()(*ServicePrincipalsItemSynchronizationAcquireAccessTokenRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationAcquireAccessTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalsItemSynchronizationRequestBuilderInternal instantiates a new SynchronizationRequestBuilder and sets the default values.
func NewServicePrincipalsItemSynchronizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemSynchronizationRequestBuilder) {
    m := &ServicePrincipalsItemSynchronizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsItemSynchronizationRequestBuilder instantiates a new SynchronizationRequestBuilder and sets the default values.
func NewServicePrincipalsItemSynchronizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemSynchronizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsItemSynchronizationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property synchronization for servicePrincipals
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get synchronization from servicePrincipals
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property synchronization in servicePrincipals
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable, requestConfiguration *ServicePrincipalsItemSynchronizationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property synchronization for servicePrincipals
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get synchronization from servicePrincipals
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable), nil
}
// Jobs provides operations to manage the jobs property of the microsoft.graph.synchronization entity.
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) Jobs()(*ServicePrincipalsItemSynchronizationJobsRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JobsById provides operations to manage the jobs property of the microsoft.graph.synchronization entity.
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) JobsById(id string)(*ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationJob%2Did"] = id
    }
    return NewServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property synchronization in servicePrincipals
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable, requestConfiguration *ServicePrincipalsItemSynchronizationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Synchronizationable), nil
}
// Ping provides operations to call the Ping method.
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) Ping()(*ServicePrincipalsItemSynchronizationPingRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationPingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Templates provides operations to manage the templates property of the microsoft.graph.synchronization entity.
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) Templates()(*ServicePrincipalsItemSynchronizationTemplatesRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplatesById provides operations to manage the templates property of the microsoft.graph.synchronization entity.
func (m *ServicePrincipalsItemSynchronizationRequestBuilder) TemplatesById(id string)(*ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationTemplate%2Did"] = id
    }
    return NewServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
