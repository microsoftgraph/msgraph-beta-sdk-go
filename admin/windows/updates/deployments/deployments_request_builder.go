package deployments

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
    id8b32e020077ef38faed8c15c55faf2a027ee380c94ea635f30aaaa71996469c "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/count"
)

// DeploymentsRequestBuilder provides operations to manage the deployments property of the microsoft.graph.windowsUpdates.updates entity.
type DeploymentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeploymentsRequestBuilderGetOptions options for Get
type DeploymentsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeploymentsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeploymentsRequestBuilderGetQueryParameters deployments created using the deployment service. Read-only.
type DeploymentsRequestBuilderGetQueryParameters struct {
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
// DeploymentsRequestBuilderPostOptions options for Post
type DeploymentsRequestBuilderPostOptions struct {
    // 
    Body ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.Deploymentable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeploymentsRequestBuilderInternal instantiates a new DeploymentsRequestBuilder and sets the default values.
func NewDeploymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeploymentsRequestBuilder) {
    m := &DeploymentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeploymentsRequestBuilder instantiates a new DeploymentsRequestBuilder and sets the default values.
func NewDeploymentsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeploymentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeploymentsRequestBuilder) Count()(*id8b32e020077ef38faed8c15c55faf2a027ee380c94ea635f30aaaa71996469c.CountRequestBuilder) {
    return id8b32e020077ef38faed8c15c55faf2a027ee380c94ea635f30aaaa71996469c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation deployments created using the deployment service. Read-only.
func (m *DeploymentsRequestBuilder) CreateGetRequestInformation(options *DeploymentsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to deployments for admin
func (m *DeploymentsRequestBuilder) CreatePostRequestInformation(options *DeploymentsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get deployments created using the deployment service. Read-only.
func (m *DeploymentsRequestBuilder) Get(options *DeploymentsRequestBuilderGetOptions)(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.CreateDeploymentCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentCollectionResponseable), nil
}
// Post create new navigation property to deployments for admin
func (m *DeploymentsRequestBuilder) Post(options *DeploymentsRequestBuilderPostOptions)(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.Deploymentable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.CreateDeploymentFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.Deploymentable), nil
}
