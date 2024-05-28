package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
type GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters the group policy presentations associated with the definition.
type GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters
}
// GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal instantiates a new GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder and sets the default values.
func NewGrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    m := &GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/previousVersionDefinition/nextVersionDefinition/presentations/{groupPolicyPresentation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder instantiates a new GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder and sets the default values.
func NewGrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.groupPolicyPresentation entity.
// returns a *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsItemDefinitionRequestBuilder when successful
func (m *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Definition()(*GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsItemDefinitionRequestBuilder) {
    return NewGrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsItemDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property presentations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
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
func (m *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
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
func (m *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder when successful
func (m *GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) WithUrl(rawUrl string)(*GrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    return NewGrouppolicydefinitionsItemPreviousversiondefinitionNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
