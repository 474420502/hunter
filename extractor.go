package hunter

import (
	"log"

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

// XPaths multi xpath extractor
func (etor *Extractor) XPaths(exp string) (*XPath, error) {
	result, err := etor.doc.Find(exp)
	return &XPath{result: []types.XPathResult{result}, errorFlags: ERROR_SKIP}, err
}

// XPathResult libxml2 xpathresult
func (etor *Extractor) XPathResult(exp string) (result types.XPathResult, err error) {
	result, err = etor.doc.Find(exp)
	return
}

type ErrorFlags int

const (
	ERROR_SKIP  ErrorFlags = 1
	ERROR_BREAK ErrorFlags = 2
)

// XPath for easy extractor data
type XPath struct {
	result     []types.XPathResult
	errorFlags ErrorFlags
}

// GetXPathResults Get Current XPath Results
func (xp *XPath) GetXPathResults() []types.XPathResult {
	return xp.result
}

// ForEachString Each String
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

// ForEachText all result get TextContent
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

// ForEachType all result get XMLNodeType
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

// ForEachValue all results get NodeValue
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

// ForEachAttr all results get NodeAttribute
func (xp *XPath) ForEachAttr(exp string) (attributes []types.Attribute, errorlist []error) {

	inames, errlist := xp.ForEachEx(exp, func(result types.XPathResult) []interface{} {
		var ir []interface{}
		for iter := result.NodeIter(); iter.Next(); {
			ele := iter.Node().(types.Element)

			attribute, err := ele.Attributes()
			if err != nil {
				log.Println(err)
			}
			for _, attr := range attribute {
				ir = append(ir, attr)
			}

		}

		return ir
	})

	for _, i := range inames {
		attributes = append(attributes, i.(types.Attribute))
	}

	return attributes, errlist
}

// ForEachAttrKeys all results get NodeAttribute Key
func (xp *XPath) ForEachAttrKeys(exp string) (keyslist [][]string, errorlist []error) {

	inames, errlist := xp.ForEachEx(exp, func(result types.XPathResult) []interface{} {
		var ir []interface{}
		for iter := result.NodeIter(); iter.Next(); {
			ele := iter.Node().(types.Element)
			attributes, err := ele.Attributes()

			var keys []string
			for _, attr := range attributes {
				if err != nil {
					log.Println(err)
				}
				keys = append(keys, attr.NodeName())
			}
			ir = append(ir, keys)
		}
		return ir
	})

	for _, i := range inames {
		keyslist = append(keyslist, i.([]string))
	}

	return keyslist, errlist
}

// ForEachAttrValue all results get NodeAttribute
func (xp *XPath) ForEachAttrValue(exp string, attributes ...string) (values []string, errorlist []error) {

	inames, errlist := xp.ForEachEx(exp, func(result types.XPathResult) []interface{} {
		var ir []interface{}
		for iter := result.NodeIter(); iter.Next(); {
			ele := iter.Node().(types.Element)
			for _, attr := range attributes {
				attribute, err := ele.GetAttribute(attr)
				if err != nil {
					log.Println(err)
				}
				ir = append(ir, attribute.Value())
			}
		}
		return ir
	})

	for _, i := range inames {
		values = append(values, i.(string))
	}

	return values, errlist
}

// ForEachName all result get NodeName
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

// ForEachEx foreach with do funciton
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

// ForEach new XPath( every result xpath get results )
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
