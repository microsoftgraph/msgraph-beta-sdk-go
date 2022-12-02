package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder provides operations to manage the schema property of the microsoft.graph.synchronizationTemplate entity.
type ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderGetQueryParameters default synchronization schema for the jobs based on this template.
type ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderGetQueryParameters
}
// ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderInternal instantiates a new SchemaRequestBuilder and sets the default values.
func NewServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) {
    m := &ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/templates/{synchronizationTemplate%2Did}/schema{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder instantiates a new SchemaRequestBuilder and sets the default values.
func NewServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schema for servicePrincipals
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation default synchronization schema for the jobs based on this template.
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schema in servicePrincipals
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property schema for servicePrincipals
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderDeleteRequestConfiguration)(error) {
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
// Directories provides operations to manage the directories property of the microsoft.graph.synchronizationSchema entity.
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) Directories()(*ServicePrincipalsItemSynchronizationTemplatesItemSchemaDirectoriesRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationTemplatesItemSchemaDirectoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoriesById provides operations to manage the directories property of the microsoft.graph.synchronizationSchema entity.
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) DirectoriesById(id string)(*ServicePrincipalsItemSynchronizationTemplatesItemSchemaDirectoriesDirectoryDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryDefinition%2Did"] = id
    }
    return NewServicePrincipalsItemSynchronizationTemplatesItemSchemaDirectoriesDirectoryDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FilterOperators provides operations to call the filterOperators method.
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) FilterOperators()(*ServicePrincipalsItemSynchronizationTemplatesItemSchemaFilterOperatorsRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationTemplatesItemSchemaFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Functions provides operations to call the functions method.
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) Functions()(*ServicePrincipalsItemSynchronizationTemplatesItemSchemaFunctionsRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationTemplatesItemSchemaFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get default synchronization schema for the jobs based on this template.
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationSchemaFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable), nil
}
// ParseExpression provides operations to call the parseExpression method.
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) ParseExpression()(*ServicePrincipalsItemSynchronizationTemplatesItemSchemaParseExpressionRequestBuilder) {
    return NewServicePrincipalsItemSynchronizationTemplatesItemSchemaParseExpressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property schema in servicePrincipals
func (m *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, requestConfiguration *ServicePrincipalsItemSynchronizationTemplatesItemSchemaRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationSchemaFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable), nil
}
