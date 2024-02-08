// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"zh": &dictionary{index: zhIndex, data: zhData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"# Create a Kubernetes deployment. Provide common variables.":                              1,
	"# Create a resource group, virtual network, subnet and virtual machine on alibaba cloud.": 3,
	"# Deploy an ELK stack using helm chart.":                                                  5,
	"Create a Kubernetes deployment":                                                           0,
	"Create an alibaba cloud virtual machine":                                                  2,
	"Deploy an ELK stack":                                                                      4,
	"Explanation of the fixes.":                                                                10,
	"Please Check and fix the given terraform module if there's any mistake.\nStrictly respond a valid JSON in the following format:": 8,
	"Please explain the given terraform module.":                                          7,
	"Terraform code that is fixed. Please do not explain, just write terraform HCL code.": 9,
	"You are translating natural language to a Terraform module. Please do not explain, just write pure terraform HCL code. Please do not explain, just write pure terraform HCL code. Please do not explain, just write pure terraform HCL code.": 6,
}

var enIndex = []uint32{ // 12 elements.
	0x00000000, 0x0000001f, 0x0000005b, 0x00000083,
	0x000000dc, 0x000000f0, 0x00000118, 0x00000205,
	0x00000230, 0x000002af, 0x00000303, 0x0000031d,
} // Size: 72 bytes.

const enData string = "" + // Size: 797 bytes.
	"\x02Create a Kubernetes deployment\x02# Create a Kubernetes deployment. " +
	"Provide common variables.\x02Create an alibaba cloud virtual machine\x02" +
	"# Create a resource group, virtual network, subnet and virtual machine o" +
	"n alibaba cloud.\x02Deploy an ELK stack\x02# Deploy an ELK stack using h" +
	"elm chart.\x02You are translating natural language to a Terraform module" +
	". Please do not explain, just write pure terraform HCL code. Please do n" +
	"ot explain, just write pure terraform HCL code. Please do not explain, j" +
	"ust write pure terraform HCL code.\x02Please explain the given terraform" +
	" module.\x02Please Check and fix the given terraform module if there's a" +
	"ny mistake.\x0aStrictly respond a valid JSON in the following format:" +
	"\x02Terraform code that is fixed. Please do not explain, just write terr" +
	"aform HCL code.\x02Explanation of the fixes."

var zhIndex = []uint32{ // 12 elements.
	0x00000000, 0x00000022, 0x00000061, 0x00000080,
	0x000000ce, 0x000000de, 0x00000103, 0x000001f0,
	0x00000216, 0x0000027a, 0x000002ce, 0x000002f6,
} // Size: 72 bytes.

const zhData string = "" + // Size: 758 bytes.
	"\x02创建一个Kubernetes deployment\x02# 创建一个Kubernetes deployment，提供常用的变量。" +
	"\x02创建一个阿里云虚拟机\x02# 在阿里云上创建一个资源组，虚拟网络，子网和虚拟机。\x02部署ELK组件\x02# 使用Helm cha" +
	"rt部署ELK组件。\x02You are translating natural language to a Terraform module" +
	". Please do not explain, just write pure terraform HCL code. Please do n" +
	"ot explain, just write pure terraform HCL code. Please do not explain, j" +
	"ust write pure terraform HCL code.\x02请解释提供的Terraform module。\x02请检查和修复提" +
	"供的Terraform module。\x0a严格遵循以下格式回复一个合法的JSON：\x02Terraform code that is " +
	"fixed. Please do not explain, just write terraform HCL code.\x02对代码修复的中文" +
	"解释说明。"

	// Total table size 1699 bytes (1KiB); checksum: 8F5B9935.
