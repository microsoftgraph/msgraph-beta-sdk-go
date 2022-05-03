package sensitivitylabels

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i34c102c60dfd05b9ee4dfef0c0a8de53f12e0a083bcf1faa7bb59e70324250d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/evaluateapplication"
    i67477300d714391f3d8c96f2fb0d3f81f396d1cd261552bd92262fd91c37a3bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/count"
    ib65e4e341d54d15e131b7de82843fefab4eb8856a5a0ae6eb0a8406f56fed314 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/evaluateclassificationresults"
    ica62016c962f92b30229c778cbb076d6d246aadf7143abdcad3a76504b235726 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/extractcontentlabel"
    id8c6caa17cc80769edd1fce7e298996794e66f8c9a61b17a52901919cd49dde1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/evaluateremoval"
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
// SensitivityLabelsRequestBuilderGetQueryParameters get sensitivityLabels from users
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
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
// Count the count property
func (m *SensitivityLabelsRequestBuilder) Count()(*i67477300d714391f3d8c96f2fb0d3f81f396d1cd261552bd92262fd91c37a3bd.CountRequestBuilder) {
    return i67477300d714391f3d8c96f2fb0d3f81f396d1cd261552bd92262fd91c37a3bd.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformationWithRequestConfiguration get sensitivityLabels from users
func (m *SensitivityLabelsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get sensitivityLabels from users
func (m *SensitivityLabelsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SensitivityLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to sensitivityLabels for users
func (m *SensitivityLabelsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to sensitivityLabels for users
func (m *SensitivityLabelsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *SensitivityLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// EvaluateApplication the evaluateApplication property
func (m *SensitivityLabelsRequestBuilder) EvaluateApplication()(*i34c102c60dfd05b9ee4dfef0c0a8de53f12e0a083bcf1faa7bb59e70324250d5.EvaluateApplicationRequestBuilder) {
    return i34c102c60dfd05b9ee4dfef0c0a8de53f12e0a083bcf1faa7bb59e70324250d5.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateClassificationResults the evaluateClassificationResults property
func (m *SensitivityLabelsRequestBuilder) EvaluateClassificationResults()(*ib65e4e341d54d15e131b7de82843fefab4eb8856a5a0ae6eb0a8406f56fed314.EvaluateClassificationResultsRequestBuilder) {
    return ib65e4e341d54d15e131b7de82843fefab4eb8856a5a0ae6eb0a8406f56fed314.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateRemoval the evaluateRemoval property
func (m *SensitivityLabelsRequestBuilder) EvaluateRemoval()(*id8c6caa17cc80769edd1fce7e298996794e66f8c9a61b17a52901919cd49dde1.EvaluateRemovalRequestBuilder) {
    return id8c6caa17cc80769edd1fce7e298996794e66f8c9a61b17a52901919cd49dde1.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtractContentLabel the extractContentLabel property
func (m *SensitivityLabelsRequestBuilder) ExtractContentLabel()(*ica62016c962f92b30229c778cbb076d6d246aadf7143abdcad3a76504b235726.ExtractContentLabelRequestBuilder) {
    return ica62016c962f92b30229c778cbb076d6d246aadf7143abdcad3a76504b235726.NewExtractContentLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler get sensitivityLabels from users
func (m *SensitivityLabelsRequestBuilder) GetWithResponseHandler(requestConfiguration *SensitivityLabelsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelCollectionResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler get sensitivityLabels from users
func (m *SensitivityLabelsRequestBuilder) GetWithResponseHandler(requestConfiguration *SensitivityLabelsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensitivityLabelCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelCollectionResponseable), nil
}
// PostWithResponseHandler create new navigation property to sensitivityLabels for users
func (m *SensitivityLabelsRequestBuilder) PostWithResponseHandler(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *SensitivityLabelsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler create new navigation property to sensitivityLabels for users
func (m *SensitivityLabelsRequestBuilder) PostWithResponseHandler(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *SensitivityLabelsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensitivityLabelFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable), nil
}
