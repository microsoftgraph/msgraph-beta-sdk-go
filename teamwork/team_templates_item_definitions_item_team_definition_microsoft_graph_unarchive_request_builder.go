package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder provides operations to call the unarchive method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilderInternal instantiates a new MicrosoftGraphUnarchiveRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/microsoft.graph.unarchive";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder instantiates a new MicrosoftGraphUnarchiveRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore an archived team. This restores users' ability to send messages and edit the team, abiding by tenant and team settings. Teams are archived using the archive API. Unarchiving is an async operation. A team is unarchived once the async operation completes successfully, which may occur subsequent to a response from this API.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/team-unarchive?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder) Post(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation restore an archived team. This restores users' ability to send messages and edit the team, abiding by tenant and team settings. Teams are archived using the archive API. Unarchiving is an async operation. A team is unarchived once the async operation completes successfully, which may occur subsequent to a response from this API.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
