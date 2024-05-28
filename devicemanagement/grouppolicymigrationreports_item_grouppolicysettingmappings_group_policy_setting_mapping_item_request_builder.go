package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder provides operations to manage the groupPolicySettingMappings property of the microsoft.graph.groupPolicyMigrationReport entity.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderGetQueryParameters a list of group policy settings to MDM/Intune mappings.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderGetQueryParameters
}
// GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderInternal instantiates a new GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) {
    m := &GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/groupPolicySettingMappings/{groupPolicySettingMapping%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder instantiates a new GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupPolicySettingMappings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a list of group policy settings to MDM/Intune mappings.
// returns a GroupPolicySettingMappingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property groupPolicySettingMappings in deviceManagement
// returns a GroupPolicySettingMappingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property groupPolicySettingMappings for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a list of group policy settings to MDM/Intune mappings.
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupPolicySettingMappings in deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, requestConfiguration *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder when successful
func (m *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) WithUrl(rawUrl string)(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder) {
    return NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
