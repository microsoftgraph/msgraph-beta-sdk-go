package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder provides operations to count the resources in the collection.
type ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder) {
    m := &ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/reviewSets/{reviewSet%2Did}/queries/$count";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the number of the resource
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "text/plain"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the number of the resource
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemReviewSetsItemQueriesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
