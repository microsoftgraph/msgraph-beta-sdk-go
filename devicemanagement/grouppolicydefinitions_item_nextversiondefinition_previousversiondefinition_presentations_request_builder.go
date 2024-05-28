package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
type GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderGetQueryParameters the group policy presentations associated with the definition.
type GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderGetQueryParameters struct {
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
// GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderGetQueryParameters
}
// GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGroupPolicyPresentationId provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) ByGroupPolicyPresentationId(groupPolicyPresentationId string)(*GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyPresentationId != "" {
        urlTplParams["groupPolicyPresentation%2Did"] = groupPolicyPresentationId
    }
    return NewGrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderInternal instantiates a new GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder and sets the default values.
func NewGrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) {
    m := &GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/nextVersionDefinition/previousVersionDefinition/presentations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder instantiates a new GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder and sets the default values.
func NewGrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsCountRequestBuilder when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) Count()(*GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsCountRequestBuilder) {
    return NewGrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the group policy presentations associated with the definition.
// returns a GroupPolicyPresentationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationCollectionResponseable), nil
}
// Post create new navigation property to presentations for deviceManagement
// returns a GroupPolicyPresentationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the group policy presentations associated with the definition.
// returns a *RequestInformation when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to presentations for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) WithUrl(rawUrl string)(*GrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder) {
    return NewGrouppolicydefinitionsItemNextversiondefinitionPreviousversiondefinitionPresentationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
