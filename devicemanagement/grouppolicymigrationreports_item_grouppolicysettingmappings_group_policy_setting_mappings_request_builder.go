package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder provides operations to manage the groupPolicySettingMappings property of the microsoft.graph.groupPolicyMigrationReport entity.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderGetQueryParameters a list of group policy settings to MDM/Intune mappings.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderGetQueryParameters struct {
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
// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderGetQueryParameters
}
// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGroupPolicySettingMappingId provides operations to manage the groupPolicySettingMappings property of the microsoft.graph.groupPolicyMigrationReport entity.
// returns a *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) ByGroupPolicySettingMappingId(groupPolicySettingMappingId string)(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicySettingMappingId != "" {
        urlTplParams["groupPolicySettingMapping%2Did"] = groupPolicySettingMappingId
    }
    return NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderInternal instantiates a new GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) {
    m := &GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/groupPolicySettingMappings{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder instantiates a new GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GrouppolicymigrationreportsItemGrouppolicysettingmappingsCountRequestBuilder when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) Count()(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsCountRequestBuilder) {
    return NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of group policy settings to MDM/Intune mappings.
// returns a GroupPolicySettingMappingCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicySettingMappingCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingCollectionResponseable), nil
}
// Post create new navigation property to groupPolicySettingMappings for deviceManagement
// returns a GroupPolicySettingMappingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicySettingMappingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable), nil
}
// ToGetRequestInformation a list of group policy settings to MDM/Intune mappings.
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to groupPolicySettingMappings for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) WithUrl(rawUrl string)(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) {
    return NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
