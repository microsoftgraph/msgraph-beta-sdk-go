package item

import (
    "context"
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
// B2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xIdentityUserFlowItemRequestBuilderGetQueryParameters represents entry point for B2X and self-service sign-up identity userflows.
type B2xIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xIdentityUserFlowItemRequestBuilderGetQueryParameters
}
// B2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
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
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property b2xUserFlows for identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *B2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents entry point for B2X and self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration represents entry point for B2X and self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *B2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property b2xUserFlows in identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2xIdentityUserFlowable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property b2xUserFlows in identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2xIdentityUserFlowable, requestConfiguration *B2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property b2xUserFlows for identity
func (m *B2xIdentityUserFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents entry point for B2X and self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2xIdentityUserFlowable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateB2xIdentityUserFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
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
func (m *B2xIdentityUserFlowItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.B2xIdentityUserFlowable, requestConfiguration *B2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(error) {
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
