package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphSecurityRunHuntingQueryRequestBuilder provides operations to call the runHuntingQuery method.
type MicrosoftGraphSecurityRunHuntingQueryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosoftGraphSecurityRunHuntingQueryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphSecurityRunHuntingQueryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphSecurityRunHuntingQueryRequestBuilderInternal instantiates a new MicrosoftGraphSecurityRunHuntingQueryRequestBuilder and sets the default values.
func NewMicrosoftGraphSecurityRunHuntingQueryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphSecurityRunHuntingQueryRequestBuilder) {
    m := &MicrosoftGraphSecurityRunHuntingQueryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/microsoft.graph.security.runHuntingQuery", pathParameters),
    }
    return m
}
// NewMicrosoftGraphSecurityRunHuntingQueryRequestBuilder instantiates a new MicrosoftGraphSecurityRunHuntingQueryRequestBuilder and sets the default values.
func NewMicrosoftGraphSecurityRunHuntingQueryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphSecurityRunHuntingQueryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphSecurityRunHuntingQueryRequestBuilderInternal(urlParams, requestAdapter)
}
// Post queries a specified set of event, activity, or entity data supported by Microsoft 365 Defender to proactively look for specific threats in your environment. This is the method for advanced hunting in Microsoft 365 Defender. This method includes a query in Kusto Query Language (KQL). It specifies a data table in the advanced hunting schema and a piped sequence of operators to filter or search that data, and format the query output in specific ways.  Find out more about hunting for threats across devices, emails, apps, and identities. Learn about KQL. For information on using advanced hunting in the Microsoft 365 Defender portal, see Proactively hunt for threats with advanced hunting in Microsoft 365 Defender.
// returns a HuntingQueryResultsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosoftGraphSecurityRunHuntingQueryRequestBuilder) Post(ctx context.Context, body MicrosoftGraphSecurityRunHuntingQueryRunHuntingQueryPostRequestBodyable, requestConfiguration *MicrosoftGraphSecurityRunHuntingQueryRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HuntingQueryResultsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateHuntingQueryResultsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HuntingQueryResultsable), nil
}
// ToPostRequestInformation queries a specified set of event, activity, or entity data supported by Microsoft 365 Defender to proactively look for specific threats in your environment. This is the method for advanced hunting in Microsoft 365 Defender. This method includes a query in Kusto Query Language (KQL). It specifies a data table in the advanced hunting schema and a piped sequence of operators to filter or search that data, and format the query output in specific ways.  Find out more about hunting for threats across devices, emails, apps, and identities. Learn about KQL. For information on using advanced hunting in the Microsoft 365 Defender portal, see Proactively hunt for threats with advanced hunting in Microsoft 365 Defender.
// returns a *RequestInformation when successful
func (m *MicrosoftGraphSecurityRunHuntingQueryRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosoftGraphSecurityRunHuntingQueryRunHuntingQueryPostRequestBodyable, requestConfiguration *MicrosoftGraphSecurityRunHuntingQueryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MicrosoftGraphSecurityRunHuntingQueryRequestBuilder when successful
func (m *MicrosoftGraphSecurityRunHuntingQueryRequestBuilder) WithUrl(rawUrl string)(*MicrosoftGraphSecurityRunHuntingQueryRequestBuilder) {
    return NewMicrosoftGraphSecurityRunHuntingQueryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
