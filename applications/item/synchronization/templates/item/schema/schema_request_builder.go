package schema

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3f8397838cdcbf421c948cec12bb6c8ae29443adfbfffafafcb0800257df0249 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/directories"
    i5b8f0ae1eb8c05092bdc285776ceda0a64242dbf7a6bd4a2dea538cb5922a020 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/parseexpression"
    i60bb4f0c9eb8dde47c85b8b4f412b0bca85a017fe9fc37056e7b1b55f179d459 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/filteroperators"
    ic3d9837872d528f66b24cce114afdbd85a16641424441ba4d77e26ac34ff3528 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/functions"
    iffe98a2db6c23deb1db6d3a5137091a707788c003479b60e45e086d4e05d1658 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/directories/item"
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
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/synchronization/templates/{synchronizationTemplate%2Did}/schema{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property schema for applications
func (m *SchemaRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property schema for applications
func (m *SchemaRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *SchemaRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SchemaRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schema in applications
func (m *SchemaRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property schema in applications
func (m *SchemaRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, requestConfiguration *SchemaRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property schema for applications
func (m *SchemaRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchemaRequestBuilderDeleteRequestConfiguration)(error) {
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
// Directories the directories property
func (m *SchemaRequestBuilder) Directories()(*i3f8397838cdcbf421c948cec12bb6c8ae29443adfbfffafafcb0800257df0249.DirectoriesRequestBuilder) {
    return i3f8397838cdcbf421c948cec12bb6c8ae29443adfbfffafafcb0800257df0249.NewDirectoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.synchronization.templates.item.schema.directories.item collection
func (m *SchemaRequestBuilder) DirectoriesById(id string)(*iffe98a2db6c23deb1db6d3a5137091a707788c003479b60e45e086d4e05d1658.DirectoryDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryDefinition%2Did"] = id
    }
    return iffe98a2db6c23deb1db6d3a5137091a707788c003479b60e45e086d4e05d1658.NewDirectoryDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FilterOperators provides operations to call the filterOperators method.
func (m *SchemaRequestBuilder) FilterOperators()(*i60bb4f0c9eb8dde47c85b8b4f412b0bca85a017fe9fc37056e7b1b55f179d459.FilterOperatorsRequestBuilder) {
    return i60bb4f0c9eb8dde47c85b8b4f412b0bca85a017fe9fc37056e7b1b55f179d459.NewFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Functions provides operations to call the functions method.
func (m *SchemaRequestBuilder) Functions()(*ic3d9837872d528f66b24cce114afdbd85a16641424441ba4d77e26ac34ff3528.FunctionsRequestBuilder) {
    return ic3d9837872d528f66b24cce114afdbd85a16641424441ba4d77e26ac34ff3528.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// ParseExpression the parseExpression property
func (m *SchemaRequestBuilder) ParseExpression()(*i5b8f0ae1eb8c05092bdc285776ceda0a64242dbf7a6bd4a2dea538cb5922a020.ParseExpressionRequestBuilder) {
    return i5b8f0ae1eb8c05092bdc285776ceda0a64242dbf7a6bd4a2dea538cb5922a020.NewParseExpressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property schema in applications
func (m *SchemaRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSchemaable, requestConfiguration *SchemaRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
