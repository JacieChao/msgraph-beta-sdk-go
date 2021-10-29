package authentication

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i031886e282a72eeb426bbd544e45b4457ce2e270fc8c1a19fa2b007b18eeae51 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i09a669c0839eb56b1267ec6c47053774b28212601399fa8e394fe3659d37cbe8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/operations"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7181776de67abe7bec5b6b136451c674def2604be35b184f0aa413473fcc5701 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods"
    i9d172c880291778c29968c982a1fde24417fe56d76705022dd57555d79c61a0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/fido2methods"
    i9fafed747e4fd8dc14a1cd08e98cf200425cadc1d857dc77905604b4d814cebc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods"
    ic0756de5303fb580740539d17e229ce1a3ba8b519f526de32096715f3410b0dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/temporaryaccesspassmethods"
    ic150efac30f986d5b426bfc0c6529ba6d28832fce2cd65152e990426885bceda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods"
    ic1f96216136db886ff1afedeb281a1126c04793184c3ea4371f0f38018c9112d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/phonemethods"
    id91dbdbdfcc1f5a69fdff3c77ff512c8b4d1c46f0a74fb12686474b14fc391bf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordmethods"
    iefd4f63a7e855b21e1b4c41ac37f80e63563cce5d2ef96a65386065e5ddc1e68 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/emailmethods"
    i0059c2e4e23bd64f21f59deef18c5332cc205989dbcfed5623a379da011bf37c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/fido2methods/item"
    i01fb633a3822873540b4a37a19ffef4b3877b2e3724ff2a9e0ef15556e878183 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item"
    i1d6e1ee1e40c5bb5120f696f065400747d9f8b7c528e9db461c049031f562672 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item"
    i2b7ed9d4e22b5e330477a489ae98a8b1036a0bf39a69812e6c8f86f7902e6637 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/temporaryaccesspassmethods/item"
    i522f450607780460342b3cadbe08d49ca77462560e1b365ab8ce05addc5fb48e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/phonemethods/item"
    i96111e0c5ad54467d200eedd5128b62e6efcead28ff274fda8911bca83e2901f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item"
    id34bbd90b1fa28ce673d316987b63e1a63ca88a556236a01b0fc6f3f0dda42e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/operations/item"
    idb337e103dafe8aa86e240ec88da925669064777080064e8250bd1af4600efc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods/item"
    ide04a99d52e0e569423a7b87e2dd9af319adae8ce3b7091b35326abddb082980 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/emailmethods/item"
    iec467b36e4fdcc4d4a86d09fbaf10732c444a78b7a0400959fa67b6b0ca78e8e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordmethods/item"
)

type AuthenticationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type AuthenticationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationRequestBuilder) {
    m := &AuthenticationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/authentication{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewAuthenticationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation(q func (value *AuthenticationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AuthenticationRequestBuilderGetQueryParameters)
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
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Authentication, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AuthenticationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationRequestBuilder) EmailMethods()(*iefd4f63a7e855b21e1b4c41ac37f80e63563cce5d2ef96a65386065e5ddc1e68.EmailMethodsRequestBuilder) {
    return iefd4f63a7e855b21e1b4c41ac37f80e63563cce5d2ef96a65386065e5ddc1e68.NewEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) EmailMethodsById(id string)(*ide04a99d52e0e569423a7b87e2dd9af319adae8ce3b7091b35326abddb082980.EmailAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod_id"] = id
    }
    return ide04a99d52e0e569423a7b87e2dd9af319adae8ce3b7091b35326abddb082980.NewEmailAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) Fido2Methods()(*i9d172c880291778c29968c982a1fde24417fe56d76705022dd57555d79c61a0e.Fido2MethodsRequestBuilder) {
    return i9d172c880291778c29968c982a1fde24417fe56d76705022dd57555d79c61a0e.NewFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*i0059c2e4e23bd64f21f59deef18c5332cc205989dbcfed5623a379da011bf37c.Fido2AuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod_id"] = id
    }
    return i0059c2e4e23bd64f21f59deef18c5332cc205989dbcfed5623a379da011bf37c.NewFido2AuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) Get(q func (value *AuthenticationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Authentication, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAuthentication() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Authentication), nil
}
func (m *AuthenticationRequestBuilder) Methods()(*i7181776de67abe7bec5b6b136451c674def2604be35b184f0aa413473fcc5701.MethodsRequestBuilder) {
    return i7181776de67abe7bec5b6b136451c674def2604be35b184f0aa413473fcc5701.NewMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*idb337e103dafe8aa86e240ec88da925669064777080064e8250bd1af4600efc7.AuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod_id"] = id
    }
    return idb337e103dafe8aa86e240ec88da925669064777080064e8250bd1af4600efc7.NewAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*i9fafed747e4fd8dc14a1cd08e98cf200425cadc1d857dc77905604b4d814cebc.MicrosoftAuthenticatorMethodsRequestBuilder) {
    return i9fafed747e4fd8dc14a1cd08e98cf200425cadc1d857dc77905604b4d814cebc.NewMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*i01fb633a3822873540b4a37a19ffef4b3877b2e3724ff2a9e0ef15556e878183.MicrosoftAuthenticatorAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod_id"] = id
    }
    return i01fb633a3822873540b4a37a19ffef4b3877b2e3724ff2a9e0ef15556e878183.NewMicrosoftAuthenticatorAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) Operations()(*i09a669c0839eb56b1267ec6c47053774b28212601399fa8e394fe3659d37cbe8.OperationsRequestBuilder) {
    return i09a669c0839eb56b1267ec6c47053774b28212601399fa8e394fe3659d37cbe8.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) OperationsById(id string)(*id34bbd90b1fa28ce673d316987b63e1a63ca88a556236a01b0fc6f3f0dda42e7.LongRunningOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation_id"] = id
    }
    return id34bbd90b1fa28ce673d316987b63e1a63ca88a556236a01b0fc6f3f0dda42e7.NewLongRunningOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethods()(*i031886e282a72eeb426bbd544e45b4457ce2e270fc8c1a19fa2b007b18eeae51.PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    return i031886e282a72eeb426bbd544e45b4457ce2e270fc8c1a19fa2b007b18eeae51.NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethodsById(id string)(*i1d6e1ee1e40c5bb5120f696f065400747d9f8b7c528e9db461c049031f562672.PasswordlessMicrosoftAuthenticatorAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordlessMicrosoftAuthenticatorAuthenticationMethod_id"] = id
    }
    return i1d6e1ee1e40c5bb5120f696f065400747d9f8b7c528e9db461c049031f562672.NewPasswordlessMicrosoftAuthenticatorAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) PasswordMethods()(*id91dbdbdfcc1f5a69fdff3c77ff512c8b4d1c46f0a74fb12686474b14fc391bf.PasswordMethodsRequestBuilder) {
    return id91dbdbdfcc1f5a69fdff3c77ff512c8b4d1c46f0a74fb12686474b14fc391bf.NewPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) PasswordMethodsById(id string)(*iec467b36e4fdcc4d4a86d09fbaf10732c444a78b7a0400959fa67b6b0ca78e8e.PasswordAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod_id"] = id
    }
    return iec467b36e4fdcc4d4a86d09fbaf10732c444a78b7a0400959fa67b6b0ca78e8e.NewPasswordAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Authentication, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *AuthenticationRequestBuilder) PhoneMethods()(*ic1f96216136db886ff1afedeb281a1126c04793184c3ea4371f0f38018c9112d.PhoneMethodsRequestBuilder) {
    return ic1f96216136db886ff1afedeb281a1126c04793184c3ea4371f0f38018c9112d.NewPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) PhoneMethodsById(id string)(*i522f450607780460342b3cadbe08d49ca77462560e1b365ab8ce05addc5fb48e.PhoneAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod_id"] = id
    }
    return i522f450607780460342b3cadbe08d49ca77462560e1b365ab8ce05addc5fb48e.NewPhoneAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethods()(*ic0756de5303fb580740539d17e229ce1a3ba8b519f526de32096715f3410b0dc.TemporaryAccessPassMethodsRequestBuilder) {
    return ic0756de5303fb580740539d17e229ce1a3ba8b519f526de32096715f3410b0dc.NewTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*i2b7ed9d4e22b5e330477a489ae98a8b1036a0bf39a69812e6c8f86f7902e6637.TemporaryAccessPassAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod_id"] = id
    }
    return i2b7ed9d4e22b5e330477a489ae98a8b1036a0bf39a69812e6c8f86f7902e6637.NewTemporaryAccessPassAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*ic150efac30f986d5b426bfc0c6529ba6d28832fce2cd65152e990426885bceda.WindowsHelloForBusinessMethodsRequestBuilder) {
    return ic150efac30f986d5b426bfc0c6529ba6d28832fce2cd65152e990426885bceda.NewWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*i96111e0c5ad54467d200eedd5128b62e6efcead28ff274fda8911bca83e2901f.WindowsHelloForBusinessAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod_id"] = id
    }
    return i96111e0c5ad54467d200eedd5128b62e6efcead28ff274fda8911bca83e2901f.NewWindowsHelloForBusinessAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}