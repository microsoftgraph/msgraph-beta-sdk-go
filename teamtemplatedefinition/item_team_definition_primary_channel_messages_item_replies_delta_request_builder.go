package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder provides operations to call the delta method.
type ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters invoke function delta
type ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters
}
// NewItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal instantiates a new ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) {
    m := &ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/messages/{chatMessage%2Did}/replies/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder instantiates a new ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function delta
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaResponseable), nil
}
// GetAsDeltaGetResponse invoke function delta
// returns a ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaGetResponseable), nil
}
// ToGetRequestInformation invoke function delta
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelMessagesItemRepliesDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
