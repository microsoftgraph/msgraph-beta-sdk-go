package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i574815dbbb1e9c51a965e25539eed078ec9ab55919033c1c32dc721a519d9f21 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userattributeassignments"
    i7dcf98ba71a64ad6e1a7f1d9cbe306c7616a2b4fbec2dbc655c658192d35a36b "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages"
    i961eb1bb9c95ad9a1c53819fc521c8a4e4e3dd72967d787f44f08b9b8c796ad2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/identityproviders"
    icd6b8fe789e0ddefec348d941d6505a9be0443397887095b870d0a0a621d081a "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userflowidentityproviders"
    i45862c76d997425b3529626dabb7258b739baf239b9d5edccd7c60df97c16804 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item"
    i7d194a5a848cdd478db3ec5b0473df2d333b7d840995b2db46e55e8b8b38b213 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/identityproviders/item"
    i8706143552eeb09a05c35bc53efe01f6fe2a821a3d0568b7c706289ea9919d1d "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userattributeassignments/item"
    ic9e278b4a19aaf9a41eabab943955d599d211f64b6406753a09626e656545741 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userflowidentityproviders/item"
)

// B2xIdentityUserFlowItemRequestBuilder provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
type B2xIdentityUserFlowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// B2xIdentityUserFlowItemRequestBuilderDeleteOptions options for Delete
type B2xIdentityUserFlowItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// B2xIdentityUserFlowItemRequestBuilderGetOptions options for Get
type B2xIdentityUserFlowItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xIdentityUserFlowItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// B2xIdentityUserFlowItemRequestBuilderGetQueryParameters represents entry point for B2X/self-service sign-up identity userflows.
type B2xIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xIdentityUserFlowItemRequestBuilderPatchOptions options for Patch
type B2xIdentityUserFlowItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2xIdentityUserFlowable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewB2xIdentityUserFlowItemRequestBuilderInternal instantiates a new B2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xIdentityUserFlowItemRequestBuilder) {
    m := &B2xIdentityUserFlowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewB2xIdentityUserFlowItemRequestBuilder instantiates a new B2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xIdentityUserFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xIdentityUserFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property b2xUserFlows for identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateGetRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property b2xUserFlows in identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property b2xUserFlows for identity
func (m *B2xIdentityUserFlowItemRequestBuilder) Delete(options *B2xIdentityUserFlowItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowItemRequestBuilder) Get(options *B2xIdentityUserFlowItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2xIdentityUserFlowable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateB2xIdentityUserFlowFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2xIdentityUserFlowable), nil
}
// IdentityProviders the identityProviders property
func (m *B2xIdentityUserFlowItemRequestBuilder) IdentityProviders()(*i961eb1bb9c95ad9a1c53819fc521c8a4e4e3dd72967d787f44f08b9b8c796ad2.IdentityProvidersRequestBuilder) {
    return i961eb1bb9c95ad9a1c53819fc521c8a4e4e3dd72967d787f44f08b9b8c796ad2.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.identityProviders.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) IdentityProvidersById(id string)(*i7d194a5a848cdd478db3ec5b0473df2d333b7d840995b2db46e55e8b8b38b213.IdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider%2Did"] = id
    }
    return i7d194a5a848cdd478db3ec5b0473df2d333b7d840995b2db46e55e8b8b38b213.NewIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages the languages property
func (m *B2xIdentityUserFlowItemRequestBuilder) Languages()(*i7dcf98ba71a64ad6e1a7f1d9cbe306c7616a2b4fbec2dbc655c658192d35a36b.LanguagesRequestBuilder) {
    return i7dcf98ba71a64ad6e1a7f1d9cbe306c7616a2b4fbec2dbc655c658192d35a36b.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.languages.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) LanguagesById(id string)(*i45862c76d997425b3529626dabb7258b739baf239b9d5edccd7c60df97c16804.UserFlowLanguageConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration%2Did"] = id
    }
    return i45862c76d997425b3529626dabb7258b739baf239b9d5edccd7c60df97c16804.NewUserFlowLanguageConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property b2xUserFlows in identity
func (m *B2xIdentityUserFlowItemRequestBuilder) Patch(options *B2xIdentityUserFlowItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserAttributeAssignments the userAttributeAssignments property
func (m *B2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignments()(*i574815dbbb1e9c51a965e25539eed078ec9ab55919033c1c32dc721a519d9f21.UserAttributeAssignmentsRequestBuilder) {
    return i574815dbbb1e9c51a965e25539eed078ec9ab55919033c1c32dc721a519d9f21.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserAttributeAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.userAttributeAssignments.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignmentsById(id string)(*i8706143552eeb09a05c35bc53efe01f6fe2a821a3d0568b7c706289ea9919d1d.IdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment%2Did"] = id
    }
    return i8706143552eeb09a05c35bc53efe01f6fe2a821a3d0568b7c706289ea9919d1d.NewIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserFlowIdentityProviders the userFlowIdentityProviders property
func (m *B2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProviders()(*icd6b8fe789e0ddefec348d941d6505a9be0443397887095b870d0a0a621d081a.UserFlowIdentityProvidersRequestBuilder) {
    return icd6b8fe789e0ddefec348d941d6505a9be0443397887095b870d0a0a621d081a.NewUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowIdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.userFlowIdentityProviders.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProvidersById(id string)(*ic9e278b4a19aaf9a41eabab943955d599d211f64b6406753a09626e656545741.IdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return ic9e278b4a19aaf9a41eabab943955d599d211f64b6406753a09626e656545741.NewIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
