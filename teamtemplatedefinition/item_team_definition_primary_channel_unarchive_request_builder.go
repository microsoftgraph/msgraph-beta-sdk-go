package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder provides operations to call the unarchive method.
type ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilderInternal instantiates a new ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder) {
    m := &ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/unarchive", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder instantiates a new ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore an archived channel. Unarchiving restores the ability for users to send messages and edit the channel. Channels are archived via the archive API. Unarchiving is an asynchronous operation; a channel is unarchived when the asynchronous unarchive operation completes successfully, which might occur after this method responds.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-unarchive?view=graph-rest-1.0
func (m *ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation restore an archived channel. Unarchiving restores the ability for users to send messages and edit the channel. Channels are archived via the archive API. Unarchiving is an asynchronous operation; a channel is unarchived when the asynchronous unarchive operation completes successfully, which might occur after this method responds.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelUnarchiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
