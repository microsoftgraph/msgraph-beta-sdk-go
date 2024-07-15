package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsValidateBulkResizeRequestBuilder provides operations to call the validateBulkResize method.
type ItemCloudPCsValidateBulkResizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudPCsValidateBulkResizeRequestBuilderInternal instantiates a new ItemCloudPCsValidateBulkResizeRequestBuilder and sets the default values.
func NewItemCloudPCsValidateBulkResizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsValidateBulkResizeRequestBuilder) {
    m := &ItemCloudPCsValidateBulkResizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/validateBulkResize", pathParameters),
    }
    return m
}
// NewItemCloudPCsValidateBulkResizeRequestBuilder instantiates a new ItemCloudPCsValidateBulkResizeRequestBuilder and sets the default values.
func NewItemCloudPCsValidateBulkResizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsValidateBulkResizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsValidateBulkResizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate that a set of cloudPC devices meet the requirements to be bulk resized.
// Deprecated: This method is obsolete. Use PostAsValidateBulkResizePostResponse instead.
// returns a ItemCloudPCsValidateBulkResizeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-validatebulkresize?view=graph-rest-beta
func (m *ItemCloudPCsValidateBulkResizeRequestBuilder) Post(ctx context.Context, body ItemCloudPCsValidateBulkResizePostRequestBodyable, requestConfiguration *ItemCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration)(ItemCloudPCsValidateBulkResizeResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsValidateBulkResizeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsValidateBulkResizeResponseable), nil
}
// PostAsValidateBulkResizePostResponse validate that a set of cloudPC devices meet the requirements to be bulk resized.
// returns a ItemCloudPCsValidateBulkResizePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-validatebulkresize?view=graph-rest-beta
func (m *ItemCloudPCsValidateBulkResizeRequestBuilder) PostAsValidateBulkResizePostResponse(ctx context.Context, body ItemCloudPCsValidateBulkResizePostRequestBodyable, requestConfiguration *ItemCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration)(ItemCloudPCsValidateBulkResizePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsValidateBulkResizePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsValidateBulkResizePostResponseable), nil
}
// ToPostRequestInformation validate that a set of cloudPC devices meet the requirements to be bulk resized.
// returns a *RequestInformation when successful
func (m *ItemCloudPCsValidateBulkResizeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCloudPCsValidateBulkResizePostRequestBodyable, requestConfiguration *ItemCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudPCsValidateBulkResizeRequestBuilder when successful
func (m *ItemCloudPCsValidateBulkResizeRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsValidateBulkResizeRequestBuilder) {
    return NewItemCloudPCsValidateBulkResizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
