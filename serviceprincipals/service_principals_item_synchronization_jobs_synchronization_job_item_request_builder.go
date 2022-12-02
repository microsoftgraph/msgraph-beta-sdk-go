package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder provides operations to manage the jobs property of the microsoft.graph.synchronization entity.
type ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetQueryParameters performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
type ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetQueryParameters
}
// ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderInternal instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) {
    m := &ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/jobs/{synchronizationJob%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property jobs for servicePrincipals
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property jobs in servicePrincipals
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, requestConfiguration *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property jobs for servicePrincipals
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable), nil
}
// Patch update the navigation property jobs in servicePrincipals
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, requestConfiguration *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable), nil
}
// Pause provides operations to call the pause method.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Pause()(*ServicePrincipalsItemSynchronizationJobsItemPauseRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationJobsItemPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisionOnDemand provides operations to call the provisionOnDemand method.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) ProvisionOnDemand()(*ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restart provides operations to call the restart method.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Restart()(*ServicePrincipalsItemSynchronizationJobsItemRestartRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationJobsItemRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schema provides operations to manage the schema property of the microsoft.graph.synchronizationJob entity.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Schema()(*ServicePrincipalsItemSynchronizationJobsItemSchemaRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationJobsItemSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Start provides operations to call the start method.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Start()(*ServicePrincipalsItemSynchronizationJobsItemStartRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationJobsItemStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Stop()(*ServicePrincipalsItemSynchronizationJobsItemStopRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationJobsItemStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateCredentials provides operations to call the validateCredentials method.
func (m *ServicePrincipalsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) ValidateCredentials()(*ServicePrincipalsItemSynchronizationJobsItemValidateCredentialsRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationJobsItemValidateCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
