package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
    i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/usersources"
    i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/sitesources"
    iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/unifiedgroupsources"
    i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/usersources/item"
    ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/sitesources/item"
    ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/unifiedgroupsources/item"
)

// LegalHoldItemRequestBuilder provides operations to manage the legalHolds property of the microsoft.graph.ediscovery.case entity.
type LegalHoldItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LegalHoldItemRequestBuilderDeleteOptions options for Delete
type LegalHoldItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// LegalHoldItemRequestBuilderGetOptions options for Get
type LegalHoldItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LegalHoldItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// LegalHoldItemRequestBuilderGetQueryParameters returns a list of case legalHold objects for this case.  Nullable.
type LegalHoldItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LegalHoldItemRequestBuilderPatchOptions options for Patch
type LegalHoldItemRequestBuilderPatchOptions struct {
    // 
    Body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewLegalHoldItemRequestBuilderInternal instantiates a new LegalHoldItemRequestBuilder and sets the default values.
func NewLegalHoldItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LegalHoldItemRequestBuilder) {
    m := &LegalHoldItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/legalHolds/{legalHold%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLegalHoldItemRequestBuilder instantiates a new LegalHoldItemRequestBuilder and sets the default values.
func NewLegalHoldItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LegalHoldItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLegalHoldItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property legalHolds for compliance
func (m *LegalHoldItemRequestBuilder) CreateDeleteRequestInformation(options *LegalHoldItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation returns a list of case legalHold objects for this case.  Nullable.
func (m *LegalHoldItemRequestBuilder) CreateGetRequestInformation(options *LegalHoldItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property legalHolds in compliance
func (m *LegalHoldItemRequestBuilder) CreatePatchRequestInformation(options *LegalHoldItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property legalHolds for compliance
func (m *LegalHoldItemRequestBuilder) Delete(options *LegalHoldItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get returns a list of case legalHold objects for this case.  Nullable.
func (m *LegalHoldItemRequestBuilder) Get(options *LegalHoldItemRequestBuilderGetOptions)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateLegalHoldFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable), nil
}
// Patch update the navigation property legalHolds in compliance
func (m *LegalHoldItemRequestBuilder) Patch(options *LegalHoldItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SiteSources the siteSources property
func (m *LegalHoldItemRequestBuilder) SiteSources()(*i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b.SiteSourcesRequestBuilder) {
    return i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b.NewSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item.siteSources.item collection
func (m *LegalHoldItemRequestBuilder) SiteSourcesById(id string)(*ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936.SiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource%2Did"] = id
    }
    return ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936.NewSiteSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnifiedGroupSources the unifiedGroupSources property
func (m *LegalHoldItemRequestBuilder) UnifiedGroupSources()(*iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8.UnifiedGroupSourcesRequestBuilder) {
    return iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8.NewUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnifiedGroupSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item.unifiedGroupSources.item collection
func (m *LegalHoldItemRequestBuilder) UnifiedGroupSourcesById(id string)(*ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760.UnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource%2Did"] = id
    }
    return ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760.NewUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserSources the userSources property
func (m *LegalHoldItemRequestBuilder) UserSources()(*i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315.UserSourcesRequestBuilder) {
    return i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315.NewUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item.userSources.item collection
func (m *LegalHoldItemRequestBuilder) UserSourcesById(id string)(*i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060.UserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource%2Did"] = id
    }
    return i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060.NewUserSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
