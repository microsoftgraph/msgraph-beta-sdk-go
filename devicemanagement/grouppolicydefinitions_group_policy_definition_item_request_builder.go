package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
type GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters the available group policy definitions for this account.
type GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters
}
// GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Category provides operations to manage the category property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GrouppolicydefinitionsItemCategoryRequestBuilder when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) Category()(*GrouppolicydefinitionsItemCategoryRequestBuilder) {
    return NewGrouppolicydefinitionsItemCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderInternal instantiates a new GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewGrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    m := &GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder instantiates a new GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewGrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DefinitionFile provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GrouppolicydefinitionsItemDefinitionfileDefinitionFileRequestBuilder when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) DefinitionFile()(*GrouppolicydefinitionsItemDefinitionfileDefinitionFileRequestBuilder) {
    return NewGrouppolicydefinitionsItemDefinitionfileDefinitionFileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property groupPolicyDefinitions for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available group policy definitions for this account.
// returns a GroupPolicyDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
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
// NextVersionDefinition provides operations to manage the nextVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GrouppolicydefinitionsItemNextversiondefinitionNextVersionDefinitionRequestBuilder when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) NextVersionDefinition()(*GrouppolicydefinitionsItemNextversiondefinitionNextVersionDefinitionRequestBuilder) {
    return NewGrouppolicydefinitionsItemNextversiondefinitionNextVersionDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property groupPolicyDefinitions in deviceManagement
// returns a GroupPolicyDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
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
// returns a *GrouppolicydefinitionsItemPresentationsRequestBuilder when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) Presentations()(*GrouppolicydefinitionsItemPresentationsRequestBuilder) {
    return NewGrouppolicydefinitionsItemPresentationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreviousVersionDefinition provides operations to manage the previousVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GrouppolicydefinitionsItemPreviousversiondefinitionPreviousVersionDefinitionRequestBuilder when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) PreviousVersionDefinition()(*GrouppolicydefinitionsItemPreviousversiondefinitionPreviousVersionDefinitionRequestBuilder) {
    return NewGrouppolicydefinitionsItemPreviousversiondefinitionPreviousVersionDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property groupPolicyDefinitions for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the available group policy definitions for this account.
// returns a *RequestInformation when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupPolicyDefinitions in deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder when successful
func (m *GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*GrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    return NewGrouppolicydefinitionsGroupPolicyDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
