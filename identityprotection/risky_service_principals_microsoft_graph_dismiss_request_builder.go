package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder provides operations to call the dismiss method.
type RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyServicePrincipalsMicrosoftGraphDismissRequestBuilderInternal instantiates a new MicrosoftGraphDismissRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsMicrosoftGraphDismissRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder) {
    m := &RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyServicePrincipals/microsoft.graph.dismiss";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewRiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder instantiates a new MicrosoftGraphDismissRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyServicePrincipalsMicrosoftGraphDismissRequestBuilderInternal(urlParams, requestAdapter)
}
// Post dismiss the risk of one or more riskyServicePrincipal objects. This action sets the targeted service principal account's risk level to `none`. You can dismiss up to 60 service principal accounts in one request.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/riskyserviceprincipal-dismiss?view=graph-rest-1.0
func (m *RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder) Post(ctx context.Context, body RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBodyable, requestConfiguration *RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation dismiss the risk of one or more riskyServicePrincipal objects. This action sets the targeted service principal account's risk level to `none`. You can dismiss up to 60 service principal accounts in one request.
func (m *RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilder) ToPostRequestInformation(ctx context.Context, body RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBodyable, requestConfiguration *RiskyServicePrincipalsMicrosoftGraphDismissRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
