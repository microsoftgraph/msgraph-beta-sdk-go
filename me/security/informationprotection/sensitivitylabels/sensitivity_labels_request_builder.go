package sensitivitylabels

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4621b23c1a4155914a831bad97217b88c3ce3b737da6d18246c93060a04baed8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/count"
    i5cad7b98159f40ed5a5c934595269dacf1b687ffeb3281ebb743bd111e0a847b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/evaluateremoval"
    i88759b56304a762698295808a5bfb06f29418867560c74fec3dc6fd888ff3f25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/evaluateapplication"
    icc10ec374205c05b4532d9c89d0c36886f133046713bc767e65df773ff64ed0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/evaluateclassificationresults"
    ie0b8e402d3912c7159a7b1635da1d709d914855376e51d3fa858ac845bf853a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/extractcontentlabel"
)

// SensitivityLabelsRequestBuilder provides operations to manage the sensitivityLabels property of the microsoft.graph.security.informationProtection entity.
type SensitivityLabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SensitivityLabelsRequestBuilderGetQueryParameters get a list of sensitivityLabel objects associated with a user or organization.
type SensitivityLabelsRequestBuilderGetQueryParameters struct {
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
// SensitivityLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SensitivityLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SensitivityLabelsRequestBuilderGetQueryParameters
}
// SensitivityLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SensitivityLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSensitivityLabelsRequestBuilderInternal instantiates a new SensitivityLabelsRequestBuilder and sets the default values.
func NewSensitivityLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SensitivityLabelsRequestBuilder) {
    m := &SensitivityLabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/security/informationProtection/sensitivityLabels{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSensitivityLabelsRequestBuilder instantiates a new SensitivityLabelsRequestBuilder and sets the default values.
func NewSensitivityLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SensitivityLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSensitivityLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *SensitivityLabelsRequestBuilder) Count()(*i4621b23c1a4155914a831bad97217b88c3ce3b737da6d18246c93060a04baed8.CountRequestBuilder) {
    return i4621b23c1a4155914a831bad97217b88c3ce3b737da6d18246c93060a04baed8.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a list of sensitivityLabel objects associated with a user or organization.
func (m *SensitivityLabelsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SensitivityLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to sensitivityLabels for me
func (m *SensitivityLabelsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *SensitivityLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SensitivityLabelsRequestBuilder) EvaluateApplication()(*i88759b56304a762698295808a5bfb06f29418867560c74fec3dc6fd888ff3f25.EvaluateApplicationRequestBuilder) {
    return i88759b56304a762698295808a5bfb06f29418867560c74fec3dc6fd888ff3f25.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateClassificationResults provides operations to call the evaluateClassificationResults method.
func (m *SensitivityLabelsRequestBuilder) EvaluateClassificationResults()(*icc10ec374205c05b4532d9c89d0c36886f133046713bc767e65df773ff64ed0c.EvaluateClassificationResultsRequestBuilder) {
    return icc10ec374205c05b4532d9c89d0c36886f133046713bc767e65df773ff64ed0c.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateRemoval provides operations to call the evaluateRemoval method.
func (m *SensitivityLabelsRequestBuilder) EvaluateRemoval()(*i5cad7b98159f40ed5a5c934595269dacf1b687ffeb3281ebb743bd111e0a847b.EvaluateRemovalRequestBuilder) {
    return i5cad7b98159f40ed5a5c934595269dacf1b687ffeb3281ebb743bd111e0a847b.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtractContentLabel provides operations to call the extractContentLabel method.
func (m *SensitivityLabelsRequestBuilder) ExtractContentLabel()(*ie0b8e402d3912c7159a7b1635da1d709d914855376e51d3fa858ac845bf853a0.ExtractContentLabelRequestBuilder) {
    return ie0b8e402d3912c7159a7b1635da1d709d914855376e51d3fa858ac845bf853a0.NewExtractContentLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get a list of sensitivityLabel objects associated with a user or organization.
func (m *SensitivityLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *SensitivityLabelsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensitivityLabelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelCollectionResponseable), nil
}
// Post create new navigation property to sensitivityLabels for me
func (m *SensitivityLabelsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *SensitivityLabelsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensitivityLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable), nil
}
