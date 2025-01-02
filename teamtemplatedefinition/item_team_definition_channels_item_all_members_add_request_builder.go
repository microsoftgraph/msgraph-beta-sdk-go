package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder provides operations to call the add method.
type ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionChannelsItemAllMembersAddRequestBuilderInternal instantiates a new ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemAllMembersAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder) {
    m := &ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/allMembers/add", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder instantiates a new ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsItemAllMembersAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// Deprecated: This method is obsolete. Use PostAsAddPostResponse instead.
// returns a ItemTeamDefinitionChannelsItemAllMembersAddResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-beta
func (m *ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder) Post(ctx context.Context, body ItemTeamDefinitionChannelsItemAllMembersAddPostRequestBodyable, requestConfiguration *ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilderPostRequestConfiguration)(ItemTeamDefinitionChannelsItemAllMembersAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionChannelsItemAllMembersAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionChannelsItemAllMembersAddResponseable), nil
}
// PostAsAddPostResponse add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a ItemTeamDefinitionChannelsItemAllMembersAddPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-beta
func (m *ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder) PostAsAddPostResponse(ctx context.Context, body ItemTeamDefinitionChannelsItemAllMembersAddPostRequestBodyable, requestConfiguration *ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilderPostRequestConfiguration)(ItemTeamDefinitionChannelsItemAllMembersAddPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionChannelsItemAllMembersAddPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionChannelsItemAllMembersAddPostResponseable), nil
}
// ToPostRequestInformation add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamDefinitionChannelsItemAllMembersAddPostRequestBodyable, requestConfiguration *ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemAllMembersAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
