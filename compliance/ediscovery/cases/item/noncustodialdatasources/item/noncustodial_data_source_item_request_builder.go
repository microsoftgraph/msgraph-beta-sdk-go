package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
    i13511b56967fa857f8b5da964db611671b641a5d8b23fc7cb2b961c55a1d34d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item/release"
    i65b909cbf966f7b29e6a2a449760a7ec248fc588d2629192781a50c72895f9f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item/updateindex"
    i7d73475b564a11cabb4f39c6fa186fcdc061a39838a1de682dc6cd0f76e978c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item/removehold"
    i9c1ab7c9ccb2300a006ec15529235f2a2582d6c7dde8c7c5ead9837cbc9312da "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item/applyhold"
    ia8798dbf43801d2aed7bc1af6c6151eba951b020eb6583ac3cce31a2a853d3ba "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item/datasource"
)

// NoncustodialDataSourceItemRequestBuilder provides operations to manage the noncustodialDataSources property of the microsoft.graph.ediscovery.case entity.
type NoncustodialDataSourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NoncustodialDataSourceItemRequestBuilderGetQueryParameters returns a list of case noncustodialDataSource objects for this case.  Nullable.
type NoncustodialDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// NoncustodialDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NoncustodialDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *NoncustodialDataSourceItemRequestBuilderGetQueryParameters
}
// NoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyHold the applyHold property
func (m *NoncustodialDataSourceItemRequestBuilder) ApplyHold()(*i9c1ab7c9ccb2300a006ec15529235f2a2582d6c7dde8c7c5ead9837cbc9312da.ApplyHoldRequestBuilder) {
    return i9c1ab7c9ccb2300a006ec15529235f2a2582d6c7dde8c7c5ead9837cbc9312da.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewNoncustodialDataSourceItemRequestBuilderInternal instantiates a new NoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewNoncustodialDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NoncustodialDataSourceItemRequestBuilder) {
    m := &NoncustodialDataSourceItemRequestBuilder{
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
// NewNoncustodialDataSourceItemRequestBuilder instantiates a new NoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewNoncustodialDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NoncustodialDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNoncustodialDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property noncustodialDataSources for compliance
func (m *NoncustodialDataSourceItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property noncustodialDataSources for compliance
func (m *NoncustodialDataSourceItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *NoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *NoncustodialDataSourceItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *NoncustodialDataSourceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *NoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *NoncustodialDataSourceItemRequestBuilder) CreatePatchRequestInformation(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property noncustodialDataSources in compliance
func (m *NoncustodialDataSourceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, requestConfiguration *NoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DataSource the dataSource property
func (m *NoncustodialDataSourceItemRequestBuilder) DataSource()(*ia8798dbf43801d2aed7bc1af6c6151eba951b020eb6583ac3cce31a2a853d3ba.DataSourceRequestBuilder) {
    return ia8798dbf43801d2aed7bc1af6c6151eba951b020eb6583ac3cce31a2a853d3ba.NewDataSourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property noncustodialDataSources for compliance
func (m *NoncustodialDataSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *NoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *NoncustodialDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *NoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *NoncustodialDataSourceItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, requestConfiguration *NoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Release the release property
func (m *NoncustodialDataSourceItemRequestBuilder) Release()(*i13511b56967fa857f8b5da964db611671b641a5d8b23fc7cb2b961c55a1d34d4.ReleaseRequestBuilder) {
    return i13511b56967fa857f8b5da964db611671b641a5d8b23fc7cb2b961c55a1d34d4.NewReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveHold the removeHold property
func (m *NoncustodialDataSourceItemRequestBuilder) RemoveHold()(*i7d73475b564a11cabb4f39c6fa186fcdc061a39838a1de682dc6cd0f76e978c2.RemoveHoldRequestBuilder) {
    return i7d73475b564a11cabb4f39c6fa186fcdc061a39838a1de682dc6cd0f76e978c2.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateIndex the updateIndex property
func (m *NoncustodialDataSourceItemRequestBuilder) UpdateIndex()(*i65b909cbf966f7b29e6a2a449760a7ec248fc588d2629192781a50c72895f9f0.UpdateIndexRequestBuilder) {
    return i65b909cbf966f7b29e6a2a449760a7ec248fc588d2629192781a50c72895f9f0.NewUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
