package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters the group policy presentations associated with the definition.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal instantiates a new GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    m := &GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/nextVersionDefinition/previousVersionDefinition/presentations/{groupPolicyPresentation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder instantiates a new GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.groupPolicyPresentation entity.
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Definition()(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property presentations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the group policy presentations associated with the definition.
// returns a GroupPolicyPresentationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable), nil
}
// Patch update the navigation property presentations in deviceManagement
// returns a GroupPolicyPresentationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable), nil
}
// ToDeleteRequestInformation delete navigation property presentations for deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the group policy presentations associated with the definition.
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property presentations in deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
