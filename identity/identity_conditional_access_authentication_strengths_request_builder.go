package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityConditionalAccessAuthenticationStrengthsRequestBuilder provides operations to manage the authenticationStrengths property of the microsoft.graph.conditionalAccessRoot entity.
type IdentityConditionalAccessAuthenticationStrengthsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityConditionalAccessAuthenticationStrengthsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessAuthenticationStrengthsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityConditionalAccessAuthenticationStrengthsRequestBuilderGetQueryParameters defines the authentication strength policies, valid authentication method combinations, and authentication method mode details that can be required by a conditional access policy .
type IdentityConditionalAccessAuthenticationStrengthsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityConditionalAccessAuthenticationStrengthsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessAuthenticationStrengthsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityConditionalAccessAuthenticationStrengthsRequestBuilderGetQueryParameters
}
// IdentityConditionalAccessAuthenticationStrengthsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessAuthenticationStrengthsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationMethodModes provides operations to manage the authenticationMethodModes property of the microsoft.graph.authenticationStrengthRoot entity.
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) AuthenticationMethodModes()(*IdentityConditionalAccessAuthenticationStrengthsAuthenticationMethodModesRequestBuilder) {
    return NewIdentityConditionalAccessAuthenticationStrengthsAuthenticationMethodModesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodModesById provides operations to manage the authenticationMethodModes property of the microsoft.graph.authenticationStrengthRoot entity.
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) AuthenticationMethodModesById(id string)(*IdentityConditionalAccessAuthenticationStrengthsAuthenticationMethodModesAuthenticationMethodModeDetailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethodModeDetail%2Did"] = id
    }
    return NewIdentityConditionalAccessAuthenticationStrengthsAuthenticationMethodModesAuthenticationMethodModeDetailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityConditionalAccessAuthenticationStrengthsRequestBuilderInternal instantiates a new AuthenticationStrengthsRequestBuilder and sets the default values.
func NewIdentityConditionalAccessAuthenticationStrengthsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) {
    m := &IdentityConditionalAccessAuthenticationStrengthsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess/authenticationStrengths{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityConditionalAccessAuthenticationStrengthsRequestBuilder instantiates a new AuthenticationStrengthsRequestBuilder and sets the default values.
func NewIdentityConditionalAccessAuthenticationStrengthsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityConditionalAccessAuthenticationStrengthsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property authenticationStrengths for identity
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation defines the authentication strength policies, valid authentication method combinations, and authentication method mode details that can be required by a conditional access policy .
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property authenticationStrengths in identity
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property authenticationStrengths for identity
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get defines the authentication strength policies, valid authentication method combinations, and authentication method mode details that can be required by a conditional access policy .
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable), nil
}
// Patch update the navigation property authenticationStrengths in identity
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable), nil
}
// Policies provides operations to manage the policies property of the microsoft.graph.authenticationStrengthRoot entity.
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) Policies()(*IdentityConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) {
    return NewIdentityConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PoliciesById provides operations to manage the policies property of the microsoft.graph.authenticationStrengthRoot entity.
func (m *IdentityConditionalAccessAuthenticationStrengthsRequestBuilder) PoliciesById(id string)(*IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationStrengthPolicy%2Did"] = id
    }
    return NewIdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
