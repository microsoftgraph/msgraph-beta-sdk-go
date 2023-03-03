package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder provides operations to call the removeHold method.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilderInternal instantiates a new SecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder) {
    m := &CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/security.removeHold";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder instantiates a new SecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of removing hold from eDiscovery non-custodial data sources. After the operation is created, you can get the status by retrieving the `Location` parameter from the response headers. The location provides a URL that will return an eDiscoveryHoldOperation object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverynoncustodialdatasource-removehold?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRemoveHoldPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation start the process of removing hold from eDiscovery non-custodial data sources. After the operation is created, you can get the status by retrieving the `Location` parameter from the response headers. The location provides a URL that will return an eDiscoveryHoldOperation object.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRemoveHoldPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
