package auditlogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i897c3994bb157d73a049b5f8e090a75f6296cc505d281e096f9c2f5015269d48 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/directoryprovisioning"
    iddb9d350ed2f2204a816a8b2c9737202cbbb609b2b77635734c77c461c7b11b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/directoryaudits"
    ie1f17e2e55a9a96b4c8eaf8d9efd0982978d50135d920222ccacbe442d0074b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/signins"
    ie7cd14714146bead7b23538aa08e8ec73dbbd6fdac83c25da131496b868f0852 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/provisioning"
    i18c1f0dab1391e55e1b65be523640813575f4930aeeb8f410364c82046e73187 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/signins/item"
    i43b93442c076ce1053480c2aaba8f405e6875171980c0abd4c67e9a1a25b53e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/directoryprovisioning/item"
    ida9b6a06c8d64bad7537948374efa593835455cc0fac39900b364fe69288a5b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/provisioning/item"
    iecdf88a9efb8a7e905d5b4313dc1bc508fc1f53e48ec460b752ecce4944a19d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/directoryaudits/item"
)

// AuditLogsRequestBuilder provides operations to manage the auditLogRoot singleton.
type AuditLogsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuditLogsRequestBuilderGetQueryParameters get auditLogs
type AuditLogsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuditLogsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditLogsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuditLogsRequestBuilderGetQueryParameters
}
// AuditLogsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditLogsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuditLogsRequestBuilderInternal instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditLogsRequestBuilder) {
    m := &AuditLogsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/auditLogs{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuditLogsRequestBuilder instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get auditLogs
func (m *AuditLogsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AuditLogsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update auditLogs
func (m *AuditLogsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable, requestConfiguration *AuditLogsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DirectoryAudits provides operations to manage the directoryAudits property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) DirectoryAudits()(*iddb9d350ed2f2204a816a8b2c9737202cbbb609b2b77635734c77c461c7b11b3.DirectoryAuditsRequestBuilder) {
    return iddb9d350ed2f2204a816a8b2c9737202cbbb609b2b77635734c77c461c7b11b3.NewDirectoryAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryAuditsById provides operations to manage the directoryAudits property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) DirectoryAuditsById(id string)(*iecdf88a9efb8a7e905d5b4313dc1bc508fc1f53e48ec460b752ecce4944a19d3.DirectoryAuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryAudit%2Did"] = id
    }
    return iecdf88a9efb8a7e905d5b4313dc1bc508fc1f53e48ec460b752ecce4944a19d3.NewDirectoryAuditItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectoryProvisioning provides operations to manage the directoryProvisioning property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) DirectoryProvisioning()(*i897c3994bb157d73a049b5f8e090a75f6296cc505d281e096f9c2f5015269d48.DirectoryProvisioningRequestBuilder) {
    return i897c3994bb157d73a049b5f8e090a75f6296cc505d281e096f9c2f5015269d48.NewDirectoryProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryProvisioningById provides operations to manage the directoryProvisioning property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) DirectoryProvisioningById(id string)(*i43b93442c076ce1053480c2aaba8f405e6875171980c0abd4c67e9a1a25b53e5.ProvisioningObjectSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["provisioningObjectSummary%2Did"] = id
    }
    return i43b93442c076ce1053480c2aaba8f405e6875171980c0abd4c67e9a1a25b53e5.NewProvisioningObjectSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get auditLogs
func (m *AuditLogsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditLogsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuditLogRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable), nil
}
// Patch update auditLogs
func (m *AuditLogsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable, requestConfiguration *AuditLogsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuditLogRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable), nil
}
// Provisioning provides operations to manage the provisioning property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) Provisioning()(*ie7cd14714146bead7b23538aa08e8ec73dbbd6fdac83c25da131496b868f0852.ProvisioningRequestBuilder) {
    return ie7cd14714146bead7b23538aa08e8ec73dbbd6fdac83c25da131496b868f0852.NewProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisioningById provides operations to manage the provisioning property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) ProvisioningById(id string)(*ida9b6a06c8d64bad7537948374efa593835455cc0fac39900b364fe69288a5b5.ProvisioningObjectSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["provisioningObjectSummary%2Did"] = id
    }
    return ida9b6a06c8d64bad7537948374efa593835455cc0fac39900b364fe69288a5b5.NewProvisioningObjectSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SignIns provides operations to manage the signIns property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) SignIns()(*ie1f17e2e55a9a96b4c8eaf8d9efd0982978d50135d920222ccacbe442d0074b7.SignInsRequestBuilder) {
    return ie1f17e2e55a9a96b4c8eaf8d9efd0982978d50135d920222ccacbe442d0074b7.NewSignInsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SignInsById provides operations to manage the signIns property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) SignInsById(id string)(*i18c1f0dab1391e55e1b65be523640813575f4930aeeb8f410364c82046e73187.SignInItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["signIn%2Did"] = id
    }
    return i18c1f0dab1391e55e1b65be523640813575f4930aeeb8f410364c82046e73187.NewSignInItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
