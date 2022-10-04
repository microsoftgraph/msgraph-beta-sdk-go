package labels

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i13bad00e070fa1dd33a2d812272260fd137db1c553faa0f03c4be0b28414ed5b "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/evaluateclassificationresults"
    i1f4566448b9b69a4c5aa53f28eeb9d10043b88f2acc3d1b58f45685b466b0bc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/evaluateapplication"
    ia970da0fc5eba52aeb4709c9bf149f352a3b6c08a985fcc0e3a4bbeb86d1b14f "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/evaluateremoval"
    iaf0346e770fe66c23422dac3f4da0d29a352d82c7de428316a95811f101ee76c "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/count"
    ibfb726105da3710efbd5bdd0afd208c0efb3e76fca696a1e9b866eee417fbb3d "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/extractlabel"
)

// LabelsRequestBuilder provides operations to manage the labels property of the microsoft.graph.informationProtectionPolicy entity.
type LabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LabelsRequestBuilderGetQueryParameters get a collection of information protection labels available to the user or to the organization.
type LabelsRequestBuilderGetQueryParameters struct {
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
// LabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LabelsRequestBuilderGetQueryParameters
}
// LabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLabelsRequestBuilderInternal instantiates a new LabelsRequestBuilder and sets the default values.
func NewLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRequestBuilder) {
    m := &LabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/policy/labels{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLabelsRequestBuilder instantiates a new LabelsRequestBuilder and sets the default values.
func NewLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *LabelsRequestBuilder) Count()(*iaf0346e770fe66c23422dac3f4da0d29a352d82c7de428316a95811f101ee76c.CountRequestBuilder) {
    return iaf0346e770fe66c23422dac3f4da0d29a352d82c7de428316a95811f101ee76c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a collection of information protection labels available to the user or to the organization.
func (m *LabelsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *LabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to labels for informationProtection
func (m *LabelsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelable, requestConfiguration *LabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// EvaluateApplication the evaluateApplication property
func (m *LabelsRequestBuilder) EvaluateApplication()(*i1f4566448b9b69a4c5aa53f28eeb9d10043b88f2acc3d1b58f45685b466b0bc6.EvaluateApplicationRequestBuilder) {
    return i1f4566448b9b69a4c5aa53f28eeb9d10043b88f2acc3d1b58f45685b466b0bc6.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateClassificationResults the evaluateClassificationResults property
func (m *LabelsRequestBuilder) EvaluateClassificationResults()(*i13bad00e070fa1dd33a2d812272260fd137db1c553faa0f03c4be0b28414ed5b.EvaluateClassificationResultsRequestBuilder) {
    return i13bad00e070fa1dd33a2d812272260fd137db1c553faa0f03c4be0b28414ed5b.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateRemoval the evaluateRemoval property
func (m *LabelsRequestBuilder) EvaluateRemoval()(*ia970da0fc5eba52aeb4709c9bf149f352a3b6c08a985fcc0e3a4bbeb86d1b14f.EvaluateRemovalRequestBuilder) {
    return ia970da0fc5eba52aeb4709c9bf149f352a3b6c08a985fcc0e3a4bbeb86d1b14f.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtractLabel the extractLabel property
func (m *LabelsRequestBuilder) ExtractLabel()(*ibfb726105da3710efbd5bdd0afd208c0efb3e76fca696a1e9b866eee417fbb3d.ExtractLabelRequestBuilder) {
    return ibfb726105da3710efbd5bdd0afd208c0efb3e76fca696a1e9b866eee417fbb3d.NewExtractLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get a collection of information protection labels available to the user or to the organization.
func (m *LabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *LabelsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionLabelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelCollectionResponseable), nil
}
// Post create new navigation property to labels for informationProtection
func (m *LabelsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelable, requestConfiguration *LabelsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelable), nil
}
