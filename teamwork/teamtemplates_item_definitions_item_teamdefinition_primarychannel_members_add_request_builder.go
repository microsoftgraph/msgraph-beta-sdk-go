package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder provides operations to call the add method.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilderInternal instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/members/add", pathParameters),
    }
    return m
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// Deprecated: This method is obsolete. Use PostAsAddPostResponse instead.
// returns a TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-beta
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder) Post(ctx context.Context, body TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostRequestBodyable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilderPostRequestConfiguration)(TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponseable), nil
}
// PostAsAddPostResponse add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-beta
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder) PostAsAddPostResponse(ctx context.Context, body TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostRequestBodyable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilderPostRequestConfiguration)(TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponseable), nil
}
// ToPostRequestInformation add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostRequestBodyable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder) WithUrl(rawUrl string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
