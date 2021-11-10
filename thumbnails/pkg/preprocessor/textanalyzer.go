package preprocessor

import (
	"unicode"
)

// Default list of scripts to be analyzed within the string.
//
// Scripts that aren't present in the list will be considered as part
// of the last "known" script. For example, if "Avestan" script (which isn't
// present) is preceeded by "Arabic" script, then the "Avestan" script will
// be considered as "Arabic"
//
// Punctuation symbols are usually considered part of the "Common" script
var DefaultScripts = []string{
	"Arabic",
	"Common",
	"Devanagari",
	"Han",
	"Hangul",
	"Hiragana",
	"Inherited",
	"Katakana",
	"Latin",
}

// Convenient map[string]map[string]string type used to merge multiple
// scripts into one. This is mainly used for japanese language which uses
// "Han", "Hiragana" and "Katakana" scripts.
//
// The map contains the expected previous script as first key, the expected
// current script as second key, and the resulting script (if both keys
// match) as value
type MergeMap map[string]map[string]string

// The default mergeMap containing info for the japanese scripts
var DefaultMergeMap = MergeMap{
	"Han": map[string]string{
		"Hiragana": "Hiragana",
		"Katakana": "Katakana",
	},
	"Hiragana": map[string]string{
		"Han":      "Hiragana",
		"Katakana": "Hiragana",
	},
	"Katakana": map[string]string{
		"Han":      "Katakana",
		"Hiragana": "Hiragana",
	},
}

// Analysis options.
type AnalysisOpts struct {
	UseMergeMap bool
	MergeMap    MergeMap
}

// A script range. The range should be attached to a string which could contain
// multiple scripts. The "TargetScript" will go from bytes "Low" to "High"
// (both inclusive), and contains a "RuneCount" number of runes or chars
// (mostly for debugging purposes).
// The Space contains the bytes (inside the range) that are considered as
// white space.
type ScriptRange struct {
	Low, High    int
	Spaces       []int
	TargetScript string
	RuneCount    int
}

// The result of a text analysis. It contains the analyzed text, a list of
// script ranges (see the ScriptRange type) and a map containing how many
// runes have been detected for a particular script.
type TextAnalysis struct {
	ScriptRanges []ScriptRange
	RuneCount    map[string]int
	Text         string
}

// The TextAnalyzer object contains private members. It should be created via
// "NewTextAnalyzer" function.
type TextAnalyzer struct {
	scripts         map[string]*unicode.RangeTable
	scriptListCache []string
}

// Create a new TextAnalyzer. A list of scripts must be provided.
// You can use the "DefaultScripts" variable for a default list,
// although it doesn't contain all the available scripts.
// See the unicode.Scripts variable (in the unicode package) for a
// full list. Note that using invalid scripts will cause an undefined
// behavior
func NewTextAnalyzer(scriptList []string) TextAnalyzer {
	scriptRanges := make(map[string]*unicode.RangeTable, len(scriptList))
	for _, script := range scriptList {
		scriptRanges[script] = unicode.Scripts[script]
	}
	return TextAnalyzer{
		scripts:         scriptRanges,
		scriptListCache: scriptList,
	}
}

// Analyze the target string using the specified options.
// A TextAnalysis will be returned with the result of the analysis.
func (ta *TextAnalyzer) AnalyzeString(word string, opts AnalysisOpts) TextAnalysis {
	analysis := TextAnalysis{
		RuneCount: make(map[string]int),
		Text:      word,
	}
	var lastRange *ScriptRange

	runeCount := 0
	for wordIndex, char := range word {
		script := "_unknown"
		for scriptIndex, scriptFound := range ta.scriptListCache {
			// if we can't match with a known script, do nothing and jump to the next char
			if unicode.Is(ta.scripts[scriptFound], char) {
				if scriptIndex > 3 {
					// we might expect more chars with the same script
					// so move the script first to match it faster next time
					ta.reorderScriptList(scriptFound)
				}
				script = scriptFound
			}
		}

		isWhiteSpace := unicode.Is(unicode.White_Space, char)
		if lastRange == nil {
			runeCount = 1
			lastRange = &ScriptRange{
				Low:          wordIndex,
				Spaces:       make([]int, 0),
				TargetScript: script,
			}
		} else {
			if script != lastRange.TargetScript {
				if opts.UseMergeMap {
					// This option mainly target japanese chars; multiple scripts can be used
					// in the same piece of text (Han, Hiragana and Katakana)
					// Instead of starting a new range, adjust the target script of the last range
					if expCurrent, currentOk := opts.MergeMap[lastRange.TargetScript]; currentOk {
						if expFinal, finalOk := expCurrent[script]; finalOk {
							lastRange.TargetScript = expFinal
							if isWhiteSpace {
								// TODO: Check if this is dead code.
								// whitespace should be part of the "Common" script, and the Common
								// script shouldn't be part of a mergeMap
								lastRange.Spaces = append(lastRange.Spaces, wordIndex)
							}
							runeCount++
							continue
						}
					}
				}

				lastRange.High = wordIndex - 1
				lastRange.RuneCount = runeCount
				analysis.ScriptRanges = append(analysis.ScriptRanges, *lastRange)
				if _, exists := analysis.RuneCount[lastRange.TargetScript]; !exists {
					analysis.RuneCount[lastRange.TargetScript] = 0
				}
				analysis.RuneCount[lastRange.TargetScript] += runeCount
				lastRange = &ScriptRange{
					Low:          wordIndex,
					Spaces:       make([]int, 0),
					TargetScript: script,
				}
				runeCount = 1
			} else {
				runeCount++
			}
		}
		if isWhiteSpace {
			lastRange.Spaces = append(lastRange.Spaces, wordIndex)
		}
	}

	if lastRange != nil {
		// close the last range
		lastRange.High = len(word) - 1
		lastRange.RuneCount = runeCount
		analysis.RuneCount[lastRange.TargetScript] += runeCount
		analysis.ScriptRanges = append(analysis.ScriptRanges, *lastRange)
	}
	return analysis
}

// Reorder the scriptListCache in the TextAnalyzer in order to speed up
// the next script searches. A "Latin" script is expected to be surrounded
// by "Latin" chars, although "Common" script chars might be present too
func (ta *TextAnalyzer) reorderScriptList(matchedScript string) {
	for index, script := range ta.scriptListCache {
		if script == matchedScript {
			if index != 0 {
				// move the script to the first position for a faster matching
				newList := append([]string{script}, ta.scriptListCache[:index]...)
				ta.scriptListCache = append(newList, ta.scriptListCache[index+1:]...)
			}
			// if index == 0 there is nothing to do: the element is already the first
			break
		}
	}
}

// Change the "Common" script to the one used in the previous script range.
// The ranges will be readjusted and merged if they're adjacent.
// This naive approach should be good enough for normal use cases
//
// The MergeMap is needed in case of the japanese language: the ranges
// "Han"-"Common"-"Katakana" might be replaced to "Han"-"Hiragana"-"Katakana"
// However, the ranges should be merged together into a big "Hiragana" range.
// If the MergeMap isn't needed, use an empty one
func (tr *TextAnalysis) MergeCommon(mergeMap MergeMap) {
	var finalRanges []ScriptRange
	var previousRange *ScriptRange
	for _, sRange := range tr.ScriptRanges {
		if previousRange != nil {
			if previousRange.TargetScript == sRange.TargetScript {
				previousRange.High = sRange.High
				previousRange.Spaces = append(previousRange.Spaces, sRange.Spaces...)
				previousRange.RuneCount += sRange.RuneCount
			} else if sRange.TargetScript == "Common" || sRange.TargetScript == "Inherited" {
				// new range will be absorbed into the previous one
				previousRange.High = sRange.High
				previousRange.Spaces = append(previousRange.Spaces, sRange.Spaces...)
				previousRange.RuneCount += sRange.RuneCount
				tr.RuneCount[previousRange.TargetScript] += sRange.RuneCount
				tr.RuneCount[sRange.TargetScript] -= sRange.RuneCount
			} else if previousRange.TargetScript == "Common" || previousRange.TargetScript == "Inherited" {
				// might happen if the text starts with a Common script
				previousRange.High = sRange.High
				previousRange.Spaces = append(previousRange.Spaces, sRange.Spaces...)
				tr.RuneCount[sRange.TargetScript] += previousRange.RuneCount
				tr.RuneCount[previousRange.TargetScript] -= previousRange.RuneCount
				previousRange.RuneCount += sRange.RuneCount
				previousRange.TargetScript = sRange.TargetScript
			} else {
				if expCurrent, currentOk := mergeMap[previousRange.TargetScript]; currentOk {
					if expFinal, finalOk := expCurrent[sRange.TargetScript]; finalOk {
						if sRange.TargetScript == expFinal {
							// the previous range has changed the target script
							tr.RuneCount[previousRange.TargetScript] -= previousRange.RuneCount
							tr.RuneCount[sRange.TargetScript] += previousRange.RuneCount
						} else {
							// new range has been absorbed
							tr.RuneCount[sRange.TargetScript] -= sRange.RuneCount
							tr.RuneCount[previousRange.TargetScript] += sRange.RuneCount
						}
						previousRange.TargetScript = expFinal
						previousRange.High = sRange.High
						previousRange.Spaces = append(previousRange.Spaces, sRange.Spaces...)
						previousRange.RuneCount += sRange.RuneCount
						continue
					}
				}
				finalRanges = append(finalRanges, *previousRange)
				*previousRange = sRange
			}
		} else {
			previousRange = &ScriptRange{}
			*previousRange = sRange
		}
	}

	finalRanges = append(finalRanges, *previousRange)
	tr.ScriptRanges = finalRanges
	delete(tr.RuneCount, "Common")
	delete(tr.RuneCount, "Inherited")
	for index, rCount := range tr.RuneCount {
		if rCount == 0 {
			delete(tr.RuneCount, index)
		}
	}
}
