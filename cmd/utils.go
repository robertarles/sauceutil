package cmd

import (
	"fmt"
	"reflect"
	"strings"
)

// TODO: Update this to handle printing nested structure fields, specifically consider the uploads response which is hides in a 'file' object

// OPrintStruct prints formatted header or formatted header+values of an arbitrary structure based on JSON field names
func OPrintStruct(reqFieldNames []string, strct interface{}, printHeader bool) error {

	sType := reflect.TypeOf(strct)
	sVals := reflect.Indirect(reflect.ValueOf(strct))
	oVals := []string{}
	eHeaders := []string{}

	// clean up fullTag by removing 'json:' and the quotes
	replacements := strings.NewReplacer("json:", "", "\"", "")

	// get a list of json name equiv for each of the struct field names
	allTypeFieldTags := []string{}
	allTypeFieldNames := []string{}
	fieldTagsToNames := make(map[string]string)
	for i := 0; i < sType.NumField(); i++ {
		jsonName := replacements.Replace(fmt.Sprintf("%v", sType.Field(i).Tag))
		allTypeFieldTags = append(allTypeFieldTags, jsonName)
		fieldTagsToNames[jsonName] = sType.Field(i).Name
		allTypeFieldNames = append(allTypeFieldNames, sType.Field(i).Name)
	}

	// get a list of the types fieldnames that match those requested for output
	outTypeFieldTags := []string{}  // the json field name
	outTypeFieldNames := []string{} // the json field name
	for i := 0; i < len(reqFieldNames); i++ {
		if ArrayContains(allTypeFieldTags, reqFieldNames[i]) {
			outTypeFieldTags = append(outTypeFieldTags, reqFieldNames[i])
			outTypeFieldNames = append(outTypeFieldNames, fieldTagsToNames[reqFieldNames[i]]) // TODO: need to find the the structure field name here
		}
	}

	// in the outTypeFieldNames requested order, store the data for output
	for i := range outTypeFieldNames {
		val := fmt.Sprintf("%v", sVals.FieldByName(outTypeFieldNames[i]))
		oVals = append(oVals, val)
		eHeaders = append(eHeaders, outTypeFieldTags[i])
	}

	// create a list of max field widths so we can format the output columns
	maxLens := []int{}
	for i := 0; i < len(oVals); i++ {
		if len(oVals[i]) > len(eHeaders[i]) {
			maxLens = append(maxLens, len(oVals[i]))
		} else {
			maxLens = append(maxLens, len(eHeaders[i]))
		}
	}

	// output
	if printHeader {
		header := ""
		for i := 0; i < len(eHeaders); i++ {
			header += fmt.Sprintf("%s", rightPad2Len(eHeaders[i], " ", maxLens[i]+2))
		}
		fmt.Printf("%s\n", header)
	}
	row := ""
	for i := 0; i < len(oVals); i++ {
		row += fmt.Sprintf("%s", rightPad2Len(oVals[i], " ", maxLens[i]+2))
	}
	fmt.Printf("%s\n", row)

	return nil
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// ArrayContains returns true if []string contains the target string
func ArrayContains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
