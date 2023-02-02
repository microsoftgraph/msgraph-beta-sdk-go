package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder provides operations to call the completeMigration method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilderInternal instantiates a new CompleteMigrationRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/microsoft.graph.completeMigration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder instantiates a new CompleteMigrationRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post complete the message migration process by removing `migration mode` from a team. `Migration mode` is a special state where certain operations are barred, like message POST and membership operations during the data migration process. After a **completeMigration** request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/team-completemigration?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder) Post(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation complete the message migration process by removing `migration mode` from a team. `Migration mode` is a special state where certain operations are barred, like message POST and membership operations during the data migration process. After a **completeMigration** request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphCompleteMigrationCompleteMigrationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
