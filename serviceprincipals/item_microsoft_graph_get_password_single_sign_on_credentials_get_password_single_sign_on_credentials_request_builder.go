package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder provides operations to call the getPasswordSingleSignOnCredentials method.
type ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new GetPasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) {
    m := &ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/microsoft.graph.getPasswordSingleSignOnCredentials";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder instantiates a new GetPasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get a list of single sign-on credentials using a password for a user or group.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceprincipal-getpasswordsinglesignoncredentials?view=graph-rest-1.0
func (m *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) Post(ctx context.Context, body ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordSingleSignOnCredentialSetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePasswordSingleSignOnCredentialSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordSingleSignOnCredentialSetable), nil
}
// ToPostRequestInformation get a list of single sign-on credentials using a password for a user or group.
func (m *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
