package main

import (
	"bufio"
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"sort"
	"testing"
	"time"
	"unicode"
)

func TestMethod(t *testing.T) {
	fmt.Println("Testing message")
	testRootPEM := `
-----BEGIN CERTIFICATE-----
MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
/iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
-----END CERTIFICATE-----`

	testCertPEM := `
-----BEGIN CERTIFICATE-----
MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
TZZyzmEOyclm3UGFYe82P/iDFt+CeQ3NpmBg+GoaVCuWAARJN/KfglbLyyYygcQq
0SgeDh8dRKUiaW3HQSoYvTvdTuqzwK4CXsr3b5/dAOY8uMuG/IAR3FgwTbZ1dtoW
RvOTa8hYiU6A475WuZKyEHcwnGYe57u2I2KbMgcKjPniocj4QzgYsVAVKW3IwaOh
yE+vPxsiUkvQHdO2fojCkY8jg70jxM+gu59tPDNbw3Uh/2Ij310FgTHsnGQMyA==
-----END CERTIFICATE-----`

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(testRootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(testCertPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	currentTime := time.Date(2014, time.May, 1, 0, 0, 0, 0, time.UTC)
	opts := x509.VerifyOptions{
		DNSName:     "mail.google.com",
		Roots:       roots,
		CurrentTime: currentTime,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}

	fmt.Println("Test passed!")

}

func TestMethod3(t *testing.T) {
	fmt.Println("Testing message")
	testRootPEM := `
-----BEGIN CERTIFICATE-----
MIIF+zCCA+OgAwIBAgIJAJgFPit5jlhyMA0GCSqGSIb3DQEBCwUAMIGTMQswCQYD
VQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJUGFsbyBBbHRv
MRUwEwYDVQQKDAxWTXdhcmUsIEluYy4xHjAcBgNVBAsMFUNvbnRhaW5lcnMgb24g
dlNwaGVyZTEkMCIGA1UEAwwbU2VsZi1zaWduZWQgYnkgVk13YXJlLCBJbmMuMB4X
DTE4MDExMDE2Mzk0OVoXDTIxMDEwOTE2Mzk0OVowgZMxCzAJBgNVBAYTAlVTMRMw
EQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlQYWxvIEFsdG8xFTATBgNVBAoM
DFZNd2FyZSwgSW5jLjEeMBwGA1UECwwVQ29udGFpbmVycyBvbiB2U3BoZXJlMSQw
IgYDVQQDDBtTZWxmLXNpZ25lZCBieSBWTXdhcmUsIEluYy4wggIiMA0GCSqGSIb3
DQEBAQUAA4ICDwAwggIKAoICAQCthA+rCkux2yQUaGatOIIw+QvNWK2+pp/+2qeJ
75EBc/QMnKjWmANfh3yOteRomf6nL5mu2zl0U6wleiqn5H/gGJ/LdeBWQPp7MfPV
AwrCCoX7zBMtq72oYDhnJIpRntMIi0cN8DzDJHKDBdn/wMl3+XmFqf5Unk86D0/A
nEVFiHcuvkAeDjPFsuezsK1kNpiC3CyT0li3doK6fAEF6QG36mvX4/enkTtcu9AZ
dw4Qm13cEZzg3+J6qMO0jc8MMF4zj5OXqJRvQGx6EJbvT/QV6h4RBoYKAJZuaipz
xrEEZATtDzTGs7SeynaFjSfW/c7hiicRLj2fmUn/nXii4GjuAn/8vtWRDPZVKjYM
rdFJzPmQfyFtmLZKI0jLw2Utkbhaz2uHu0viH7gX7YvjWYwiZBT7NbUvs50l58bp
RV7ff35E9XlH2xjVBeY+cu9SAd7mKZfLl506qQhQAq10dLyo0G4sXyiqhIPswj+o
L3Oh/s38RtZbQl2aNefVqyujW9eD3TF0lyG1ZhsjtcdFUndRNeezeRdCRuJ9WqSL
kFeKf5p8a24SijP33FRz9sZKKNwLBdwGq+eycLtxCMDQ69w6JJYOFljqVBz2dFvj
ICZal0orq5Cl9Pq+Bg73YzAM5dZxh99Ykiltrr6deRnMpLraKE/sHOLe6ZWr9aai
AsrirwIDAQABo1AwTjAdBgNVHQ4EFgQUqhxDXsn2wnZpf4Qpv7QdkBvRD3cwHwYD
VR0jBBgwFoAUqhxDXsn2wnZpf4Qpv7QdkBvRD3cwDAYDVR0TBAUwAwEB/zANBgkq
hkiG9w0BAQsFAAOCAgEAU67C37VunxUFn1rpvODLVMitWYI0f8ugqYyNEPXHNTv/
wKrhdmrQWSuGThX7UNSTP3xflVCzAHj5T4OdhUIKOr07VW9WX4mTEpXiqee4Idly
56FHwfe9IjXwQx1Gsh+4J178rRuLpJMgNBTHAov7L1qYQyj0P+hJ+8eefzs1ZTaH
uaO1HvoMtMtSCVfOUewaCt8t19YwV9JQoSUeZUqxP6PUuXwPZuRkYFzSAmJ+NzEU
o1LVvqLbGXU/BruVGZ6tDUKXlJTev7WpQ/p3GXGm2yX492ClQ9bJLhWz77rv29JT
5rv148gsdk91Zt1jYrZGad3oPZLBasgWXsywNEt6CwyYspyCO4YL2NnmuoGo9T5N
4vFKWic9tZ+7Hkc/ifsWnm1K2RtlTwRFRyz33RBGoumKV/in+VCYXbbqG8fGam1h
dmH8AD3ac5FLhmW2H3vqK6j/Rbr8pXtFAqniYyrCBF3kB4hekLXkg5B9dbF3lflG
t0t1H7uQTjrl7stZ9FQCWli/cLaBx1k0c0T4efM6GSVdaj7SLHFlMYhvPTCFD76x
DEwJt0Dw8z/IlvTdjYnrIzH21wN7lJWXB6RkxhdAqz3Ibn4uZSkqmF24wErShwL+
M7P3UFXo3NSQ3p01HramFNEZrMWRmm0ab7d8BayyJzHr7uv/PEheem4kEuj3VIw=
-----END CERTIFICATE-----`

	testCertPEM := `
-----BEGIN CERTIFICATE-----
MIIFrzCCA5egAwIBAgIJAOulY8CKAGd/MA0GCSqGSIb3DQEBCwUAMIGTMQswCQYD
VQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJUGFsbyBBbHRv
MRUwEwYDVQQKDAxWTXdhcmUsIEluYy4xHjAcBgNVBAsMFUNvbnRhaW5lcnMgb24g
dlNwaGVyZTEkMCIGA1UEAwwbU2VsZi1zaWduZWQgYnkgVk13YXJlLCBJbmMuMB4X
DTE4MDExMDE2Mzk1MFoXDTIxMDEwOTE2Mzk1MFowgYQxCzAJBgNVBAYTAlVTMRMw
EQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlQYWxvIEFsdG8xDzANBgNVBAoM
BlZNd2FyZTEeMBwGA1UECwwVQ29udGFpbmVycyBvbiB2U3BoZXJlMRswGQYDVQQD
DBJ2aWMtdG9tLmNvcnAubG9jYWwwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK
AoICAQDGJNfsfr2+ArMiWqk6dx9D7by8KDPw/XoT/6RTcUzstdZVjPp32bEniCDQ
VTCJyADBgTcm/GLdKNZT0ASPuG2PhLHFamM8sh01qzwRO5xoIfmJvZSZxIoKqS4D
BAxyqBYf+pzH8srCKCQLxXyw+0RAS2ndiNoP0WtpkEyNqpEgak/jsh+V543hi74G
HGQmtr35yvQmXIuSZFM+k4K/rQ/jdQsovzQeeWr+eTD4kDTRse+WYp335UPB88Zi
e44cUbhl+K5azJBtYAjvQJBP5uBeH57uidmO0C9+qpCizH++kO3tq56VBzTywchD
wKXWGkThgo6aypmjJRrKs7KaqDetdf9iIYvnOoIFRbwu2YdK1zbA56VCWpJ0juMZ
W1Fxji3woS5zsnl5neG1OlH7uWcGdRhcivxle8G6FAQVwSneB+bRd8iT5Y1K1PCu
Q9C/DXUoILgGeEEEDOyYH92z06rlQx2Kg2PBr7S7UtkbFPzJU5bY0yPmoS+yWfan
ebIVMwqbMr9FwREeVyCancxzkg8nXTAHBkYkRVL13Fyit7eWpU/sECqKcKkoMDsc
U7KQLXo0RdFPgovt2h0+CkEQlyiSpB4kbxwEnE2EjGrMa64/H2EU+udovwk5IqGe
nIz1LVEr32bobPRCDqUMTU78b5qAemhpsuIllhGXeMTJlBiHbwIDAQABoxMwETAP
BgNVHREECDAGhwQKnswjMA0GCSqGSIb3DQEBCwUAA4ICAQA/dG7XJg7jgJGZTe8T
1Uc8trqSQQnYANcBlOcyRy/Uh8BS7mz26h3RozHaMjpLtCqYRboDYQYm+GRTO3Ym
tZz2cmz82VU26yDlTAjnEPoeZ+Za0loSC7kG40+vzsZqUkcboAgQCllyB4mlBl5s
rrf6LK5OuVhek4KKc24M5X5FV7Ak2ylWWjpLaEew9IoN0yg/fKI9ffqug/eYqUrf
hMyvFpodIIPZHuBQ9nzZyEM2pVXOIo4kS/s4KMdJjhRvX5HlQXI6v37gyVgkw1JH
fsw1nmiiuk56GmJStLv6aF1tkQaRFb3tnxExR81CNitbzH9qYFNKJg/6hQZhdUOS
Dkyjfip1uB3X+1wbUfJbj2ldULx7BJ8xemdmU+aGXwDUPL+xqn+m/8XlYG4FlF9g
mjpn5ThWC+Xx4/RU+SqIv75m3c2pO233oMQCw+gKr2uVSFH+vHskvwVP21FVRT/t
WQecX6tAE8Hi7KttEOOuLn9iOMJCUSdvGHEfLcrl0efpPPOe60mpb+1vs/8SbIm1
YptyoNv73l3yZT1wK0HR5af0bELKvEjd6ySETXJU4THWaACCRdRaeGYGXTESmXqJ
oUf5ef0rLe08ujyWUfgzBlMrJofOHPhUg25pD/vCg0pHvaqbKxbbCehqH1FYCbI4
KvwmXmto4EWXIh621yZrs1AbkQ==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIF+zCCA+OgAwIBAgIJAJgFPit5jlhyMA0GCSqGSIb3DQEBCwUAMIGTMQswCQYD
VQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJUGFsbyBBbHRv
MRUwEwYDVQQKDAxWTXdhcmUsIEluYy4xHjAcBgNVBAsMFUNvbnRhaW5lcnMgb24g
dlNwaGVyZTEkMCIGA1UEAwwbU2VsZi1zaWduZWQgYnkgVk13YXJlLCBJbmMuMB4X
DTE4MDExMDE2Mzk0OVoXDTIxMDEwOTE2Mzk0OVowgZMxCzAJBgNVBAYTAlVTMRMw
EQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlQYWxvIEFsdG8xFTATBgNVBAoM
DFZNd2FyZSwgSW5jLjEeMBwGA1UECwwVQ29udGFpbmVycyBvbiB2U3BoZXJlMSQw
IgYDVQQDDBtTZWxmLXNpZ25lZCBieSBWTXdhcmUsIEluYy4wggIiMA0GCSqGSIb3
DQEBAQUAA4ICDwAwggIKAoICAQCthA+rCkux2yQUaGatOIIw+QvNWK2+pp/+2qeJ
75EBc/QMnKjWmANfh3yOteRomf6nL5mu2zl0U6wleiqn5H/gGJ/LdeBWQPp7MfPV
AwrCCoX7zBMtq72oYDhnJIpRntMIi0cN8DzDJHKDBdn/wMl3+XmFqf5Unk86D0/A
nEVFiHcuvkAeDjPFsuezsK1kNpiC3CyT0li3doK6fAEF6QG36mvX4/enkTtcu9AZ
dw4Qm13cEZzg3+J6qMO0jc8MMF4zj5OXqJRvQGx6EJbvT/QV6h4RBoYKAJZuaipz
xrEEZATtDzTGs7SeynaFjSfW/c7hiicRLj2fmUn/nXii4GjuAn/8vtWRDPZVKjYM
rdFJzPmQfyFtmLZKI0jLw2Utkbhaz2uHu0viH7gX7YvjWYwiZBT7NbUvs50l58bp
RV7ff35E9XlH2xjVBeY+cu9SAd7mKZfLl506qQhQAq10dLyo0G4sXyiqhIPswj+o
L3Oh/s38RtZbQl2aNefVqyujW9eD3TF0lyG1ZhsjtcdFUndRNeezeRdCRuJ9WqSL
kFeKf5p8a24SijP33FRz9sZKKNwLBdwGq+eycLtxCMDQ69w6JJYOFljqVBz2dFvj
ICZal0orq5Cl9Pq+Bg73YzAM5dZxh99Ykiltrr6deRnMpLraKE/sHOLe6ZWr9aai
AsrirwIDAQABo1AwTjAdBgNVHQ4EFgQUqhxDXsn2wnZpf4Qpv7QdkBvRD3cwHwYD
VR0jBBgwFoAUqhxDXsn2wnZpf4Qpv7QdkBvRD3cwDAYDVR0TBAUwAwEB/zANBgkq
hkiG9w0BAQsFAAOCAgEAU67C37VunxUFn1rpvODLVMitWYI0f8ugqYyNEPXHNTv/
wKrhdmrQWSuGThX7UNSTP3xflVCzAHj5T4OdhUIKOr07VW9WX4mTEpXiqee4Idly
56FHwfe9IjXwQx1Gsh+4J178rRuLpJMgNBTHAov7L1qYQyj0P+hJ+8eefzs1ZTaH
uaO1HvoMtMtSCVfOUewaCt8t19YwV9JQoSUeZUqxP6PUuXwPZuRkYFzSAmJ+NzEU
o1LVvqLbGXU/BruVGZ6tDUKXlJTev7WpQ/p3GXGm2yX492ClQ9bJLhWz77rv29JT
5rv148gsdk91Zt1jYrZGad3oPZLBasgWXsywNEt6CwyYspyCO4YL2NnmuoGo9T5N
4vFKWic9tZ+7Hkc/ifsWnm1K2RtlTwRFRyz33RBGoumKV/in+VCYXbbqG8fGam1h
dmH8AD3ac5FLhmW2H3vqK6j/Rbr8pXtFAqniYyrCBF3kB4hekLXkg5B9dbF3lflG
t0t1H7uQTjrl7stZ9FQCWli/cLaBx1k0c0T4efM6GSVdaj7SLHFlMYhvPTCFD76x
DEwJt0Dw8z/IlvTdjYnrIzH21wN7lJWXB6RkxhdAqz3Ibn4uZSkqmF24wErShwL+
M7P3UFXo3NSQ3p01HramFNEZrMWRmm0ab7d8BayyJzHr7uv/PEheem4kEuj3VIw=
-----END CERTIFICATE-----`

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(testRootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(testCertPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		DNSName: "vic-tom.corp.local",
		Roots:   roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}

	fmt.Println("Test passed!")
}

func TestMethod5(t *testing.T) {
	fmt.Println("Testing message")
	testRootPEM := `
-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`

	testCertPEM := `
-----BEGIN CERTIFICATE-----
MIIF7TCCA9WgAwIBAgIJAK8xPUiy5BTTMA0GCSqGSIb3DQEBCwUAMIGTMQswCQYD
VQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJUGFsbyBBbHRv
MRUwEwYDVQQKDAxWTXdhcmUsIEluYy4xHjAcBgNVBAsMFUNvbnRhaW5lcnMgb24g
dlNwaGVyZTEkMCIGA1UEAwwbU2VsZi1zaWduZWQgYnkgVk13YXJlLCBJbmMuMB4X
DTE4MDIwMTA1MzcwOFoXDTIxMDEzMTA1MzcwOFowgZkxCzAJBgNVBAYTAlVTMRMw
EQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlQYWxvIEFsdG8xDzANBgNVBAoM
BlZNd2FyZTEeMBwGA1UECwwVQ29udGFpbmVycyBvbiB2U3BoZXJlMTAwLgYDVQQD
DCdzYy1yZG9wcy12bTEwLWRoY3AtMzUtNTIuZW5nLnZtd2FyZS5jb20wggIiMA0G
CSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCyclPLkcP75rFSVoOG2+7bCuzKuf/5
IJNB0FPoj8+4stCLfKnxi0nK/ktWQWHRtS30nQCc/BYs33oMiz/R+MSusCe/FWtT
58AUFJCWDTHm/IpSDOZ/jqT6529JTtgWXDOo/euTk0ibD0yu68pfSLDT0S6nEekK
7Kgr1B3AwF/Kbtrh6xDeUGVW2R/FUcCzW+wz20YPsnuyy2xnlat6ipYNS4ahlHjq
7WBwyf/WMVhOd08asGB33CnW1m0xvgDjxskumeiGlreDHYTvVgdPwB+5amBYxr9V
2x/lBErjcUGHVVpQi/2kW3A/bCrd6Dnzhquf0FwJSMW19G+c0XTgJyvX2h4yTlUq
IacwpoNveRzFS5XkzqllZReehlUIBdegV3ctwG8TER0tt4vlSlUmV26eACvHsI/e
Wuk31M5W373mW6wrrxXzL4gHxP7RZiB6t5Ann3tFAp6tjAzXZf8gXmWG4TMvCbFm
m/B4PuakF98zzcLo4DFCIx1uJ6TfcXh1CsaFC6jPq17L3RBVQjrGvOumz6PLjGCZ
Hb/Zmg7WeR2PtknoLfITt8+wlt7+MpgQ32yB1EhR/M4ON8aPp4ejQof2EotiddIl
sNVDwz93sFAAPVJTVrZETKQ/yci+JvHGr2vA3s72yk5Dvb29Fu03s+Qu6uPsZwGJ
DX1iYxE8ZohtxwIDAQABozwwOjA4BgNVHREEMTAvgidzYy1yZG9wcy12bTEwLWRo
Y3AtMzUtNTIuZW5nLnZtd2FyZS5jb22HBAqhIzQwDQYJKoZIhvcNAQELBQADggIB
ACApZ10Yk786bU4gQMOuPuVNnihYeZSQMvf+ECWP8RI3jbTx5YqMWmXFwERmi1Ud
qdfOlmYK6WDDNxPheAYpQXwvnIjd9d+AEBGdF5gncd2/QR4r8nAnzjv4Ozm3bD99
d1vKzUE94k3D2rBBIMeO5y4fRGoExRykHNAFzeTyBkkfj1lOjGJiNyjr23nh8aJm
YJzygTPx64Lzlb/6zipKUp7ZaQvFKESy+3WqlESbDa1NfLVTIimKFyvpTAeh7/Rk
x0I84RJz8rPVpFusk3dsULDKochLF6m8X9j4lEkJhgjDH6qzoX2IQ9byczdCtuyl
5GPdA/WQ07NLQWr5NglwxSjlAGPdGR+emfKVMZH/PjcB4JrXcDXKohfa75J6+0DA
AZWgcTxzsv9ByynxmBEbSg4epg5PNAgBPfvMziMEchfvjNSzwvkDrBVRJpnETEYy
j1fWOTrUuPrCBPN5wGNf/oAFjYMKNi9bfZg5CrjG67QRKSqNa9d+ab+O5pkwiHaW
8CjyECiIhItkOsox8Wmb8vFYfhyvpXUKG5nfQAE2MEKC8wV1RqI4YLmbb49xa4U5
i84T/mn2mlqgGpClixPk0IPpStLdaeaBDZd6rc4jPG5g2AS+JsVi8b3IV4VfsUr4
DtgjOr08/9PUXC1Q4WzzqQk7tolm+uH8FKs6s1EKs2cc
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(testRootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(testCertPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		DNSName: "sc-rdops-vm10-dhcp-35-52.eng.vmware.com",
		Roots:   roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}

	opts = x509.VerifyOptions{
		DNSName: "10.161.35.52",
		Roots:   roots,
	}

	if _, err = cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}

	fmt.Println("Test passed!")
}

func TestBytes(t *testing.T) {
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб")))
	fmt.Printf("%s", bytes.Title([]byte("her royal highness")))
}

func TestMethodCompare(t *testing.T) {
	// Interpret Compare's result by comparing it to zero.
	var a, b []byte
	a = []byte("sample")
	b = []byte("foo")
	if bytes.Compare(a, b) < 0 {
		fmt.Println("a less b")
		// a less b
	}
	if bytes.Compare(a, b) <= 0 {
		// a less or equal b
		fmt.Println("a less or equals b")
	}
	if bytes.Compare(a, b) > 0 {
		// a greater b
		fmt.Println("a greater b")
	}
	if bytes.Compare(a, b) >= 0 {
		// a greater or equal b
		fmt.Println("a greater or equal b")
	}

	// Prefer Equal to Compare for equality comparisons.
	if bytes.Equal(a, b) {
		// a equal b
		fmt.Println("a equal b")
	}
	if !bytes.Equal(a, b) {
		// a not equal b
		fmt.Println("a not equal b")
	}
}

func TestMethodSearch(t *testing.T) {
	// Binary search to find a matching byte slice.
	var needle = []byte("dig")
	var haystack = [][]byte{
		[]byte("apple"),
		[]byte("bear"),
		[]byte("code"),
		[]byte("dear"),
		[]byte("egg"),
		[]byte("frank"),
		[]byte("sample"),
	} // Assume sorted
	i := sort.Search(len(haystack), func(i int) bool {
		// Return haystack[i] >= needle.
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		// Found it!
		fmt.Println("Found it")
	} else {
		fmt.Printf("Not found, suggest index:%v", i)
	}
}

func TestMethodContains(t *testing.T) {
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))
}

func TestMethodContainsAny(t *testing.T) {
	fmt.Println(bytes.ContainsAny([]byte("张道军"), "人民道路"))
	fmt.Println(bytes.ContainsAny([]byte("张道军"), "人民路"))
}

func TestMethodCount(t *testing.T) {
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("人民道路"), []byte(""))) // before & after each
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))
}
func TestMethodFiels(t *testing.T) {
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
}

func TestMethodFieldFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))
}

func TestMethodHasPrefix(t *testing.T) {
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))
}

func TestMethodHasSuffix(t *testing.T) {
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("go")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("O")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("")))
}

func TestMethodIndex(t *testing.T) {
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))
	fmt.Println(bytes.Index([]byte("chicken"), []byte("dmr")))
}

func TestMethodIndexAny(t *testing.T) {
	fmt.Println(bytes.IndexAny([]byte("chicken"), "aeiouy"))
	fmt.Println(bytes.IndexAny([]byte("crwth"), "aeiouy"))
}

func TestMethodIndexFunc(t *testing.T) {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f))
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f))
}

func TestMethodJoin(t *testing.T) {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(", ")))
}

func TestMethodMap(t *testing.T) {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))
}

func TestMethodRepeat(t *testing.T) {
	fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2))
}

func TestMethodReplace(t *testing.T) {
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
}

func TestMethodSplit(t *testing.T) {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
}

func TestMethodSplitAfterN(t *testing.T) {
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2))
}

func TestMethodToUpper(t *testing.T) {
	fmt.Printf("%s", bytes.ToUpper([]byte("Gopher")))
}

func TestMethodTrim(t *testing.T) {
	fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))
}

func TestMethodReadFile(t *testing.T) {
	file, err := os.Open("/Users/daojunz/temp/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestMethodWriteToFile(t *testing.T) {
	file, err := os.Create("/Users/daojunz/temp/sample2.txt")
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush()
}

func TestMethodOpenFile(t *testing.T) {
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
