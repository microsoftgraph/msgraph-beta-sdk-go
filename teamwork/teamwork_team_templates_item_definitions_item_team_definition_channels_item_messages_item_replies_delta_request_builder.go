package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder provides operations to call the delta method.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderGetQueryParameters invoke function delta
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderGetQueryParameters
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderInternal instantiates a new DeltaRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/microsoft.graph.delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder instantiates a new DeltaRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function delta
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function delta
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponseable), nil
}
