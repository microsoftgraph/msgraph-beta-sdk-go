package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder provides operations to manage the jobs property of the microsoft.graph.synchronization entity.
type ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetQueryParameters performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
type ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetQueryParameters
}
// ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderInternal instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) {
    m := &ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property jobs for applications
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property jobs in applications
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, requestConfiguration *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property jobs for applications
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, error) {
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
// Patch update the navigation property jobs in applications
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, requestConfiguration *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, error) {
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
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Pause()(*ApplicationsItemSynchronizationJobsItemPauseRequestBuilder) {
    return NewApplicationsItemSynchronizationJobsItemPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisionOnDemand provides operations to call the provisionOnDemand method.
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) ProvisionOnDemand()(*ApplicationsItemSynchronizationJobsItemProvisionOnDemandRequestBuilder) {
    return NewApplicationsItemSynchronizationJobsItemProvisionOnDemandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restart provides operations to call the restart method.
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Restart()(*ApplicationsItemSynchronizationJobsItemRestartRequestBuilder) {
    return NewApplicationsItemSynchronizationJobsItemRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schema provides operations to manage the schema property of the microsoft.graph.synchronizationJob entity.
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Schema()(*ApplicationsItemSynchronizationJobsItemSchemaRequestBuilder) {
    return NewApplicationsItemSynchronizationJobsItemSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Start provides operations to call the start method.
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Start()(*ApplicationsItemSynchronizationJobsItemStartRequestBuilder) {
    return NewApplicationsItemSynchronizationJobsItemStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) Stop()(*ApplicationsItemSynchronizationJobsItemStopRequestBuilder) {
    return NewApplicationsItemSynchronizationJobsItemStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateCredentials provides operations to call the validateCredentials method.
func (m *ApplicationsItemSynchronizationJobsSynchronizationJobItemRequestBuilder) ValidateCredentials()(*ApplicationsItemSynchronizationJobsItemValidateCredentialsRequestBuilder) {
    return NewApplicationsItemSynchronizationJobsItemValidateCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
