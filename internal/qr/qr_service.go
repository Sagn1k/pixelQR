package qr

import (
    "encoding/base64"
    "github.com/skip2/go-qrcode"
)

func GenerateQRCodeBase64(payload string) (string, error) {
    var png []byte
    png, err := qrcode.Encode(payload, qrcode.Medium, 256)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(png), nil
}
