package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder provides operations to call the creationOptions method.
type ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderGetQueryParameters retrieve creation options required to generate and register a Microsoft Entra ID-compatible passkey. Self-service operations aren't supported.  
type ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderGetQueryParameters struct {
    // Usage: challengeTimeoutInMinutes=@challengeTimeoutInMinutes
    ChallengeTimeoutInMinutes *int32 `uriparametername:"challengeTimeoutInMinutes"`
}
// ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderGetQueryParameters
}
// NewItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderInternal instantiates a new ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder and sets the default values.
func NewItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder) {
    m := &ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/fido2Methods/creationOptions(challengeTimeoutInMinutes=@challengeTimeoutInMinutes){?challengeTimeoutInMinutes*}", pathParameters),
    }
    return m
}
// NewItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder instantiates a new ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder and sets the default values.
func NewItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve creation options required to generate and register a Microsoft Entra ID-compatible passkey. Self-service operations aren't supported.  
// returns a WebauthnCredentialCreationOptionsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/fido2authenticationmethod-creationoptions?view=graph-rest-beta
func (m *ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebauthnCredentialCreationOptionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWebauthnCredentialCreationOptionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebauthnCredentialCreationOptionsable), nil
}
// ToGetRequestInformation retrieve creation options required to generate and register a Microsoft Entra ID-compatible passkey. Self-service operations aren't supported.  
// returns a *RequestInformation when successful
func (m *ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder when successful
func (m *ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder) {
    return NewItemAuthenticationFido2MethodsCreationOptionsWithChallengeTimeoutInMinutesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
