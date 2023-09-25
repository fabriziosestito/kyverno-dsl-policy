package kubeutil

import _ "unsafe"

type Embedme interface{}

func DryRunCreate(objBytes interface{}, namespace interface{}) (interface{}, interface{}) {
 panic("stub")
}

func StrategicMergePatch(objBytes, patchBytes interface{}, namespace interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetApplyPatchBytes(manifestBytes interface{}, objNamespace interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

type podGetterFunc func(obj interface{}) (interface{}, interface{})

type ImageObject struct {PodName interface{}; ContainerName interface{}; ImageID interface{}; ResourceBundleRef interface{}; Digest interface{}}

func GetInClusterConfig() (interface{}, interface{}) {
 panic("stub")
}

func IsInCluster() interface{} {
 panic("stub")
}

func GetOutOfClusterConfig() (interface{}, interface{}) {
 panic("stub")
}

func GetKubeConfig() (interface{}, interface{}) {
 panic("stub")
}

func SetKubeConfig(conf interface{}) {
 panic("stub")
}

func MatchLabels(obj interface{}, labelSelector interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetAPIResources() (interface{}, interface{}) {
 panic("stub")
}

func GetNamespaces() (interface{}, interface{}) {
 panic("stub")
}

func ParseObjectRefInCluster(objRef interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func ParseObjectRefInClusterWithKind(objRef interface{}) (interface{}, interface{}, interface{}, interface{}) {
 panic("stub")
}

func GetSecret(namespace, name interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetResource(apiVersion, kind, namespace, name interface{}) (interface{}, interface{}) {
 panic("stub")
}

func ListResources(apiVersion, kind, namespace interface{}) (interface{}, interface{}) {
 panic("stub")
}

func contains(all interface{}, one interface{}) interface{} {
 panic("stub")
}

func GetAllImagesFromObject(obj interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetAllPodsFromObject(obj interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getPodsNotImplemented(obj interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getPodsFromDeployment(obj interface{}) (interface{}, interface{}) {
 panic("stub")
}

