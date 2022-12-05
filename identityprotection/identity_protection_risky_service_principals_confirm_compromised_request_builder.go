package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder provides operations to call the confirmCompromised method.
type IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilderInternal instantiates a new ConfirmCompromisedRequestBuilder and sets the default values.
func NewIdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder) {
    m := &IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyServicePrincipals/microsoft.graph.confirmCompromised";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder instantiates a new ConfirmCompromisedRequestBuilder and sets the default values.
func NewIdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation confirm one or more riskyServicePrincipal objects as compromised. This action sets the targeted service principal account's risk level to `high`. When the risk level of the service principal is confirmed as compromised, the service principal object is disabled and its **disabledByMicrosoftStatus** property is updated.
func (m *IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProtectionRiskyServicePrincipalsConfirmCompromisedPostRequestBodyable, requestConfiguration *IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post confirm one or more riskyServicePrincipal objects as compromised. This action sets the targeted service principal account's risk level to `high`. When the risk level of the service principal is confirmed as compromised, the service principal object is disabled and its **disabledByMicrosoftStatus** property is updated.
func (m *IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProtectionRiskyServicePrincipalsConfirmCompromisedPostRequestBodyable, requestConfiguration *IdentityProtectionRiskyServicePrincipalsConfirmCompromisedRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
