package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder provides operations to manage the templates property of the microsoft.graph.synchronization entity.
type ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderGetQueryParameters pre-configured synchronization settings for a particular application.
type ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderGetQueryParameters
}
// ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderInternal instantiates a new SynchronizationTemplateItemRequestBuilder and sets the default values.
func NewServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) {
    m := &ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/templates/{synchronizationTemplate%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder instantiates a new SynchronizationTemplateItemRequestBuilder and sets the default values.
func NewServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property templates for servicePrincipals
func (m *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation pre-configured synchronization settings for a particular application.
func (m *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property templates in servicePrincipals
func (m *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationTemplateable, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property templates for servicePrincipals
func (m *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get pre-configured synchronization settings for a particular application.
func (m *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationTemplateable), nil
}
// Patch update the navigation property templates in servicePrincipals
func (m *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationTemplateable, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationTemplateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationTemplateable), nil
}
// Schema provides operations to manage the schema property of the microsoft.graph.synchronizationTemplate entity.
func (m *ServicePrincipalsItemSynchronizationTemplatesSynchronizationTemplateItemRequestBuilder) Schema()(*ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
