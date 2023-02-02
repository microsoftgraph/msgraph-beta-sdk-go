package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder provides operations to call the addToReviewSet method.
type CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilderInternal instantiates a new AddToReviewSetRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/microsoft.graph.security.addToReviewSet";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder instantiates a new AddToReviewSetRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of adding a collection from Microsoft 365 services to a review set. After the operation is created, you can get the status of the operation by retrieving the `Location` parameter from the response headers. The location provides a URL that will return a Add to review set operation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoveryreviewset-addtoreviewset?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation start the process of adding a collection from Microsoft 365 services to a review set. After the operation is created, you can get the status of the operation by retrieving the `Location` parameter from the response headers. The location provides a URL that will return a Add to review set operation.
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
