package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/identityproviders"
    ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userattributeassignments"
    ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userflowidentityproviders"
    ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages"
    i49a1823495467514804483eeb49dca5d9b4265432b5f032c862b8625c52a4c44 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/identityproviders/item"
    i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item"
    i9ea52c3311e45f7594a4b33eff17673fff89c980f459b1cb921c89a7ec842263 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userflowidentityproviders/item"
    ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userattributeassignments/item"
)

// B2cIdentityUserFlowItemRequestBuilder provides operations to manage the b2cUserFlows property of the microsoft.graph.identityContainer entity.
type B2cIdentityUserFlowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// B2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2cIdentityUserFlowItemRequestBuilderGetQueryParameters represents entry point for B2C identity userflows.
type B2cIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cIdentityUserFlowItemRequestBuilderGetQueryParameters
}
// B2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cIdentityUserFlowItemRequestBuilderInternal instantiates a new B2cIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2cIdentityUserFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cIdentityUserFlowItemRequestBuilder) {
    m := &B2cIdentityUserFlowItemRequestBuilder{
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
// NewB2cIdentityUserFlowItemRequestBuilder instantiates a new B2cIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2cIdentityUserFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cIdentityUserFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cIdentityUserFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property b2cUserFlows for identity
func (m *B2cIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property b2cUserFlows for identity
func (m *B2cIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *B2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *B2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property b2cUserFlows in identity
func (m *B2cIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property b2cUserFlows in identity
func (m *B2cIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, requestConfiguration *B2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property b2cUserFlows for identity
func (m *B2cIdentityUserFlowItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *B2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property b2cUserFlows for identity
func (m *B2cIdentityUserFlowItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *B2cIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) GetWithResponseHandler(requestConfiguration *B2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) GetWithResponseHandler(requestConfiguration *B2cIdentityUserFlowItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateB2cIdentityUserFlowFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable), nil
}
// IdentityProviders the identityProviders property
func (m *B2cIdentityUserFlowItemRequestBuilder) IdentityProviders()(*i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df.IdentityProvidersRequestBuilder) {
    return i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.identityProviders.item collection
func (m *B2cIdentityUserFlowItemRequestBuilder) IdentityProvidersById(id string)(*i49a1823495467514804483eeb49dca5d9b4265432b5f032c862b8625c52a4c44.IdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider%2Did"] = id
    }
    return i49a1823495467514804483eeb49dca5d9b4265432b5f032c862b8625c52a4c44.NewIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages the languages property
func (m *B2cIdentityUserFlowItemRequestBuilder) Languages()(*ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be.LanguagesRequestBuilder) {
    return ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.languages.item collection
func (m *B2cIdentityUserFlowItemRequestBuilder) LanguagesById(id string)(*i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c.UserFlowLanguageConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration%2Did"] = id
    }
    return i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c.NewUserFlowLanguageConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property b2cUserFlows in identity
func (m *B2cIdentityUserFlowItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, requestConfiguration *B2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property b2cUserFlows in identity
func (m *B2cIdentityUserFlowItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2cIdentityUserFlowable, requestConfiguration *B2cIdentityUserFlowItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// UserAttributeAssignments the userAttributeAssignments property
func (m *B2cIdentityUserFlowItemRequestBuilder) UserAttributeAssignments()(*ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576.UserAttributeAssignmentsRequestBuilder) {
    return ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserAttributeAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.userAttributeAssignments.item collection
func (m *B2cIdentityUserFlowItemRequestBuilder) UserAttributeAssignmentsById(id string)(*ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888.IdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment%2Did"] = id
    }
    return ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888.NewIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserFlowIdentityProviders the userFlowIdentityProviders property
func (m *B2cIdentityUserFlowItemRequestBuilder) UserFlowIdentityProviders()(*ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed.UserFlowIdentityProvidersRequestBuilder) {
    return ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed.NewUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowIdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.userFlowIdentityProviders.item collection
func (m *B2cIdentityUserFlowItemRequestBuilder) UserFlowIdentityProvidersById(id string)(*i9ea52c3311e45f7594a4b33eff17673fff89c980f459b1cb921c89a7ec842263.IdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return i9ea52c3311e45f7594a4b33eff17673fff89c980f459b1cb921c89a7ec842263.NewIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
