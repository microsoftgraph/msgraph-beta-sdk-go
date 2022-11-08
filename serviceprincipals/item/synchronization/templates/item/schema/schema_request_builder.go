package schema

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/directories"
    i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/functions"
    i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/filteroperators"
    idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/parseexpression"
    i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/directories/item"
)

// SchemaRequestBuilder provides operations to manage the schema property of the microsoft.graph.synchronizationTemplate entity.
type SchemaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SchemaRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SchemaRequestBuilderGetQueryParameters default synchronization schema for the jobs based on this template.
type SchemaRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SchemaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchemaRequestBuilderGetQueryParameters
}
// SchemaRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaRequestBuilderInternal instantiates a new SchemaRequestBuilder and sets the default values.
func NewSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaRequestBuilder) {
    m := &SchemaRequestBuilder{
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
// NewSchemaRequestBuilder instantiates a new SchemaRequestBuilder and sets the default values.
func NewSchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schema for servicePrincipals
func (m *SchemaRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SchemaRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SchemaRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SchemaRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, requestConfiguration *SchemaRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SchemaRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchemaRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *SchemaRequestBuilder) Directories()(*i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a.DirectoriesRequestBuilder) {
    return i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a.NewDirectoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoriesById provides operations to manage the directories property of the microsoft.graph.synchronizationSchema entity.
func (m *SchemaRequestBuilder) DirectoriesById(id string)(*i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350.DirectoryDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryDefinition%2Did"] = id
    }
    return i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350.NewDirectoryDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FilterOperators provides operations to call the filterOperators method.
func (m *SchemaRequestBuilder) FilterOperators()(*i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178.FilterOperatorsRequestBuilder) {
    return i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178.NewFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Functions provides operations to call the functions method.
func (m *SchemaRequestBuilder) Functions()(*i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27.FunctionsRequestBuilder) {
    return i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, error) {
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
func (m *SchemaRequestBuilder) ParseExpression()(*idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e.ParseExpressionRequestBuilder) {
    return idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e.NewParseExpressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property schema in servicePrincipals
func (m *SchemaRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, requestConfiguration *SchemaRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, error) {
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
