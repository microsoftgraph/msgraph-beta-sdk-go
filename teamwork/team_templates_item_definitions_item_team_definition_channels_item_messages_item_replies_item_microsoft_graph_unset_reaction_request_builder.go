package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder provides operations to call the unsetReaction method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilderInternal instantiates a new MicrosoftGraphUnsetReactionRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/microsoft.graph.unsetReaction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder instantiates a new MicrosoftGraphUnsetReactionRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unsetReaction
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder) Post(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action unsetReaction
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
