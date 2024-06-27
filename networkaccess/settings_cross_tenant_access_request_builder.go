package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// SettingsCrossTenantAccessRequestBuilder provides operations to manage the crossTenantAccess property of the microsoft.graph.networkaccess.settings entity.
type SettingsCrossTenantAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SettingsCrossTenantAccessRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsCrossTenantAccessRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SettingsCrossTenantAccessRequestBuilderGetQueryParameters retrieve the cross-tenant access settings, which include network packet tagging to enforce Tenant Restrictions Policies (TRv2 Policies) aimed at preventing data exfiltration to external tenants.
type SettingsCrossTenantAccessRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SettingsCrossTenantAccessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsCrossTenantAccessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SettingsCrossTenantAccessRequestBuilderGetQueryParameters
}
// SettingsCrossTenantAccessRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsCrossTenantAccessRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSettingsCrossTenantAccessRequestBuilderInternal instantiates a new SettingsCrossTenantAccessRequestBuilder and sets the default values.
func NewSettingsCrossTenantAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SettingsCrossTenantAccessRequestBuilder) {
    m := &SettingsCrossTenantAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/settings/crossTenantAccess{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSettingsCrossTenantAccessRequestBuilder instantiates a new SettingsCrossTenantAccessRequestBuilder and sets the default values.
func NewSettingsCrossTenantAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SettingsCrossTenantAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSettingsCrossTenantAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property crossTenantAccess for networkAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SettingsCrossTenantAccessRequestBuilder) Delete(ctx context.Context, requestConfiguration *SettingsCrossTenantAccessRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the cross-tenant access settings, which include network packet tagging to enforce Tenant Restrictions Policies (TRv2 Policies) aimed at preventing data exfiltration to external tenants.
// returns a CrossTenantAccessSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/networkaccess-crosstenantaccesssettings-get?view=graph-rest-beta
func (m *SettingsCrossTenantAccessRequestBuilder) Get(ctx context.Context, requestConfiguration *SettingsCrossTenantAccessRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CrossTenantAccessSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateCrossTenantAccessSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CrossTenantAccessSettingsable), nil
}
// Patch update the cross-tenant access settings to include network packet tagging for enforcing Tenant Restrictions Policies (TRv2 Policies) that prevent data exfiltration to external tenants.
// returns a CrossTenantAccessSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/networkaccess-crosstenantaccesssettings-update?view=graph-rest-beta
func (m *SettingsCrossTenantAccessRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CrossTenantAccessSettingsable, requestConfiguration *SettingsCrossTenantAccessRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CrossTenantAccessSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateCrossTenantAccessSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CrossTenantAccessSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property crossTenantAccess for networkAccess
// returns a *RequestInformation when successful
func (m *SettingsCrossTenantAccessRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SettingsCrossTenantAccessRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the cross-tenant access settings, which include network packet tagging to enforce Tenant Restrictions Policies (TRv2 Policies) aimed at preventing data exfiltration to external tenants.
// returns a *RequestInformation when successful
func (m *SettingsCrossTenantAccessRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SettingsCrossTenantAccessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the cross-tenant access settings to include network packet tagging for enforcing Tenant Restrictions Policies (TRv2 Policies) that prevent data exfiltration to external tenants.
// returns a *RequestInformation when successful
func (m *SettingsCrossTenantAccessRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CrossTenantAccessSettingsable, requestConfiguration *SettingsCrossTenantAccessRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SettingsCrossTenantAccessRequestBuilder when successful
func (m *SettingsCrossTenantAccessRequestBuilder) WithUrl(rawUrl string)(*SettingsCrossTenantAccessRequestBuilder) {
    return NewSettingsCrossTenantAccessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
