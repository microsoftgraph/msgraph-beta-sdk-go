package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsBulkresizeBulkResizeRequestBuilder provides operations to call the bulkResize method.
type ItemCloudpcsBulkresizeBulkResizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsBulkresizeBulkResizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsBulkresizeBulkResizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudpcsBulkresizeBulkResizeRequestBuilderInternal instantiates a new ItemCloudpcsBulkresizeBulkResizeRequestBuilder and sets the default values.
func NewItemCloudpcsBulkresizeBulkResizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsBulkresizeBulkResizeRequestBuilder) {
    m := &ItemCloudpcsBulkresizeBulkResizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/bulkResize", pathParameters),
    }
    return m
}
// NewItemCloudpcsBulkresizeBulkResizeRequestBuilder instantiates a new ItemCloudpcsBulkresizeBulkResizeRequestBuilder and sets the default values.
func NewItemCloudpcsBulkresizeBulkResizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsBulkresizeBulkResizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsBulkresizeBulkResizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform a bulk resize action to resize a group of cloudPCs that successfully pass validation. If any devices can't be resized, those devices indicate 'resize failed'. The remaining devices are provisioned for the resize process.
// Deprecated: This method is obsolete. Use PostAsBulkResizePostResponse instead.
// returns a ItemCloudpcsBulkresizeBulkResizeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-bulkresize?view=graph-rest-beta
func (m *ItemCloudpcsBulkresizeBulkResizeRequestBuilder) Post(ctx context.Context, body ItemCloudpcsBulkresizeBulkResizePostRequestBodyable, requestConfiguration *ItemCloudpcsBulkresizeBulkResizeRequestBuilderPostRequestConfiguration)(ItemCloudpcsBulkresizeBulkResizeResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudpcsBulkresizeBulkResizeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudpcsBulkresizeBulkResizeResponseable), nil
}
// PostAsBulkResizePostResponse perform a bulk resize action to resize a group of cloudPCs that successfully pass validation. If any devices can't be resized, those devices indicate 'resize failed'. The remaining devices are provisioned for the resize process.
// Deprecated: The bulkResize action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkResize
// returns a ItemCloudpcsBulkresizeBulkResizePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-bulkresize?view=graph-rest-beta
func (m *ItemCloudpcsBulkresizeBulkResizeRequestBuilder) PostAsBulkResizePostResponse(ctx context.Context, body ItemCloudpcsBulkresizeBulkResizePostRequestBodyable, requestConfiguration *ItemCloudpcsBulkresizeBulkResizeRequestBuilderPostRequestConfiguration)(ItemCloudpcsBulkresizeBulkResizePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudpcsBulkresizeBulkResizePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudpcsBulkresizeBulkResizePostResponseable), nil
}
// ToPostRequestInformation perform a bulk resize action to resize a group of cloudPCs that successfully pass validation. If any devices can't be resized, those devices indicate 'resize failed'. The remaining devices are provisioned for the resize process.
// Deprecated: The bulkResize action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkResize
// returns a *RequestInformation when successful
func (m *ItemCloudpcsBulkresizeBulkResizeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCloudpcsBulkresizeBulkResizePostRequestBodyable, requestConfiguration *ItemCloudpcsBulkresizeBulkResizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The bulkResize action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkResize
// returns a *ItemCloudpcsBulkresizeBulkResizeRequestBuilder when successful
func (m *ItemCloudpcsBulkresizeBulkResizeRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsBulkresizeBulkResizeRequestBuilder) {
    return NewItemCloudpcsBulkresizeBulkResizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
