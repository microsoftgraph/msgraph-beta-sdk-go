package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder provides operations to call the getMemberObjects method.
type DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal instantiates a new DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder and sets the default values.
func NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    m := &DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/deletedItems/{directoryObject%2Did}/getMemberObjects", pathParameters),
    }
    return m
}
// NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder instantiates a new DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder and sets the default values.
func NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getMemberObjects
// Deprecated: This method is obsolete. Use PostAsGetMemberObjectsPostResponse instead.
// returns a DeleteditemsItemGetmemberobjectsGetMemberObjectsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) Post(ctx context.Context, body DeleteditemsItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(DeleteditemsItemGetmemberobjectsGetMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsItemGetmemberobjectsGetMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsItemGetmemberobjectsGetMemberObjectsResponseable), nil
}
// PostAsGetMemberObjectsPostResponse invoke action getMemberObjects
// returns a DeleteditemsItemGetmemberobjectsGetMemberObjectsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) PostAsGetMemberObjectsPostResponse(ctx context.Context, body DeleteditemsItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(DeleteditemsItemGetmemberobjectsGetMemberObjectsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsItemGetmemberobjectsGetMemberObjectsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsItemGetmemberobjectsGetMemberObjectsPostResponseable), nil
}
// ToPostRequestInformation invoke action getMemberObjects
// returns a *RequestInformation when successful
func (m *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeleteditemsItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder when successful
func (m *DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) WithUrl(rawUrl string)(*DeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    return NewDeleteditemsItemGetmemberobjectsGetMemberObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
