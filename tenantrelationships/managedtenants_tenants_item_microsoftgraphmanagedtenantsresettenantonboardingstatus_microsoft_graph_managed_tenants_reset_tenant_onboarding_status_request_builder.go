package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder provides operations to call the resetTenantOnboardingStatus method.
type ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderInternal instantiates a new ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder and sets the default values.
func NewManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) {
    m := &ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenants/{tenant%2Did}/microsoft.graph.managedTenants.resetTenantOnboardingStatus", pathParameters),
    }
    return m
}
// NewManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder instantiates a new ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder and sets the default values.
func NewManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post carries out the appropriate procedures to reset the onboarding status for the managed tenant that was removed from the multitenant management platform using the offboardTenant action. By invoking this action the platform attempts to onboard the managed tenant for management.
// returns a Tenantable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-tenant-resettenantonboardingstatus?view=graph-rest-beta
func (m *ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) Post(ctx context.Context, requestConfiguration *ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.Tenantable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation carries out the appropriate procedures to reset the onboarding status for the managed tenant that was removed from the multitenant management platform using the offboardTenant action. By invoking this action the platform attempts to onboard the managed tenant for management.
// returns a *RequestInformation when successful
func (m *ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder when successful
func (m *ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder) {
    return NewManagedtenantsTenantsItemMicrosoftgraphmanagedtenantsresettenantonboardingstatusMicrosoftGraphManagedTenantsResetTenantOnboardingStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
