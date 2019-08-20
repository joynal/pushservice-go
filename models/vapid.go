package models

// Site model type
type VapidDetails struct {
  Subject    string
  PublicKey  string    `bson:"publicKey"`
  PrivateKey string    `bson:"privateKey"`
}
