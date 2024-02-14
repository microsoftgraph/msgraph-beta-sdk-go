package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder provides operations to call the copyToDefaultContentLocation method.
type ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal instantiates a new ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder) {
    m := &ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/{contentType%2Did}/copyToDefaultContentLocation", pathParameters),
    }
    return m
}
// NewItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder instantiates a new ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post copy a file to a default content location in a [content type][contentType]. The file can then be added as a default file or template via a POST operation.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-copytodefaultcontentlocation?view=graph-rest-1.0
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder) Post(ctx context.Context, body ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyable, requestConfiguration *ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation copy a file to a default content location in a [content type][contentType]. The file can then be added as a default file or template via a POST operation.
// returns a *RequestInformation when successful
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyable, requestConfiguration *ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder when successful
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder) {
    return NewItemSitesItemContentTypesItemCopyToDefaultContentLocationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
