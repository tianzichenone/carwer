package parser

import (
	"regexp"
	"carwer/enginee"
	"fmt"
	"carwer/model"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([^<]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var statusRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var eduRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var placeRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)

func ParserProfile(b []byte) enginee.ParserResult {
	parserResult := enginee.ParserResult{}
	profile := model.Profile{}
	age, err := strconv.Atoi(fetchProfileInfo(ageRe, b))
	if err != nil {
		age = 0
	}
	profile.Age = age
	profile.Height = fetchProfileInfo(heightRe, b)
	profile.Income = fetchProfileInfo(incomeRe, b)
	profile.Status = fetchProfileInfo(statusRe, b)
	profile.Edu = fetchProfileInfo(eduRe, b)
	profile.Occupation = fetchProfileInfo(occupationRe, b)
	profile.Place = fetchProfileInfo(placeRe, b)
	profile.Hukou = fetchProfileInfo(hukouRe, b)
	profile.Constellation = fetchProfileInfo(constellationRe, b)
	parserResult.Items = append(parserResult.Items, profile)
	fmt.Printf("%v\n", profile)
	return parserResult
}

func fetchProfileInfo(re *regexp.Regexp, contents []byte) string {
	match := re.FindSubmatch(contents)
	return string(match[1])
}
