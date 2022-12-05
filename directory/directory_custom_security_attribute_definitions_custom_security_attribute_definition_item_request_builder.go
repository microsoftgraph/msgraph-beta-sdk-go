package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder provides operations to manage the customSecurityAttributeDefinitions property of the microsoft.graph.directory entity.
type DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderGetQueryParameters schema of a custom security attributes (key-value pairs).
type DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderGetQueryParameters
}
// DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllowedValues provides operations to manage the allowedValues property of the microsoft.graph.customSecurityAttributeDefinition entity.
func (m *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) AllowedValues()(*DirectoryCustomSecurityAttributeDefinitionsItemAllowedValuesRequestBuilder) {
    return NewDirectoryCustomSecurityAttributeDefinitionsItemAllowedValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllowedValuesById provides operations to manage the allowedValues property of the microsoft.graph.customSecurityAttributeDefinition entity.
func (m *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) AllowedValuesById(id string)(*DirectoryCustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["allowedValue%2Did"] = id
    }
    return NewDirectoryCustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderInternal instantiates a new CustomSecurityAttributeDefinitionItemRequestBuilder and sets the default values.
func NewDirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) {
    m := &DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/customSecurityAttributeDefinitions/{customSecurityAttributeDefinition%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder instantiates a new CustomSecurityAttributeDefinitionItemRequestBuilder and sets the default values.
func NewDirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customSecurityAttributeDefinitions for directory
func (m *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation schema of a custom security attributes (key-value pairs).
func (m *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property customSecurityAttributeDefinitions in directory
func (m *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomSecurityAttributeDefinitionable, requestConfiguration *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property customSecurityAttributeDefinitions for directory
func (m *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get schema of a custom security attributes (key-value pairs).
func (m *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomSecurityAttributeDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomSecurityAttributeDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomSecurityAttributeDefinitionable), nil
}
// Patch update the navigation property customSecurityAttributeDefinitions in directory
func (m *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomSecurityAttributeDefinitionable, requestConfiguration *DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomSecurityAttributeDefinitionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomSecurityAttributeDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomSecurityAttributeDefinitionable), nil
}
