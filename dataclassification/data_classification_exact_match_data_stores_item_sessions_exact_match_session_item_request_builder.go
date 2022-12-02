package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder provides operations to manage the sessions property of the microsoft.graph.exactMatchDataStore entity.
type DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetQueryParameters get sessions from dataClassification
type DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetQueryParameters
}
// DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Cancel()(*DataClassificationExactMatchDataStoresItemSessionsItemCancelRequestBuilder) {
    return NewDataClassificationExactMatchDataStoresItemSessionsItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Commit provides operations to call the commit method.
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Commit()(*DataClassificationExactMatchDataStoresItemSessionsItemCommitRequestBuilder) {
    return NewDataClassificationExactMatchDataStoresItemSessionsItemCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderInternal instantiates a new ExactMatchSessionItemRequestBuilder and sets the default values.
func NewDataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) {
    m := &DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder instantiates a new ExactMatchSessionItemRequestBuilder and sets the default values.
func NewDataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sessions for dataClassification
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get sessions from dataClassification
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sessions in dataClassification
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, requestConfiguration *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sessions for dataClassification
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get sessions from dataClassification
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable), nil
}
// Patch update the navigation property sessions in dataClassification
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, requestConfiguration *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable), nil
}
// Renew provides operations to call the renew method.
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Renew()(*DataClassificationExactMatchDataStoresItemSessionsItemRenewRequestBuilder) {
    return NewDataClassificationExactMatchDataStoresItemSessionsItemRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadAgent provides operations to manage the uploadAgent property of the microsoft.graph.exactMatchSession entity.
func (m *DataClassificationExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) UploadAgent()(*DataClassificationExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) {
    return NewDataClassificationExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
