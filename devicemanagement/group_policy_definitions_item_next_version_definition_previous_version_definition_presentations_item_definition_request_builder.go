package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder provides operations to manage the definition property of the microsoft.graph.groupPolicyPresentation entity.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetQueryParameters the group policy definition associated with the presentation.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetQueryParameters
}
// NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal instantiates a new DefinitionRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) {
    m := &GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/nextVersionDefinition/previousVersionDefinition/presentations/{groupPolicyPresentation%2Did}/definition{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder instantiates a new DefinitionRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the group policy definition associated with the presentation.
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable), nil
}
// ToGetRequestInformation the group policy definition associated with the presentation.
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
