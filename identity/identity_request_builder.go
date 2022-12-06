package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i11f22886c33eed99dd85e4d242bbf0976ec471e9e441cfa6a58b5864040bd094 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/apiconnectors"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5f1ee64209aad8370602f56501934370194f1248216e73d51b8543369c558033 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/continuousaccessevaluationpolicy"
    i6635ff8f2aa1c074739a0c7312e2c31642186b9aadaa7290a91096bd632e192d "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows"
    i96c0428619bb9ada314a888efda4bc4c3bb5c55513e2cfaa28f956e608f98c45 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/userflows"
    iaa68fab70e5038f618ea9044bb8a9fbc651b67c06a35357c7b18c20d63880c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/userflowattributes"
    ic34e66a722321196953b295f6b75e587885dec11e8a616d1a3f940a14b88d50e "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess"
    icdc246cecb641b30952ec0f2ae5ae755c0633e4f5bb05dbfce622920c1245205 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/customauthenticationextensions"
    ie614621c7e356610a3cda4707ec623ea542f5d84eaede8bc6ac868a655447bc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows"
    ief24c82e2119211d350c80136b97d6bf10b3d0e78a9681394a143089f017cd94 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/authenticationeventlisteners"
    if8e62443e000757279f6963237e640c3a2f6bf54d567fb126a582ea7bac48a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/identityproviders"
    i056de954ae8d2bd3c37cd62ab74e2a653c76cc631e79df2abe7d5423b433ad7f "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item"
    i4064b0ed26ff7ac321cce776d8d90a08dbeca0d51fec9f75ba0ee4621fd09ffa "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/identityproviders/item"
    i5948d5a09fb22b5b488a63253dae6b777e5fbf71f13cd243e618c0155cad10c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/customauthenticationextensions/item"
    i66a685596e55f332f9dea7615db3653de9c5945d4f7b70187e76b70ed8344761 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item"
    i7551a396b27b3ddf77a579c78cd05a221708f29814c2e58a262ae86fdfccebc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/authenticationeventlisteners/item"
    i80b7576b721756e8211e557cb66927cff088fae6e30a6b5990e694f8c8938b4d "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/userflows/item"
    i96fbb19ece75424f3a1291715def826b0b7091a7b004df122cb970e740300385 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/userflowattributes/item"
    id9b7138b09b41c1638c96829912584afe356128a86a38484a5041a342124094b "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/apiconnectors/item"
)

// IdentityRequestBuilder provides operations to manage the identityContainer singleton.
type IdentityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityRequestBuilderGetQueryParameters get identity
type IdentityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityRequestBuilderGetQueryParameters
}
// IdentityRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApiConnectors provides operations to manage the apiConnectors property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ApiConnectors()(*i11f22886c33eed99dd85e4d242bbf0976ec471e9e441cfa6a58b5864040bd094.ApiConnectorsRequestBuilder) {
    return i11f22886c33eed99dd85e4d242bbf0976ec471e9e441cfa6a58b5864040bd094.NewApiConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiConnectorsById provides operations to manage the apiConnectors property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ApiConnectorsById(id string)(*id9b7138b09b41c1638c96829912584afe356128a86a38484a5041a342124094b.IdentityApiConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityApiConnector%2Did"] = id
    }
    return id9b7138b09b41c1638c96829912584afe356128a86a38484a5041a342124094b.NewIdentityApiConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationEventListeners provides operations to manage the authenticationEventListeners property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) AuthenticationEventListeners()(*ief24c82e2119211d350c80136b97d6bf10b3d0e78a9681394a143089f017cd94.AuthenticationEventListenersRequestBuilder) {
    return ief24c82e2119211d350c80136b97d6bf10b3d0e78a9681394a143089f017cd94.NewAuthenticationEventListenersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationEventListenersById provides operations to manage the authenticationEventListeners property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) AuthenticationEventListenersById(id string)(*i7551a396b27b3ddf77a579c78cd05a221708f29814c2e58a262ae86fdfccebc3.AuthenticationEventListenerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationEventListener%2Did"] = id
    }
    return i7551a396b27b3ddf77a579c78cd05a221708f29814c2e58a262ae86fdfccebc3.NewAuthenticationEventListenerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// B2cUserFlows provides operations to manage the b2cUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2cUserFlows()(*i6635ff8f2aa1c074739a0c7312e2c31642186b9aadaa7290a91096bd632e192d.B2cUserFlowsRequestBuilder) {
    return i6635ff8f2aa1c074739a0c7312e2c31642186b9aadaa7290a91096bd632e192d.NewB2cUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// B2cUserFlowsById provides operations to manage the b2cUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2cUserFlowsById(id string)(*i66a685596e55f332f9dea7615db3653de9c5945d4f7b70187e76b70ed8344761.B2cIdentityUserFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["b2cIdentityUserFlow%2Did"] = id
    }
    return i66a685596e55f332f9dea7615db3653de9c5945d4f7b70187e76b70ed8344761.NewB2cIdentityUserFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// B2xUserFlows provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2xUserFlows()(*ie614621c7e356610a3cda4707ec623ea542f5d84eaede8bc6ac868a655447bc5.B2xUserFlowsRequestBuilder) {
    return ie614621c7e356610a3cda4707ec623ea542f5d84eaede8bc6ac868a655447bc5.NewB2xUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// B2xUserFlowsById provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2xUserFlowsById(id string)(*i056de954ae8d2bd3c37cd62ab74e2a653c76cc631e79df2abe7d5423b433ad7f.B2xIdentityUserFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["b2xIdentityUserFlow%2Did"] = id
    }
    return i056de954ae8d2bd3c37cd62ab74e2a653c76cc631e79df2abe7d5423b433ad7f.NewB2xIdentityUserFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccess provides operations to manage the conditionalAccess property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ConditionalAccess()(*ic34e66a722321196953b295f6b75e587885dec11e8a616d1a3f940a14b88d50e.ConditionalAccessRequestBuilder) {
    return ic34e66a722321196953b295f6b75e587885dec11e8a616d1a3f940a14b88d50e.NewConditionalAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityRequestBuilderInternal instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityRequestBuilder) {
    m := &IdentityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityRequestBuilder instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityRequestBuilderInternal(urlParams, requestAdapter)
}
// ContinuousAccessEvaluationPolicy provides operations to manage the continuousAccessEvaluationPolicy property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ContinuousAccessEvaluationPolicy()(*i5f1ee64209aad8370602f56501934370194f1248216e73d51b8543369c558033.ContinuousAccessEvaluationPolicyRequestBuilder) {
    return i5f1ee64209aad8370602f56501934370194f1248216e73d51b8543369c558033.NewContinuousAccessEvaluationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get identity
func (m *IdentityRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update identity
func (m *IdentityRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, requestConfiguration *IdentityRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CustomAuthenticationExtensions provides operations to manage the customAuthenticationExtensions property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) CustomAuthenticationExtensions()(*icdc246cecb641b30952ec0f2ae5ae755c0633e4f5bb05dbfce622920c1245205.CustomAuthenticationExtensionsRequestBuilder) {
    return icdc246cecb641b30952ec0f2ae5ae755c0633e4f5bb05dbfce622920c1245205.NewCustomAuthenticationExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomAuthenticationExtensionsById provides operations to manage the customAuthenticationExtensions property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) CustomAuthenticationExtensionsById(id string)(*i5948d5a09fb22b5b488a63253dae6b777e5fbf71f13cd243e618c0155cad10c1.CustomAuthenticationExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customAuthenticationExtension%2Did"] = id
    }
    return i5948d5a09fb22b5b488a63253dae6b777e5fbf71f13cd243e618c0155cad10c1.NewCustomAuthenticationExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get identity
func (m *IdentityRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable), nil
}
// IdentityProviders provides operations to manage the identityProviders property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) IdentityProviders()(*if8e62443e000757279f6963237e640c3a2f6bf54d567fb126a582ea7bac48a91.IdentityProvidersRequestBuilder) {
    return if8e62443e000757279f6963237e640c3a2f6bf54d567fb126a582ea7bac48a91.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById provides operations to manage the identityProviders property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) IdentityProvidersById(id string)(*i4064b0ed26ff7ac321cce776d8d90a08dbeca0d51fec9f75ba0ee4621fd09ffa.IdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return i4064b0ed26ff7ac321cce776d8d90a08dbeca0d51fec9f75ba0ee4621fd09ffa.NewIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update identity
func (m *IdentityRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, requestConfiguration *IdentityRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable), nil
}
// UserFlowAttributes provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlowAttributes()(*iaa68fab70e5038f618ea9044bb8a9fbc651b67c06a35357c7b18c20d63880c85.UserFlowAttributesRequestBuilder) {
    return iaa68fab70e5038f618ea9044bb8a9fbc651b67c06a35357c7b18c20d63880c85.NewUserFlowAttributesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowAttributesById provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlowAttributesById(id string)(*i96fbb19ece75424f3a1291715def826b0b7091a7b004df122cb970e740300385.IdentityUserFlowAttributeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttribute%2Did"] = id
    }
    return i96fbb19ece75424f3a1291715def826b0b7091a7b004df122cb970e740300385.NewIdentityUserFlowAttributeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserFlows provides operations to manage the userFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlows()(*i96c0428619bb9ada314a888efda4bc4c3bb5c55513e2cfaa28f956e608f98c45.UserFlowsRequestBuilder) {
    return i96c0428619bb9ada314a888efda4bc4c3bb5c55513e2cfaa28f956e608f98c45.NewUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowsById provides operations to manage the userFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlowsById(id string)(*i80b7576b721756e8211e557cb66927cff088fae6e30a6b5990e694f8c8938b4d.IdentityUserFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlow%2Did"] = id
    }
    return i80b7576b721756e8211e557cb66927cff088fae6e30a6b5990e694f8c8938b4d.NewIdentityUserFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
