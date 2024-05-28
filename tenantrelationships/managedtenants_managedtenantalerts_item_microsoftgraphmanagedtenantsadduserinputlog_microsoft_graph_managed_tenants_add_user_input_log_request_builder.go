package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder provides operations to call the addUserInputLog method.
type ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderInternal instantiates a new ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder and sets the default values.
func NewManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) {
    m := &ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}/microsoft.graph.managedTenants.addUserInputLog", pathParameters),
    }
    return m
}
// NewManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder instantiates a new ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder and sets the default values.
func NewManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addUserInputLog
// returns a ManagedTenantAlertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) Post(ctx context.Context, body ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogAddUserInputLogPostRequestBodyable, requestConfiguration *ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable), nil
}
// ToPostRequestInformation invoke action addUserInputLog
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogAddUserInputLogPostRequestBodyable, requestConfiguration *ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder when successful
func (m *ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) {
    return NewManagedtenantsManagedtenantalertsItemMicrosoftgraphmanagedtenantsadduserinputlogMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
