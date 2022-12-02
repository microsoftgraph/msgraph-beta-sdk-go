package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
type SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters returns a list of case ediscoveryNoncustodialDataSource objects for this case.
type SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters
}
// SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyHold provides operations to call the applyHold method.
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ApplyHold()(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemApplyHoldRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal instantiates a new EdiscoveryNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    m := &SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder instantiates a new EdiscoveryNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property noncustodialDataSources for security
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of case ediscoveryNoncustodialDataSource objects for this case.
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property noncustodialDataSources in security
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable, requestConfiguration *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DataSource provides operations to manage the dataSource property of the microsoft.graph.security.ediscoveryNoncustodialDataSource entity.
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) DataSource()(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemDataSourceRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemDataSourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property noncustodialDataSources for security
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case ediscoveryNoncustodialDataSource objects for this case.
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable), nil
}
// LastIndexOperation provides operations to manage the lastIndexOperation property of the microsoft.graph.security.ediscoveryNoncustodialDataSource entity.
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) LastIndexOperation()(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemLastIndexOperationRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemLastIndexOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property noncustodialDataSources in security
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable, requestConfiguration *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable), nil
}
// Release provides operations to call the release method.
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Release()(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemReleaseRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveHold provides operations to call the removeHold method.
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) RemoveHold()(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemRemoveHoldRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateIndex provides operations to call the updateIndex method.
func (m *SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) UpdateIndex()(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemUpdateIndexRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesItemUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
