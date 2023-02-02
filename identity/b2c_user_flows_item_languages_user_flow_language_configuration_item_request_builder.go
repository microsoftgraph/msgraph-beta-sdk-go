package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder provides operations to manage the languages property of the microsoft.graph.b2cIdentityUserFlow entity.
type B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters the languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
type B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters
}
// B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderInternal instantiates a new UserFlowLanguageConfigurationItemRequestBuilder and sets the default values.
func NewB2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) {
    m := &B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewB2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder instantiates a new UserFlowLanguageConfigurationItemRequestBuilder and sets the default values.
func NewB2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultPages provides operations to manage the defaultPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) DefaultPages()(*B2cUserFlowsItemLanguagesItemDefaultPagesRequestBuilder) {
    return NewB2cUserFlowsItemLanguagesItemDefaultPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DefaultPagesById provides operations to manage the defaultPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) DefaultPagesById(id string)(*B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage%2Did"] = id
    }
    return NewB2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete navigation property languages for identity
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguageConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable), nil
}
// OverridesPages provides operations to manage the overridesPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) OverridesPages()(*B2cUserFlowsItemLanguagesItemOverridesPagesRequestBuilder) {
    return NewB2cUserFlowsItemLanguagesItemOverridesPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OverridesPagesById provides operations to manage the overridesPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) OverridesPagesById(id string)(*B2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage%2Did"] = id
    }
    return NewB2cUserFlowsItemLanguagesItemOverridesPagesUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property languages in identity
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, requestConfiguration *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguageConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property languages for identity
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property languages in identity
func (m *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, requestConfiguration *B2cUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
