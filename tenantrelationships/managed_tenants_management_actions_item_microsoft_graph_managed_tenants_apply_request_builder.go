package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder provides operations to call the apply method.
type ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilderInternal instantiates a new MicrosoftGraphManagedTenantsApplyRequestBuilder and sets the default values.
func NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder) {
    m := &ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementActions/{managementAction%2Did}/microsoft.graph.managedTenants.apply", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder instantiates a new MicrosoftGraphManagedTenantsApplyRequestBuilder and sets the default values.
func NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post applies a management action against a specific managed tenant. By performing this operation the appropriate configurations will be made and policies created. As example when applying the require multifactor authentication for admins management action creates a Microsoft Entra Conditional Access policy that requires multifactor authentication for all users that have been assigned an administrative directory role. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-managementaction-apply?view=graph-rest-1.0
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder) Post(ctx context.Context, body ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBodyable, requestConfiguration *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionDeploymentStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementActionDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionDeploymentStatusable), nil
}
// ToPostRequestInformation applies a management action against a specific managed tenant. By performing this operation the appropriate configurations will be made and policies created. As example when applying the require multifactor authentication for admins management action creates a Microsoft Entra Conditional Access policy that requires multifactor authentication for all users that have been assigned an administrative directory role. This API is available in the following national cloud deployments.
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyApplyPostRequestBodyable, requestConfiguration *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder) {
    return NewManagedTenantsManagementActionsItemMicrosoftGraphManagedTenantsApplyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
