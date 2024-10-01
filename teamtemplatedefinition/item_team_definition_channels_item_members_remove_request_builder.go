package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder provides operations to call the remove method.
type ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionChannelsItemMembersRemoveRequestBuilderInternal instantiates a new ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemMembersRemoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder) {
    m := &ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/members/remove", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder instantiates a new ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsItemMembersRemoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action remove
// Deprecated: This method is obsolete. Use PostAsRemovePostResponse instead.
// returns a ItemTeamDefinitionChannelsItemMembersRemoveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder) Post(ctx context.Context, body ItemTeamDefinitionChannelsItemMembersRemovePostRequestBodyable, requestConfiguration *ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration)(ItemTeamDefinitionChannelsItemMembersRemoveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionChannelsItemMembersRemoveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionChannelsItemMembersRemoveResponseable), nil
}
// PostAsRemovePostResponse invoke action remove
// returns a ItemTeamDefinitionChannelsItemMembersRemovePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder) PostAsRemovePostResponse(ctx context.Context, body ItemTeamDefinitionChannelsItemMembersRemovePostRequestBodyable, requestConfiguration *ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration)(ItemTeamDefinitionChannelsItemMembersRemovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionChannelsItemMembersRemovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionChannelsItemMembersRemovePostResponseable), nil
}
// ToPostRequestInformation invoke action remove
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamDefinitionChannelsItemMembersRemovePostRequestBodyable, requestConfiguration *ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMembersRemoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
