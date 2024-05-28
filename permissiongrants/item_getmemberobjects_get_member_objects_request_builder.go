package permissiongrants

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetmemberobjectsGetMemberObjectsRequestBuilder provides operations to call the getMemberObjects method.
type ItemGetmemberobjectsGetMemberObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal instantiates a new ItemGetmemberobjectsGetMemberObjectsRequestBuilder and sets the default values.
func NewItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    m := &ItemGetmemberobjectsGetMemberObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/permissionGrants/{resourceSpecificPermissionGrant%2Did}/getMemberObjects", pathParameters),
    }
    return m
}
// NewItemGetmemberobjectsGetMemberObjectsRequestBuilder instantiates a new ItemGetmemberobjectsGetMemberObjectsRequestBuilder and sets the default values.
func NewItemGetmemberobjectsGetMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getMemberObjects
// Deprecated: This method is obsolete. Use PostAsGetMemberObjectsPostResponse instead.
// returns a ItemGetmemberobjectsGetMemberObjectsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetmemberobjectsGetMemberObjectsRequestBuilder) Post(ctx context.Context, body ItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *ItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(ItemGetmemberobjectsGetMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmemberobjectsGetMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmemberobjectsGetMemberObjectsResponseable), nil
}
// PostAsGetMemberObjectsPostResponse invoke action getMemberObjects
// returns a ItemGetmemberobjectsGetMemberObjectsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetmemberobjectsGetMemberObjectsRequestBuilder) PostAsGetMemberObjectsPostResponse(ctx context.Context, body ItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *ItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(ItemGetmemberobjectsGetMemberObjectsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmemberobjectsGetMemberObjectsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmemberobjectsGetMemberObjectsPostResponseable), nil
}
// ToPostRequestInformation invoke action getMemberObjects
// returns a *RequestInformation when successful
func (m *ItemGetmemberobjectsGetMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGetmemberobjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *ItemGetmemberobjectsGetMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGetmemberobjectsGetMemberObjectsRequestBuilder when successful
func (m *ItemGetmemberobjectsGetMemberObjectsRequestBuilder) WithUrl(rawUrl string)(*ItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    return NewItemGetmemberobjectsGetMemberObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
