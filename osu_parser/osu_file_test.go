package osu_parser_test

import (
	"fmt"
	"testing"

	"github.com/Waffle-osu/osu-parser/osu_parser"
)

func TestOsuParser(t *testing.T) {
	parsedOsuFile, _ := osu_parser.ParseFile("../cases/COOL&CREATE - サトリムソウ (Furball) [Insane].osu")

	if parsedOsuFile.Md5Hash != "0dfcc1b4a695fac58bfc15782ad65fde" {
		t.Fail()
	}

	if len(parsedOsuFile.ParserWarnings) != 0 {
		for _, warning := range parsedOsuFile.ParserWarnings {
			fmt.Printf("%s", warning)
		}
	}
}
