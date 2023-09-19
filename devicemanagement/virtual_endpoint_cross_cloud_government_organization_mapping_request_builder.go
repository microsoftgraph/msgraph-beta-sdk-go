package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder provides operations to manage the crossCloudGovernmentOrganizationMapping property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderGetQueryParameters read the properties and relationships of a cloudPcCrossCloudGovernmentOrganizationMapping object.
type VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderGetQueryParameters
}
// VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderInternal instantiates a new CrossCloudGovernmentOrganizationMappingRequestBuilder and sets the default values.
func NewVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    m := &VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/crossCloudGovernmentOrganizationMapping{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder instantiates a new CrossCloudGovernmentOrganizationMappingRequestBuilder and sets the default values.
func NewVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property crossCloudGovernmentOrganizationMapping for deviceManagement
func (m *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a cloudPcCrossCloudGovernmentOrganizationMapping object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpccrosscloudgovernmentorganizationmapping-get?view=graph-rest-1.0
func (m *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcCrossCloudGovernmentOrganizationMappingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable), nil
}
// Patch update the navigation property crossCloudGovernmentOrganizationMapping in deviceManagement
func (m *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable, requestConfiguration *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcCrossCloudGovernmentOrganizationMappingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable), nil
}
// ToDeleteRequestInformation delete navigation property crossCloudGovernmentOrganizationMapping for deviceManagement
func (m *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation read the properties and relationships of a cloudPcCrossCloudGovernmentOrganizationMapping object.
func (m *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property crossCloudGovernmentOrganizationMapping in deviceManagement
func (m *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable, requestConfiguration *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    return NewVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
