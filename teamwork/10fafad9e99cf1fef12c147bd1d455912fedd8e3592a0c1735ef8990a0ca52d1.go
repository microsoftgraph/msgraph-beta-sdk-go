package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder provides operations to call the add method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilderInternal instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/getAllMembers/add", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// Deprecated: This method is obsolete. Use PostAsAddPostResponse instead.
// returns a TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-beta
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder) Post(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilderPostRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddResponseable), nil
}
// PostAsAddPostResponse add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-beta
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder) PostAsAddPostResponse(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilderPostRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponseable), nil
}
// ToPostRequestInformation add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelGetAllMembersAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
