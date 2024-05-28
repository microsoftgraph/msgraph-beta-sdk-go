package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder provides operations to manage the definitionValues property of the microsoft.graph.groupPolicyConfiguration entity.
type GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters the list of enabled or disabled group policy definition values for the configuration.
type GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters
}
// GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderInternal instantiates a new GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder and sets the default values.
func NewGrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) {
    m := &GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}/definitionValues/{groupPolicyDefinitionValue%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder instantiates a new GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder and sets the default values.
func NewGrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.groupPolicyDefinitionValue entity.
// returns a *GrouppolicyconfigurationsItemDefinitionvaluesItemDefinitionRequestBuilder when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) Definition()(*GrouppolicyconfigurationsItemDefinitionvaluesItemDefinitionRequestBuilder) {
    return NewGrouppolicyconfigurationsItemDefinitionvaluesItemDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property definitionValues for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of enabled or disabled group policy definition values for the configuration.
// returns a GroupPolicyDefinitionValueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable), nil
}
// Patch update the navigation property definitionValues in deviceManagement
// returns a GroupPolicyDefinitionValueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable), nil
}
// PresentationValues provides operations to manage the presentationValues property of the microsoft.graph.groupPolicyDefinitionValue entity.
// returns a *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) PresentationValues()(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) {
    return NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property definitionValues for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of enabled or disabled group policy definition values for the configuration.
// returns a *RequestInformation when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property definitionValues in deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder) {
    return NewGrouppolicyconfigurationsItemDefinitionvaluesGroupPolicyDefinitionValueItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
