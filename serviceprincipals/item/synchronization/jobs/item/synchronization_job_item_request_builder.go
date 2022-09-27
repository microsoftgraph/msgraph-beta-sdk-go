package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/schema"
    i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/stop"
    i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/provisionondemand"
    i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/restart"
    ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/validatecredentials"
    id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/pause"
    ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/start"
)

// SynchronizationJobItemRequestBuilder provides operations to manage the jobs property of the microsoft.graph.synchronization entity.
type SynchronizationJobItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SynchronizationJobItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationJobItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SynchronizationJobItemRequestBuilderGetQueryParameters performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
type SynchronizationJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SynchronizationJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SynchronizationJobItemRequestBuilderGetQueryParameters
}
// SynchronizationJobItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationJobItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationJobItemRequestBuilderInternal instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewSynchronizationJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationJobItemRequestBuilder) {
    m := &SynchronizationJobItemRequestBuilder{
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
// NewSynchronizationJobItemRequestBuilder instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewSynchronizationJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property jobs for servicePrincipals
func (m *SynchronizationJobItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property jobs for servicePrincipals
func (m *SynchronizationJobItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *SynchronizationJobItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SynchronizationJobItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *SynchronizationJobItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SynchronizationJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SynchronizationJobItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property jobs in servicePrincipals
func (m *SynchronizationJobItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, requestConfiguration *SynchronizationJobItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property jobs for servicePrincipals
func (m *SynchronizationJobItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SynchronizationJobItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *SynchronizationJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SynchronizationJobItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *SynchronizationJobItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, requestConfiguration *SynchronizationJobItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Pause the pause property
func (m *SynchronizationJobItemRequestBuilder) Pause()(*id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7.PauseRequestBuilder) {
    return id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisionOnDemand the provisionOnDemand property
func (m *SynchronizationJobItemRequestBuilder) ProvisionOnDemand()(*i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5.ProvisionOnDemandRequestBuilder) {
    return i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5.NewProvisionOnDemandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restart the restart property
func (m *SynchronizationJobItemRequestBuilder) Restart()(*i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb.RestartRequestBuilder) {
    return i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb.NewRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schema the schema property
func (m *SynchronizationJobItemRequestBuilder) Schema()(*i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24.SchemaRequestBuilder) {
    return i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Start the start property
func (m *SynchronizationJobItemRequestBuilder) Start()(*ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643.StartRequestBuilder) {
    return ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop the stop property
func (m *SynchronizationJobItemRequestBuilder) Stop()(*i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7.StopRequestBuilder) {
    return i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateCredentials the validateCredentials property
func (m *SynchronizationJobItemRequestBuilder) ValidateCredentials()(*ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b.ValidateCredentialsRequestBuilder) {
    return ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b.NewValidateCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
