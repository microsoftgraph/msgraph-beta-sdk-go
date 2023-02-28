package auditlogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuditLogsRequestBuilderGetQueryParameters
}
// AuditLogsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditLogsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
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
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAuditLogsRequestBuilder instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// DirectoryAudits provides operations to manage the directoryAudits property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) DirectoryAudits()(*DirectoryAuditsRequestBuilder) {
    return NewDirectoryAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectoryAuditsById provides operations to manage the directoryAudits property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) DirectoryAuditsById(id string)(*DirectoryAuditsDirectoryAuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryAudit%2Did"] = id
    }
    return NewDirectoryAuditsDirectoryAuditItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DirectoryProvisioning provides operations to manage the directoryProvisioning property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) DirectoryProvisioning()(*DirectoryProvisioningRequestBuilder) {
    return NewDirectoryProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectoryProvisioningById provides operations to manage the directoryProvisioning property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) DirectoryProvisioningById(id string)(*DirectoryProvisioningProvisioningObjectSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["provisioningObjectSummary%2Did"] = id
    }
    return NewDirectoryProvisioningProvisioningObjectSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get auditLogs
func (m *AuditLogsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditLogsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuditLogRootFromDiscriminatorValue, errorMapping)
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
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuditLogRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable), nil
}
// Provisioning provides operations to manage the provisioning property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) Provisioning()(*ProvisioningRequestBuilder) {
    return NewProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ProvisioningById provides operations to manage the provisioning property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) ProvisioningById(id string)(*ProvisioningProvisioningObjectSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["provisioningObjectSummary%2Did"] = id
    }
    return NewProvisioningProvisioningObjectSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SignIns provides operations to manage the signIns property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) SignIns()(*SignInsRequestBuilder) {
    return NewSignInsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SignInsById provides operations to manage the signIns property of the microsoft.graph.auditLogRoot entity.
func (m *AuditLogsRequestBuilder) SignInsById(id string)(*SignInsSignInItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["signIn%2Did"] = id
    }
    return NewSignInsSignInItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation get auditLogs
func (m *AuditLogsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuditLogsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update auditLogs
func (m *AuditLogsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuditLogRootable, requestConfiguration *AuditLogsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
