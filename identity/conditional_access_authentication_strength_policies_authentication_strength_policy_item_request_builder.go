package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder provides operations to manage the policies property of the microsoft.graph.authenticationStrengthRoot entity.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters
}
// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CombinationConfigurations provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CombinationConfigurations()(*ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CombinationConfigurationsById provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CombinationConfigurationsById(id string)(*ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationCombinationConfiguration%2Did"] = id
    }
    return NewConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal instantiates a new AuthenticationStrengthPolicyItemRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    m := &ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder instantiates a new AuthenticationStrengthPolicyItemRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property policies for identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable), nil
}
// MicrosoftGraphUpdateAllowedCombinations provides operations to call the updateAllowedCombinations method.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) MicrosoftGraphUpdateAllowedCombinations()(*ConditionalAccessAuthenticationStrengthPoliciesItemMicrosoftGraphUpdateAllowedCombinationsRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthPoliciesItemMicrosoftGraphUpdateAllowedCombinationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUsage provides operations to call the usage method.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) MicrosoftGraphUsage()(*ConditionalAccessAuthenticationStrengthPoliciesItemMicrosoftGraphUsageRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthPoliciesItemMicrosoftGraphUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property policies in identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property policies for identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property policies in identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
