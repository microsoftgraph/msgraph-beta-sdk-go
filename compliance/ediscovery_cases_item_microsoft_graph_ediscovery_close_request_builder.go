package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder provides operations to call the close method.
type EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilderInternal instantiates a new MicrosoftGraphEdiscoveryCloseRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder) {
    m := &EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/microsoft.graph.ediscovery.close";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder instantiates a new MicrosoftGraphEdiscoveryCloseRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post close an eDiscovery case. For details, see Close a case.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/ediscovery-case-close?view=graph-rest-1.0
func (m *EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation close an eDiscovery case. For details, see Close a case.
func (m *EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
