package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder provides operations to manage the defaultPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
type IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
type IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters
}
// IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal instantiates a new UserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewIdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) {
    m := &IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/{userFlowLanguagePage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder instantiates a new UserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewIdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the identityContainer entity.
func (m *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Content()(*IdentityB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) {
    return NewIdentityB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property defaultPages for identity
func (m *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
func (m *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property defaultPages in identity
func (m *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, requestConfiguration *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property defaultPages for identity
func (m *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
func (m *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguagePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable), nil
}
// Patch update the navigation property defaultPages in identity
func (m *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, requestConfiguration *IdentityB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguagePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable), nil
}
