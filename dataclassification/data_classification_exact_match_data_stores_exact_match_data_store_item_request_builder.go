package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder provides operations to manage the exactMatchDataStores property of the microsoft.graph.dataClassificationService entity.
type DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderGetQueryParameters get exactMatchDataStores from dataClassification
type DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderGetQueryParameters
}
// DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderInternal instantiates a new ExactMatchDataStoreItemRequestBuilder and sets the default values.
func NewDataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) {
    m := &DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder instantiates a new ExactMatchDataStoreItemRequestBuilder and sets the default values.
func NewDataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property exactMatchDataStores for dataClassification
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get exactMatchDataStores from dataClassification
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property exactMatchDataStores in dataClassification
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable, requestConfiguration *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property exactMatchDataStores for dataClassification
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get exactMatchDataStores from dataClassification
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchDataStoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable), nil
}
// Lookup provides operations to call the lookup method.
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) Lookup()(*DataClassificationExactMatchDataStoresItemLookupRequestBuilder) {
    return NewDataClassificationExactMatchDataStoresItemLookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property exactMatchDataStores in dataClassification
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable, requestConfiguration *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchDataStoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable), nil
}
// Sessions provides operations to manage the sessions property of the microsoft.graph.exactMatchDataStore entity.
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) Sessions()(*DataClassificationExactMatchDataStoresItemSessionsRequestBuilder) {
    return NewDataClassificationExactMatchDataStoresItemSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SessionsById provides operations to manage the sessions property of the microsoft.graph.exactMatchDataStore entity.
func (m *DataClassificationExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) SessionsById(id string)(*DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["exactMatchSession%2Did"] = id
    }
    return NewDataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
