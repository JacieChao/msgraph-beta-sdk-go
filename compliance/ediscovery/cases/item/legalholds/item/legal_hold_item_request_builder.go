package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/usersources"
    i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/sitesources"
    iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/unifiedgroupsources"
    i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/usersources/item"
    ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/sitesources/item"
    ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/unifiedgroupsources/item"
)

// LegalHoldItemRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\legalHolds\{legalHold-id}
type LegalHoldItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// LegalHoldItemRequestBuilderDeleteOptions options for Delete
type LegalHoldItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// LegalHoldItemRequestBuilderGetOptions options for Get
type LegalHoldItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *LegalHoldItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// LegalHoldItemRequestBuilderGetQueryParameters returns a list of case legalHold objects for this case.  Nullable.
type LegalHoldItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// LegalHoldItemRequestBuilderPatchOptions options for Patch
type LegalHoldItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LegalHold;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewLegalHoldItemRequestBuilderInternal instantiates a new LegalHoldItemRequestBuilder and sets the default values.
func NewLegalHoldItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LegalHoldItemRequestBuilder) {
    m := &LegalHoldItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/legalHolds/{legalHold_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLegalHoldItemRequestBuilder instantiates a new LegalHoldItemRequestBuilder and sets the default values.
func NewLegalHoldItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LegalHoldItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLegalHoldItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation returns a list of case legalHold objects for this case.  Nullable.
func (m *LegalHoldItemRequestBuilder) CreateDeleteRequestInformation(options *LegalHoldItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of case legalHold objects for this case.  Nullable.
func (m *LegalHoldItemRequestBuilder) CreateGetRequestInformation(options *LegalHoldItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation returns a list of case legalHold objects for this case.  Nullable.
func (m *LegalHoldItemRequestBuilder) CreatePatchRequestInformation(options *LegalHoldItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete returns a list of case legalHold objects for this case.  Nullable.
func (m *LegalHoldItemRequestBuilder) Delete(options *LegalHoldItemRequestBuilderDeleteOptions)(error) {
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
// Get returns a list of case legalHold objects for this case.  Nullable.
func (m *LegalHoldItemRequestBuilder) Get(options *LegalHoldItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LegalHold, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewLegalHold() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LegalHold), nil
}
// Patch returns a list of case legalHold objects for this case.  Nullable.
func (m *LegalHoldItemRequestBuilder) Patch(options *LegalHoldItemRequestBuilderPatchOptions)(error) {
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
func (m *LegalHoldItemRequestBuilder) SiteSources()(*i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b.SiteSourcesRequestBuilder) {
    return i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b.NewSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item.siteSources.item collection
func (m *LegalHoldItemRequestBuilder) SiteSourcesById(id string)(*ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936.SiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource_id"] = id
    }
    return ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936.NewSiteSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *LegalHoldItemRequestBuilder) UnifiedGroupSources()(*iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8.UnifiedGroupSourcesRequestBuilder) {
    return iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8.NewUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnifiedGroupSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item.unifiedGroupSources.item collection
func (m *LegalHoldItemRequestBuilder) UnifiedGroupSourcesById(id string)(*ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760.UnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource_id"] = id
    }
    return ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760.NewUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *LegalHoldItemRequestBuilder) UserSources()(*i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315.UserSourcesRequestBuilder) {
    return i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315.NewUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item.userSources.item collection
func (m *LegalHoldItemRequestBuilder) UserSourcesById(id string)(*i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060.UserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource_id"] = id
    }
    return i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060.NewUserSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}