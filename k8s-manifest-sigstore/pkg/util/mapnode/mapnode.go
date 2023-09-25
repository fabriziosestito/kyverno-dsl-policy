package mapnode

import _ "unsafe"

type Embedme interface{}

type DiffResult struct {Items interface{}}

type Difference struct {Key interface{}; Values interface{}}

type DiffPattern Difference

func (d *Difference) Equal(d2 interface{}) interface{} {
 panic("stub")
}

func (dp *DiffPattern) Match(diff interface{}) interface{} {
 panic("stub")
}

func (d *DiffResult) Keys() interface{} {
 panic("stub")
}

func (d *DiffResult) Values() interface{} {
 panic("stub")
}

func (dr *DiffResult) Size() interface{} {
 panic("stub")
}

func (dr *DiffResult) Remove(patterns interface{}) interface{} {
 panic("stub")
}

func (dr *DiffResult) Filter(maskKeys interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func (d *DiffResult) ToJson() interface{} {
 panic("stub")
}

func (d *DiffResult) String() interface{} {
 panic("stub")
}

func (d *DiffResult) KeyString() interface{} {
 panic("stub")
}

func keyExistsInList(slice interface{}, val interface{}) (interface{}, interface{}) {
 panic("stub")
}

func patternMatchString(data, pattern interface{}) interface{} {
 panic("stub")
}

func patternMatch(data, pattern interface{}) interface{} {
 panic("stub")
}

type Node struct {Value interface{}; Children interface{}}

type matchMode string

type NodeValue struct {Type interface{}; Value interface{}}

func NewNodeValue(val interface{}) interface{} {
 panic("stub")
}

func (nv *NodeValue) Interface() interface{} {
 panic("stub")
}

func (nv *NodeValue) String() interface{} {
 panic("stub")
}

func emptyNode() interface{} {
 panic("stub")
}

func NewFromBytes(rawObj interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewFromYamlBytes(rawObj interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewFromInterfaceBytes(rawObj interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewFromMap(objMap interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewNode(val interface{}) interface{} {
 panic("stub")
}

func (n *Node) KeyExists(concatKey interface{}) interface{} {
 panic("stub")
}

func (n *Node) DeepCopyInto(n2 interface{}) {
 panic("stub")
}

func (n *Node) Copy() interface{} {
 panic("stub")
}

func (n *Node) Merge(n2 interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (n *Node) Extract(filterKeys interface{}) interface{} {
 panic("stub")
}

func (t *Node) Mask(keys interface{}) interface{} {
 panic("stub")
}

func (n *Node) recursiveMask(currentPath interface{}, maskKeys interface{}) interface{} {
 panic("stub")
}

func (n *Node) IsValue() interface{} {
 panic("stub")
}

func (n *Node) IsMap() interface{} {
 panic("stub")
}

func (n *Node) IsSlice() interface{} {
 panic("stub")
}

func (n *Node) Size() interface{} {
 panic("stub")
}

func (n *Node) GetChildrenMap() interface{} {
 panic("stub")
}

func (n *Node) GetChildrenSlice() interface{} {
 panic("stub")
}

func (n *Node) GetNodeByJSONPath(jpathKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (t *Node) validateKeyList(keys interface{}) interface{} {
 panic("stub")
}

func GetConcreteKeys(keys interface{}, n interface{}) interface{} {
 panic("stub")
}

func (t *Node) generateKeyList(concatKey interface{}) interface{} {
 panic("stub")
}

func (t *Node) MultipleSubNode(concatKey interface{}) interface{} {
 panic("stub")
}

func (t *Node) GetNode(concatKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (t *Node) SubNode(concatKey interface{}) interface{} {
 panic("stub")
}

func (t *Node) GetString(concatKey interface{}) interface{} {
 panic("stub")
}

func (t *Node) GetBool(concatKey interface{}, defaultValue interface{}) interface{} {
 panic("stub")
}

func (t *Node) Get(concatKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (n *Node) GetChild(key interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (t *Node) Ravel() interface{} {
 panic("stub")
}

func (n *Node) recursiveRavel(currentPath interface{}, m interface{}) interface{} {
 panic("stub")
}

func (n *Node) Interface() interface{} {
 panic("stub")
}

func (t *Node) ToMap() interface{} {
 panic("stub")
}

func (n *Node) ToJson() interface{} {
 panic("stub")
}

func (n *Node) ToYaml() interface{} {
 panic("stub")
}

func (n *Node) String() interface{} {
 panic("stub")
}

func (t *Node) Diff(t2 interface{}) interface{} {
 panic("stub")
}

func (t *Node) DiffStrict(t2 interface{}) interface{} {
 panic("stub")
}

func (t *Node) DiffSpecificType(t2 interface{}, findTypeList interface{}) interface{} {
 panic("stub")
}

func (t *Node) FindUpdatedAndDeleted(t2 interface{}) interface{} {
 panic("stub")
}

func (t *Node) FindUpdatedAndCreated(t2 interface{}) interface{} {
 panic("stub")
}

func extractComparableMap(m1, m2 interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func FindDiffBetweenNodes(t1, t2 interface{}, findType interface{}, matchMode interface{}) interface{} {
 panic("stub")
}

func recursiveGetByKey(m interface{}, i interface{}, keyList interface{}) (interface{}, interface{}) {
 panic("stub")
}

func parseConcatKey(concatKey interface{}) interface{} {
 panic("stub")
}

func splitConcatKey(concatKey interface{}) interface{} {
 panic("stub")
}

func removeKeyDiffsInListNode(items interface{}) interface{} {
 panic("stub")
}

func removeDiffWithDefaultConversion(items interface{}) interface{} {
 panic("stub")
}

func GetValueByLongKey(m interface{}, longKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func SplitCommaSeparatedKeys(key interface{}) interface{} {
 panic("stub")
}

