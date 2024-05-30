package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder provides operations to manage the crossCloudGovernmentOrganizationMapping property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderGetQueryParameters read the properties and relationships of a cloudPcCrossCloudGovernmentOrganizationMapping object.
type VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderGetQueryParameters
}
// VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderInternal instantiates a new VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder and sets the default values.
func NewVirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    m := &VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/crossCloudGovernmentOrganizationMapping{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder instantiates a new VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder and sets the default values.
func NewVirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property crossCloudGovernmentOrganizationMapping for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a cloudPcCrossCloudGovernmentOrganizationMapping object.
// returns a CloudPcCrossCloudGovernmentOrganizationMappingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpccrosscloudgovernmentorganizationmapping-get?view=graph-rest-beta
func (m *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a CloudPcCrossCloudGovernmentOrganizationMappingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable, requestConfiguration *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a cloudPcCrossCloudGovernmentOrganizationMapping object.
// returns a *RequestInformation when successful
func (m *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property crossCloudGovernmentOrganizationMapping in deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcCrossCloudGovernmentOrganizationMappingable, requestConfiguration *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder when successful
func (m *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    return NewVirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
