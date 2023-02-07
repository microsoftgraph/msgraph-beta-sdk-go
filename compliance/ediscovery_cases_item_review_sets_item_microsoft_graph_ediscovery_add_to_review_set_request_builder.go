package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder provides operations to call the addToReviewSet method.
type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilderInternal instantiates a new MicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder) {
    m := &EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/reviewSets/{reviewSet%2Did}/microsoft.graph.ediscovery.addToReviewSet";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder instantiates a new MicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start the process of adding a collection from Microsoft 365 services to a review set. After the operation is created, you can get the status of the operation by retrieving the `Location` parameter from the response headers. The location provides a URL that will return a caseExportOperation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/ediscovery-reviewset-addtoreviewset?view=graph-rest-1.0
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder) Post(ctx context.Context, body EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation start the process of adding a collection from Microsoft 365 services to a review set. After the operation is created, you can get the status of the operation by retrieving the `Location` parameter from the response headers. The location provides a URL that will return a caseExportOperation.
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilder) ToPostRequestInformation(ctx context.Context, body EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
