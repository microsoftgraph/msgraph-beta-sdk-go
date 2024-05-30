package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder provides operations to manage the presentationValues property of the microsoft.graph.groupPolicyDefinitionValue entity.
type GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderGetQueryParameters the associated group policy presentation values with the definition value.
type GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderGetQueryParameters struct {
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
// GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderGetQueryParameters
}
// GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGroupPolicyPresentationValueId provides operations to manage the presentationValues property of the microsoft.graph.groupPolicyDefinitionValue entity.
// returns a *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesGroupPolicyPresentationValueItemRequestBuilder when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) ByGroupPolicyPresentationValueId(groupPolicyPresentationValueId string)(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesGroupPolicyPresentationValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyPresentationValueId != "" {
        urlTplParams["groupPolicyPresentationValue%2Did"] = groupPolicyPresentationValueId
    }
    return NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesGroupPolicyPresentationValueItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderInternal instantiates a new GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder and sets the default values.
func NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) {
    m := &GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}/definitionValues/{groupPolicyDefinitionValue%2Did}/presentationValues{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder instantiates a new GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder and sets the default values.
func NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) Count()(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder) {
    return NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the associated group policy presentation values with the definition value.
// returns a GroupPolicyPresentationValueCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationValueCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueCollectionResponseable), nil
}
// Post create new navigation property to presentationValues for deviceManagement
// returns a GroupPolicyPresentationValueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable), nil
}
// ToGetRequestInformation the associated group policy presentation values with the definition value.
// returns a *RequestInformation when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to presentationValues for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder) {
    return NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesPresentationValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
