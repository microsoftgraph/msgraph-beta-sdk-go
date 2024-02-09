package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder provides operations to call the export method.
type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilderInternal instantiates a new EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder) {
    m := &EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/reviewSets/{reviewSet%2Did}/microsoft.graph.ediscovery.export", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder instantiates a new EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate an export from a reviewSet.  For details, see Export documents from a review set in Advanced eDiscovery.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-reviewset-export?view=graph-rest-1.0
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder) Post(ctx context.Context, body EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation initiate an export from a reviewSet.  For details, see Export documents from a review set in Advanced eDiscovery.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder) {
    return NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
