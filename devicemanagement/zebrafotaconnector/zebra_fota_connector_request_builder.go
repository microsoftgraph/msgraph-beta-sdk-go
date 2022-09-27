package zebrafotaconnector

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i28906a226456a671e4df164729c0ea7ac015d9a6475a44a9bd06c3b189583468 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/zebrafotaconnector/approvefotaapps"
    i625d8d6bdb933bca0689e1edc8ee0dd946b6ec3f24687e0130436e8ae11a7da3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/zebrafotaconnector/connect"
    i7b37f3837e71200d1ef0aa4174eb307e460cb5bd7cc9e1761e7779beb64ae55c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/zebrafotaconnector/hasactivedeployments"
    i806c9652f6fa842e11373fedc19939a6eacb3fac760ecfad2d11e59f4380f211 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/zebrafotaconnector/disconnect"
)

// ZebraFotaConnectorRequestBuilder provides operations to manage the zebraFotaConnector property of the microsoft.graph.deviceManagement entity.
type ZebraFotaConnectorRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ZebraFotaConnectorRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ZebraFotaConnectorRequestBuilderGetQueryParameters the singleton ZebraFotaConnector associated with account.
type ZebraFotaConnectorRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ZebraFotaConnectorRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ZebraFotaConnectorRequestBuilderGetQueryParameters
}
// ZebraFotaConnectorRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApproveFotaApps the approveFotaApps property
func (m *ZebraFotaConnectorRequestBuilder) ApproveFotaApps()(*i28906a226456a671e4df164729c0ea7ac015d9a6475a44a9bd06c3b189583468.ApproveFotaAppsRequestBuilder) {
    return i28906a226456a671e4df164729c0ea7ac015d9a6475a44a9bd06c3b189583468.NewApproveFotaAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Connect the connect property
func (m *ZebraFotaConnectorRequestBuilder) Connect()(*i625d8d6bdb933bca0689e1edc8ee0dd946b6ec3f24687e0130436e8ae11a7da3.ConnectRequestBuilder) {
    return i625d8d6bdb933bca0689e1edc8ee0dd946b6ec3f24687e0130436e8ae11a7da3.NewConnectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewZebraFotaConnectorRequestBuilderInternal instantiates a new ZebraFotaConnectorRequestBuilder and sets the default values.
func NewZebraFotaConnectorRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorRequestBuilder) {
    m := &ZebraFotaConnectorRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/zebraFotaConnector{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewZebraFotaConnectorRequestBuilder instantiates a new ZebraFotaConnectorRequestBuilder and sets the default values.
func NewZebraFotaConnectorRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaConnectorRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property zebraFotaConnector for deviceManagement
func (m *ZebraFotaConnectorRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property zebraFotaConnector for deviceManagement
func (m *ZebraFotaConnectorRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ZebraFotaConnectorRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the singleton ZebraFotaConnector associated with account.
func (m *ZebraFotaConnectorRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the singleton ZebraFotaConnector associated with account.
func (m *ZebraFotaConnectorRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ZebraFotaConnectorRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property zebraFotaConnector in deviceManagement
func (m *ZebraFotaConnectorRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property zebraFotaConnector in deviceManagement
func (m *ZebraFotaConnectorRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable, requestConfiguration *ZebraFotaConnectorRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property zebraFotaConnector for deviceManagement
func (m *ZebraFotaConnectorRequestBuilder) Delete(ctx context.Context, requestConfiguration *ZebraFotaConnectorRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Disconnect the disconnect property
func (m *ZebraFotaConnectorRequestBuilder) Disconnect()(*i806c9652f6fa842e11373fedc19939a6eacb3fac760ecfad2d11e59f4380f211.DisconnectRequestBuilder) {
    return i806c9652f6fa842e11373fedc19939a6eacb3fac760ecfad2d11e59f4380f211.NewDisconnectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the singleton ZebraFotaConnector associated with account.
func (m *ZebraFotaConnectorRequestBuilder) Get(ctx context.Context, requestConfiguration *ZebraFotaConnectorRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateZebraFotaConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable), nil
}
// HasActiveDeployments the hasActiveDeployments property
func (m *ZebraFotaConnectorRequestBuilder) HasActiveDeployments()(*i7b37f3837e71200d1ef0aa4174eb307e460cb5bd7cc9e1761e7779beb64ae55c.HasActiveDeploymentsRequestBuilder) {
    return i7b37f3837e71200d1ef0aa4174eb307e460cb5bd7cc9e1761e7779beb64ae55c.NewHasActiveDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property zebraFotaConnector in deviceManagement
func (m *ZebraFotaConnectorRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable, requestConfiguration *ZebraFotaConnectorRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateZebraFotaConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable), nil
}
