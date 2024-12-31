package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder provides operations to call the remove method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderInternal instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/allMembers/remove", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// Deprecated: This method is obsolete. Use PostAsRemovePostResponse instead.
// returns a TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmember-remove?view=graph-rest-beta
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder) Post(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemovePostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderPostRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveResponseable), nil
}
// PostAsRemovePostResponse remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// returns a TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmember-remove?view=graph-rest-beta
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder) PostAsRemovePostResponse(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemovePostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderPostRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponseable), nil
}
// ToPostRequestInformation remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemovePostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
