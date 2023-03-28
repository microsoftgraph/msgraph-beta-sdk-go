package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder provides operations to call the getWebPartsByPosition method.
type ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPagesItemGetWebPartsByPositionRequestBuilderInternal instantiates a new GetWebPartsByPositionRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGetWebPartsByPositionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder) {
    m := &ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{sitePage%2Did}/getWebPartsByPosition", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder instantiates a new GetWebPartsByPositionRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGetWebPartsByPositionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getWebPartsByPosition
func (m *ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder) Post(ctx context.Context, body ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBodyable, requestConfiguration *ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilderPostRequestConfiguration)(ItemSitesItemPagesItemGetWebPartsByPositionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemPagesItemGetWebPartsByPositionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemPagesItemGetWebPartsByPositionResponseable), nil
}
// ToPostRequestInformation invoke action getWebPartsByPosition
func (m *ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBodyable, requestConfiguration *ItemSitesItemPagesItemGetWebPartsByPositionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
