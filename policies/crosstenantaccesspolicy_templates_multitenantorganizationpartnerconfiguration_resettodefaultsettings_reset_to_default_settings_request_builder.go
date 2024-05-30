package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder provides operations to call the resetToDefaultSettings method.
type CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilderInternal instantiates a new CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder) {
    m := &CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationPartnerConfiguration/resetToDefaultSettings", pathParameters),
    }
    return m
}
// NewCrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder instantiates a new CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset the cross-tenant access policy template with inbound and outbound partner configuration settings for a multitenant organization to the default values. In its reset state, the template has no impact on partner configuration settings, other than that an unconfigured partner configuration is created if needed, for every multitenant organization tenant.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganizationpartnerconfigurationtemplate-resettodefaultsettings?view=graph-rest-beta
func (m *CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder) Post(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reset the cross-tenant access policy template with inbound and outbound partner configuration settings for a multitenant organization to the default values. In its reset state, the template has no impact on partner configuration settings, other than that an unconfigured partner configuration is created if needed, for every multitenant organization tenant.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder when successful
func (m *CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder) WithUrl(rawUrl string)(*CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder) {
    return NewCrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationResettodefaultsettingsResetToDefaultSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
