package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1cd82da38eaeb9697597d203dbe05121215aa83c60f4bb16b4c4bb0ae1e829f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item/gethealthmetrictimeseries"
    i3f40885f535fa833a03343f90bdd196d1f53ba4e83ad2000066e6a2beaee2389 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item/generateserverlogcollectionrequest"
    i9c30fe64f3cadb8062e1e80b68b8d12a9b9727c1f84ad283d585cd25e965ef20 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item/createserverlogcollectionrequest"
    ia9fb0847e9b7d23ecba981329a874216f6257bc66e680123b5b5de57e3c5785f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item/gethealthmetrics"
)

// MicrosoftTunnelServerItemRequestBuilder provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
type MicrosoftTunnelServerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosoftTunnelServerItemRequestBuilderGetQueryParameters a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
type MicrosoftTunnelServerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftTunnelServerItemRequestBuilderGetQueryParameters
}
// MicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftTunnelServerItemRequestBuilderInternal instantiates a new MicrosoftTunnelServerItemRequestBuilder and sets the default values.
func NewMicrosoftTunnelServerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelServerItemRequestBuilder) {
    m := &MicrosoftTunnelServerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers/{microsoftTunnelServer%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMicrosoftTunnelServerItemRequestBuilder instantiates a new MicrosoftTunnelServerItemRequestBuilder and sets the default values.
func NewMicrosoftTunnelServerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelServerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelServerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property microsoftTunnelServers for deviceManagement
func (m *MicrosoftTunnelServerItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelServerItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property microsoftTunnelServers in deviceManagement
func (m *MicrosoftTunnelServerItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, requestConfiguration *MicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateServerLogCollectionRequest provides operations to call the createServerLogCollectionRequest method.
func (m *MicrosoftTunnelServerItemRequestBuilder) CreateServerLogCollectionRequest()(*i9c30fe64f3cadb8062e1e80b68b8d12a9b9727c1f84ad283d585cd25e965ef20.CreateServerLogCollectionRequestRequestBuilder) {
    return i9c30fe64f3cadb8062e1e80b68b8d12a9b9727c1f84ad283d585cd25e965ef20.NewCreateServerLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property microsoftTunnelServers for deviceManagement
func (m *MicrosoftTunnelServerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GenerateServerLogCollectionRequest provides operations to call the generateServerLogCollectionRequest method.
func (m *MicrosoftTunnelServerItemRequestBuilder) GenerateServerLogCollectionRequest()(*i3f40885f535fa833a03343f90bdd196d1f53ba4e83ad2000066e6a2beaee2389.GenerateServerLogCollectionRequestRequestBuilder) {
    return i3f40885f535fa833a03343f90bdd196d1f53ba4e83ad2000066e6a2beaee2389.NewGenerateServerLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelServerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable), nil
}
// GetHealthMetrics provides operations to call the getHealthMetrics method.
func (m *MicrosoftTunnelServerItemRequestBuilder) GetHealthMetrics()(*ia9fb0847e9b7d23ecba981329a874216f6257bc66e680123b5b5de57e3c5785f.GetHealthMetricsRequestBuilder) {
    return ia9fb0847e9b7d23ecba981329a874216f6257bc66e680123b5b5de57e3c5785f.NewGetHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetHealthMetricTimeSeries provides operations to call the getHealthMetricTimeSeries method.
func (m *MicrosoftTunnelServerItemRequestBuilder) GetHealthMetricTimeSeries()(*i1cd82da38eaeb9697597d203dbe05121215aa83c60f4bb16b4c4bb0ae1e829f2.GetHealthMetricTimeSeriesRequestBuilder) {
    return i1cd82da38eaeb9697597d203dbe05121215aa83c60f4bb16b4c4bb0ae1e829f2.NewGetHealthMetricTimeSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property microsoftTunnelServers in deviceManagement
func (m *MicrosoftTunnelServerItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, requestConfiguration *MicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable), nil
}
