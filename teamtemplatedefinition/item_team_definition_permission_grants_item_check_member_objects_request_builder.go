package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder provides operations to call the checkMemberObjects method.
type ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal instantiates a new CheckMemberObjectsRequestBuilder and sets the default values.
func NewItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    m := &ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/permissionGrants/{resourceSpecificPermissionGrant%2Did}/checkMemberObjects", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder instantiates a new CheckMemberObjectsRequestBuilder and sets the default values.
func NewItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkMemberObjects
// Deprecated: This method is obsolete. Use PostAsCheckMemberObjectsPostResponse instead.
func (m *ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder) Post(ctx context.Context, body ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration)(ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponseable), nil
}
// PostAsCheckMemberObjectsPostResponse invoke action checkMemberObjects
func (m *ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder) PostAsCheckMemberObjectsPostResponse(ctx context.Context, body ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration)(ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsPostResponseable), nil
}
// ToPostRequestInformation invoke action checkMemberObjects
func (m *ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    return NewItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
