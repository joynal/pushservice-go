package utils

import (
  "crypto/elliptic"
  "encoding/base64"

  "github.com/enceve/crypto/dh/ecdh"
)

// GenerateVapidKeys generates dummy vapid
func GenerateVapidKeys() (string, string, error) {
  curve := ecdh.GenericCurve(elliptic.P256())
  privateKey, publicKey, err := curve.GenerateKey(nil)
  if err != nil {
    return "", "", err
  }
  privKey := base64.RawURLEncoding.EncodeToString(privateKey)
  pubKey := base64.RawURLEncoding.EncodeToString(publicKey)
  return privKey, pubKey, nil
}
