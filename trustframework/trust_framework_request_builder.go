package trustframework

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/policies"
    i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets"
    i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item"
    i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/policies/item"
)

// TrustFrameworkRequestBuilder provides operations to manage the trustFramework singleton.
type TrustFrameworkRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TrustFrameworkRequestBuilderGetOptions options for Get
type TrustFrameworkRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TrustFrameworkRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TrustFrameworkRequestBuilderGetQueryParameters get trustFramework
type TrustFrameworkRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TrustFrameworkRequestBuilderPatchOptions options for Patch
type TrustFrameworkRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewTrustFrameworkRequestBuilderInternal instantiates a new TrustFrameworkRequestBuilder and sets the default values.
func NewTrustFrameworkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrustFrameworkRequestBuilder) {
    m := &TrustFrameworkRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/trustFramework{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTrustFrameworkRequestBuilder instantiates a new TrustFrameworkRequestBuilder and sets the default values.
func NewTrustFrameworkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrustFrameworkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrustFrameworkRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get trustFramework
func (m *TrustFrameworkRequestBuilder) CreateGetRequestInformation(options *TrustFrameworkRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update trustFramework
func (m *TrustFrameworkRequestBuilder) CreatePatchRequestInformation(options *TrustFrameworkRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get trustFramework
func (m *TrustFrameworkRequestBuilder) Get(options *TrustFrameworkRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkable), nil
}
// KeySets the keySets property
func (m *TrustFrameworkRequestBuilder) KeySets()(*i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270.KeySetsRequestBuilder) {
    return i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270.NewKeySetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// KeySetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.trustFramework.keySets.item collection
func (m *TrustFrameworkRequestBuilder) KeySetsById(id string)(*i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280.TrustFrameworkKeySetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["trustFrameworkKeySet%2Did"] = id
    }
    return i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280.NewTrustFrameworkKeySetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update trustFramework
func (m *TrustFrameworkRequestBuilder) Patch(options *TrustFrameworkRequestBuilderPatchOptions)(error) {
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
// Policies the policies property
func (m *TrustFrameworkRequestBuilder) Policies()(*i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9.PoliciesRequestBuilder) {
    return i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.trustFramework.policies.item collection
func (m *TrustFrameworkRequestBuilder) PoliciesById(id string)(*i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd.TrustFrameworkPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["trustFrameworkPolicy%2Did"] = id
    }
    return i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd.NewTrustFrameworkPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
