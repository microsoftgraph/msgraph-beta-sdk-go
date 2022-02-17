package audience

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i27c073677df2af7a178b750045cf7538776463d44b1fdb96205037112761934e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members"
    i6ca7350a454d0eb7416a3ffe6d406c2f9a011b9818e8a5fc6714fbfab983c426 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/updateaudience"
    i9d78b5d7e63632aad0dc14fc18eec733cc1a34e41c1bd899b904db2431e7bbde "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions"
    ibb560c7d9d162fcf5a50b3a20535d73efbfb6eada7792589ef7f79232af0cd28 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/updateaudiencebyid"
    icf6ee88716e5da88403b56c76157382dd11930a0110b7137be42bf2fb2425058 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members/item"
    ifa1c023a93cedbc885180b644196fd22db8e7768dfa64cd6a803d74884273eac "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item"
)

// AudienceRequestBuilder builds and executes requests for operations under \admin\windows\updates\deployments\{deployment-id}\audience
type AudienceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AudienceRequestBuilderDeleteOptions options for Delete
type AudienceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AudienceRequestBuilderGetOptions options for Get
type AudienceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AudienceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AudienceRequestBuilderGetQueryParameters specifies the audience to which content is deployed.
type AudienceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AudienceRequestBuilderPatchOptions options for Patch
type AudienceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeploymentAudience;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAudienceRequestBuilderInternal instantiates a new AudienceRequestBuilder and sets the default values.
func NewAudienceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AudienceRequestBuilder) {
    m := &AudienceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment_id}/audience{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAudienceRequestBuilder instantiates a new AudienceRequestBuilder and sets the default values.
func NewAudienceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AudienceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAudienceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation specifies the audience to which content is deployed.
func (m *AudienceRequestBuilder) CreateDeleteRequestInformation(options *AudienceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation specifies the audience to which content is deployed.
func (m *AudienceRequestBuilder) CreateGetRequestInformation(options *AudienceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation specifies the audience to which content is deployed.
func (m *AudienceRequestBuilder) CreatePatchRequestInformation(options *AudienceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete specifies the audience to which content is deployed.
func (m *AudienceRequestBuilder) Delete(options *AudienceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AudienceRequestBuilder) Exclusions()(*i9d78b5d7e63632aad0dc14fc18eec733cc1a34e41c1bd899b904db2431e7bbde.ExclusionsRequestBuilder) {
    return i9d78b5d7e63632aad0dc14fc18eec733cc1a34e41c1bd899b904db2431e7bbde.NewExclusionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExclusionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.admin.windows.updates.deployments.item.audience.exclusions.item collection
func (m *AudienceRequestBuilder) ExclusionsById(id string)(*ifa1c023a93cedbc885180b644196fd22db8e7768dfa64cd6a803d74884273eac.UpdatableAssetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset_id"] = id
    }
    return ifa1c023a93cedbc885180b644196fd22db8e7768dfa64cd6a803d74884273eac.NewUpdatableAssetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get specifies the audience to which content is deployed.
func (m *AudienceRequestBuilder) Get(options *AudienceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeploymentAudience, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeploymentAudience() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeploymentAudience), nil
}
func (m *AudienceRequestBuilder) Members()(*i27c073677df2af7a178b750045cf7538776463d44b1fdb96205037112761934e.MembersRequestBuilder) {
    return i27c073677df2af7a178b750045cf7538776463d44b1fdb96205037112761934e.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.admin.windows.updates.deployments.item.audience.members.item collection
func (m *AudienceRequestBuilder) MembersById(id string)(*icf6ee88716e5da88403b56c76157382dd11930a0110b7137be42bf2fb2425058.UpdatableAssetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset_id"] = id
    }
    return icf6ee88716e5da88403b56c76157382dd11930a0110b7137be42bf2fb2425058.NewUpdatableAssetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch specifies the audience to which content is deployed.
func (m *AudienceRequestBuilder) Patch(options *AudienceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AudienceRequestBuilder) UpdateAudience()(*i6ca7350a454d0eb7416a3ffe6d406c2f9a011b9818e8a5fc6714fbfab983c426.UpdateAudienceRequestBuilder) {
    return i6ca7350a454d0eb7416a3ffe6d406c2f9a011b9818e8a5fc6714fbfab983c426.NewUpdateAudienceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AudienceRequestBuilder) UpdateAudienceById()(*ibb560c7d9d162fcf5a50b3a20535d73efbfb6eada7792589ef7f79232af0cd28.UpdateAudienceByIdRequestBuilder) {
    return ibb560c7d9d162fcf5a50b3a20535d73efbfb6eada7792589ef7f79232af0cd28.NewUpdateAudienceByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
