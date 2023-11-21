package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder provides operations to manage the groupPolicySettingMappings property of the microsoft.graph.groupPolicyMigrationReport entity.
type GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderGetQueryParameters a list of group policy settings to MDM/Intune mappings.
type GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderGetQueryParameters
}
// GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderInternal instantiates a new GroupPolicySettingMappingItemRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) {
    m := &GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/groupPolicySettingMappings/{groupPolicySettingMapping%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewGroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder instantiates a new GroupPolicySettingMappingItemRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupPolicySettingMappings for deviceManagement
func (m *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, requestConfiguration *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicySettingMappingable, requestConfiguration *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) {
    return NewGroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
