package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyDefinition entity.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderGetQueryParameters the group policy file associated with the definition.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderGetQueryParameters
}
// NewGroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderInternal instantiates a new DefinitionFileRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder) {
    m := &GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/previousVersionDefinition/definitionFile{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewGroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder instantiates a new DefinitionFileRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the group policy file associated with the definition.
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionFileable), nil
}
// ToGetRequestInformation the group policy file associated with the definition.
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder) {
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
