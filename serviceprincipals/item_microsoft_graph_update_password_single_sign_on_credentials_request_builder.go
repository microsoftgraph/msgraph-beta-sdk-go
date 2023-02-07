package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder provides operations to call the updatePasswordSingleSignOnCredentials method.
type ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new MicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/microsoft.graph.updatePasswordSingleSignOnCredentials";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder instantiates a new MicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update single sign-on credentials using a password for a user or group.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceprincipal-updatepasswordsinglesignoncredentials?view=graph-rest-1.0
func (m *ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder) Post(ctx context.Context, body ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsUpdatePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation update single sign-on credentials using a password for a user or group.
func (m *ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsUpdatePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphUpdatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
