package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic56e48b961e57de85d10390f870dd3e57722e4a44a0964781543b92374e68259 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item/defaultpages"
    iecf50a78d9d6bd3f01a4f78210ea956d05c66f139964a7db9869f81ca435eaca "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item/overridespages"
    i2fb4388026c86a21e8b8d1776d2c5f2b64cc93696f3f4d055161caa5dff7e785 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item/defaultpages/item"
    ia8cb0ef728fb55863ea7e003c3f99fbd388a2e1017ae424b5f63e5e4b1660d67 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item/overridespages/item"
)

// UserFlowLanguageConfigurationItemRequestBuilder provides operations to manage the languages property of the microsoft.graph.b2cIdentityUserFlow entity.
type UserFlowLanguageConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters the languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
type UserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters
}
// UserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserFlowLanguageConfigurationItemRequestBuilderInternal instantiates a new UserFlowLanguageConfigurationItemRequestBuilder and sets the default values.
func NewUserFlowLanguageConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserFlowLanguageConfigurationItemRequestBuilder) {
    m := &UserFlowLanguageConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserFlowLanguageConfigurationItemRequestBuilder instantiates a new UserFlowLanguageConfigurationItemRequestBuilder and sets the default values.
func NewUserFlowLanguageConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserFlowLanguageConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserFlowLanguageConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property languages for identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property languages for identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property languages in identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property languages in identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DefaultPages the defaultPages property
func (m *UserFlowLanguageConfigurationItemRequestBuilder) DefaultPages()(*ic56e48b961e57de85d10390f870dd3e57722e4a44a0964781543b92374e68259.DefaultPagesRequestBuilder) {
    return ic56e48b961e57de85d10390f870dd3e57722e4a44a0964781543b92374e68259.NewDefaultPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultPagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.languages.item.defaultPages.item collection
func (m *UserFlowLanguageConfigurationItemRequestBuilder) DefaultPagesById(id string)(*i2fb4388026c86a21e8b8d1776d2c5f2b64cc93696f3f4d055161caa5dff7e785.UserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage%2Did"] = id
    }
    return i2fb4388026c86a21e8b8d1776d2c5f2b64cc93696f3f4d055161caa5dff7e785.NewUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property languages for identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguageConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable), nil
}
// OverridesPages the overridesPages property
func (m *UserFlowLanguageConfigurationItemRequestBuilder) OverridesPages()(*iecf50a78d9d6bd3f01a4f78210ea956d05c66f139964a7db9869f81ca435eaca.OverridesPagesRequestBuilder) {
    return iecf50a78d9d6bd3f01a4f78210ea956d05c66f139964a7db9869f81ca435eaca.NewOverridesPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OverridesPagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.languages.item.overridesPages.item collection
func (m *UserFlowLanguageConfigurationItemRequestBuilder) OverridesPagesById(id string)(*ia8cb0ef728fb55863ea7e003c3f99fbd388a2e1017ae424b5f63e5e4b1660d67.UserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage%2Did"] = id
    }
    return ia8cb0ef728fb55863ea7e003c3f99fbd388a2e1017ae424b5f63e5e4b1660d67.NewUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property languages in identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguageConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable), nil
}
