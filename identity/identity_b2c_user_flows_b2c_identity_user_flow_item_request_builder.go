package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder provides operations to manage the b2cUserFlows property of the microsoft.graph.identityContainer entity.
type IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderGetQueryParameters represents entry point for B2C identity userflows.
type IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderGetQueryParameters
}
// IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderInternal instantiates a new B2cIdentityUserFlowItemRequestBuilder and sets the default values.
func NewIdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) {
    m := &IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder instantiates a new B2cIdentityUserFlowItemRequestBuilder and sets the default values.
func NewIdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property b2cUserFlows for identity
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents entry point for B2C identity userflows.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property b2cUserFlows in identity
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, requestConfiguration *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property b2cUserFlows for identity
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents entry point for B2C identity userflows.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateB2cIdentityUserFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable), nil
}
// IdentityProviders provides operations to manage the identityProviders property of the microsoft.graph.b2cIdentityUserFlow entity.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) IdentityProviders()(*IdentityB2cUserFlowsItemIdentityProvidersRequestBuilder) {
    return NewIdentityB2cUserFlowsItemIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.identityProviders.item collection
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) IdentityProvidersById(id string)(*IdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider%2Did"] = id
    }
    return NewIdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages provides operations to manage the languages property of the microsoft.graph.b2cIdentityUserFlow entity.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) Languages()(*IdentityB2cUserFlowsItemLanguagesRequestBuilder) {
    return NewIdentityB2cUserFlowsItemLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById provides operations to manage the languages property of the microsoft.graph.b2cIdentityUserFlow entity.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) LanguagesById(id string)(*IdentityB2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration%2Did"] = id
    }
    return NewIdentityB2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property b2cUserFlows in identity
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, requestConfiguration *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateB2cIdentityUserFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable), nil
}
// UserAttributeAssignments provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2cIdentityUserFlow entity.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) UserAttributeAssignments()(*IdentityB2cUserFlowsItemUserAttributeAssignmentsRequestBuilder) {
    return NewIdentityB2cUserFlowsItemUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserAttributeAssignmentsById provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2cIdentityUserFlow entity.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) UserAttributeAssignmentsById(id string)(*IdentityB2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment%2Did"] = id
    }
    return NewIdentityB2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserFlowIdentityProviders provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2cIdentityUserFlow entity.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) UserFlowIdentityProviders()(*IdentityB2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) {
    return NewIdentityB2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowIdentityProvidersById provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2cIdentityUserFlow entity.
func (m *IdentityB2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) UserFlowIdentityProvidersById(id string)(*IdentityB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return NewIdentityB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
