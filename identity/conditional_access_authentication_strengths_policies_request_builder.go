package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder provides operations to manage the policies property of the microsoft.graph.authenticationStrengthRoot entity.
type ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderGetQueryParameters a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
type ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderGetQueryParameters
}
// ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationStrengthPolicyId provides operations to manage the policies property of the microsoft.graph.authenticationStrengthRoot entity.
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
func (m *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) ByAuthenticationStrengthPolicyId(authenticationStrengthPolicyId string)(*ConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationStrengthPolicyId != "" {
        urlTplParams["authenticationStrengthPolicy%2Did"] = authenticationStrengthPolicyId
    }
    return NewConditionalAccessAuthenticationStrengthsPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderInternal instantiates a new PoliciesRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) {
    m := &ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrengths/policies{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder instantiates a new PoliciesRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) Count()(*ConditionalAccessAuthenticationStrengthsPoliciesCountRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthsPoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FindByMethodModeWithAuthenticationMethodModes provides operations to call the findByMethodMode method.
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
func (m *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) FindByMethodModeWithAuthenticationMethodModes(authenticationMethodModes *string)(*ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, authenticationMethodModes)
}
// Get a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
func (m *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyCollectionResponseable), nil
}
// Post create new navigation property to policies for identity
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
func (m *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, requestConfiguration *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable), nil
}
// ToGetRequestInformation a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
func (m *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPostRequestInformation create new navigation property to policies for identity
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
func (m *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthPolicyable, requestConfiguration *ConditionalAccessAuthenticationStrengthsPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
