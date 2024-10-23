Started by user unknown or anonymous
Obtained Jenkinsfile from git https://github.com/palashbhasme/catfact-app.git
[Pipeline] Start of Pipeline
[Pipeline] node
Running on Jenkins in /var/jenkins_home/workspace/catfact-job2
[Pipeline] {
[Pipeline] stage
[Pipeline] { (Declarative: Checkout SCM)
[Pipeline] checkout
The recommended git tool is: git
No credentials specified
 > git rev-parse --resolve-git-dir /var/jenkins_home/workspace/catfact-job2/.git # timeout=10
Fetching changes from the remote Git repository
 > git config remote.origin.url https://github.com/palashbhasme/catfact-app.git # timeout=10
Fetching upstream changes from https://github.com/palashbhasme/catfact-app.git
 > git --version # timeout=10
 > git --version # 'git version 2.39.5'
 > git fetch --tags --force --progress -- https://github.com/palashbhasme/catfact-app.git +refs/heads/*:refs/remotes/origin/* # timeout=10
 > git rev-parse refs/remotes/origin/main^{commit} # timeout=10
Checking out Revision 9fc0467d8c629842883c2e57cdcaa5d45b043cf1 (refs/remotes/origin/main)
 > git config core.sparsecheckout # timeout=10
 > git checkout -f 9fc0467d8c629842883c2e57cdcaa5d45b043cf1 # timeout=10
Commit message: "Update Jenkinsfile"
 > git rev-list --no-walk f3497c3057ae38bcb94ef211363d699aa65aa504 # timeout=10
[Pipeline] }
[Pipeline] // stage
[Pipeline] withEnv
[Pipeline] {
[Pipeline] stage
[Pipeline] { (Build)
[Pipeline] getContext
[Pipeline] isUnix
[Pipeline] withEnv
[Pipeline] {
[Pipeline] sh
+ docker inspect -f . golang:1.23
.
[Pipeline] }
[Pipeline] // withEnv
[Pipeline] withDockerContainer
Jenkins seems to be running inside container 12ff0d3b0d6e8a6c95d9bf3dddca4b0d1305274240f851e9ce4d4c3a56e92c47
but /var/jenkins_home/workspace/catfact-job2 could not be found among []
but /var/jenkins_home/workspace/catfact-job2@tmp could not be found among []
$ docker run -t -d -u 1000:1000 -w /var/jenkins_home/workspace/catfact-job2 -v /var/jenkins_home/workspace/catfact-job2:/var/jenkins_home/workspace/catfact-job2:rw,z -v /var/jenkins_home/workspace/catfact-job2@tmp:/var/jenkins_home/workspace/catfact-job2@tmp:rw,z -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** golang:1.23 cat
$ docker top 5f29ef0f915381a521a5f2bd0f7e6432ad94fe78bffb181101ed55bc390854c7 -eo pid,comm
[Pipeline] {
[Pipeline] sh
+ export GOCACHE=/tmp/.cache
+ mkdir -p /tmp/.cache
+ chmod -R 777 /tmp/.cache
+ go version
go version go1.23.2 linux/amd64
+ go mod tidy
go: downloading go.mongodb.org/mongo-driver v1.17.1
go: downloading github.com/golang/snappy v0.0.4
go: downloading github.com/montanaflynn/stats v0.7.1
go: downloading github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78
go: downloading github.com/klauspost/compress v1.13.6
go: downloading github.com/xdg-go/scram v1.1.2
go: downloading github.com/xdg-go/stringprep v1.0.4
go: downloading golang.org/x/sync v0.8.0
go: downloading golang.org/x/crypto v0.26.0
go: downloading github.com/xdg-go/pbkdf2 v1.0.0
go: downloading golang.org/x/text v0.17.0
go: downloading github.com/google/go-cmp v0.6.0
go: downloading github.com/davecgh/go-spew v1.1.1
+ go build -v
internal/coverage/rtcov
internal/goarch
internal/godebugs
internal/unsafeheader
internal/goos
internal/goexperiment
internal/byteorder
internal/cpu
internal/profilerecord
internal/runtime/atomic
internal/runtime/syscall
internal/abi
runtime/internal/math
runtime/internal/sys
internal/race
internal/chacha8rand
sync/atomic
internal/asan
internal/itoa
internal/msan
unicode
internal/runtime/exithook
internal/bytealg
unicode/utf8
cmp
encoding
math/bits
unicode/utf16
go.mongodb.org/mongo-driver/bson/bsonoptions
go.mongodb.org/mongo-driver/bson/bsontype
crypto/internal/alias
crypto/subtle
crypto/internal/boring/sig
math
internal/nettrace
go.mongodb.org/mongo-driver/internal/handshake
go.mongodb.org/mongo-driver/internal/ptrutil
container/list
vendor/golang.org/x/crypto/cryptobyte/asn1
internal/stringslite
vendor/golang.org/x/crypto/internal/alias
log/internal
go.mongodb.org/mongo-driver/version
runtime
github.com/klauspost/compress
internal/reflectlite
internal/weak
iter
sync
maps
slices
internal/bisect
internal/singleflight
internal/testlog
runtime/cgo
errors
sort
internal/godebug
internal/oserror
io
strconv
path
vendor/golang.org/x/net/dns/dnsmessage
syscall
bytes
math/rand
strings
hash
reflect
crypto
crypto/internal/randutil
regexp/syntax
crypto/cipher
math/rand/v2
bufio
hash/crc32
internal/concurrent
crypto/internal/boring
crypto/des
unique
crypto/internal/edwards25519/field
crypto/aes
crypto/internal/nistec/fiat
net/netip
time
internal/syscall/unix
internal/syscall/execenv
regexp
crypto/sha512
crypto/internal/edwards25519
crypto/hmac
crypto/md5
vendor/golang.org/x/crypto/hkdf
crypto/rc4
crypto/sha1
crypto/sha256
vendor/golang.org/x/text/transform
net/http/internal/ascii
context
io/fs
internal/poll
golang.org/x/crypto/pbkdf2
hash/adler32
golang.org/x/text/transform
go.mongodb.org/mongo-driver/internal/aws
go.mongodb.org/mongo-driver/internal/csot
internal/filepathlite
embed
internal/fmtsort
encoding/binary
crypto/internal/nistec
os
encoding/base64
vendor/golang.org/x/crypto/chacha20
vendor/golang.org/x/crypto/internal/poly1305
go.mongodb.org/mongo-driver/internal/rand
golang.org/x/crypto/scrypt
github.com/golang/snappy
github.com/klauspost/compress/internal/snapref
github.com/klauspost/compress/zstd/internal/xxhash
encoding/pem
crypto/ecdh
github.com/xdg-go/pbkdf2
io/ioutil
go.mongodb.org/mongo-driver/internal/driverutil
path/filepath
fmt
vendor/golang.org/x/sys/cpu
net
vendor/golang.org/x/crypto/sha3
vendor/golang.org/x/crypto/chacha20poly1305
os/exec
encoding/hex
go.mongodb.org/mongo-driver/tag
log
net/url
compress/flate
encoding/json
math/big
go.mongodb.org/mongo-driver/mongo/readpref
vendor/golang.org/x/text/unicode/norm
vendor/golang.org/x/net/http2/hpack
vendor/golang.org/x/text/unicode/bidi
mime
mime/quotedprintable
compress/gzip
net/http/internal
vendor/golang.org/x/text/secure/bidirule
compress/zlib
github.com/klauspost/compress/fse
runtime/debug
golang.org/x/text/unicode/norm
go.mongodb.org/mongo-driver/internal/aws/awserr
github.com/montanaflynn/stats
golang.org/x/sync/errgroup
golang.org/x/sync/singleflight
github.com/klauspost/compress/huff0
vendor/golang.org/x/net/idna
go.mongodb.org/mongo-driver/internal/aws/credentials
crypto/elliptic
crypto/internal/boring/bbig
crypto/internal/bigmod
crypto/rand
encoding/asn1
go.mongodb.org/mongo-driver/bson/primitive
crypto/ed25519
crypto/internal/hpke
crypto/internal/mlkem768
crypto/dsa
crypto/rsa
go.mongodb.org/mongo-driver/internal/randutil
go.mongodb.org/mongo-driver/x/bsonx/bsoncore
go.mongodb.org/mongo-driver/internal/logger
go.mongodb.org/mongo-driver/internal/uuid
github.com/klauspost/compress/zstd
github.com/xdg-go/stringprep
crypto/x509/pkix
vendor/golang.org/x/crypto/cryptobyte
crypto/ecdsa
go.mongodb.org/mongo-driver/x/mongo/driver/wiremessage
go.mongodb.org/mongo-driver/mongo/readconcern
go.mongodb.org/mongo-driver/internal/csfle
go.mongodb.org/mongo-driver/bson/bsonrw
github.com/xdg-go/scram
go.mongodb.org/mongo-driver/bson/bsoncodec
go.mongodb.org/mongo-driver/mongo/address
vendor/golang.org/x/net/http/httpproxy
go.mongodb.org/mongo-driver/x/mongo/driver/dns
net/textproto
crypto/x509
go.mongodb.org/mongo-driver/bson
vendor/golang.org/x/net/http/httpguts
mime/multipart
go.mongodb.org/mongo-driver/internal/bsonutil
go.mongodb.org/mongo-driver/internal/codecutil
go.mongodb.org/mongo-driver/mongo/writeconcern
go.mongodb.org/mongo-driver/mongo/description
golang.org/x/crypto/ocsp
github.com/youmark/pkcs8
crypto/tls
go.mongodb.org/mongo-driver/event
go.mongodb.org/mongo-driver/x/mongo/driver/session
net/http/httptrace
net/http
go.mongodb.org/mongo-driver/internal/httputil
go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options
go.mongodb.org/mongo-driver/internal/aws/signer/v4
go.mongodb.org/mongo-driver/internal/credproviders
go.mongodb.org/mongo-driver/x/mongo/driver/ocsp
go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt
go.mongodb.org/mongo-driver/x/mongo/driver
go.mongodb.org/mongo-driver/x/mongo/driver/auth/creds
go.mongodb.org/mongo-driver/x/mongo/driver/operation
go.mongodb.org/mongo-driver/x/mongo/driver/auth
go.mongodb.org/mongo-driver/x/mongo/driver/connstring
go.mongodb.org/mongo-driver/mongo/options
go.mongodb.org/mongo-driver/x/mongo/driver/topology
go.mongodb.org/mongo-driver/mongo
catfact
[Pipeline] }
$ docker stop --time=1 5f29ef0f915381a521a5f2bd0f7e6432ad94fe78bffb181101ed55bc390854c7
$ docker rm -f --volumes 5f29ef0f915381a521a5f2bd0f7e6432ad94fe78bffb181101ed55bc390854c7
[Pipeline] // withDockerContainer
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Test)
[Pipeline] getContext
[Pipeline] isUnix
[Pipeline] withEnv
[Pipeline] {
[Pipeline] sh
+ docker inspect -f . golang:1.23
.
[Pipeline] }
[Pipeline] // withEnv
[Pipeline] withDockerContainer
Jenkins seems to be running inside container 12ff0d3b0d6e8a6c95d9bf3dddca4b0d1305274240f851e9ce4d4c3a56e92c47
but /var/jenkins_home/workspace/catfact-job2 could not be found among []
but /var/jenkins_home/workspace/catfact-job2@tmp could not be found among []
$ docker run -t -d -u 1000:1000 -w /var/jenkins_home/workspace/catfact-job2 -v /var/jenkins_home/workspace/catfact-job2:/var/jenkins_home/workspace/catfact-job2:rw,z -v /var/jenkins_home/workspace/catfact-job2@tmp:/var/jenkins_home/workspace/catfact-job2@tmp:rw,z -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** -e ******** golang:1.23 cat
$ docker top a8dcdd23ab7ef09f843c1cb3f22c1b4e17d8e3e7658bb78da3a87b84e5d89f5b -eo pid,comm
[Pipeline] {
[Pipeline] sh
+ export GOCACHE=/tmp/.cache
+ mkdir -p /tmp/.cache
+ chmod -R 777 /tmp/.cache
+ echo Running Tests...
Running Tests...
+ go test ./...
go: downloading go.mongodb.org/mongo-driver v1.17.1
go: downloading github.com/golang/snappy v0.0.4
go: downloading github.com/klauspost/compress v1.13.6
go: downloading github.com/xdg-go/scram v1.1.2
go: downloading github.com/xdg-go/stringprep v1.0.4
go: downloading github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78
go: downloading github.com/montanaflynn/stats v0.7.1
go: downloading golang.org/x/sync v0.8.0
go: downloading golang.org/x/crypto v0.26.0
go: downloading golang.org/x/text v0.17.0
go: downloading github.com/xdg-go/pbkdf2 v1.0.0
?   	catfact	[no test files]
[Pipeline] }
$ docker stop --time=1 a8dcdd23ab7ef09f843c1cb3f22c1b4e17d8e3e7658bb78da3a87b84e5d89f5b
$ docker rm -f --volumes a8dcdd23ab7ef09f843c1cb3f22c1b4e17d8e3e7658bb78da3a87b84e5d89f5b
[Pipeline] // withDockerContainer
[Pipeline] }
[Pipeline] // stage
[Pipeline] }
[Pipeline] // withEnv
[Pipeline] }
[Pipeline] // node
[Pipeline] End of Pipeline
Finished: SUCCESS
