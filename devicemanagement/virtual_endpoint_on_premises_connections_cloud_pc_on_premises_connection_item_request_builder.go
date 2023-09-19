package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters read the properties and relationships of the cloudPcOnPremisesConnection object.
type VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters
}
// VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal instantiates a new CloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    m := &VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder instantiates a new CloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a specific cloudPcOnPremisesConnection object. When you delete an Azure network connection, permissions to the service are removed from the specified Azure resources. You cannot delete an Azure network connection when it's in use, as indicated by the inUse property.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-delete?view=graph-rest-1.0
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of the cloudPcOnPremisesConnection object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-get?view=graph-rest-1.0
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcOnPremisesConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable), nil
}
// Patch update the properties of a cloudPcOnPremisesConnection object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-update?view=graph-rest-1.0
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, requestConfiguration *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcOnPremisesConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable), nil
}
// RunHealthChecks provides operations to call the runHealthChecks method.
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) RunHealthChecks()(*VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder) {
    return NewVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a specific cloudPcOnPremisesConnection object. When you delete an Azure network connection, permissions to the service are removed from the specified Azure resources. You cannot delete an Azure network connection when it's in use, as indicated by the inUse property.
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of the cloudPcOnPremisesConnection object.
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of a cloudPcOnPremisesConnection object.
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, requestConfiguration *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UpdateAdDomainPassword provides operations to call the updateAdDomainPassword method.
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) UpdateAdDomainPassword()(*VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder) {
    return NewVirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    return NewVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
