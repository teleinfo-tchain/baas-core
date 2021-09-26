module github.com/bif/telbaas/baas-core

go 1.15

replace xorm.io/core v0.7.3 => github.com/go-xorm/core v0.6.3

// sdk依赖了fabric-protos-go，不能替换
//replace github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23 => ./fabric-protos-go
require (
	cloud.google.com/go v0.81.0 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/Shopify/sarama v1.29.1 // indirect
	github.com/alexandrevicenzi/unchained v1.3.0
	github.com/containerd/cgroups v1.0.1 // indirect
	github.com/containerd/containerd v1.4.9 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/docker v20.10.8+incompatible // indirect
	github.com/fsouza/go-dockerclient v1.7.3 // indirect
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.4
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hyperledger/fabric v1.4.11
	github.com/hyperledger/fabric-amcl v0.0.0-20210603140002-2670f91851c8 // indirect
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/json-iterator/go v1.1.12
	github.com/kr/pretty v0.3.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/onsi/gomega v1.16.0 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/opencontainers/runc v1.0.1 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/viper v1.2.1 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602 // indirect
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.40.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.20.0-beta.1
	k8s.io/apimachinery v0.20.0-beta.1
	k8s.io/client-go v0.20.0-beta.1
	k8s.io/klog/v2 v2.9.0 // indirect
	k8s.io/utils v0.0.0-20210707171843-4b05e18ac7d9 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.2 // indirect
	xorm.io/core v0.7.3
)
