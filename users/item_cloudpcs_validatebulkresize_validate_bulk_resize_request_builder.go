package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder provides operations to call the validateBulkResize method.
type ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilderInternal instantiates a new ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder and sets the default values.
func NewItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder) {
    m := &ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/validateBulkResize", pathParameters),
    }
    return m
}
// NewItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder instantiates a new ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder and sets the default values.
func NewItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate that a set of cloudPC devices meet the requirements to be bulk resized.
// Deprecated: This method is obsolete. Use PostAsValidateBulkResizePostResponse instead.
// returns a ItemCloudpcsValidatebulkresizeValidateBulkResizeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-validatebulkresize?view=graph-rest-beta
func (m *ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder) Post(ctx context.Context, body ItemCloudpcsValidatebulkresizeValidateBulkResizePostRequestBodyable, requestConfiguration *ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilderPostRequestConfiguration)(ItemCloudpcsValidatebulkresizeValidateBulkResizeResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudpcsValidatebulkresizeValidateBulkResizeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudpcsValidatebulkresizeValidateBulkResizeResponseable), nil
}
// PostAsValidateBulkResizePostResponse validate that a set of cloudPC devices meet the requirements to be bulk resized.
// returns a ItemCloudpcsValidatebulkresizeValidateBulkResizePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-validatebulkresize?view=graph-rest-beta
func (m *ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder) PostAsValidateBulkResizePostResponse(ctx context.Context, body ItemCloudpcsValidatebulkresizeValidateBulkResizePostRequestBodyable, requestConfiguration *ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilderPostRequestConfiguration)(ItemCloudpcsValidatebulkresizeValidateBulkResizePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudpcsValidatebulkresizeValidateBulkResizePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudpcsValidatebulkresizeValidateBulkResizePostResponseable), nil
}
// ToPostRequestInformation validate that a set of cloudPC devices meet the requirements to be bulk resized.
// returns a *RequestInformation when successful
func (m *ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCloudpcsValidatebulkresizeValidateBulkResizePostRequestBodyable, requestConfiguration *ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder when successful
func (m *ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder) {
    return NewItemCloudpcsValidatebulkresizeValidateBulkResizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
