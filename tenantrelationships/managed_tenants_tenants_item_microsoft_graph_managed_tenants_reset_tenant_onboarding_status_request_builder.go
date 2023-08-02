package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder provides operations to call the resetTenantOnboardingStatus method.
type ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderInternal instantiates a new MicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder and sets the default values.
func NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) {
    m := &ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenants/{tenant%2Did}/microsoft.graph.managedTenants.resetTenantOnboardingStatus", pathParameters),
    }
    return m
}
// NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder instantiates a new MicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder and sets the default values.
func NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post carries out the appropriate procedures to reset the onboarding status for the managed tenant that was removed from the multi-tenant management platform using the offboardTenant action. By invoking this action the platform will attempt to onboard the managed tenant for management.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-tenant-resettenantonboardingstatus?view=graph-rest-1.0
func (m *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) Post(ctx context.Context, requestConfiguration *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.Tenantable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.Tenantable), nil
}
// ToPostRequestInformation carries out the appropriate procedures to reset the onboarding status for the managed tenant that was removed from the multi-tenant management platform using the offboardTenant action. By invoking this action the platform will attempt to onboard the managed tenant for management.
func (m *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
