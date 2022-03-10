package serviceprincipals

import (
    i06ff0d811046a74890f51edd3cff651bef78898b54f602513c8c7a5b32cecf38 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/count"
    i07ea82e8cdb01938ad025e6648ae29c3c96f2adbafa7de6af15c650980651744 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/delta"
    i13ad01682da78442a9cf027f302e407c5d377fb9738dde4e75f036b89203db62 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/getbyids"
    i59f9b4fdf3f63fe1a7bfbabda2130077511ff77c1f090e63ad3afc1690b2a0b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/validateproperties"
    ic1b17114e49e92b3da03ac7e6328ba8f143e419e9a0cb2b80fe833a35584283a "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/getuserownedobjects"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// ServicePrincipalsRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ServicePrincipalsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServicePrincipalsRequestBuilderGetOptions options for Get
type ServicePrincipalsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServicePrincipalsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServicePrincipalsRequestBuilderGetQueryParameters get entities from servicePrincipals
type ServicePrincipalsRequestBuilderGetQueryParameters struct {
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
// ServicePrincipalsRequestBuilderPostOptions options for Post
type ServicePrincipalsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewServicePrincipalsRequestBuilderInternal instantiates a new ServicePrincipalsRequestBuilder and sets the default values.
func NewServicePrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalsRequestBuilder) {
    m := &ServicePrincipalsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsRequestBuilder instantiates a new ServicePrincipalsRequestBuilder and sets the default values.
func NewServicePrincipalsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ServicePrincipalsRequestBuilder) Count()(*i06ff0d811046a74890f51edd3cff651bef78898b54f602513c8c7a5b32cecf38.CountRequestBuilder) {
    return i06ff0d811046a74890f51edd3cff651bef78898b54f602513c8c7a5b32cecf38.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from servicePrincipals
func (m *ServicePrincipalsRequestBuilder) CreateGetRequestInformation(options *ServicePrincipalsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to servicePrincipals
func (m *ServicePrincipalsRequestBuilder) CreatePostRequestInformation(options *ServicePrincipalsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *ServicePrincipalsRequestBuilder) Delta()(*i07ea82e8cdb01938ad025e6648ae29c3c96f2adbafa7de6af15c650980651744.DeltaRequestBuilder) {
    return i07ea82e8cdb01938ad025e6648ae29c3c96f2adbafa7de6af15c650980651744.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from servicePrincipals
func (m *ServicePrincipalsRequestBuilder) Get(options *ServicePrincipalsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateServicePrincipalCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalCollectionResponseable), nil
}
func (m *ServicePrincipalsRequestBuilder) GetByIds()(*i13ad01682da78442a9cf027f302e407c5d377fb9738dde4e75f036b89203db62.GetByIdsRequestBuilder) {
    return i13ad01682da78442a9cf027f302e407c5d377fb9738dde4e75f036b89203db62.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalsRequestBuilder) GetUserOwnedObjects()(*ic1b17114e49e92b3da03ac7e6328ba8f143e419e9a0cb2b80fe833a35584283a.GetUserOwnedObjectsRequestBuilder) {
    return ic1b17114e49e92b3da03ac7e6328ba8f143e419e9a0cb2b80fe833a35584283a.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to servicePrincipals
func (m *ServicePrincipalsRequestBuilder) Post(options *ServicePrincipalsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateServicePrincipalFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalable), nil
}
func (m *ServicePrincipalsRequestBuilder) ValidateProperties()(*i59f9b4fdf3f63fe1a7bfbabda2130077511ff77c1f090e63ad3afc1690b2a0b7.ValidatePropertiesRequestBuilder) {
    return i59f9b4fdf3f63fe1a7bfbabda2130077511ff77c1f090e63ad3afc1690b2a0b7.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
