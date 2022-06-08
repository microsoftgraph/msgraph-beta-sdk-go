package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i38edbc2c15228f0a4386c06fe1cd74f5b61b0532c060a337d0d143cee71f76dc "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item/overridespages"
    id8b4aa20110754f7a6332241da005226519c2fdc8bae4a32149df9d988ded399 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item/defaultpages"
    id97df1ec1b4e9135f6abf467520dc7657ad6787d1073dfe17aa2b9434dc66f74 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item/overridespages/item"
    if7f959c11e06f172ec88344202e39c37c579835cbcd855889fc28d546c7c6126 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item/defaultpages/item"
)

// UserFlowLanguageConfigurationItemRequestBuilder provides operations to manage the languages property of the microsoft.graph.b2xIdentityUserFlow entity.
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
// UserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign up user flow. You cannot create custom languages in self-service sign up user flows.
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
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign up user flow. You cannot create custom languages in self-service sign up user flows.
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign up user flow. You cannot create custom languages in self-service sign up user flows.
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
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DefaultPages the defaultPages property
func (m *UserFlowLanguageConfigurationItemRequestBuilder) DefaultPages()(*id8b4aa20110754f7a6332241da005226519c2fdc8bae4a32149df9d988ded399.DefaultPagesRequestBuilder) {
    return id8b4aa20110754f7a6332241da005226519c2fdc8bae4a32149df9d988ded399.NewDefaultPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultPagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.languages.item.defaultPages.item collection
func (m *UserFlowLanguageConfigurationItemRequestBuilder) DefaultPagesById(id string)(*if7f959c11e06f172ec88344202e39c37c579835cbcd855889fc28d546c7c6126.UserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage%2Did"] = id
    }
    return if7f959c11e06f172ec88344202e39c37c579835cbcd855889fc28d546c7c6126.NewUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property languages for identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property languages for identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign up user flow. You cannot create custom languages in self-service sign up user flows.
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign up user flow. You cannot create custom languages in self-service sign up user flows.
func (m *UserFlowLanguageConfigurationItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguageConfigurationFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable), nil
}
// OverridesPages the overridesPages property
func (m *UserFlowLanguageConfigurationItemRequestBuilder) OverridesPages()(*i38edbc2c15228f0a4386c06fe1cd74f5b61b0532c060a337d0d143cee71f76dc.OverridesPagesRequestBuilder) {
    return i38edbc2c15228f0a4386c06fe1cd74f5b61b0532c060a337d0d143cee71f76dc.NewOverridesPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OverridesPagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.languages.item.overridesPages.item collection
func (m *UserFlowLanguageConfigurationItemRequestBuilder) OverridesPagesById(id string)(*id97df1ec1b4e9135f6abf467520dc7657ad6787d1073dfe17aa2b9434dc66f74.UserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage%2Did"] = id
    }
    return id97df1ec1b4e9135f6abf467520dc7657ad6787d1073dfe17aa2b9434dc66f74.NewUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property languages in identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property languages in identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguageConfigurationable, requestConfiguration *UserFlowLanguageConfigurationItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
