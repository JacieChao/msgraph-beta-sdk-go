package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder casts the previous resource to application.
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.application
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderGetQueryParameters
}
// NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderInternal instantiates a new ApplicationRequestBuilder and sets the default values.
func NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder) {
    m := &AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.application{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder instantiates a new ApplicationRequestBuilder and sets the default values.
func NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.application
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.application
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}