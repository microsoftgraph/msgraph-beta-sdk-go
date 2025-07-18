// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder provides operations to manage the servicePrincipalSignInActivities property of the microsoft.graph.reportRoot entity.
type ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderGetQueryParameters get a servicePrincipalSignInActivity object that contains sign-in activity information for a service principal in a Microsoft Entra tenant. You can use a service principal as a client or resource. A service principal supports delegated or app-only authentication context.
type ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderGetQueryParameters
}
// ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderInternal instantiates a new ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder and sets the default values.
func NewServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) {
    m := &ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/servicePrincipalSignInActivities/{servicePrincipalSignInActivity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder instantiates a new ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder and sets the default values.
func NewServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property servicePrincipalSignInActivities for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get a servicePrincipalSignInActivity object that contains sign-in activity information for a service principal in a Microsoft Entra tenant. You can use a service principal as a client or resource. A service principal supports delegated or app-only authentication context.
// returns a ServicePrincipalSignInActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipalsigninactivity-get?view=graph-rest-beta
func (m *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalSignInActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable), nil
}
// Patch update the navigation property servicePrincipalSignInActivities in reports
// returns a ServicePrincipalSignInActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable, requestConfiguration *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalSignInActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable), nil
}
// ToDeleteRequestInformation delete navigation property servicePrincipalSignInActivities for reports
// returns a *RequestInformation when successful
func (m *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get a servicePrincipalSignInActivity object that contains sign-in activity information for a service principal in a Microsoft Entra tenant. You can use a service principal as a client or resource. A service principal supports delegated or app-only authentication context.
// returns a *RequestInformation when successful
func (m *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property servicePrincipalSignInActivities in reports
// returns a *RequestInformation when successful
func (m *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable, requestConfiguration *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder when successful
func (m *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) WithUrl(rawUrl string)(*ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) {
    return NewServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
