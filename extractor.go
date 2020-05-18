package hunter

import (
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/clib"
	"github.com/lestrrat-go/libxml2/types"
)

// Extractor 提取器
type Extractor struct {
	Content []byte
	doc     types.Document
}

// NewExtractor 创建提取器
func NewExtractor(content []byte) *Extractor {
	doc, err := libxml2.ParseHTML(content)
	if err != nil {
		panic(err)
	}
	return &Extractor{Content: content, doc: doc}
}

// XPath 路径提取
func (etor *Extractor) XPath(exp string) (*XPath, error) {
	result, err := etor.doc.Find(exp)
	return &XPath{result: []types.XPathResult{result}, errorFlags: ERROR_SKIP}, err
}

type ErrorFlags int

const (
	ERROR_SKIP  ErrorFlags = 1
	ERROR_BREAK ErrorFlags = 2
)

type XPath struct {
	result     []types.XPathResult
	errorFlags ErrorFlags
}

func (xp *XPath) ForEachString(exp string) (sstr []string, errorlist []error) {

	inames, errlist := xp.ForEachEx(exp, func(result types.XPathResult) []interface{} {
		var ir []interface{}
		for iter := result.NodeIter(); iter.Next(); {

			ir = append(ir, iter.Node().String())
		}
		return ir
	})

	for _, i := range inames {
		sstr = append(sstr, i.(string))
	}

	return sstr, errlist
}

func (xp *XPath) ForEachText(exp string) (texts []string, errorlist []error) {

	inames, errlist := xp.ForEachEx(exp, func(result types.XPathResult) []interface{} {
		var ir []interface{}
		for iter := result.NodeIter(); iter.Next(); {

			ir = append(ir, iter.Node().TextContent())
		}
		return ir
	})

	for _, i := range inames {
		texts = append(texts, i.(string))
	}

	return texts, errlist
}

func (xp *XPath) ForEachType(exp string) (typelist []clib.XMLNodeType, errorlist []error) {

	inames, errlist := xp.ForEachEx(exp, func(result types.XPathResult) []interface{} {
		var ir []interface{}
		for iter := result.NodeIter(); iter.Next(); {

			ir = append(ir, iter.Node().NodeType())
		}
		return ir
	})

	for _, i := range inames {
		typelist = append(typelist, i.(clib.XMLNodeType))
	}

	return typelist, errlist
}

func (xp *XPath) ForEachValue(exp string) (values []string, errorlist []error) {

	inames, errlist := xp.ForEachEx(exp, func(result types.XPathResult) []interface{} {
		var ir []interface{}
		for iter := result.NodeIter(); iter.Next(); {

			ir = append(ir, iter.Node().NodeValue())
		}
		return ir
	})

	for _, i := range inames {
		values = append(values, i.(string))
	}

	return values, errlist
}

func (xp *XPath) ForEachName(exp string) (names []string, errorlist []error) {

	inames, errlist := xp.ForEachEx(exp, func(result types.XPathResult) []interface{} {
		var ir []interface{}
		for iter := result.NodeIter(); iter.Next(); {

			ir = append(ir, iter.Node().NodeName())
		}
		return ir
	})

	for _, i := range inames {
		names = append(names, i.(string))
	}

	return names, errlist
}

func (xp *XPath) ForEachEx(exp string, do func(types.XPathResult) []interface{}) (values []interface{}, errorlist []error) {
	if len(xp.result) == 0 {
		return
	}

	for _, xpresult := range xp.result {

		iter := xpresult.NodeIter()
		for iter.Next() {
			node := iter.Node()
			result, err := node.Find(exp)
			iresult := do(result)
			if err != nil {
				if xp.errorFlags == ERROR_SKIP {
					errorlist = append(errorlist, err)
				} else {
					break
				}
			}
			values = append(values, iresult...)
		}
	}

	return
}

func (xp *XPath) ForEach(exp string) (newxpath *XPath, errorlist []error) {
	if len(xp.result) == 0 {
		return
	}

	newxpath = &XPath{errorFlags: xp.errorFlags}

	for _, xpresult := range xp.result {

		iter := xpresult.NodeIter()
		for iter.Next() {
			node := iter.Node()
			result, err := node.Find(exp)
			if err != nil {
				if xp.errorFlags == ERROR_SKIP {
					errorlist = append(errorlist, err)
				} else {
					break
				}
			}
			newxpath.result = append(newxpath.result, result)
		}
	}

	return
}
