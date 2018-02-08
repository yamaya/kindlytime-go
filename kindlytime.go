/*
	kindlytimeパッケージは時刻の簡易表記を扱うパッケージです

	"1 days ago" という文字列からtime.Time構造体を得ることができます。

	t := kindlytime.Parse("1 days ago")

	可能な表記は
	- "days" or "d"
	- "hours" or "h"
	- "minutes" or "m"
	- "seconds" or "s"
	です。各々、単数系を記述する事も可能です（e.g. "1 day"）。

	また、"1 hour 30 minutes" というような組み合わせも可能です。
	数値と単位の間の空白はなくても構いません。よって、"1h 30m" という表記も可能です。

	NOTE: 現状、文字列からTime構造体への変換機能しかありません。
*/
package kindlytime

import (
	"regexp"
	"time"
	"strconv"
	"strings"
	"errors"
	"fmt"
)

const verbose = false

// cellの種類を表すenum
type cellType int
const (
	Unknown cellType = iota + 1
	Day
	Hour
	Minute
	Second
)

// cellTypeの文字列表現を返す関数
func (this cellType) String() string {
	switch this {
	case Day: return "d"
	case Hour: return "h"
	case Minute: return "m"
	case Second: return "s"
	}
	return "unknown"
}

// cellTypeに対応するtime.Duration値を返す関数
func (this cellType) toDuration() time.Duration {
	switch this {
	case Day: return time.Hour * 24
	case Hour: return time.Hour
	case Minute: return time.Minute
	case Second: return time.Second
	}
	return 0
}

// 文字列からcellTypeを生成する関数
func NewCellFromString(s string) (cellType, error) {
	switch s[0] {
	case 'd': return Day, nil
	case 'h': return Hour, nil
	case 'm': return Minute, nil
	case 's': return Second, nil
	}
	return Unknown, errors.New("Unrecognized type description")
}

// cell型
type cell struct {
	Type  cellType 	// 種類 (day|hour|minute|second)
	Value int64		// 値
}

// cell型の文字列表現を返す関数
func (this cell) String() string {
	return fmt.Sprintf("cell{Type:%s, Value:%d}", this.Type, this.Value)
}

// cell型をtime.Durationに変換する関数
func (this cell) toDuration() time.Duration {
	return this.Type.toDuration() * time.Duration(this.Value)
}

// パース関数
func Parse(input string, origin time.Time) (time.Time, error) {
	result := origin
	input = strings.TrimSpace(input)

	if input == "now" {
		return time.Now(), nil
	}

	tensePattern := regexp.MustCompile(`\s+(ago|later)$`)
	tense := tensePattern.FindAllStringSubmatch(input, -1)
	if verbose {
		fmt.Printf("tensePattern=%#v\n", tense)
	}
	past := (len(tense) == 1 && len(tense[0]) == 2 && tense[0][1] == "ago") // デフォルトはlater
	if verbose {
		fmt.Printf("ago: %v\n", past)
	}

	cellPattern := regexp.MustCompile(`(\d+)\s*(d(?:ays?)?|h(?:ours?)?|m(?:inutes?)?|s(?:econds?)?)`)
	groups := cellPattern.FindAllStringSubmatch(input, -1)
	if len(groups) != 0 {
		cells := []cell{}
		for i, group := range(groups) {
			if len(group) == 3 {
				t, _ := NewCellFromString(group[2])
				v, _ := strconv.ParseInt(group[1], 10, 64)
				if past {
					v = v * -1
				}
				cells = append(cells, cell{t, v})
			} else {
				if verbose {
					fmt.Printf("\tUnrecognized match. [%d] = `%#v`\n", i, group)
				}
				return result, fmt.Errorf("Unrecognized input \"%s\"", input)
			}
		}
		if verbose {
			fmt.Printf("\tcells: %v\n", cells)
		}
		for _, cell := range(cells) {
			result = result.Add(cell.toDuration())
		}
	} else {
		if verbose {
			fmt.Println("\tDoes not match")
		}
		r, err := time.Parse(time.RFC3339, input)
		if err != nil {
			fmt.Errorf("Unrecognized input \"%s\" -- %s", input, err)
		}
		if verbose {
			fmt.Println(r)
		}
		result = r
		return result, err
	}

	return result, nil
}

// パース関数（現在時刻を使用）
func ParseBaseOnCurrentTime(input string) (time.Time, error) {
	return Parse(input, time.Now())
}
