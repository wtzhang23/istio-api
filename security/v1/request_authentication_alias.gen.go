// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1

import "istio.io/api/security/v1beta1"

// <!-- crd generation tags
// +cue-gen:RequestAuthentication:groupName:security.istio.io
// +cue-gen:RequestAuthentication:versions:v1beta1,v1
// +cue-gen:RequestAuthentication:storageVersion
// +cue-gen:RequestAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:RequestAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:RequestAuthentication:subresource:status
// +cue-gen:RequestAuthentication:scope:Namespaced
// +cue-gen:RequestAuthentication:resource:categories=istio-io,security-istio-io,shortNames=ra
// +cue-gen:RequestAuthentication:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// +kubebuilder:validation:XValidation:message="only one of targetRefs or selector can be set",rule="oneof(self.selector, self.targetRef, self.targetRefs)"
type RequestAuthentication = v1beta1.RequestAuthentication

// JSON Web Token (JWT) token format for authentication as defined by
// [RFC 7519](https://tools.ietf.org/html/rfc7519). See [OAuth 2.0](https://tools.ietf.org/html/rfc6749) and
// [OIDC 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// Examples:
//
// Spec for a JWT that is issued by `https://example.com`, with the audience claims must be either
// `bookstore_android.apps.example.com` or `bookstore_web.apps.example.com`.
// The token should be presented at the `Authorization` header (default). The JSON Web Key Set (JWKS)
// will be discovered following OpenID Connect protocol.
//
// ```yaml
// issuer: https://example.com
// audiences:
//   - bookstore_android.apps.example.com
//     bookstore_web.apps.example.com
//
// ```
//
// This example specifies a token in a non-default location (`x-goog-iap-jwt-assertion` header). It also
// defines the URI to fetch JWKS explicitly.
//
// ```yaml
// issuer: https://example.com
// jwksUri: https://example.com/.secret/jwks.json
// fromHeaders:
// - "x-goog-iap-jwt-assertion"
// ```
// +kubebuilder:validation:XValidation:message="only one of jwks or jwksUri can be set",rule="oneof(self.jwksUri, self.jwks_uri, self.jwks)"
type JWTRule = v1beta1.JWTRule

// This message specifies a header location to extract JWT token.
type JWTHeader = v1beta1.JWTHeader

// This message specifies the detail for copying claim to header.
type ClaimToHeader = v1beta1.ClaimToHeader
