package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder provides operations to call the getRoleScopeTagsById method.
type RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderInternal instantiates a new RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder and sets the default values.
func NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder) {
    m := &RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/roleScopeTags/getRoleScopeTagsById", pathParameters),
    }
    return m
}
// NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder instantiates a new RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder and sets the default values.
func NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getRoleScopeTagsById
// Deprecated: This method is obsolete. Use PostAsGetRoleScopeTagsByIdPostResponse instead.
// returns a RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder) Post(ctx context.Context, body RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostRequestBodyable, requestConfiguration *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration)(RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdResponseable), nil
}
// PostAsGetRoleScopeTagsByIdPostResponse invoke action getRoleScopeTagsById
// returns a RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder) PostAsGetRoleScopeTagsByIdPostResponse(ctx context.Context, body RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostRequestBodyable, requestConfiguration *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration)(RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostResponseable), nil
}
// ToPostRequestInformation invoke action getRoleScopeTagsById
// returns a *RequestInformation when successful
func (m *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdPostRequestBodyable, requestConfiguration *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder when successful
func (m *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder) WithUrl(rawUrl string)(*RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder) {
    return NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
