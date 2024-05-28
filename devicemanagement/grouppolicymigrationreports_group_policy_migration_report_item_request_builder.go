package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
type GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderGetQueryParameters a list of Group Policy migration reports.
type GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderGetQueryParameters
}
// GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderInternal instantiates a new GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) {
    m := &GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder instantiates a new GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupPolicyMigrationReports for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a list of Group Policy migration reports.
// returns a GroupPolicyMigrationReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyMigrationReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable), nil
}
// GroupPolicySettingMappings provides operations to manage the groupPolicySettingMappings property of the microsoft.graph.groupPolicyMigrationReport entity.
// returns a *GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder when successful
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) GroupPolicySettingMappings()(*GrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilder) {
    return NewGrouppolicymigrationreportsItemGrouppolicysettingmappingsGroupPolicySettingMappingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property groupPolicyMigrationReports in deviceManagement
// returns a GroupPolicyMigrationReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable, requestConfiguration *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyMigrationReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable), nil
}
// ToDeleteRequestInformation delete navigation property groupPolicyMigrationReports for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a list of Group Policy migration reports.
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupPolicyMigrationReports in deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable, requestConfiguration *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UnsupportedGroupPolicyExtensions provides operations to manage the unsupportedGroupPolicyExtensions property of the microsoft.graph.groupPolicyMigrationReport entity.
// returns a *GrouppolicymigrationreportsItemUnsupportedgrouppolicyextensionsUnsupportedGroupPolicyExtensionsRequestBuilder when successful
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) UnsupportedGroupPolicyExtensions()(*GrouppolicymigrationreportsItemUnsupportedgrouppolicyextensionsUnsupportedGroupPolicyExtensionsRequestBuilder) {
    return NewGrouppolicymigrationreportsItemUnsupportedgrouppolicyextensionsUnsupportedGroupPolicyExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateScopeTags provides operations to call the updateScopeTags method.
// returns a *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder when successful
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) UpdateScopeTags()(*GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder) {
    return NewGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder when successful
func (m *GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) WithUrl(rawUrl string)(*GrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder) {
    return NewGrouppolicymigrationreportsGroupPolicyMigrationReportItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
