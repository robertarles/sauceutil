package cmd

import (
	"fmt"
	"reflect"
	"strings"
)

// Iprint experimenting with handling arbitrary structs
func Oprint(strct interface{}, reqFieldNames []string) error {

	sType := reflect.TypeOf(strct)
	sVals := reflect.Indirect(reflect.ValueOf(strct))
	oVals := []string{}
	eHeaders := []string{}

	// clean up fullTag by removing 'json:' and the quotes
	replacements := strings.NewReplacer("json:", "", "\"", "")

	// get a list of json name equiv for each of the struct field names
	allTypeFieldTags := []string{}
	allTypeFieldNames := []string{}
	for i := 0; i < sType.NumField(); i++ {
		jsonName := replacements.Replace(fmt.Sprintf("%v", sType.Field(i).Tag))
		allTypeFieldTags = append(allTypeFieldTags, jsonName)
		allTypeFieldNames = append(allTypeFieldNames, sType.Field(i).Name)
	}

	// get a list of the types fieldnames that match those requested for output
	outTypeFieldTags := []string{}  // the json field name
	outTypeFieldNames := []string{} // the json field name
	for i := 0; i < sType.NumField(); i++ {
		if ArrayContains(reqFieldNames, allTypeFieldTags[i]) {
			outTypeFieldTags = append(outTypeFieldTags, allTypeFieldTags[i])
			outTypeFieldNames = append(outTypeFieldNames, allTypeFieldNames[i])
		}
	}

	// in the outFieldNames requested order, store the data for output
	for i := range reqFieldNames {
		if ArrayContains(outTypeFieldTags, reqFieldNames[i]) {
			val := fmt.Sprintf("%v", sVals.FieldByName(outTypeFieldNames[i]))
			oVals = append(oVals, val)
			eHeaders = append(eHeaders, reqFieldNames[i])
		}
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
	header := ""
	for i := 0; i < len(eHeaders); i++ {
		header += fmt.Sprintf("%s", rightPad2Len(eHeaders[i], " ", maxLens[i]+2))
	}
	fmt.Printf("%s\n", header)
	row := ""
	for i := 0; i < len(oVals); i++ {
		row += fmt.Sprintf("%s", rightPad2Len(oVals[i], " ", maxLens[i]+2))
	}
	fmt.Printf("%s\n", row)

	return nil
}
