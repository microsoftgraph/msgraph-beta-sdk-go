package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder provides operations to manage the endpointPrivilegeManagementProvisioningStatus property of the microsoft.graph.deviceManagement entity.
type EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderGetQueryParameters endpoint privilege management (EPM) tenant provisioning status contains tenant level license and onboarding state information.
type EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderGetQueryParameters
}
// EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderInternal instantiates a new EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder and sets the default values.
func NewEndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) {
    m := &EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/endpointPrivilegeManagementProvisioningStatus{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder instantiates a new EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder and sets the default values.
func NewEndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property endpointPrivilegeManagementProvisioningStatus for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) Delete(ctx context.Context, requestConfiguration *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get endpoint privilege management (EPM) tenant provisioning status contains tenant level license and onboarding state information.
// returns a EndpointPrivilegeManagementProvisioningStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndpointPrivilegeManagementProvisioningStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEndpointPrivilegeManagementProvisioningStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndpointPrivilegeManagementProvisioningStatusable), nil
}
// Patch update the navigation property endpointPrivilegeManagementProvisioningStatus in deviceManagement
// returns a EndpointPrivilegeManagementProvisioningStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndpointPrivilegeManagementProvisioningStatusable, requestConfiguration *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndpointPrivilegeManagementProvisioningStatusable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEndpointPrivilegeManagementProvisioningStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndpointPrivilegeManagementProvisioningStatusable), nil
}
// ToDeleteRequestInformation delete navigation property endpointPrivilegeManagementProvisioningStatus for deviceManagement
// returns a *RequestInformation when successful
func (m *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation endpoint privilege management (EPM) tenant provisioning status contains tenant level license and onboarding state information.
// returns a *RequestInformation when successful
func (m *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property endpointPrivilegeManagementProvisioningStatus in deviceManagement
// returns a *RequestInformation when successful
func (m *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndpointPrivilegeManagementProvisioningStatusable, requestConfiguration *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder when successful
func (m *EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) WithUrl(rawUrl string)(*EndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder) {
    return NewEndpointprivilegemanagementprovisioningstatusEndpointPrivilegeManagementProvisioningStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
