package labels

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i815af1b5d2ba5bd7e5f98cda3daead6e66346dba70155ed163a163a6576409e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/evaluateclassificationresults"
    i899f24d803ea2e82f81f82cb89b445d890b85146bdccb8d94b936dd1f2312fd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/count"
    i8e9df7f67db474152e4a3949333d711503813269c2462b73b7c3fb524a46b0d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/evaluateremoval"
    id5310ddb400c09ed63cd8fa70c110b08804c18ed5839cdc5def5070ee8be4a53 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/extractlabel"
    ife1e0da99801d9046825648970ce58a1561136c8db63462e111d9f32f40f7682 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/evaluateapplication"
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
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/informationProtection/policy/labels{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
// Count provides operations to count the resources in the collection.
func (m *LabelsRequestBuilder) Count()(*i899f24d803ea2e82f81f82cb89b445d890b85146bdccb8d94b936dd1f2312fd9.CountRequestBuilder) {
    return i899f24d803ea2e82f81f82cb89b445d890b85146bdccb8d94b936dd1f2312fd9.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreatePostRequestInformation create new navigation property to labels for users
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
// EvaluateApplication provides operations to call the evaluateApplication method.
func (m *LabelsRequestBuilder) EvaluateApplication()(*ife1e0da99801d9046825648970ce58a1561136c8db63462e111d9f32f40f7682.EvaluateApplicationRequestBuilder) {
    return ife1e0da99801d9046825648970ce58a1561136c8db63462e111d9f32f40f7682.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateClassificationResults provides operations to call the evaluateClassificationResults method.
func (m *LabelsRequestBuilder) EvaluateClassificationResults()(*i815af1b5d2ba5bd7e5f98cda3daead6e66346dba70155ed163a163a6576409e5.EvaluateClassificationResultsRequestBuilder) {
    return i815af1b5d2ba5bd7e5f98cda3daead6e66346dba70155ed163a163a6576409e5.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateRemoval provides operations to call the evaluateRemoval method.
func (m *LabelsRequestBuilder) EvaluateRemoval()(*i8e9df7f67db474152e4a3949333d711503813269c2462b73b7c3fb524a46b0d5.EvaluateRemovalRequestBuilder) {
    return i8e9df7f67db474152e4a3949333d711503813269c2462b73b7c3fb524a46b0d5.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtractLabel provides operations to call the extractLabel method.
func (m *LabelsRequestBuilder) ExtractLabel()(*id5310ddb400c09ed63cd8fa70c110b08804c18ed5839cdc5def5070ee8be4a53.ExtractLabelRequestBuilder) {
    return id5310ddb400c09ed63cd8fa70c110b08804c18ed5839cdc5def5070ee8be4a53.NewExtractLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Post create new navigation property to labels for users
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
