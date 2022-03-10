package sensitivitylabels

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/security"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i34c102c60dfd05b9ee4dfef0c0a8de53f12e0a083bcf1faa7bb59e70324250d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/evaluateapplication"
    i67477300d714391f3d8c96f2fb0d3f81f396d1cd261552bd92262fd91c37a3bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/count"
    ib65e4e341d54d15e131b7de82843fefab4eb8856a5a0ae6eb0a8406f56fed314 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/evaluateclassificationresults"
    ica62016c962f92b30229c778cbb076d6d246aadf7143abdcad3a76504b235726 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/extractcontentlabel"
    id8c6caa17cc80769edd1fce7e298996794e66f8c9a61b17a52901919cd49dde1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security/informationprotection/sensitivitylabels/evaluateremoval"
)

// SensitivityLabelsRequestBuilder provides operations to manage the sensitivityLabels property of the microsoft.graph.security.informationProtection entity.
type SensitivityLabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SensitivityLabelsRequestBuilderGetOptions options for Get
type SensitivityLabelsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SensitivityLabelsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SensitivityLabelsRequestBuilderGetQueryParameters get sensitivityLabels from users
type SensitivityLabelsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// SensitivityLabelsRequestBuilderPostOptions options for Post
type SensitivityLabelsRequestBuilderPostOptions struct {
    // 
    Body i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.SensitivityLabelable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSensitivityLabelsRequestBuilderInternal instantiates a new SensitivityLabelsRequestBuilder and sets the default values.
func NewSensitivityLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SensitivityLabelsRequestBuilder) {
    m := &SensitivityLabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/security/informationProtection/sensitivityLabels{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSensitivityLabelsRequestBuilder instantiates a new SensitivityLabelsRequestBuilder and sets the default values.
func NewSensitivityLabelsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SensitivityLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSensitivityLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SensitivityLabelsRequestBuilder) Count()(*i67477300d714391f3d8c96f2fb0d3f81f396d1cd261552bd92262fd91c37a3bd.CountRequestBuilder) {
    return i67477300d714391f3d8c96f2fb0d3f81f396d1cd261552bd92262fd91c37a3bd.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get sensitivityLabels from users
func (m *SensitivityLabelsRequestBuilder) CreateGetRequestInformation(options *SensitivityLabelsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to sensitivityLabels for users
func (m *SensitivityLabelsRequestBuilder) CreatePostRequestInformation(options *SensitivityLabelsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *SensitivityLabelsRequestBuilder) EvaluateApplication()(*i34c102c60dfd05b9ee4dfef0c0a8de53f12e0a083bcf1faa7bb59e70324250d5.EvaluateApplicationRequestBuilder) {
    return i34c102c60dfd05b9ee4dfef0c0a8de53f12e0a083bcf1faa7bb59e70324250d5.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) EvaluateClassificationResults()(*ib65e4e341d54d15e131b7de82843fefab4eb8856a5a0ae6eb0a8406f56fed314.EvaluateClassificationResultsRequestBuilder) {
    return ib65e4e341d54d15e131b7de82843fefab4eb8856a5a0ae6eb0a8406f56fed314.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) EvaluateRemoval()(*id8c6caa17cc80769edd1fce7e298996794e66f8c9a61b17a52901919cd49dde1.EvaluateRemovalRequestBuilder) {
    return id8c6caa17cc80769edd1fce7e298996794e66f8c9a61b17a52901919cd49dde1.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) ExtractContentLabel()(*ica62016c962f92b30229c778cbb076d6d246aadf7143abdcad3a76504b235726.ExtractContentLabelRequestBuilder) {
    return ica62016c962f92b30229c778cbb076d6d246aadf7143abdcad3a76504b235726.NewExtractContentLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get sensitivityLabels from users
func (m *SensitivityLabelsRequestBuilder) Get(options *SensitivityLabelsRequestBuilderGetOptions)(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.SensitivityLabelCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.CreateSensitivityLabelCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.SensitivityLabelCollectionResponseable), nil
}
// Post create new navigation property to sensitivityLabels for users
func (m *SensitivityLabelsRequestBuilder) Post(options *SensitivityLabelsRequestBuilderPostOptions)(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.SensitivityLabelable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.CreateSensitivityLabelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.SensitivityLabelable), nil
}
