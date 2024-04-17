package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder provides operations to manage the nextVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
type GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderGetQueryParameters definition of the next version of this definition
type GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderGetQueryParameters
}
// GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Category provides operations to manage the category property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionCategoryRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) Category()(*GroupPolicyDefinitionsItemNextVersionDefinitionCategoryRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderInternal instantiates a new GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) {
    m := &GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/nextVersionDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder instantiates a new GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// DefinitionFile provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionDefinitionFileRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) DefinitionFile()(*GroupPolicyDefinitionsItemNextVersionDefinitionDefinitionFileRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionDefinitionFileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property nextVersionDefinition for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get definition of the next version of this definition
// returns a GroupPolicyDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property nextVersionDefinition in deviceManagement
// returns a GroupPolicyDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Presentations provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionPresentationsRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) Presentations()(*GroupPolicyDefinitionsItemNextVersionDefinitionPresentationsRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPresentationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreviousVersionDefinition provides operations to manage the previousVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) PreviousVersionDefinition()(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property nextVersionDefinition for deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation definition of the next version of this definition
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property nextVersionDefinition in deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
