package directory

import (
    i099820decfc9bea0a9207fe8eabbdbec38a02bc9f29b8c032c0f3dc2b1de2037 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits"
    i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/sharedemaildomains"
    i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems"
    i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item"
    i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/sharedemaildomains/item"
    i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item"
    ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item"
)

type DirectoryRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DirectoryRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *DirectoryRequestBuilder) AdministrativeUnits()(*i099820decfc9bea0a9207fe8eabbdbec38a02bc9f29b8c032c0f3dc2b1de2037.AdministrativeUnitsRequestBuilder) {
    return i099820decfc9bea0a9207fe8eabbdbec38a02bc9f29b8c032c0f3dc2b1de2037.NewAdministrativeUnitsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) AdministrativeUnitsById(id string)(*i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a.AdministrativeUnitRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["administrativeUnit_id"] = id
    }
    return i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a.NewAdministrativeUnitRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/directory{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DirectoryRequestBuilder) CreateGetRequestInformation(q func (value *DirectoryRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DirectoryRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DirectoryRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directory, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DirectoryRequestBuilder) DeletedItems()(*i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a.DeletedItemsRequestBuilder) {
    return i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a.NewDeletedItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) DeletedItemsById(id string)(*ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b.DirectoryObjectRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b.NewDirectoryObjectRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) FeatureRolloutPolicies()(*i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf.FeatureRolloutPoliciesRequestBuilder) {
    return i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf.NewFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) FeatureRolloutPoliciesById(id string)(*i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4.FeatureRolloutPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy_id"] = id
    }
    return i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4.NewFeatureRolloutPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) Get(q func (value *DirectoryRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directory, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectory() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directory), nil
}
func (m *DirectoryRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directory, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DirectoryRequestBuilder) SharedEmailDomains()(*i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d.SharedEmailDomainsRequestBuilder) {
    return i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d.NewSharedEmailDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) SharedEmailDomainsById(id string)(*i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a.SharedEmailDomainRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["sharedEmailDomain_id"] = id
    }
    return i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a.NewSharedEmailDomainRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
