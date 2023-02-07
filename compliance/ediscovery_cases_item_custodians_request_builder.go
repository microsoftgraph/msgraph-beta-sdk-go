package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemCustodiansRequestBuilder provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
type EdiscoveryCasesItemCustodiansRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters get a list of the custodian objects and their properties.
type EdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemCustodiansRequestBuilderInternal instantiates a new CustodiansRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansRequestBuilder) {
    m := &EdiscoveryCasesItemCustodiansRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEdiscoveryCasesItemCustodiansRequestBuilder instantiates a new CustodiansRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemCustodiansRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EdiscoveryCasesItemCustodiansRequestBuilder) Count()(*EdiscoveryCasesItemCustodiansCountRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get a list of the custodian objects and their properties.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/ediscovery-case-list-custodians?view=graph-rest-1.0
func (m *EdiscoveryCasesItemCustodiansRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CustodianCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CustodianCollectionResponseable), nil
}
// MicrosoftGraphEdiscoveryApplyHold provides operations to call the applyHold method.
func (m *EdiscoveryCasesItemCustodiansRequestBuilder) MicrosoftGraphEdiscoveryApplyHold()(*EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEdiscoveryRemoveHold provides operations to call the removeHold method.
func (m *EdiscoveryCasesItemCustodiansRequestBuilder) MicrosoftGraphEdiscoveryRemoveHold()(*EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Post create a new custodian object. After the custodian object is created, you will need to create the custodian's userSource to reference their mailbox and OneDrive for Business site.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/ediscovery-case-post-custodians?view=graph-rest-1.0
func (m *EdiscoveryCasesItemCustodiansRequestBuilder) Post(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *EdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable), nil
}
// ToGetRequestInformation get a list of the custodian objects and their properties.
func (m *EdiscoveryCasesItemCustodiansRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new custodian object. After the custodian object is created, you will need to create the custodian's userSource to reference their mailbox and OneDrive for Business site.
func (m *EdiscoveryCasesItemCustodiansRequestBuilder) ToPostRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *EdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
