package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder provides operations to call the checkMemberObjects method.
type DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal instantiates a new DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder and sets the default values.
func NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    m := &DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/deletedItems/{directoryObject%2Did}/checkMemberObjects", pathParameters),
    }
    return m
}
// NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder instantiates a new DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder and sets the default values.
func NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkMemberObjects
// Deprecated: This method is obsolete. Use PostAsCheckMemberObjectsPostResponse instead.
// returns a DeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) Post(ctx context.Context, body DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration)(DeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponseable), nil
}
// PostAsCheckMemberObjectsPostResponse invoke action checkMemberObjects
// returns a DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) PostAsCheckMemberObjectsPostResponse(ctx context.Context, body DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration)(DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponseable), nil
}
// ToPostRequestInformation invoke action checkMemberObjects
// returns a *RequestInformation when successful
func (m *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder when successful
func (m *DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) WithUrl(rawUrl string)(*DeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    return NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
