package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder provides operations to call the getCloudPcLaunchInfo method.
type VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder) {
    m := &VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/getCloudPcLaunchInfo()", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder instantiates a new VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the cloudPCLaunchInfo for the signed-in user.
// returns a CloudPcLaunchInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getcloudpclaunchinfo?view=graph-rest-beta
func (m *VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcLaunchInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcLaunchInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcLaunchInfoable), nil
}
// ToGetRequestInformation get the cloudPCLaunchInfo for the signed-in user.
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder) {
    return NewVirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
