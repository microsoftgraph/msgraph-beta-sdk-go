package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder provides operations to manage the policies property of the microsoft.graph.authenticationStrengthRoot entity.
type IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
type IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters
}
// IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CombinationConfigurations provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CombinationConfigurations()(*IdentityConditionalAccessAuthenticationStrengthsPoliciesItemCombinationConfigurationsRequestBuilder) {
    return NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemCombinationConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CombinationConfigurationsById provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CombinationConfigurationsById(id string)(*IdentityConditionalAccessAuthenticationStrengthsPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationCombinationConfiguration%2Did"] = id
    }
    return NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal instantiates a new AuthenticationStrengthPolicyItemRequestBuilder and sets the default values.
func NewIdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    m := &IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess/authenticationStrengths/policies/{authenticationStrengthPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder instantiates a new AuthenticationStrengthPolicyItemRequestBuilder and sets the default values.
func NewIdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property policies for identity
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property policies in identity
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property policies for identity
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable), nil
}
// Patch update the navigation property policies in identity
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable), nil
}
// UpdateAllowedCombinations provides operations to call the updateAllowedCombinations method.
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) UpdateAllowedCombinations()(*IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder) {
    return NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Usage provides operations to call the usage method.
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Usage()(*IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUsageRequestBuilder) {
    return NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
