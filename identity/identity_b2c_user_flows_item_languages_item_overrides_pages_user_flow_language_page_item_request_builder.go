package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder provides operations to manage the overridesPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
type IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
type IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters
}
// IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderInternal instantiates a new UserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewIdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) {
    m := &IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/overridesPages/{userFlowLanguagePage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder instantiates a new UserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewIdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the identityContainer entity.
func (m *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) Content()(*IdentityB2cUserFlowsItemLanguagesItemOverridesPagesItemValueContentRequestBuilder) {
    return NewIdentityB2cUserFlowsItemLanguagesItemOverridesPagesItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property overridesPages for identity
func (m *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
func (m *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property overridesPages in identity
func (m *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, requestConfiguration *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property overridesPages for identity
func (m *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
func (m *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, error) {
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
// Patch update the navigation property overridesPages in identity
func (m *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, requestConfiguration *IdentityB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, error) {
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
