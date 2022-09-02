package sensitivitylabels

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2166bc7b786fde2109d924f08f6296c65e77de9464021480f8892fbdf99b3f35 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/extractcontentlabel"
    i29fcfaae13ef83638781424dc061005bcd763ce8206749ffada5d6a890ec642d "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/evaluateclassificationresults"
    i59437693d0d5b6e9457415ed26971e0b814d891ab580bef2e27a0da41fc85dc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/count"
    i957182eebe9c6b496009074d85aad428d49f0594ec0d6b6159474b3f4773707a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/evaluateapplication"
    ibcbf9e301f23628c55f07a925c00eab304609586bd89ac974cf3dc82de40422e "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/evaluateremoval"
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
// SensitivityLabelsRequestBuilderGetQueryParameters read the Microsoft Purview Information Protection labels for the user or organization.
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
    m.urlTemplate = "{+baseurl}/security/informationProtection/sensitivityLabels{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
// Count the Count property
func (m *SensitivityLabelsRequestBuilder) Count()(*i59437693d0d5b6e9457415ed26971e0b814d891ab580bef2e27a0da41fc85dc2.CountRequestBuilder) {
    return i59437693d0d5b6e9457415ed26971e0b814d891ab580bef2e27a0da41fc85dc2.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation read the Microsoft Purview Information Protection labels for the user or organization.
func (m *SensitivityLabelsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration read the Microsoft Purview Information Protection labels for the user or organization.
func (m *SensitivityLabelsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SensitivityLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to sensitivityLabels for security
func (m *SensitivityLabelsRequestBuilder) CreatePostRequestInformation(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to sensitivityLabels for security
func (m *SensitivityLabelsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *SensitivityLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// EvaluateApplication the evaluateApplication property
func (m *SensitivityLabelsRequestBuilder) EvaluateApplication()(*i957182eebe9c6b496009074d85aad428d49f0594ec0d6b6159474b3f4773707a.EvaluateApplicationRequestBuilder) {
    return i957182eebe9c6b496009074d85aad428d49f0594ec0d6b6159474b3f4773707a.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateClassificationResults the evaluateClassificationResults property
func (m *SensitivityLabelsRequestBuilder) EvaluateClassificationResults()(*i29fcfaae13ef83638781424dc061005bcd763ce8206749ffada5d6a890ec642d.EvaluateClassificationResultsRequestBuilder) {
    return i29fcfaae13ef83638781424dc061005bcd763ce8206749ffada5d6a890ec642d.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateRemoval the evaluateRemoval property
func (m *SensitivityLabelsRequestBuilder) EvaluateRemoval()(*ibcbf9e301f23628c55f07a925c00eab304609586bd89ac974cf3dc82de40422e.EvaluateRemovalRequestBuilder) {
    return ibcbf9e301f23628c55f07a925c00eab304609586bd89ac974cf3dc82de40422e.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtractContentLabel the extractContentLabel property
func (m *SensitivityLabelsRequestBuilder) ExtractContentLabel()(*i2166bc7b786fde2109d924f08f6296c65e77de9464021480f8892fbdf99b3f35.ExtractContentLabelRequestBuilder) {
    return i2166bc7b786fde2109d924f08f6296c65e77de9464021480f8892fbdf99b3f35.NewExtractContentLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read the Microsoft Purview Information Protection labels for the user or organization.
func (m *SensitivityLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *SensitivityLabelsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Post create new navigation property to sensitivityLabels for security
func (m *SensitivityLabelsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *SensitivityLabelsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
