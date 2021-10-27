package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// OPrintFormatted prints formatted header or formatted header+values of a json string based on json field names
func OPrintFormatted(reqFieldNames []string, jsonString string, printHeader bool) error {

	errMessage := ""           // error messages, because I like to save things (okay, for printing later)
	maxLens := []int{}         // max names/values lengths, for printing column widths
	reportRows := [][]string{} // each row will be an array of values (a report row) based on a resultMap item

	// flag, if converion to an array is successful (else it's a map)
	resultIsArray := false

	// attempt conversion of result to an ARRAY of maps
	var resultArray []map[string]interface{}
	errToArray := json.Unmarshal([]byte(jsonString), &resultArray) // convert/unmarshal
	if errToArray == nil {
		//success?, flag as array
		resultIsArray = true
	}

	// attempt conversion to a map
	var resultMap map[string]interface{} // result is going to be the map object created from the json string
	errToMap := json.Unmarshal([]byte(jsonString), &resultMap)
	if errToMap == nil {
		// success? flag NOT array
		resultIsArray = false
	}

	// did both unmarshal attempts fail?, quit here
	if errToMap != nil && errToArray != nil {
		return errors.New("failed to unmarshal API response")
	}

	// determine if the map contains an uploads API files array response
	// if so, we need to extract the ARRAY of maps from this map
	_, ok := resultMap["files"]
	if ok {
		if len(resultMap) == 1 {
			var resultMapMap map[string][]map[string]interface{} // note that this is an array of maps, EMBEDED in a map
			errToMap := json.Unmarshal([]byte(jsonString), &resultMapMap)
			if errToMap != nil {
				return errors.New("failed to unmarshal uploads API response into an array of maps")
			}
			resultArray = resultMapMap["files"]
			// we'd flagged this as NOT an array, but we have now extracted an array of maps.
			resultIsArray = true
		}
	}

	// if API response was an array
	if resultIsArray {
		// get the available field names based on requested field names (user may typo or ask for non existant fields)
		reportHeaders, err := ArrayToMapNamesIntersection(reqFieldNames, resultArray[0])
		if err != nil {
			errMessage += err.Error()
		}
		// set the initial max col lengths to the column/field names
		for _, name := range reportHeaders {
			maxLens = append(maxLens, len(fmt.Sprintf("%v", name)))
		}

		for _, resultMap := range resultArray {
			// store a report row here
			row := []string{}
			for i, name := range reportHeaders {
				// store this value in its report row
				row = append(row, fmt.Sprintf("%v", resultMap[name]))
				// update the record of longest items in a column
				if len(fmt.Sprintf("%v", resultMap[name])) > maxLens[i] {
					maxLens[i] = len(fmt.Sprintf("%v", resultMap[name]))
				}
			}
			reportRows = append(reportRows, row)
		}
		printReport(reportHeaders, reportRows, maxLens)
		// if API response was a map
	} else {

		// get the available field names based on requested field names (user may typo or ask for non existant fields)
		reportHeaders, err := ArrayToMapNamesIntersection(reqFieldNames, resultMap)
		if err != nil {
			errMessage += err.Error()
		}
		// set the initial max col lengths to the column/field names
		for _, name := range reportHeaders {
			maxLens = append(maxLens, len(fmt.Sprintf("%v", name)))
		}
		row := []string{}
		for i, name := range reportHeaders {
			// store this value in its report row
			row = append(row, fmt.Sprintf("%v", resultMap[name]))
			// update the record of longest items in a column
			if len(fmt.Sprintf("%v", resultMap[name])) > maxLens[i] {
				maxLens[i] = len(fmt.Sprintf("%v", resultMap[name]))
			}
		}
		reportRows = append(reportRows, row)
		printReport(reportHeaders, reportRows, maxLens)
	}

	// TODO: Get the valid field names and set the initial values for maxLens

	// use sampleitem to validate that the requested field names exist, build reportFieldNames (validated requested names)
	// IF reqFieldName is in sampleItem
	// THEN
	//		append(reportFieldNames, reqFieldName)
	// 		set maxLens as field names are validated
	// ELSE
	//		append(errMessage...)

	// TODO: step through the json item(s) and build each report line to print
	// IF 'result' is an array
	// THEN
	// 		step through each and build a line of each result[reportFiledNames] value
	//		if the len(value) > maxLens[i], set maxLens[i] to value
	// ELSE just build a line from 'result'

	// check api response for the requested fieldnames, validate that they exist
	// also, set the initial maxLens of each column, at least wide enough for the column names

	// Print the header row
	// For each reportLine, fmt.Printf("%s ", rightPad2Len(line, " ", maxLens[i]+2))

	return nil
}

// ArrayToMapNamesIntersection returns the intersection of two arrays
func ArrayToMapNamesIntersection(reqFieldNames []string, resultMap map[string]interface{}) ([]string, error) {
	intersectionNames := []string{}
	errorMessage := ""
	for _, name := range reqFieldNames {
		_, ok := resultMap[name]
		if ok {
			intersectionNames = append(intersectionNames, name)
		} else {
			errorMessage += fmt.Sprintf("could not find field name '%s' in API response\n", name)
		}
	}
	return intersectionNames, nil
}

func printReport(reportHeaders []string, reportRows [][]string, maxColLengths []int) {
	for i, header := range reportHeaders {
		fmt.Printf(fmt.Sprintf("%s", rightPad2Len(header, " ", maxColLengths[i]+2)))
	}
	fmt.Println("")

	for _, row := range reportRows {

		for i, value := range row {
			fmt.Printf(fmt.Sprintf("%s", rightPad2Len(value, " ", maxColLengths[i]+2)))
		}

		fmt.Println("")
	}
}

// TODO: Update this to handle printing nested structure fields, specifically consider the uploads response which it hides in a 'file' object

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
