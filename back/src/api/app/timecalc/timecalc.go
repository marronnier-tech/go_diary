package timecalc

import (
	"fmt"
	stc "strconv"
	"time"

	"github.com/lib/pq"
)

func DifftoNow(previous pq.NullTime) (res string) {

	if previous.Valid == false {
		res = "達成日はありません"
	} else {
		duration := time.Now().YearDay() - previous.Time.YearDay()

		if duration == 0 {
			res = "今日"
		} else {
			res = fmt.Sprintf("%s日前", stc.Itoa(duration))
		}

	}

	return

}
