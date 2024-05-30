package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder provides operations to call the generateDownloadUri method.
type AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilderInternal instantiates a new AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder and sets the default values.
func NewAccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder) {
    m := &AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/historyDefinitions/{accessReviewHistoryDefinition%2Did}/instances/{accessReviewHistoryInstance%2Did}/generateDownloadUri", pathParameters),
    }
    return m
}
// NewAccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder instantiates a new AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder and sets the default values.
func NewAccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generates a URI for an accessReviewHistoryInstance object the status for which is done. Each URI can be used to retrieve the instance's review history data. Each URI is valid for 24 hours and can be retrieved by fetching the downloadUri property from the accessReviewHistoryInstance object.
// returns a AccessReviewHistoryInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewhistoryinstance-generatedownloaduri?view=graph-rest-beta
func (m *AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewHistoryInstanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewHistoryInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewHistoryInstanceable), nil
}
// ToPostRequestInformation generates a URI for an accessReviewHistoryInstance object the status for which is done. Each URI can be used to retrieve the instance's review history data. Each URI is valid for 24 hours and can be retrieved by fetching the downloadUri property from the accessReviewHistoryInstance object.
// returns a *RequestInformation when successful
func (m *AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder when successful
func (m *AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder) {
    return NewAccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
