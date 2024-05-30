package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderGetQueryParameters read the properties and relationships of a cloudPcExternalPartnerSetting object.
type VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderGetQueryParameters
}
// VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderInternal instantiates a new VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder and sets the default values.
func NewVirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) {
    m := &VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/externalPartnerSettings/{cloudPcExternalPartnerSetting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder instantiates a new VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder and sets the default values.
func NewVirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property externalPartnerSettings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a cloudPcExternalPartnerSetting object.
// returns a CloudPcExternalPartnerSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcexternalpartnersetting-get?view=graph-rest-beta
func (m *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcExternalPartnerSettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcExternalPartnerSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcExternalPartnerSettingable), nil
}
// Patch update the properties of a cloudPcExternalPartnerSetting object.
// returns a CloudPcExternalPartnerSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcexternalpartnersetting-update?view=graph-rest-beta
func (m *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcExternalPartnerSettingable, requestConfiguration *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcExternalPartnerSettingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcExternalPartnerSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcExternalPartnerSettingable), nil
}
// ToDeleteRequestInformation delete navigation property externalPartnerSettings for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a cloudPcExternalPartnerSetting object.
// returns a *RequestInformation when successful
func (m *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a cloudPcExternalPartnerSetting object.
// returns a *RequestInformation when successful
func (m *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcExternalPartnerSettingable, requestConfiguration *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder when successful
func (m *VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder) {
    return NewVirtualendpointExternalpartnersettingsCloudPcExternalPartnerSettingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
