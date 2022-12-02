package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder provides operations to manage the noncustodialDataSources property of the microsoft.graph.ediscovery.case entity.
type ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters returns a list of case noncustodialDataSource objects for this case.  Nullable.
type ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters
}
// ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyHold provides operations to call the applyHold method.
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) ApplyHold()(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemApplyHoldRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderInternal instantiates a new NoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) {
    m := &ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/noncustodialDataSources/{noncustodialDataSource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder instantiates a new NoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property noncustodialDataSources for compliance
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property noncustodialDataSources in compliance
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, requestConfiguration *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DataSource provides operations to manage the dataSource property of the microsoft.graph.ediscovery.noncustodialDataSource entity.
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) DataSource()(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemDataSourceRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemDataSourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property noncustodialDataSources for compliance
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable), nil
}
// Patch update the navigation property noncustodialDataSources in compliance
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, requestConfiguration *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable), nil
}
// Release provides operations to call the release method.
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) Release()(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemReleaseRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveHold provides operations to call the removeHold method.
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) RemoveHold()(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemRemoveHoldRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateIndex provides operations to call the updateIndex method.
func (m *ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) UpdateIndex()(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemUpdateIndexRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesItemUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
