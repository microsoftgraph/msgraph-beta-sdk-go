package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder provides operations to call the updateIndex method.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal instantiates a new CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    m := &CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/microsoft.graph.security.updateIndex", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder instantiates a new CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateIndex
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexUpdateIndexPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action updateIndex
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexUpdateIndexPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphSecurityUpdateIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
