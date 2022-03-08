package riskyserviceprincipals

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i418a1a5558948b4662e488386dc90986e3fc661f4e22322afbb353d02d78d298 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/dismiss"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i85dee32ed1d0f963b95f0e6653515c81e7dc0dc00db2e6cc11adfd6b4476e28f "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/confirmcompromised"
    ib48ad338b29836e614d9379952d94ca511eae67832181629704a82189ddf853c "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/count"
)

// RiskyServicePrincipalsRequestBuilder provides operations to manage the riskyServicePrincipals property of the microsoft.graph.identityProtectionRoot entity.
type RiskyServicePrincipalsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RiskyServicePrincipalsRequestBuilderGetOptions options for Get
type RiskyServicePrincipalsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RiskyServicePrincipalsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RiskyServicePrincipalsRequestBuilderGetQueryParameters azure AD service principals that are at risk.
type RiskyServicePrincipalsRequestBuilderGetQueryParameters struct {
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
// RiskyServicePrincipalsRequestBuilderPostOptions options for Post
type RiskyServicePrincipalsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RiskyServicePrincipalable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RiskyServicePrincipalsRequestBuilder) ConfirmCompromised()(*i85dee32ed1d0f963b95f0e6653515c81e7dc0dc00db2e6cc11adfd6b4476e28f.ConfirmCompromisedRequestBuilder) {
    return i85dee32ed1d0f963b95f0e6653515c81e7dc0dc00db2e6cc11adfd6b4476e28f.NewConfirmCompromisedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRiskyServicePrincipalsRequestBuilderInternal instantiates a new RiskyServicePrincipalsRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RiskyServicePrincipalsRequestBuilder) {
    m := &RiskyServicePrincipalsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyServicePrincipals{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRiskyServicePrincipalsRequestBuilder instantiates a new RiskyServicePrincipalsRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RiskyServicePrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyServicePrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RiskyServicePrincipalsRequestBuilder) Count()(*ib48ad338b29836e614d9379952d94ca511eae67832181629704a82189ddf853c.CountRequestBuilder) {
    return ib48ad338b29836e614d9379952d94ca511eae67832181629704a82189ddf853c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation azure AD service principals that are at risk.
func (m *RiskyServicePrincipalsRequestBuilder) CreateGetRequestInformation(options *RiskyServicePrincipalsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to riskyServicePrincipals for identityProtection
func (m *RiskyServicePrincipalsRequestBuilder) CreatePostRequestInformation(options *RiskyServicePrincipalsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *RiskyServicePrincipalsRequestBuilder) Dismiss()(*i418a1a5558948b4662e488386dc90986e3fc661f4e22322afbb353d02d78d298.DismissRequestBuilder) {
    return i418a1a5558948b4662e488386dc90986e3fc661f4e22322afbb353d02d78d298.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get azure AD service principals that are at risk.
func (m *RiskyServicePrincipalsRequestBuilder) Get(options *RiskyServicePrincipalsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RiskyServicePrincipalCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateRiskyServicePrincipalCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RiskyServicePrincipalCollectionResponseable), nil
}
// Post create new navigation property to riskyServicePrincipals for identityProtection
func (m *RiskyServicePrincipalsRequestBuilder) Post(options *RiskyServicePrincipalsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RiskyServicePrincipalable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateRiskyServicePrincipalFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RiskyServicePrincipalable), nil
}
