package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia3d1361f46182d3550e67aa93c4815afd1bb111a192b4fae107e1737a85945f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/oemwarrantyinformationonboarding/item/enable"
    ia4101f7c03671dbe9e186fb7780ddc57b13b32ffba5901a1e744ec3b67d3b442 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/oemwarrantyinformationonboarding/item/disable"
)

// OemWarrantyInformationOnboardingItemRequestBuilder builds and executes requests for operations under \deviceManagement\oemWarrantyInformationOnboarding\{oemWarrantyInformationOnboarding-id}
type OemWarrantyInformationOnboardingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OemWarrantyInformationOnboardingItemRequestBuilderDeleteOptions options for Delete
type OemWarrantyInformationOnboardingItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OemWarrantyInformationOnboardingItemRequestBuilderGetOptions options for Get
type OemWarrantyInformationOnboardingItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters list of OEM Warranty Statuses
type OemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OemWarrantyInformationOnboardingItemRequestBuilderPatchOptions options for Patch
type OemWarrantyInformationOnboardingItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OemWarrantyInformationOnboarding;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOemWarrantyInformationOnboardingItemRequestBuilderInternal instantiates a new OemWarrantyInformationOnboardingItemRequestBuilder and sets the default values.
func NewOemWarrantyInformationOnboardingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OemWarrantyInformationOnboardingItemRequestBuilder) {
    m := &OemWarrantyInformationOnboardingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/oemWarrantyInformationOnboarding/{oemWarrantyInformationOnboarding_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOemWarrantyInformationOnboardingItemRequestBuilder instantiates a new OemWarrantyInformationOnboardingItemRequestBuilder and sets the default values.
func NewOemWarrantyInformationOnboardingItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OemWarrantyInformationOnboardingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOemWarrantyInformationOnboardingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation list of OEM Warranty Statuses
func (m *OemWarrantyInformationOnboardingItemRequestBuilder) CreateDeleteRequestInformation(options *OemWarrantyInformationOnboardingItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation list of OEM Warranty Statuses
func (m *OemWarrantyInformationOnboardingItemRequestBuilder) CreateGetRequestInformation(options *OemWarrantyInformationOnboardingItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation list of OEM Warranty Statuses
func (m *OemWarrantyInformationOnboardingItemRequestBuilder) CreatePatchRequestInformation(options *OemWarrantyInformationOnboardingItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete list of OEM Warranty Statuses
func (m *OemWarrantyInformationOnboardingItemRequestBuilder) Delete(options *OemWarrantyInformationOnboardingItemRequestBuilderDeleteOptions)(error) {
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
func (m *OemWarrantyInformationOnboardingItemRequestBuilder) Disable()(*ia4101f7c03671dbe9e186fb7780ddc57b13b32ffba5901a1e744ec3b67d3b442.DisableRequestBuilder) {
    return ia4101f7c03671dbe9e186fb7780ddc57b13b32ffba5901a1e744ec3b67d3b442.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OemWarrantyInformationOnboardingItemRequestBuilder) Enable()(*ia3d1361f46182d3550e67aa93c4815afd1bb111a192b4fae107e1737a85945f2.EnableRequestBuilder) {
    return ia3d1361f46182d3550e67aa93c4815afd1bb111a192b4fae107e1737a85945f2.NewEnableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get list of OEM Warranty Statuses
func (m *OemWarrantyInformationOnboardingItemRequestBuilder) Get(options *OemWarrantyInformationOnboardingItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OemWarrantyInformationOnboarding, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOemWarrantyInformationOnboarding() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OemWarrantyInformationOnboarding), nil
}
// Patch list of OEM Warranty Statuses
func (m *OemWarrantyInformationOnboardingItemRequestBuilder) Patch(options *OemWarrantyInformationOnboardingItemRequestBuilderPatchOptions)(error) {
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