package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder provides operations to call the updateAllowedCombinations method.
type IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilderInternal instantiates a new UpdateAllowedCombinationsRequestBuilder and sets the default values.
func NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder) {
    m := &IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess/authenticationStrengths/policies/{authenticationStrengthPolicy%2Did}/microsoft.graph.updateAllowedCombinations";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder instantiates a new UpdateAllowedCombinationsRequestBuilder and sets the default values.
func NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation update the allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBodyable, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post update the allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
func (m *IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBodyable, requestConfiguration *IdentityConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateAllowedCombinationsResultable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUpdateAllowedCombinationsResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateAllowedCombinationsResultable), nil
}
